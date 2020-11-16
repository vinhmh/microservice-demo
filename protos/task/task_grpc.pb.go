// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package task

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TaskSvcClient is the client API for TaskSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskSvcClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
}

type taskSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskSvcClient(cc grpc.ClientConnInterface) TaskSvcClient {
	return &taskSvcClient{cc}
}

func (c *taskSvcClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/demo_task.TaskSvc/createTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskSvcClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/demo_task.TaskSvc/updateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskSvcServer is the server API for TaskSvc service.
// All implementations must embed UnimplementedTaskSvcServer
// for forward compatibility
type TaskSvcServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*TaskResponse, error)
	mustEmbedUnimplementedTaskSvcServer()
}

// UnimplementedTaskSvcServer must be embedded to have forward compatible implementations.
type UnimplementedTaskSvcServer struct {
}

func (UnimplementedTaskSvcServer) CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskSvcServer) UpdateTask(context.Context, *UpdateTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskSvcServer) mustEmbedUnimplementedTaskSvcServer() {}

// UnsafeTaskSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskSvcServer will
// result in compilation errors.
type UnsafeTaskSvcServer interface {
	mustEmbedUnimplementedTaskSvcServer()
}

func RegisterTaskSvcServer(s grpc.ServiceRegistrar, srv TaskSvcServer) {
	s.RegisterService(&_TaskSvc_serviceDesc, srv)
}

func _TaskSvc_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskSvcServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_task.TaskSvc/createTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskSvcServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskSvc_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskSvcServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_task.TaskSvc/updateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskSvcServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo_task.TaskSvc",
	HandlerType: (*TaskSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createTask",
			Handler:    _TaskSvc_CreateTask_Handler,
		},
		{
			MethodName: "updateTask",
			Handler:    _TaskSvc_UpdateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/task/task.proto",
}
