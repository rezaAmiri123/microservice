package grpc

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	api "github.com/rezaAmiri123/microservice/proglog/api/v1"
	"github.com/rezaAmiri123/microservice/proglog/internal/distribution"
	"github.com/rezaAmiri123/microservice/proglog/internal/domain"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"time"
)

type (
	Config struct {
		Log     domain.Log
		Servers distribution.GetServers
	}
	grpcServer struct {
		*Config
		api.UnimplementedLogServer
	}
)

func NewGRPCServer(config *Config, grpcOpts ...grpc.ServerOption) (*grpc.Server, error) {
	logger := zap.L().Named("server")
	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		return nil, err
	}

	grpcOpts = append(grpcOpts,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger, zapOpts...),
			//grpc_auth.StreamServerInterceptor(authenticate),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger, zapOpts...),
			//grpc_auth.UnaryServerInterceptor(authenticate),
		)),
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)

	gsrv := grpc.NewServer(grpcOpts...)
	srv := &grpcServer{Config: config}
	api.RegisterLogServer(gsrv, srv)

	return gsrv, nil
}

func (s *grpcServer) Produce(ctx context.Context, req *api.ProduceRequest) (*api.ProduceResponse, error) {
	offset, err := s.Log.Append(recordToDomain(req.Record))
	if err != nil {
		return nil, err
	}
	return &api.ProduceResponse{Offset: offset}, nil
}

func (s *grpcServer) Consume(ctx context.Context, req *api.ConsumeRequest) (*api.ConsumeResponse, error) {
	record, err := s.Log.Read(req.GetOffset())
	if err != nil {
		return nil, err
	}
	return &api.ConsumeResponse{Record: recordToAPI(record)}, nil
}
func (s *grpcServer) GetServers(context.Context, *api.GetServersRequest) (*api.GetServersResponse, error) {
	gotServers, err := s.Servers.GetServers()
	if err != nil {
		return nil, err
	}
	var servers []*api.Server
	for _, server := range gotServers {
		servers = append(servers, &api.Server{
			Id:       server.Id,
			RpcAddr:  server.RpcAddr,
			IsLeader: server.IsLeader,
		})
	}
	return &api.GetServersResponse{Servers: servers}, nil

}
func recordToDomain(in *api.Record) *domain.Record {
	return &domain.Record{
		Offset: in.Offset,
		Value:  in.Value,
		Term:   in.Term,
		Type:   in.Type,
	}
}

func recordToAPI(in *domain.Record) *api.Record {
	return &api.Record{
		Offset: in.Offset,
		Value:  in.Value,
		Term:   in.Term,
		Type:   in.Type,
	}
}
