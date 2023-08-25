package grpc

import (
	"context"
	"flag"
	api "github.com/rezaAmiri123/microservice/proglog/api/v1"
	"github.com/rezaAmiri123/microservice/proglog/internal/distribution"
	"github.com/rezaAmiri123/microservice/proglog/internal/domain"
	"github.com/stretchr/testify/require"
	"go.opencensus.io/examples/exporter"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"testing"
	"time"
)

var debug = flag.Bool("debug", false, "Enable observability for debugging")

// go test -v -debug=true
// metrics log file: /tmp/metrics-{{random string}}.log
// traces log file: /tmp/traces-{{random string}}.log
func TestMain(m *testing.M) {
	flag.Parse()
	if *debug {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		zap.ReplaceGlobals(logger)
	}
	os.Exit(m.Run())
}

func TestServer(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T, client api.LogClient, config *Config){
		"produce/consume a message to/from the log succeeds": testProduceConsume,
		"consume past log boundary fails":                    testConsumePastBoundary,
	} {
		t.Run(scenario, func(t *testing.T) {
			client, cfg, teardown := setupTest(t, nil)
			defer teardown()
			fn(t, client, cfg)
		})
	}
}

func setupTest(t *testing.T, fn func(*Config)) (client api.LogClient, cfg *Config, teardown func()) {
	t.Helper()

	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	conn, err := grpc.Dial(l.Addr().String(), grpc.WithInsecure())
	require.NoError(t, err)

	client = api.NewLogClient(conn)

	dir, err := os.MkdirTemp("", "server-test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	logc, err := domain.NewLog(dir, domain.Config{})
	require.NoError(t, err)

	var telemetryExpoter *exporter.LogExporter
	if *debug {
		metricsLogFile, err := os.CreateTemp("", "metrics-*.log")
		require.NoError(t, err)
		t.Logf("metrics log file: %s", metricsLogFile.Name())

		tracesLogFile, err := os.CreateTemp("", "tracer-*.log")
		require.NoError(t, err)
		t.Logf("traces log file: %s", tracesLogFile.Name())

		telemetryExpoter, err = exporter.NewLogExporter(exporter.Options{
			MetricsLogFile:    metricsLogFile.Name(),
			TracesLogFile:     tracesLogFile.Name(),
			ReportingInterval: time.Second,
		})
		require.NoError(t, err)

		err = telemetryExpoter.Start()
		require.NoError(t, err)
	}
	cfg = &Config{
		Log:     logc,
		Servers: distribution.NewMockGetServers(t),
	}

	if fn != nil {
		fn(cfg)
	}
	var opts []grpc.ServerOption
	server, err := NewGRPCServer(cfg, opts...)
	require.NoError(t, err)

	go func() {
		server.Serve(l)
	}()

	return client, cfg, func() {
		server.Stop()
		conn.Close()
		l.Close()
		if telemetryExpoter != nil {
			time.Sleep(1500 * time.Millisecond)
			telemetryExpoter.Stop()
			telemetryExpoter.Close()
		}
	}
}

func testProduceConsume(t *testing.T, client api.LogClient, config *Config) {
	ctx := context.Background()
	want := &domain.Record{Value: []byte("hello world")}
	produce, err := client.Produce(ctx, &api.ProduceRequest{Record: recordToAPI(want)})
	require.NoError(t, err)

	consume, err := client.Consume(ctx, &api.ConsumeRequest{Offset: produce.GetOffset()})
	require.NoError(t, err)
	require.Equal(t, want.GetValue(), consume.GetRecord().GetValue())
	require.Equal(t, want.GetOffset(), consume.GetRecord().GetOffset())
}

func testConsumePastBoundary(t *testing.T, client api.LogClient, config *Config) {
	ctx := context.Background()
	want := &domain.Record{Value: []byte("hello world")}
	produce, err := client.Produce(ctx, &api.ProduceRequest{Record: recordToAPI(want)})
	require.NoError(t, err)

	consume, err := client.Consume(ctx, &api.ConsumeRequest{Offset: produce.GetOffset() + 1})
	require.Error(t, err)
	require.Nil(t, consume)

	gotErr := status.Code(err)
	wantErr := status.Code(domain.ErrOffsetOutOfRange)
	require.Equal(t, gotErr, wantErr)
}
