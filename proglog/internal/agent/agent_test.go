package agent_test

import (
	"context"
	"fmt"
	api "github.com/rezaAmiri123/microservice/proglog/api/v1"
	"github.com/rezaAmiri123/microservice/proglog/internal/agent"
	"github.com/rezaAmiri123/microservice/proglog/internal/constants"
	"github.com/stretchr/testify/require"
	"github.com/travisjeffery/go-dynaport"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//func TestAgent(t *testing.T) {
//	var agents []*agent.Agent
//
//	for i := 0; i < 3; i++ {
//		ports := dynaport.Get(2)
//		bindAddr := fmt.Sprintf("%s:%d", "127.0.0.1", ports[0])
//		rpcPort := ports[1]
//
//		dataDir, err := os.MkdirTemp("", "agent-test-log")
//		require.NoError(t, err)
//
//		var startJonAddr []string
//		if i != 0 {
//			startJonAddr = append(startJonAddr, agents[0].Config.BindAddr)
//		}
//
//		agent, err := agent.New(agent.Config{
//			NodeName:       fmt.Sprintf("%d", i),
//			Bootstrap:      i == 0,
//			StartJoinAddrs: startJonAddr,
//			BindAddr:       bindAddr,
//			RPCPort:        rpcPort,
//			DataDir:        dataDir,
//		})
//		require.NoError(t, err)
//
//		agents = append(agents, agent)
//	}
//	defer func() {
//		for _, agent := range agents {
//			_ = agent.Shutdown()
//			err := os.RemoveAll(agent.Config.DataDir)
//			require.NoError(t, err)
//		}
//	}()
//
//	time.Sleep(3 * time.Second)
//
//	record := &api.Record{Value: []byte("foo")}
//	ctx := context.Background()
//	leaderClient := client(t, agents[0])
//	producerResponse, err := leaderClient.Produce(ctx, &api.ProduceRequest{Record: record})
//	require.NoError(t, err)
//
//	// wait until replication has finished
//	time.Sleep(3 * time.Second)
//
//	consumeResponse, err := leaderClient.Consume(ctx, &api.ConsumeRequest{Offset: producerResponse.GetOffset()})
//	require.NoError(t, err)
//	require.Equal(t, consumeResponse.GetRecord(), record)
//}
//
//func client(t *testing.T, agent *agent.Agent) api.LogClient {
//	opts := []grpc.DialOption{
//		//grpc.WithTransportCredentials(insecure.NewCredentials()),
//		grpc.WithInsecure(),
//	}
//	rpcAddr, err := agent.Config.RPCAddr()
//	require.NoError(t, err)
//
//	addr := fmt.Sprintf("%s:///%s", constants.ServiceName, rpcAddr)
//	conn, err := grpc.Dial(addr, opts...)
//	require.NoError(t, err)
//
//	client := api.NewLogClient(conn)
//
//	return client
//}

//##################################################

//func TestAgent(t *testing.T) {
//	var agents []*agent.Agent
//
//	for i := 0; i < 3; i++ {
//		ports := dynaport.Get(2)
//		bindAddr := fmt.Sprintf("%s:%d", "127.0.0.1", ports[0])
//		rpcPort := ports[1]
//
//		dataDir, err := ioutil.TempDir("", "agent-test-log")
//		require.NoError(t, err)
//
//		var startJoinAddrs []string
//		if i != 0 {
//			startJoinAddrs = append(startJoinAddrs, agents[0].Config.BindAddr)
//		}
//
//		agent, err := agent.New(agent.Config{
//			NodeName:       fmt.Sprintf("%d", i),
//			Bootstrap:      i == 0,
//			StartJoinAddrs: startJoinAddrs,
//			BindAddr:       bindAddr,
//			RPCPort:        rpcPort,
//			DataDir:        dataDir,
//		})
//		require.NoError(t, err)
//
//		agents = append(agents, agent)
//	}
//	defer func() {
//		for _, agent := range agents {
//			_ = agent.Shutdown()
//			require.NoError(
//				t,
//				os.RemoveAll(agent.Config.DataDir),
//			)
//		}
//	}()
//
//	// wait until agents have joined the cluster
//	time.Sleep(3 * time.Second)
//
//	leaderClient := client(t, agents[0])
//	produceResponse, err := leaderClient.Produce(
//		context.Background(),
//		&api.ProduceRequest{
//			Record: &api.Record{
//				Value: []byte("foo"),
//			},
//		},
//	)
//	require.NoError(t, err)
//
//	// START: test_change
//	// wait until replication has finished
//	time.Sleep(3 * time.Second)
//
//	//START: leader_check
//	consumeResponse, err := leaderClient.Consume( // <label id="produce" />
//		context.Background(),
//		&api.ConsumeRequest{
//			Offset: produceResponse.Offset,
//		},
//	)
//	require.NoError(t, err)
//	require.Equal(t, consumeResponse.Record.Value, []byte("foo"))
//
//	followerClient := client(t, agents[1])
//	consumeResponse, err = followerClient.Consume( // <label id="follower" />
//		context.Background(),
//		&api.ConsumeRequest{
//			Offset: produceResponse.Offset,
//		},
//	)
//	require.NoError(t, err)
//	require.Equal(t, consumeResponse.Record.Value, []byte("foo"))
//	// END: test_change
//}
//
//// START: client
//func client(
//	t *testing.T,
//	agent *agent.Agent,
//	// tlsConfig *tls.Config,
//) api.LogClient {
//	//tlsCreds := credentials.NewTLS(tlsConfig)
//	opts := []grpc.DialOption{
//		//grpc.WithTransportCredentials(tlsCreds),
//		grpc.WithInsecure(),
//	}
//	rpcAddr, err := agent.Config.RPCAddr()
//	require.NoError(t, err)
//	// START_HIGHLIGHT
//	conn, err := grpc.Dial(fmt.Sprintf(
//		"%s:///%s",
//		constants.ServiceName,
//		rpcAddr,
//	), opts...)
//	// END_HIGHLIGHT
//	require.NoError(t, err)
//	client := api.NewLogClient(conn)
//	return client
//}
//
//// END: client

//##################################################

func TestAgent(t *testing.T) {
	var agents []*agent.Agent

	//serverTLSConfig, err := config.SetupTLSConfig(config.TLSConfig{
	//	CertFile:      config.ServerCertFile,
	//	KeyFile:       config.ServerKeyFile,
	//	CAFile:        config.CAFile,
	//	Server:        true,
	//	ServerAddress: "127.0.0.1",
	//})
	//require.NoError(t, err)
	//
	//peerTLSConfig, err := config.SetupTLSConfig(config.TLSConfig{
	//	CertFile:      config.RootClientCertFile,
	//	KeyFile:       config.RootClientKeyFile,
	//	CAFile:        config.CAFile,
	//	Server:        false,
	//	ServerAddress: "127.0.0.1",
	//})
	//require.NoError(t, err)

	for i := 0; i < 3; i++ {
		ports := dynaport.Get(2)
		bindAddr := fmt.Sprintf("%s:%d", "127.0.0.1", ports[0])
		rpcPort := ports[1]

		dataDir, err := ioutil.TempDir("", "agent-test-log")
		require.NoError(t, err)

		var startJoinAddrs []string
		if i != 0 {
			startJoinAddrs = append(startJoinAddrs, agents[0].Config.BindAddr)
		}

		agent, err := agent.New(agent.Config{
			NodeName:       fmt.Sprintf("%d", i),
			Bootstrap:      i == 0,
			StartJoinAddrs: startJoinAddrs,
			BindAddr:       bindAddr,
			RPCPort:        rpcPort,
			DataDir:        dataDir,
			//ACLModelFile:    config.ACLModelFile,
			//ACLPolicyFile:   config.ACLPolicyFile,
			//ServerTLSConfig: serverTLSConfig,
			//PeerTLSConfig:   peerTLSConfig,
		})
		require.NoError(t, err)

		agents = append(agents, agent)
	}
	defer func() {
		for _, agent := range agents {
			_ = agent.Shutdown()
			require.NoError(
				t,
				os.RemoveAll(agent.Config.DataDir),
			)
		}
	}()

	// wait until agents have joined the cluster
	time.Sleep(3 * time.Second)

	leaderClient := client(t, agents[0])
	produceResponse, err := leaderClient.Produce(
		context.Background(),
		&api.ProduceRequest{
			Record: &api.Record{
				Value: []byte("foo"),
			},
		},
	)
	require.NoError(t, err)

	// START: test_change
	// wait until replication has finished
	time.Sleep(3 * time.Second)

	//START: leader_check
	consumeResponse, err := leaderClient.Consume( // <label id="produce" />
		context.Background(),
		&api.ConsumeRequest{
			Offset: produceResponse.Offset,
		},
	)
	require.NoError(t, err)
	require.Equal(t, consumeResponse.Record.Value, []byte("foo"))

	followerClient := client(t, agents[1])
	consumeResponse, err = followerClient.Consume( // <label id="follower" />
		context.Background(),
		&api.ConsumeRequest{
			Offset: produceResponse.Offset,
		},
	)
	require.NoError(t, err)
	require.Equal(t, consumeResponse.Record.Value, []byte("foo"))
	// END: test_change
}

// START: client
func client(
	t *testing.T,
	agent *agent.Agent,
	// tlsConfig *tls.Config,
) api.LogClient {
	//tlsCreds := credentials.NewTLS(tlsConfig)
	opts := []grpc.DialOption{
		//grpc.WithTransportCredentials(tlsCreds),
		grpc.WithInsecure(),
	}
	rpcAddr, err := agent.Config.RPCAddr()
	require.NoError(t, err)
	// START_HIGHLIGHT
	conn, err := grpc.Dial(fmt.Sprintf(
		"%s:///%s",
		constants.ServiceName,
		rpcAddr,
	), opts...)
	// END_HIGHLIGHT
	require.NoError(t, err)
	client := api.NewLogClient(conn)
	return client
}

// END: client
