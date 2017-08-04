// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apiserver/apiserver.proto

/*
Package apiserver is a generated protocol buffer package.

It is generated from these files:
	pkg/apiserver/apiserver.proto

It has these top-level messages:
	WorkflowIdentifier
	SearchWorkflowRequest
	SearchWorkflowResponse
	WorkflowInvocationIdentifier
	WorkflowInvocationQuery
	WorkflowInvocationList
	WorkflowInvocationGetReponse
	Health
*/
package apiserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/fission/fission-workflow/pkg/types"
import eventstore "github.com/fission/fission-workflow/pkg/eventstore"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowIdentifier struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WorkflowIdentifier) Reset()                    { *m = WorkflowIdentifier{} }
func (m *WorkflowIdentifier) String() string            { return proto.CompactTextString(m) }
func (*WorkflowIdentifier) ProtoMessage()               {}
func (*WorkflowIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WorkflowIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchWorkflowRequest struct {
}

func (m *SearchWorkflowRequest) Reset()                    { *m = SearchWorkflowRequest{} }
func (m *SearchWorkflowRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchWorkflowRequest) ProtoMessage()               {}
func (*SearchWorkflowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SearchWorkflowResponse struct {
	Workflows []string `protobuf:"bytes,1,rep,name=workflows" json:"workflows,omitempty"`
}

func (m *SearchWorkflowResponse) Reset()                    { *m = SearchWorkflowResponse{} }
func (m *SearchWorkflowResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchWorkflowResponse) ProtoMessage()               {}
func (*SearchWorkflowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SearchWorkflowResponse) GetWorkflows() []string {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type WorkflowInvocationIdentifier struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WorkflowInvocationIdentifier) Reset()                    { *m = WorkflowInvocationIdentifier{} }
func (m *WorkflowInvocationIdentifier) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationIdentifier) ProtoMessage()               {}
func (*WorkflowInvocationIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WorkflowInvocationIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type WorkflowInvocationQuery struct {
	Subject string `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
}

func (m *WorkflowInvocationQuery) Reset()                    { *m = WorkflowInvocationQuery{} }
func (m *WorkflowInvocationQuery) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationQuery) ProtoMessage()               {}
func (*WorkflowInvocationQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WorkflowInvocationQuery) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type WorkflowInvocationList struct {
	Invocations []string `protobuf:"bytes,1,rep,name=invocations" json:"invocations,omitempty"`
}

func (m *WorkflowInvocationList) Reset()                    { *m = WorkflowInvocationList{} }
func (m *WorkflowInvocationList) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationList) ProtoMessage()               {}
func (*WorkflowInvocationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WorkflowInvocationList) GetInvocations() []string {
	if m != nil {
		return m.Invocations
	}
	return nil
}

type WorkflowInvocationGetReponse struct {
	Events []*eventstore.Event `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *WorkflowInvocationGetReponse) Reset()                    { *m = WorkflowInvocationGetReponse{} }
func (m *WorkflowInvocationGetReponse) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationGetReponse) ProtoMessage()               {}
func (*WorkflowInvocationGetReponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WorkflowInvocationGetReponse) GetEvents() []*eventstore.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Health struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Health) Reset()                    { *m = Health{} }
func (m *Health) String() string            { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()               {}
func (*Health) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowIdentifier)(nil), "apiserver.WorkflowIdentifier")
	proto.RegisterType((*SearchWorkflowRequest)(nil), "apiserver.SearchWorkflowRequest")
	proto.RegisterType((*SearchWorkflowResponse)(nil), "apiserver.SearchWorkflowResponse")
	proto.RegisterType((*WorkflowInvocationIdentifier)(nil), "apiserver.WorkflowInvocationIdentifier")
	proto.RegisterType((*WorkflowInvocationQuery)(nil), "apiserver.WorkflowInvocationQuery")
	proto.RegisterType((*WorkflowInvocationList)(nil), "apiserver.WorkflowInvocationList")
	proto.RegisterType((*WorkflowInvocationGetReponse)(nil), "apiserver.WorkflowInvocationGetReponse")
	proto.RegisterType((*Health)(nil), "apiserver.Health")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkflowAPI service

type WorkflowAPIClient interface {
	Create(ctx context.Context, in *types.WorkflowSpec, opts ...grpc.CallOption) (*WorkflowIdentifier, error)
	List(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*SearchWorkflowResponse, error)
	Get(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*types.Workflow, error)
	Validate(ctx context.Context, in *types.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type workflowAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowAPIClient(cc *grpc.ClientConn) WorkflowAPIClient {
	return &workflowAPIClient{cc}
}

func (c *workflowAPIClient) Create(ctx context.Context, in *types.WorkflowSpec, opts ...grpc.CallOption) (*WorkflowIdentifier, error) {
	out := new(WorkflowIdentifier)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowAPI/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) List(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*SearchWorkflowResponse, error) {
	out := new(SearchWorkflowResponse)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Get(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*types.Workflow, error) {
	out := new(types.Workflow)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Validate(ctx context.Context, in *types.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowAPI/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowAPI service

type WorkflowAPIServer interface {
	Create(context.Context, *types.WorkflowSpec) (*WorkflowIdentifier, error)
	List(context.Context, *google_protobuf2.Empty) (*SearchWorkflowResponse, error)
	Get(context.Context, *WorkflowIdentifier) (*types.Workflow, error)
	Validate(context.Context, *types.WorkflowSpec) (*google_protobuf2.Empty, error)
}

func RegisterWorkflowAPIServer(s *grpc.Server, srv WorkflowAPIServer) {
	s.RegisterService(&_WorkflowAPI_serviceDesc, srv)
}

func _WorkflowAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Create(ctx, req.(*types.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).List(ctx, req.(*google_protobuf2.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Get(ctx, req.(*WorkflowIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Validate(ctx, req.(*types.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver.WorkflowAPI",
	HandlerType: (*WorkflowAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WorkflowAPI_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowAPI_Get_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WorkflowAPI_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for WorkflowInvocationAPI service

type WorkflowInvocationAPIClient interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(ctx context.Context, in *types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*WorkflowInvocationIdentifier, error)
	InvokeSync(ctx context.Context, in *types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*types.WorkflowInvocationContainer, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	List(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*types.WorkflowInvocationContainer, error)
}

type workflowInvocationAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowInvocationAPIClient(cc *grpc.ClientConn) WorkflowInvocationAPIClient {
	return &workflowInvocationAPIClient{cc}
}

func (c *workflowInvocationAPIClient) Invoke(ctx context.Context, in *types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*WorkflowInvocationIdentifier, error) {
	out := new(WorkflowInvocationIdentifier)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowInvocationAPI/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) InvokeSync(ctx context.Context, in *types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*types.WorkflowInvocationContainer, error) {
	out := new(types.WorkflowInvocationContainer)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowInvocationAPI/InvokeSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Cancel(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowInvocationAPI/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) List(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*WorkflowInvocationList, error) {
	out := new(WorkflowInvocationList)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowInvocationAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Get(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*types.WorkflowInvocationContainer, error) {
	out := new(types.WorkflowInvocationContainer)
	err := grpc.Invoke(ctx, "/apiserver.WorkflowInvocationAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowInvocationAPI service

type WorkflowInvocationAPIServer interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(context.Context, *types.WorkflowInvocationSpec) (*WorkflowInvocationIdentifier, error)
	InvokeSync(context.Context, *types.WorkflowInvocationSpec) (*types.WorkflowInvocationContainer, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(context.Context, *WorkflowInvocationIdentifier) (*google_protobuf2.Empty, error)
	List(context.Context, *google_protobuf2.Empty) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(context.Context, *WorkflowInvocationIdentifier) (*types.WorkflowInvocationContainer, error)
}

func RegisterWorkflowInvocationAPIServer(s *grpc.Server, srv WorkflowInvocationAPIServer) {
	s.RegisterService(&_WorkflowInvocationAPI_serviceDesc, srv)
}

func _WorkflowInvocationAPI_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowInvocationAPI/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, req.(*types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_InvokeSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowInvocationAPI/InvokeSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, req.(*types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowInvocationIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowInvocationAPI/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, req.(*WorkflowInvocationIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowInvocationAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).List(ctx, req.(*google_protobuf2.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowInvocationIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.WorkflowInvocationAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, req.(*WorkflowInvocationIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowInvocationAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver.WorkflowInvocationAPI",
	HandlerType: (*WorkflowInvocationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _WorkflowInvocationAPI_Invoke_Handler,
		},
		{
			MethodName: "InvokeSync",
			Handler:    _WorkflowInvocationAPI_InvokeSync_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _WorkflowInvocationAPI_Cancel_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowInvocationAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowInvocationAPI_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for AdminAPI service

type AdminAPIClient interface {
	Status(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*Health, error)
}

type adminAPIClient struct {
	cc *grpc.ClientConn
}

func NewAdminAPIClient(cc *grpc.ClientConn) AdminAPIClient {
	return &adminAPIClient{cc}
}

func (c *adminAPIClient) Status(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := grpc.Invoke(ctx, "/apiserver.AdminAPI/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdminAPI service

type AdminAPIServer interface {
	Status(context.Context, *google_protobuf2.Empty) (*Health, error)
}

func RegisterAdminAPIServer(s *grpc.Server, srv AdminAPIServer) {
	s.RegisterService(&_AdminAPI_serviceDesc, srv)
}

func _AdminAPI_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAPIServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.AdminAPI/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAPIServer).Status(ctx, req.(*google_protobuf2.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver.AdminAPI",
	HandlerType: (*AdminAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AdminAPI_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for FunctionEnvApi service

type FunctionEnvApiClient interface {
	// Debug purposes
	Invoke(ctx context.Context, in *types.FunctionInvocationSpec, opts ...grpc.CallOption) (*types.FunctionInvocation, error)
}

type functionEnvApiClient struct {
	cc *grpc.ClientConn
}

func NewFunctionEnvApiClient(cc *grpc.ClientConn) FunctionEnvApiClient {
	return &functionEnvApiClient{cc}
}

func (c *functionEnvApiClient) Invoke(ctx context.Context, in *types.FunctionInvocationSpec, opts ...grpc.CallOption) (*types.FunctionInvocation, error) {
	out := new(types.FunctionInvocation)
	err := grpc.Invoke(ctx, "/apiserver.FunctionEnvApi/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FunctionEnvApi service

type FunctionEnvApiServer interface {
	// Debug purposes
	Invoke(context.Context, *types.FunctionInvocationSpec) (*types.FunctionInvocation, error)
}

func RegisterFunctionEnvApiServer(s *grpc.Server, srv FunctionEnvApiServer) {
	s.RegisterService(&_FunctionEnvApi_serviceDesc, srv)
}

func _FunctionEnvApi_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FunctionInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionEnvApiServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.FunctionEnvApi/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionEnvApiServer).Invoke(ctx, req.(*types.FunctionInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _FunctionEnvApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver.FunctionEnvApi",
	HandlerType: (*FunctionEnvApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _FunctionEnvApi_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

func init() { proto.RegisterFile("pkg/apiserver/apiserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6a, 0xdb, 0x4c,
	0x10, 0x46, 0x4e, 0x50, 0xa2, 0x31, 0xf1, 0x9f, 0x6c, 0xfe, 0xd8, 0xc6, 0x75, 0x8a, 0xbb, 0x14,
	0x9a, 0x1a, 0x2a, 0x81, 0x03, 0x85, 0xfa, 0x50, 0x08, 0x6e, 0x9a, 0x06, 0x72, 0x48, 0xed, 0x92,
	0x42, 0x0f, 0x05, 0x59, 0x1e, 0x3b, 0xdb, 0x38, 0xbb, 0xaa, 0x76, 0xe5, 0x60, 0x4a, 0x2f, 0x7d,
	0x85, 0x42, 0x9f, 0xaa, 0xb7, 0xbe, 0x42, 0x1f, 0xa4, 0x68, 0x25, 0xd9, 0x6a, 0x24, 0x87, 0xf4,
	0x62, 0x6b, 0x67, 0x76, 0xbf, 0xef, 0x9b, 0x99, 0x6f, 0x60, 0xdf, 0xbf, 0x9a, 0x38, 0xae, 0xcf,
	0x24, 0x06, 0x33, 0x0c, 0x96, 0x5f, 0xb6, 0x1f, 0x08, 0x25, 0x88, 0xb5, 0x08, 0x34, 0x5e, 0x4c,
	0x98, 0xba, 0x0c, 0x87, 0xb6, 0x27, 0xae, 0x9d, 0x31, 0x93, 0x92, 0x09, 0x9e, 0xfe, 0x3f, 0xbb,
	0x11, 0xc1, 0xd5, 0x78, 0x2a, 0x6e, 0x9c, 0x08, 0x4d, 0xcd, 0x7d, 0x94, 0xf1, 0x6f, 0x8c, 0xd2,
	0xe8, 0xdd, 0xf7, 0x29, 0xce, 0x90, 0x2b, 0xa9, 0x44, 0x80, 0x99, 0xcf, 0x04, 0xe4, 0xc1, 0x44,
	0x88, 0xc9, 0x14, 0x1d, 0x7d, 0x1a, 0x86, 0x63, 0x07, 0xaf, 0x7d, 0x35, 0x4f, 0x92, 0xcd, 0x24,
	0xe9, 0xfa, 0xcc, 0x71, 0x39, 0x17, 0xca, 0x55, 0x4c, 0xf0, 0x84, 0x9f, 0x3e, 0x06, 0xf2, 0x3e,
	0xa1, 0x39, 0x1d, 0x21, 0x57, 0x6c, 0xcc, 0x30, 0x20, 0x15, 0x28, 0xb1, 0x51, 0xdd, 0x68, 0x19,
	0x07, 0x56, 0xbf, 0xc4, 0x46, 0xb4, 0x06, 0x7b, 0x03, 0x74, 0x03, 0xef, 0x32, 0xbd, 0xdb, 0xc7,
	0xcf, 0x21, 0x4a, 0x45, 0x9f, 0x43, 0xf5, 0x76, 0x42, 0xfa, 0x82, 0x4b, 0x24, 0x4d, 0xb0, 0x52,
	0xfd, 0xb2, 0x6e, 0xb4, 0xd6, 0x0e, 0xac, 0xfe, 0x32, 0x40, 0x6d, 0x68, 0x2e, 0x68, 0xf9, 0x4c,
	0x78, 0x5a, 0xd3, 0x1d, 0x02, 0x0e, 0xa1, 0x96, 0xbf, 0xff, 0x36, 0xc4, 0x60, 0x4e, 0xea, 0xb0,
	0x21, 0xc3, 0xe1, 0x27, 0xf4, 0x54, 0x72, 0x3f, 0x3d, 0xd2, 0x2e, 0x54, 0xf3, 0x8f, 0xce, 0x98,
	0x54, 0xa4, 0x05, 0x65, 0xb6, 0x88, 0xa4, 0xf2, 0xb2, 0x21, 0xfa, 0xb2, 0x48, 0xe0, 0x09, 0xaa,
	0x3e, 0xc6, 0xe5, 0x3d, 0x04, 0x33, 0x1e, 0x83, 0x7e, 0x5c, 0xee, 0x98, 0xf6, 0x71, 0x74, 0xec,
	0x27, 0x51, 0xda, 0x02, 0xf3, 0x0d, 0xba, 0x53, 0x75, 0x49, 0xaa, 0x60, 0x4a, 0xe5, 0xaa, 0x50,
	0x26, 0xf2, 0x92, 0x53, 0xe7, 0x67, 0x09, 0xca, 0x29, 0xc5, 0xd1, 0xf9, 0x29, 0x39, 0x03, 0xb3,
	0x17, 0xa0, 0xab, 0x90, 0x6c, 0xd9, 0x69, 0x7c, 0xe0, 0xa3, 0xd7, 0xd8, 0xb7, 0x97, 0xd6, 0xcb,
	0xcf, 0x8a, 0xfe, 0xff, 0xed, 0xd7, 0xef, 0xef, 0xa5, 0x0a, 0xb5, 0x9c, 0xb4, 0xbd, 0x5d, 0xa3,
	0x4d, 0xde, 0xc1, 0xba, 0xae, 0xb4, 0x6a, 0xc7, 0xe3, 0xb7, 0x53, 0x6f, 0xd8, 0xc7, 0x91, 0x37,
	0x1a, 0x8f, 0x32, 0xa0, 0xc5, 0x13, 0xa4, 0x3b, 0x1a, 0xb8, 0x4c, 0x96, 0xc0, 0xe4, 0x04, 0xd6,
	0x4e, 0x50, 0x91, 0xbb, 0x15, 0x35, 0xac, 0x45, 0x90, 0x56, 0x35, 0xc6, 0x36, 0xa9, 0x2c, 0x30,
	0x9c, 0x2f, 0x6c, 0xf4, 0x95, 0x9c, 0xc3, 0xe6, 0x85, 0x3b, 0x65, 0xa3, 0x82, 0x72, 0x57, 0x28,
	0xa6, 0xfb, 0x1a, 0xaa, 0x46, 0xc9, 0x12, 0x6a, 0x96, 0x40, 0x74, 0x8d, 0x76, 0xe7, 0xc7, 0x3a,
	0xec, 0xe5, 0x27, 0x16, 0x35, 0x76, 0x08, 0x66, 0x14, 0xb8, 0x42, 0x52, 0xb3, 0xf3, 0x37, 0x34,
	0xe7, 0x93, 0xa2, 0x82, 0x0a, 0x7c, 0x99, 0xd6, 0x43, 0xcb, 0xce, 0xd2, 0x2c, 0x51, 0xbb, 0x15,
	0x40, 0xcc, 0x31, 0x98, 0x73, 0x6f, 0x35, 0x4f, 0xb3, 0x20, 0xd1, 0x13, 0x5c, 0xb9, 0x8c, 0x63,
	0x40, 0x1d, 0x0d, 0xfe, 0x94, 0x6e, 0x67, 0xc0, 0x1d, 0x39, 0xe7, 0x5e, 0xd7, 0x68, 0x7f, 0x20,
	0x24, 0x17, 0x26, 0x1e, 0x98, 0x3d, 0x97, 0x7b, 0x38, 0x25, 0xf7, 0x2d, 0x60, 0x65, 0x77, 0xeb,
	0x9a, 0x9b, 0xb4, 0xff, 0x22, 0xd1, 0xa3, 0xba, 0xf8, 0x07, 0x27, 0x15, 0xaf, 0x1b, 0xdd, 0xd5,
	0xe0, 0x5b, 0x24, 0xdb, 0x35, 0x32, 0x8e, 0xbd, 0x74, 0x6f, 0xe5, 0x77, 0xf7, 0x2e, 0xd1, 0x4f,
	0x72, 0xfa, 0x3b, 0xe7, 0xb0, 0x79, 0x34, 0xba, 0x66, 0xda, 0x0a, 0xaf, 0xc0, 0x1c, 0xe8, 0xed,
	0x5b, 0x59, 0xcd, 0x4e, 0x46, 0x4e, 0xbc, 0xc0, 0xf4, 0x3f, 0x0d, 0x6d, 0x91, 0x0d, 0x27, 0xd9,
	0xdc, 0x8f, 0x50, 0x79, 0x1d, 0x72, 0x2f, 0xa2, 0x38, 0xe6, 0xb3, 0x23, 0x9f, 0x45, 0xbb, 0xbb,
	0xb0, 0x58, 0x9a, 0xba, 0x35, 0xfa, 0xdd, 0x82, 0x44, 0x66, 0x77, 0xc7, 0x49, 0xb2, 0x6b, 0xb4,
	0x87, 0xa6, 0xd6, 0x74, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x52, 0xed, 0x2a, 0x81, 0x06,
	0x00, 0x00,
}
