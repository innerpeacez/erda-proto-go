// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: project.proto

package pb

import (
	transport "github.com/erda-project/erda-infra/pkg/transport"
	reflect "reflect"
)

// RegisterProjectServiceImp project.proto
func RegisterProjectServiceImp(regester transport.Register, srv ProjectServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterProjectServiceHandler(regester, ProjectServiceHandler(srv), _ops.HTTP...)
	RegisterProjectServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.msp.apm.project.ProjectService",
	)
}

var (
	projectServiceClientType  = reflect.TypeOf((*ProjectServiceClient)(nil)).Elem()
	projectServiceServerType  = reflect.TypeOf((*ProjectServiceServer)(nil)).Elem()
	projectServiceHandlerType = reflect.TypeOf((*ProjectServiceHandler)(nil)).Elem()
)

// ProjectServiceClientType .
func ProjectServiceClientType() reflect.Type { return projectServiceClientType }

// ProjectServiceServerType .
func ProjectServiceServerType() reflect.Type { return projectServiceServerType }

// ProjectServiceHandlerType .
func ProjectServiceHandlerType() reflect.Type { return projectServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		projectServiceClientType,
		// server types
		projectServiceServerType,
		// handler types
		projectServiceHandlerType,
	}
}
