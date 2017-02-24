package tapdance

import "sync/atomic"

type counter_uint64 struct {
	value uint64
}

func (c *counter_uint64) inc() uint64 {
	atomic.AddUint64(&c.value, uint64(1))
	return c.value
}

func (c *counter_uint64) dec() uint64 {
	atomic.AddUint64(&c.value, ^uint64(0))
	return c.value
}

func (c *counter_uint64) get() uint64 {
	return c.value
}
