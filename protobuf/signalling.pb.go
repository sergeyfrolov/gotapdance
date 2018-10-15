// Code generated by protoc-gen-go.
// source: signalling.proto
// DO NOT EDIT!

/*
Package tdproto is a generated protocol buffer package.

It is generated from these files:
	signalling.proto

It has these top-level messages:
	PubKey
	TLSDecoySpec
	ClientConf
	DecoyList
	StationToClient
	ClientToStation
	SessionStats
*/
package tdproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyType int32

const (
	KeyType_AES_GCM_128 KeyType = 90
	KeyType_AES_GCM_256 KeyType = 91
)

var KeyType_name = map[int32]string{
	90: "AES_GCM_128",
	91: "AES_GCM_256",
}
var KeyType_value = map[string]int32{
	"AES_GCM_128": 90,
	"AES_GCM_256": 91,
}

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}
func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}
func (x *KeyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(KeyType_value, data, "KeyType")
	if err != nil {
		return err
	}
	*x = KeyType(value)
	return nil
}
func (KeyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// State transitions of the client
type C2S_Transition int32

const (
	C2S_Transition_C2S_NO_CHANGE                C2S_Transition = 0
	C2S_Transition_C2S_SESSION_INIT             C2S_Transition = 1
	C2S_Transition_C2S_SESSION_COVERT_INIT      C2S_Transition = 11
	C2S_Transition_C2S_EXPECT_RECONNECT         C2S_Transition = 2
	C2S_Transition_C2S_SESSION_CLOSE            C2S_Transition = 3
	C2S_Transition_C2S_YIELD_UPLOAD             C2S_Transition = 4
	C2S_Transition_C2S_ACQUIRE_UPLOAD           C2S_Transition = 5
	C2S_Transition_C2S_EXPECT_UPLOADONLY_RECONN C2S_Transition = 6
	C2S_Transition_C2S_ERROR                    C2S_Transition = 255
)

var C2S_Transition_name = map[int32]string{
	0:   "C2S_NO_CHANGE",
	1:   "C2S_SESSION_INIT",
	11:  "C2S_SESSION_COVERT_INIT",
	2:   "C2S_EXPECT_RECONNECT",
	3:   "C2S_SESSION_CLOSE",
	4:   "C2S_YIELD_UPLOAD",
	5:   "C2S_ACQUIRE_UPLOAD",
	6:   "C2S_EXPECT_UPLOADONLY_RECONN",
	255: "C2S_ERROR",
}
var C2S_Transition_value = map[string]int32{
	"C2S_NO_CHANGE":                0,
	"C2S_SESSION_INIT":             1,
	"C2S_SESSION_COVERT_INIT":      11,
	"C2S_EXPECT_RECONNECT":         2,
	"C2S_SESSION_CLOSE":            3,
	"C2S_YIELD_UPLOAD":             4,
	"C2S_ACQUIRE_UPLOAD":           5,
	"C2S_EXPECT_UPLOADONLY_RECONN": 6,
	"C2S_ERROR":                    255,
}

func (x C2S_Transition) Enum() *C2S_Transition {
	p := new(C2S_Transition)
	*p = x
	return p
}
func (x C2S_Transition) String() string {
	return proto.EnumName(C2S_Transition_name, int32(x))
}
func (x *C2S_Transition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(C2S_Transition_value, data, "C2S_Transition")
	if err != nil {
		return err
	}
	*x = C2S_Transition(value)
	return nil
}
func (C2S_Transition) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// State transitions of the server
type S2C_Transition int32

const (
	S2C_Transition_S2C_NO_CHANGE           S2C_Transition = 0
	S2C_Transition_S2C_SESSION_INIT        S2C_Transition = 1
	S2C_Transition_S2C_SESSION_COVERT_INIT S2C_Transition = 11
	S2C_Transition_S2C_CONFIRM_RECONNECT   S2C_Transition = 2
	S2C_Transition_S2C_SESSION_CLOSE       S2C_Transition = 3
	// TODO should probably also allow EXPECT_RECONNECT here, for DittoTap
	S2C_Transition_S2C_ERROR S2C_Transition = 255
)

var S2C_Transition_name = map[int32]string{
	0:   "S2C_NO_CHANGE",
	1:   "S2C_SESSION_INIT",
	11:  "S2C_SESSION_COVERT_INIT",
	2:   "S2C_CONFIRM_RECONNECT",
	3:   "S2C_SESSION_CLOSE",
	255: "S2C_ERROR",
}
var S2C_Transition_value = map[string]int32{
	"S2C_NO_CHANGE":           0,
	"S2C_SESSION_INIT":        1,
	"S2C_SESSION_COVERT_INIT": 11,
	"S2C_CONFIRM_RECONNECT":   2,
	"S2C_SESSION_CLOSE":       3,
	"S2C_ERROR":               255,
}

func (x S2C_Transition) Enum() *S2C_Transition {
	p := new(S2C_Transition)
	*p = x
	return p
}
func (x S2C_Transition) String() string {
	return proto.EnumName(S2C_Transition_name, int32(x))
}
func (x *S2C_Transition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(S2C_Transition_value, data, "S2C_Transition")
	if err != nil {
		return err
	}
	*x = S2C_Transition(value)
	return nil
}
func (S2C_Transition) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Should accompany all S2C_ERROR messages.
type ErrorReasonS2C int32

const (
	ErrorReasonS2C_NO_ERROR         ErrorReasonS2C = 0
	ErrorReasonS2C_COVERT_STREAM    ErrorReasonS2C = 1
	ErrorReasonS2C_CLIENT_REPORTED  ErrorReasonS2C = 2
	ErrorReasonS2C_CLIENT_PROTOCOL  ErrorReasonS2C = 3
	ErrorReasonS2C_STATION_INTERNAL ErrorReasonS2C = 4
	ErrorReasonS2C_DECOY_OVERLOAD   ErrorReasonS2C = 5
	ErrorReasonS2C_CLIENT_STREAM    ErrorReasonS2C = 100
	ErrorReasonS2C_CLIENT_TIMEOUT   ErrorReasonS2C = 101
)

var ErrorReasonS2C_name = map[int32]string{
	0:   "NO_ERROR",
	1:   "COVERT_STREAM",
	2:   "CLIENT_REPORTED",
	3:   "CLIENT_PROTOCOL",
	4:   "STATION_INTERNAL",
	5:   "DECOY_OVERLOAD",
	100: "CLIENT_STREAM",
	101: "CLIENT_TIMEOUT",
}
var ErrorReasonS2C_value = map[string]int32{
	"NO_ERROR":         0,
	"COVERT_STREAM":    1,
	"CLIENT_REPORTED":  2,
	"CLIENT_PROTOCOL":  3,
	"STATION_INTERNAL": 4,
	"DECOY_OVERLOAD":   5,
	"CLIENT_STREAM":    100,
	"CLIENT_TIMEOUT":   101,
}

func (x ErrorReasonS2C) Enum() *ErrorReasonS2C {
	p := new(ErrorReasonS2C)
	*p = x
	return p
}
func (x ErrorReasonS2C) String() string {
	return proto.EnumName(ErrorReasonS2C_name, int32(x))
}
func (x *ErrorReasonS2C) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorReasonS2C_value, data, "ErrorReasonS2C")
	if err != nil {
		return err
	}
	*x = ErrorReasonS2C(value)
	return nil
}
func (ErrorReasonS2C) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PubKey struct {
	// A public key, as used by the station.
	Key              []byte   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Type             *KeyType `protobuf:"varint,2,opt,name=type,enum=tapdance.KeyType" json:"type,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PubKey) Reset()                    { *m = PubKey{} }
func (m *PubKey) String() string            { return proto.CompactTextString(m) }
func (*PubKey) ProtoMessage()               {}
func (*PubKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PubKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PubKey) GetType() KeyType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return KeyType_AES_GCM_128
}

type TLSDecoySpec struct {
	// The hostname/SNI to use for this host
	//
	// The hostname is the only required field, although other
	// fields are expected to be present in most cases.
	Hostname *string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	// The 32-bit ipv4 address, in network byte order
	//
	// If the IPv4 address is absent, then it may be resolved via
	// DNS by the client, or the client may discard this decoy spec
	// if local DNS is untrusted, or the service may be multihomed.
	Ipv4Addr *uint32 `protobuf:"fixed32,2,opt,name=ipv4addr" json:"ipv4addr,omitempty"`
	// The Tapdance station public key to use when contacting this
	// decoy
	//
	// If omitted, the default station public key (if any) is used.
	Pubkey *PubKey `protobuf:"bytes,3,opt,name=pubkey" json:"pubkey,omitempty"`
	// The maximum duration, in milliseconds, to maintain an open
	// connection to this decoy (because the decoy may close the
	// connection itself after this length of time)
	//
	// If omitted, a default of 30,000 milliseconds is assumed.
	Timeout *uint32 `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	// The maximum TCP window size to attempt to use for this decoy.
	//
	// If omitted, a default of 15360 is assumed.
	//
	// TODO: the default is based on the current heuristic of only
	// using decoys that permit windows of 15KB or larger.  If this
	// heuristic changes, then this default doesn't make sense.
	Tcpwin           *uint32 `protobuf:"varint,5,opt,name=tcpwin" json:"tcpwin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TLSDecoySpec) Reset()                    { *m = TLSDecoySpec{} }
func (m *TLSDecoySpec) String() string            { return proto.CompactTextString(m) }
func (*TLSDecoySpec) ProtoMessage()               {}
func (*TLSDecoySpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TLSDecoySpec) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *TLSDecoySpec) GetIpv4Addr() uint32 {
	if m != nil && m.Ipv4Addr != nil {
		return *m.Ipv4Addr
	}
	return 0
}

func (m *TLSDecoySpec) GetPubkey() *PubKey {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *TLSDecoySpec) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *TLSDecoySpec) GetTcpwin() uint32 {
	if m != nil && m.Tcpwin != nil {
		return *m.Tcpwin
	}
	return 0
}

type ClientConf struct {
	DecoyList        *DecoyList `protobuf:"bytes,1,opt,name=decoy_list" json:"decoy_list,omitempty"`
	Generation       *uint32    `protobuf:"varint,2,opt,name=generation" json:"generation,omitempty"`
	DefaultPubkey    *PubKey    `protobuf:"bytes,3,opt,name=default_pubkey" json:"default_pubkey,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ClientConf) Reset()                    { *m = ClientConf{} }
func (m *ClientConf) String() string            { return proto.CompactTextString(m) }
func (*ClientConf) ProtoMessage()               {}
func (*ClientConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientConf) GetDecoyList() *DecoyList {
	if m != nil {
		return m.DecoyList
	}
	return nil
}

func (m *ClientConf) GetGeneration() uint32 {
	if m != nil && m.Generation != nil {
		return *m.Generation
	}
	return 0
}

func (m *ClientConf) GetDefaultPubkey() *PubKey {
	if m != nil {
		return m.DefaultPubkey
	}
	return nil
}

type DecoyList struct {
	TlsDecoys        []*TLSDecoySpec `protobuf:"bytes,1,rep,name=tls_decoys" json:"tls_decoys,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DecoyList) Reset()                    { *m = DecoyList{} }
func (m *DecoyList) String() string            { return proto.CompactTextString(m) }
func (*DecoyList) ProtoMessage()               {}
func (*DecoyList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DecoyList) GetTlsDecoys() []*TLSDecoySpec {
	if m != nil {
		return m.TlsDecoys
	}
	return nil
}

type StationToClient struct {
	// Should accompany (at least) SESSION_INIT and CONFIRM_RECONNECT.
	ProtocolVersion *uint32 `protobuf:"varint,1,opt,name=protocol_version" json:"protocol_version,omitempty"`
	// There might be a state transition. May be absent; absence should be
	// treated identically to NO_CHANGE.
	StateTransition *S2C_Transition `protobuf:"varint,2,opt,name=state_transition,enum=tapdance.S2C_Transition" json:"state_transition,omitempty"`
	// The station can send client config info piggybacked
	// on any message, as it sees fit
	ConfigInfo *ClientConf `protobuf:"bytes,3,opt,name=config_info" json:"config_info,omitempty"`
	// If state_transition == S2C_ERROR, this field is the explanation.
	ErrReason *ErrorReasonS2C `protobuf:"varint,4,opt,name=err_reason,enum=tapdance.ErrorReasonS2C" json:"err_reason,omitempty"`
	// Signals client to stop connecting for following amount of seconds
	TmpBackoff *uint32 `protobuf:"varint,5,opt,name=tmp_backoff" json:"tmp_backoff,omitempty"`
	// Sent in SESSION_INIT, identifies the station that picked up
	StationId *string `protobuf:"bytes,6,opt,name=station_id" json:"station_id,omitempty"`
	// Random-sized junk to defeat packet size fingerprinting.
	Padding          []byte `protobuf:"bytes,100,opt,name=padding" json:"padding,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *StationToClient) Reset()                    { *m = StationToClient{} }
func (m *StationToClient) String() string            { return proto.CompactTextString(m) }
func (*StationToClient) ProtoMessage()               {}
func (*StationToClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StationToClient) GetProtocolVersion() uint32 {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return 0
}

func (m *StationToClient) GetStateTransition() S2C_Transition {
	if m != nil && m.StateTransition != nil {
		return *m.StateTransition
	}
	return S2C_Transition_S2C_NO_CHANGE
}

func (m *StationToClient) GetConfigInfo() *ClientConf {
	if m != nil {
		return m.ConfigInfo
	}
	return nil
}

func (m *StationToClient) GetErrReason() ErrorReasonS2C {
	if m != nil && m.ErrReason != nil {
		return *m.ErrReason
	}
	return ErrorReasonS2C_NO_ERROR
}

func (m *StationToClient) GetTmpBackoff() uint32 {
	if m != nil && m.TmpBackoff != nil {
		return *m.TmpBackoff
	}
	return 0
}

func (m *StationToClient) GetStationId() string {
	if m != nil && m.StationId != nil {
		return *m.StationId
	}
	return ""
}

func (m *StationToClient) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

type ClientToStation struct {
	ProtocolVersion *uint32 `protobuf:"varint,1,opt,name=protocol_version" json:"protocol_version,omitempty"`
	// The client reports its decoy list's version number here, which the
	// station can use to decide whether to send an updated one. The station
	// should always send a list if this field is set to 0.
	DecoyListGeneration *uint32         `protobuf:"varint,2,opt,name=decoy_list_generation" json:"decoy_list_generation,omitempty"`
	StateTransition     *C2S_Transition `protobuf:"varint,3,opt,name=state_transition,enum=tapdance.C2S_Transition" json:"state_transition,omitempty"`
	// The position in the overall session's upload sequence where the current
	// YIELD=>ACQUIRE switchover is happening.
	UploadSync *uint64 `protobuf:"varint,4,opt,name=upload_sync" json:"upload_sync,omitempty"`
	// List of decoys that client have unsuccessfully tried in current session.
	// Could be sent in chunks
	FailedDecoys []string      `protobuf:"bytes,10,rep,name=failed_decoys" json:"failed_decoys,omitempty"`
	Stats        *SessionStats `protobuf:"bytes,11,opt,name=stats" json:"stats,omitempty"`
	// Station is only required to check this variable during session initialization.
	// If set, station must facilitate connection to said target by itself, i.e. write into squid
	// socket an HTTP/SOCKS/any other connection request.
	// covert_address must have exactly one ':' colon, that separates host (literal IP address or
	// resolvable hostname) and port
	// TODO: make it required for initialization, and stop connecting any client straight to squid?
	CovertAddress *string `protobuf:"bytes,20,opt,name=covert_address" json:"covert_address,omitempty"`
	// Random-sized junk to defeat packet size fingerprinting.
	Padding          []byte `protobuf:"bytes,100,opt,name=padding" json:"padding,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientToStation) Reset()                    { *m = ClientToStation{} }
func (m *ClientToStation) String() string            { return proto.CompactTextString(m) }
func (*ClientToStation) ProtoMessage()               {}
func (*ClientToStation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClientToStation) GetProtocolVersion() uint32 {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return 0
}

func (m *ClientToStation) GetDecoyListGeneration() uint32 {
	if m != nil && m.DecoyListGeneration != nil {
		return *m.DecoyListGeneration
	}
	return 0
}

func (m *ClientToStation) GetStateTransition() C2S_Transition {
	if m != nil && m.StateTransition != nil {
		return *m.StateTransition
	}
	return C2S_Transition_C2S_NO_CHANGE
}

func (m *ClientToStation) GetUploadSync() uint64 {
	if m != nil && m.UploadSync != nil {
		return *m.UploadSync
	}
	return 0
}

func (m *ClientToStation) GetFailedDecoys() []string {
	if m != nil {
		return m.FailedDecoys
	}
	return nil
}

func (m *ClientToStation) GetStats() *SessionStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *ClientToStation) GetCovertAddress() string {
	if m != nil && m.CovertAddress != nil {
		return *m.CovertAddress
	}
	return ""
}

func (m *ClientToStation) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

type SessionStats struct {
	FailedDecoys *uint32 `protobuf:"varint,20,opt,name=failed_decoys" json:"failed_decoys,omitempty"`
	// Applicable to whole session:
	TotalTimeToConnect *int64 `protobuf:"varint,31,opt,name=total_time_to_connect" json:"total_time_to_connect,omitempty"`
	// Last (i.e. successful) decoy:
	RttToStation     *int64 `protobuf:"varint,33,opt,name=rtt_to_station" json:"rtt_to_station,omitempty"`
	TlsToDecoy       *int64 `protobuf:"varint,38,opt,name=tls_to_decoy" json:"tls_to_decoy,omitempty"`
	TcpToDecoy       *int64 `protobuf:"varint,39,opt,name=tcp_to_decoy" json:"tcp_to_decoy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SessionStats) Reset()                    { *m = SessionStats{} }
func (m *SessionStats) String() string            { return proto.CompactTextString(m) }
func (*SessionStats) ProtoMessage()               {}
func (*SessionStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SessionStats) GetFailedDecoys() uint32 {
	if m != nil && m.FailedDecoys != nil {
		return *m.FailedDecoys
	}
	return 0
}

func (m *SessionStats) GetTotalTimeToConnect() int64 {
	if m != nil && m.TotalTimeToConnect != nil {
		return *m.TotalTimeToConnect
	}
	return 0
}

func (m *SessionStats) GetRttToStation() int64 {
	if m != nil && m.RttToStation != nil {
		return *m.RttToStation
	}
	return 0
}

func (m *SessionStats) GetTlsToDecoy() int64 {
	if m != nil && m.TlsToDecoy != nil {
		return *m.TlsToDecoy
	}
	return 0
}

func (m *SessionStats) GetTcpToDecoy() int64 {
	if m != nil && m.TcpToDecoy != nil {
		return *m.TcpToDecoy
	}
	return 0
}

func init() {
	proto.RegisterType((*PubKey)(nil), "tapdance.PubKey")
	proto.RegisterType((*TLSDecoySpec)(nil), "tapdance.TLSDecoySpec")
	proto.RegisterType((*ClientConf)(nil), "tapdance.ClientConf")
	proto.RegisterType((*DecoyList)(nil), "tapdance.DecoyList")
	proto.RegisterType((*StationToClient)(nil), "tapdance.StationToClient")
	proto.RegisterType((*ClientToStation)(nil), "tapdance.ClientToStation")
	proto.RegisterType((*SessionStats)(nil), "tapdance.SessionStats")
	proto.RegisterEnum("tapdance.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("tapdance.C2S_Transition", C2S_Transition_name, C2S_Transition_value)
	proto.RegisterEnum("tapdance.S2C_Transition", S2C_Transition_name, S2C_Transition_value)
	proto.RegisterEnum("tapdance.ErrorReasonS2C", ErrorReasonS2C_name, ErrorReasonS2C_value)
}

func init() { proto.RegisterFile("signalling.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 861 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0x37, 0x6d, 0xda, 0x1c, 0x27, 0xce, 0x74, 0x9a, 0x16, 0x23, 0x40, 0x1b, 0x22, 0xc1,
	0x86, 0x82, 0x2a, 0x61, 0xc1, 0xc2, 0x6d, 0xe4, 0x0c, 0x4b, 0xb4, 0xa9, 0x1d, 0x6c, 0x17, 0x51,
	0xb8, 0x18, 0xb9, 0xf6, 0x24, 0x58, 0xeb, 0x7a, 0x2c, 0xcf, 0xa4, 0x28, 0x6f, 0xc0, 0x23, 0xf0,
	0x10, 0xbc, 0x17, 0xf7, 0x5c, 0x72, 0xc3, 0x6a, 0x6c, 0xa7, 0xf9, 0xe9, 0x6a, 0x2f, 0xfd, 0xcd,
	0x9c, 0xef, 0xef, 0x8c, 0x01, 0x89, 0x64, 0x91, 0x85, 0x69, 0x9a, 0x64, 0x8b, 0xab, 0xbc, 0xe0,
	0x92, 0xe3, 0x13, 0x19, 0xe6, 0x71, 0x98, 0x45, 0x6c, 0xf0, 0x12, 0x9a, 0xb3, 0xe5, 0xdd, 0x6b,
	0xb6, 0xc2, 0x3a, 0x34, 0xde, 0xb0, 0x95, 0xa9, 0xf5, 0xb5, 0x61, 0x1b, 0x3f, 0x87, 0x43, 0xb9,
	0xca, 0x99, 0x79, 0xd0, 0xd7, 0x86, 0x86, 0x75, 0x7a, 0xb5, 0xbe, 0x7f, 0xf5, 0x9a, 0xad, 0x82,
	0x55, 0xce, 0x06, 0x4b, 0x68, 0x07, 0x53, 0x7f, 0xcc, 0x22, 0xbe, 0xf2, 0x73, 0x16, 0x61, 0x04,
	0x27, 0xbf, 0x73, 0x21, 0xb3, 0xf0, 0x9e, 0x95, 0x14, 0x2d, 0x85, 0x24, 0xf9, 0xc3, 0x37, 0x61,
	0x1c, 0x17, 0x25, 0xcd, 0x31, 0xee, 0x43, 0x33, 0x5f, 0xde, 0x29, 0x91, 0x46, 0x5f, 0x1b, 0xea,
	0x16, 0xda, 0xd0, 0xd6, 0x1e, 0xba, 0x70, 0x2c, 0x93, 0x7b, 0xc6, 0x97, 0xd2, 0x3c, 0xec, 0x6b,
	0xc3, 0x0e, 0x36, 0xa0, 0x29, 0xa3, 0xfc, 0x8f, 0x24, 0x33, 0x8f, 0xd4, 0xf7, 0x40, 0x00, 0xd8,
	0x69, 0xc2, 0x32, 0x69, 0xf3, 0x6c, 0x8e, 0x5f, 0x00, 0xc4, 0xca, 0x01, 0x4d, 0x13, 0x21, 0x4b,
	0x59, 0xdd, 0x3a, 0xdb, 0x90, 0x96, 0xee, 0xa6, 0x89, 0x90, 0x18, 0x03, 0x2c, 0x58, 0xc6, 0x8a,
	0x50, 0x26, 0x3c, 0x2b, 0xdd, 0x74, 0xf0, 0x10, 0x8c, 0x98, 0xcd, 0xc3, 0x65, 0x2a, 0xe9, 0xfb,
	0x5d, 0x0d, 0xbe, 0x83, 0xd6, 0x86, 0xea, 0x12, 0x40, 0xa6, 0x82, 0x96, 0xba, 0xc2, 0xd4, 0xfa,
	0x8d, 0xa1, 0x6e, 0x5d, 0x6c, 0x46, 0xb6, 0x4b, 0x19, 0xfc, 0xab, 0x41, 0xd7, 0x97, 0xa5, 0x68,
	0xc0, 0x2b, 0xdf, 0xd8, 0x04, 0x54, 0xee, 0x20, 0xe2, 0x29, 0x7d, 0x60, 0x85, 0x50, 0x86, 0xb4,
	0xd2, 0x90, 0x05, 0x48, 0xc8, 0x50, 0x32, 0x2a, 0x8b, 0x30, 0x13, 0xc9, 0xa3, 0x55, 0xc3, 0x32,
	0x37, 0xfc, 0xbe, 0x65, 0xd3, 0xe0, 0xf1, 0x1c, 0x7f, 0x01, 0x7a, 0xc4, 0xb3, 0x79, 0xb2, 0xa0,
	0x49, 0x36, 0xe7, 0x75, 0x82, 0xde, 0xe6, 0xfa, 0x56, 0x59, 0x5f, 0x01, 0xb0, 0xa2, 0xa0, 0x05,
	0x0b, 0x05, 0xcf, 0xca, 0x7a, 0x77, 0x88, 0x49, 0x51, 0xf0, 0xc2, 0x2b, 0x0f, 0x7d, 0xcb, 0xc6,
	0x67, 0xa0, 0xcb, 0xfb, 0x9c, 0xde, 0x85, 0xd1, 0x1b, 0x3e, 0x9f, 0x57, 0xed, 0xab, 0x1a, 0x45,
	0x15, 0x87, 0x26, 0xb1, 0xd9, 0x2c, 0xd7, 0xdc, 0x85, 0xe3, 0x3c, 0x8c, 0xe3, 0x24, 0x5b, 0x98,
	0xb1, 0x7a, 0x3a, 0x83, 0xff, 0x34, 0xe8, 0x56, 0xb2, 0x01, 0xaf, 0xc3, 0xbf, 0x27, 0xf4, 0x27,
	0x70, 0xbe, 0x59, 0x21, 0x7d, 0xb2, 0xa4, 0x77, 0x75, 0xd2, 0xd8, 0xb7, 0x6e, 0x5b, 0xfe, 0x76,
	0x27, 0x67, 0xa0, 0x2f, 0xf3, 0x94, 0x87, 0x31, 0x15, 0xab, 0x2c, 0x2a, 0x93, 0x1e, 0xe2, 0x73,
	0xe8, 0xcc, 0xc3, 0x24, 0x65, 0xf1, 0x7a, 0x73, 0xd0, 0x6f, 0x0c, 0x5b, 0xf8, 0x33, 0x38, 0x52,
	0xfc, 0xc2, 0xd4, 0xcb, 0xe6, 0xb6, 0x16, 0xe9, 0x33, 0xa1, 0x6c, 0xaa, 0x04, 0x02, 0x5f, 0x80,
	0x11, 0xf1, 0x07, 0x56, 0x48, 0xaa, 0x9e, 0x33, 0x13, 0xc2, 0xec, 0xbd, 0x3b, 0xfc, 0x9f, 0x1a,
	0xb4, 0x77, 0x26, 0x9f, 0xe8, 0xf6, 0xd6, 0xb1, 0x25, 0x97, 0x61, 0x4a, 0xd5, 0x73, 0xa7, 0x92,
	0xd3, 0x88, 0x67, 0x19, 0x8b, 0xa4, 0xf9, 0xbc, 0xaf, 0x0d, 0x1b, 0x4a, 0xaf, 0x90, 0x52, 0xe1,
	0x75, 0xdf, 0xe6, 0xa7, 0x25, 0xde, 0x83, 0xb6, 0x7a, 0x7c, 0x92, 0x57, 0x6c, 0xe6, 0xe7, 0x8f,
	0x68, 0x94, 0x6f, 0xd0, 0x17, 0x0a, 0xbd, 0xfc, 0x12, 0x8e, 0xeb, 0x9f, 0x15, 0x77, 0x41, 0x1f,
	0x11, 0x9f, 0xbe, 0xb2, 0xaf, 0xe9, 0xd7, 0xd6, 0xf7, 0xe8, 0xd7, 0x6d, 0xc0, 0xfa, 0xf6, 0x25,
	0xfa, 0xed, 0xf2, 0x1f, 0x0d, 0x8c, 0xbd, 0x1a, 0x4f, 0xa1, 0xa3, 0x10, 0xc7, 0xa5, 0xf6, 0x8f,
	0x23, 0xe7, 0x15, 0x41, 0xcf, 0x70, 0x0f, 0x90, 0x82, 0x7c, 0xe2, 0xfb, 0x13, 0xd7, 0xa1, 0x13,
	0x67, 0x12, 0x20, 0x0d, 0x7f, 0x04, 0x1f, 0x6c, 0xa3, 0xb6, 0xfb, 0x33, 0xf1, 0x82, 0xea, 0x50,
	0xc7, 0x26, 0xf4, 0xd4, 0x21, 0xf9, 0x65, 0x46, 0xec, 0x80, 0x7a, 0xc4, 0x76, 0x1d, 0x87, 0xd8,
	0x01, 0x3a, 0xc0, 0xe7, 0x70, 0xba, 0x33, 0x36, 0x75, 0x7d, 0x82, 0x1a, 0x6b, 0x8d, 0xdb, 0x09,
	0x99, 0x8e, 0xe9, 0xcd, 0x6c, 0xea, 0x8e, 0xc6, 0xe8, 0x10, 0x5f, 0x00, 0x56, 0xe8, 0xc8, 0xfe,
	0xe9, 0x66, 0xe2, 0x91, 0x35, 0x7e, 0x84, 0xfb, 0xf0, 0xf1, 0x16, 0x7d, 0x05, 0xbb, 0xce, 0xf4,
	0xb6, 0x56, 0x42, 0x4d, 0x6c, 0x40, 0xab, 0xbc, 0xe1, 0x79, 0xae, 0x87, 0xfe, 0xd7, 0x2e, 0xff,
	0xd2, 0xc0, 0xd8, 0xfb, 0x89, 0x4e, 0xa1, 0xa3, 0x90, 0xbd, 0xa4, 0x0a, 0x7a, 0x9a, 0x74, 0x1b,
	0xdd, 0x4d, 0xfa, 0x21, 0x9c, 0xab, 0x43, 0xdb, 0x75, 0x7e, 0x98, 0x78, 0xd7, 0xfb, 0x51, 0x77,
	0xe6, 0xea, 0xa8, 0x06, 0xb4, 0x14, 0xfc, 0x68, 0xed, 0x6f, 0x0d, 0x8c, 0xbd, 0xdf, 0xb0, 0x0d,
	0x27, 0x8e, 0x5b, 0xdf, 0x78, 0x56, 0xae, 0xa4, 0xd2, 0xf4, 0x03, 0x8f, 0x8c, 0xae, 0x91, 0x86,
	0xcf, 0xa0, 0x6b, 0x4f, 0x27, 0xc4, 0x51, 0xdd, 0xce, 0x5c, 0x2f, 0x20, 0x63, 0x74, 0xb0, 0x05,
	0xce, 0x3c, 0x37, 0x70, 0x6d, 0x77, 0x5a, 0x15, 0xeb, 0x07, 0xa3, 0xa0, 0x8a, 0x13, 0x10, 0xcf,
	0x19, 0x4d, 0xd1, 0x21, 0xc6, 0x60, 0x8c, 0x89, 0xed, 0xde, 0x52, 0xc5, 0x5b, 0x97, 0xaa, 0x64,
	0xaa, 0xf1, 0x5a, 0x26, 0x56, 0xd7, 0x6a, 0x28, 0x98, 0x5c, 0x13, 0xf7, 0x26, 0x40, 0xec, 0x6d,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x32, 0x40, 0x07, 0x57, 0x06, 0x00, 0x00,
}
