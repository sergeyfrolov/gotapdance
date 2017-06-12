package tdproxy

import (
	"errors"
	"github.com/SergeyFrolov/gotapdance/tapdance"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

// Connection-oriented state
type tapDanceFlow struct {
	// tunnel index and start time
	id      uint64
	startMs time.Time

	// reference to global proxy
	proxy *TapDanceProxy

	servConn net.Conn // can cast to tapdance.Conn but don't need to
	userConn net.Conn
}

func makeTapDanceFlow(proxy *TapDanceProxy, id uint64) *tapDanceFlow {
	state := new(tapDanceFlow)

	state.proxy = proxy
	state.id = id

	state.startMs = time.Now()

	Logger.Debugf("Created new TDState ", state)
	return state
}

func (TDstate *tapDanceFlow) redirect() error {
	conn, err := tapdance.DialSOCKS()
	if err != nil {
		TDstate.userConn.Close()
		return err
	}
	TDstate.servConn = conn
	errChan := make(chan error)
	defer func() {
		TDstate.userConn.Close()
		TDstate.servConn.Close()
		_ = <-errChan // wait for second goroutine to close
	}()

	forwardFromServerToClient := func() {
		n, _err := io.Copy(TDstate.userConn, TDstate.servConn)
		Logger.Debugf("{tapDanceFlow} forwardFromServerToClient returns, bytes sent: " +
			strconv.FormatUint(uint64(n), 10))
		if _err == nil {
			_err = errors.New("server returned without error")
		}
		errChan <- _err
		return
	}

	forwardFromClientToServer := func() {
		n, _err := io.Copy(TDstate.servConn, TDstate.userConn)
		Logger.Debugf("{tapDanceFlow} forwardFromClientToServer returns, bytes sent: " +
			strconv.FormatUint(uint64(n), 10))
		if _err == nil {
			_err = errors.New("StoppedByUser")
		}
		errChan <- _err
		return
	}

	go forwardFromServerToClient()
	go forwardFromClientToServer()

	if err = <-errChan; err != nil {
		if err.Error() == "MSG_CLOSE" || err.Error() == "StoppedByUser" {
			Logger.Debugln("[Session " + strconv.FormatUint(uint64(TDstate.id), 10) +
				" Redirect function returns gracefully: " + err.Error())
			TDstate.proxy.closedGracefully.Inc()
			err = nil
		} else {
			str_err := err.Error()

			// statistics
			if strings.Contains(str_err, "TapDance station didn't pick up the request") {
				TDstate.proxy.notPickedUp.Inc()
			} else if strings.Contains(str_err, ": i/o timeout") {
				TDstate.proxy.timedOut.Inc()
			} else {
				TDstate.proxy.unexpectedError.Inc()
			}
		}
	}
	return err
}
