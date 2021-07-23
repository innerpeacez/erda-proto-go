// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: microservice_scope.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/scope/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// MicroserviceScopeService microservice_scope.proto
	MicroserviceScopeService() pb.MicroserviceScopeServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		microserviceScopeService: pb.NewMicroserviceScopeServiceClient(cc),
	}
}

type serviceClients struct {
	microserviceScopeService pb.MicroserviceScopeServiceClient
}

func (c *serviceClients) MicroserviceScopeService() pb.MicroserviceScopeServiceClient {
	return c.microserviceScopeService
}

type microserviceScopeServiceWrapper struct {
	client pb.MicroserviceScopeServiceClient
	opts   []grpc1.CallOption
}

func (s *microserviceScopeServiceWrapper) EncryptMicroserviceScope(ctx context.Context, req *pb.EncryptMicroserviceScopeRequest) (*pb.EncryptMicroserviceScopeResponse, error) {
	return s.client.EncryptMicroserviceScope(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *microserviceScopeServiceWrapper) DecryptMicroserviceScope(ctx context.Context, req *pb.DecryptMicroserviceScopeRequest) (*pb.DecryptMicroserviceScopeResponse, error) {
	return s.client.DecryptMicroserviceScope(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
