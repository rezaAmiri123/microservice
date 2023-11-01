package distribution

type (
	GetServers interface {
		GetServers() ([]*Server, error)
	}
	Server struct {
		Id       string
		RpcAddr  string
		IsLeader bool
	}
)
