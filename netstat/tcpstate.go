package main

type TcpState int

const (
	Closed      = TcpState(_MIB_TCP_STATE_CLOSED)
	Listen      = TcpState(_MIB_TCP_STATE_LISTEN)
	SynSent     = TcpState(_MIB_TCP_STATE_SYN_SENT)
	SynReceived = TcpState(_MIB_TCP_STATE_SYN_RCVD)
	Established = TcpState(_MIB_TCP_STATE_ESTAB)
	FinWait1    = TcpState(_MIB_TCP_STATE_FIN_WAIT1)
	FinWait2    = TcpState(_MIB_TCP_STATE_FIN_WAIT2)
	CloseWait   = TcpState(_MIB_TCP_STATE_CLOSE_WAIT)
	Closing     = TcpState(_MIB_TCP_STATE_CLOSING)
	LastAck     = TcpState(_MIB_TCP_STATE_LAST_ACK)
	TimeWait    = TcpState(_MIB_TCP_STATE_TIME_WAIT)
	DeleteTCB   = TcpState(_MIB_TCP_STATE_DELETE_TCB)
)

func (s TcpState) String() string {
	switch s {
	case Closed:
		return "CLOSED"
	case Listen:
		return "LISTEN"
	case SynSent:
		return "SYN_SENT"
	case SynReceived:
		return "SYN_RCVD"
	case Established:
		return "ESTABLISHED"
	case FinWait1:
		return "FIN_WAIT1"
	case FinWait2:
		return "FIN_WAIT2"
	case CloseWait:
		return "CLOSE_WAIT"
	case Closing:
		return "CLOSING"
	case LastAck:
		return "LAST_ACK"
	case TimeWait:
		return "TIME_WAIT"
	case DeleteTCB:
		return "DELETE_TCB"
	default:
		return "UNKNOWN"
	}
}
