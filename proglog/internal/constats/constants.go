package constats

const (
	EventRPCAddress = "rpc_addr"
)
const RaftRPC = 1

type RequestType uint8

const (
	AppendRequestType RequestType = 0
)
