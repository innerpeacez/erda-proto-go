// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: query.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/logs/query/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// LogQueryService query.proto
	LogQueryService() pb.LogQueryServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		logQueryService: pb.NewLogQueryServiceClient(cc),
	}
}

type serviceClients struct {
	logQueryService pb.LogQueryServiceClient
}

func (c *serviceClients) LogQueryService() pb.LogQueryServiceClient {
	return c.logQueryService
}

type logQueryServiceWrapper struct {
	client pb.LogQueryServiceClient
	opts   []grpc1.CallOption
}

func (s *logQueryServiceWrapper) GetLog(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogResponse, error) {
	return s.client.GetLog(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *logQueryServiceWrapper) GetLogByRuntime(ctx context.Context, req *pb.GetLogByRuntimeRequest) (*pb.GetLogByRuntimeResponse, error) {
	return s.client.GetLogByRuntime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *logQueryServiceWrapper) GetLogByOrganization(ctx context.Context, req *pb.GetLogByOrganizationRequest) (*pb.GetLogByOrganizationResponse, error) {
	return s.client.GetLogByOrganization(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
