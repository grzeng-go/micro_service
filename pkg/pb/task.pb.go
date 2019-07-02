// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetListTaskReq struct {
	ProcessVariables     string   `protobuf:"bytes,1,opt,name=processVariables,proto3" json:"processVariables,omitempty"`
	ProcessDefinitionKey string   `protobuf:"bytes,2,opt,name=processDefinitionKey,proto3" json:"processDefinitionKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListTaskReq) Reset()         { *m = GetListTaskReq{} }
func (m *GetListTaskReq) String() string { return proto.CompactTextString(m) }
func (*GetListTaskReq) ProtoMessage()    {}
func (*GetListTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *GetListTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListTaskReq.Unmarshal(m, b)
}
func (m *GetListTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListTaskReq.Marshal(b, m, deterministic)
}
func (m *GetListTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListTaskReq.Merge(m, src)
}
func (m *GetListTaskReq) XXX_Size() int {
	return xxx_messageInfo_GetListTaskReq.Size(m)
}
func (m *GetListTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetListTaskReq proto.InternalMessageInfo

func (m *GetListTaskReq) GetProcessVariables() string {
	if m != nil {
		return m.ProcessVariables
	}
	return ""
}

func (m *GetListTaskReq) GetProcessDefinitionKey() string {
	if m != nil {
		return m.ProcessDefinitionKey
	}
	return ""
}

type GetListTaskResp struct {
	Tasks                []*TaskItem   `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Code                 int64         `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Err                  *CamundaError `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetListTaskResp) Reset()         { *m = GetListTaskResp{} }
func (m *GetListTaskResp) String() string { return proto.CompactTextString(m) }
func (*GetListTaskResp) ProtoMessage()    {}
func (*GetListTaskResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *GetListTaskResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListTaskResp.Unmarshal(m, b)
}
func (m *GetListTaskResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListTaskResp.Marshal(b, m, deterministic)
}
func (m *GetListTaskResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListTaskResp.Merge(m, src)
}
func (m *GetListTaskResp) XXX_Size() int {
	return xxx_messageInfo_GetListTaskResp.Size(m)
}
func (m *GetListTaskResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListTaskResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetListTaskResp proto.InternalMessageInfo

func (m *GetListTaskResp) GetTasks() []*TaskItem {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *GetListTaskResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetListTaskResp) GetErr() *CamundaError {
	if m != nil {
		return m.Err
	}
	return nil
}

type TaskItem struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Assignee             string   `protobuf:"bytes,3,opt,name=assignee,proto3" json:"assignee,omitempty"`
	Created              string   `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Due                  string   `protobuf:"bytes,5,opt,name=due,proto3" json:"due,omitempty"`
	FollowUp             string   `protobuf:"bytes,6,opt,name=followUp,proto3" json:"followUp,omitempty"`
	DelegationState      string   `protobuf:"bytes,7,opt,name=delegationState,proto3" json:"delegationState,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	ExecutionId          string   `protobuf:"bytes,9,opt,name=executionId,proto3" json:"executionId,omitempty"`
	Owner                string   `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
	ParentTaskId         string   `protobuf:"bytes,11,opt,name=parentTaskId,proto3" json:"parentTaskId,omitempty"`
	Priority             int64    `protobuf:"varint,12,opt,name=priority,proto3" json:"priority,omitempty"`
	ProcessDefinitionId  string   `protobuf:"bytes,13,opt,name=processDefinitionId,proto3" json:"processDefinitionId,omitempty"`
	ProcessInstanceId    string   `protobuf:"bytes,14,opt,name=processInstanceId,proto3" json:"processInstanceId,omitempty"`
	TaskDefinitionKey    string   `protobuf:"bytes,15,opt,name=taskDefinitionKey,proto3" json:"taskDefinitionKey,omitempty"`
	CaseExecutionId      string   `protobuf:"bytes,16,opt,name=caseExecutionId,proto3" json:"caseExecutionId,omitempty"`
	CaseInstanceId       string   `protobuf:"bytes,17,opt,name=caseInstanceId,proto3" json:"caseInstanceId,omitempty"`
	CaseDefinitionId     string   `protobuf:"bytes,18,opt,name=caseDefinitionId,proto3" json:"caseDefinitionId,omitempty"`
	Suspended            bool     `protobuf:"varint,19,opt,name=suspended,proto3" json:"suspended,omitempty"`
	FormKey              string   `protobuf:"bytes,20,opt,name=formKey,proto3" json:"formKey,omitempty"`
	TenantId             string   `protobuf:"bytes,21,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskItem) Reset()         { *m = TaskItem{} }
func (m *TaskItem) String() string { return proto.CompactTextString(m) }
func (*TaskItem) ProtoMessage()    {}
func (*TaskItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TaskItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskItem.Unmarshal(m, b)
}
func (m *TaskItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskItem.Marshal(b, m, deterministic)
}
func (m *TaskItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskItem.Merge(m, src)
}
func (m *TaskItem) XXX_Size() int {
	return xxx_messageInfo_TaskItem.Size(m)
}
func (m *TaskItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskItem.DiscardUnknown(m)
}

var xxx_messageInfo_TaskItem proto.InternalMessageInfo

func (m *TaskItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskItem) GetAssignee() string {
	if m != nil {
		return m.Assignee
	}
	return ""
}

func (m *TaskItem) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *TaskItem) GetDue() string {
	if m != nil {
		return m.Due
	}
	return ""
}

func (m *TaskItem) GetFollowUp() string {
	if m != nil {
		return m.FollowUp
	}
	return ""
}

func (m *TaskItem) GetDelegationState() string {
	if m != nil {
		return m.DelegationState
	}
	return ""
}

func (m *TaskItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskItem) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *TaskItem) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TaskItem) GetParentTaskId() string {
	if m != nil {
		return m.ParentTaskId
	}
	return ""
}

func (m *TaskItem) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TaskItem) GetProcessDefinitionId() string {
	if m != nil {
		return m.ProcessDefinitionId
	}
	return ""
}

func (m *TaskItem) GetProcessInstanceId() string {
	if m != nil {
		return m.ProcessInstanceId
	}
	return ""
}

func (m *TaskItem) GetTaskDefinitionKey() string {
	if m != nil {
		return m.TaskDefinitionKey
	}
	return ""
}

func (m *TaskItem) GetCaseExecutionId() string {
	if m != nil {
		return m.CaseExecutionId
	}
	return ""
}

func (m *TaskItem) GetCaseInstanceId() string {
	if m != nil {
		return m.CaseInstanceId
	}
	return ""
}

func (m *TaskItem) GetCaseDefinitionId() string {
	if m != nil {
		return m.CaseDefinitionId
	}
	return ""
}

func (m *TaskItem) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

func (m *TaskItem) GetFormKey() string {
	if m != nil {
		return m.FormKey
	}
	return ""
}

func (m *TaskItem) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

type GetTaskReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskReq) Reset()         { *m = GetTaskReq{} }
func (m *GetTaskReq) String() string { return proto.CompactTextString(m) }
func (*GetTaskReq) ProtoMessage()    {}
func (*GetTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{3}
}

func (m *GetTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskReq.Unmarshal(m, b)
}
func (m *GetTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskReq.Marshal(b, m, deterministic)
}
func (m *GetTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskReq.Merge(m, src)
}
func (m *GetTaskReq) XXX_Size() int {
	return xxx_messageInfo_GetTaskReq.Size(m)
}
func (m *GetTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskReq proto.InternalMessageInfo

func (m *GetTaskReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTaskResp struct {
	Task                 *TaskItem     `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Code                 int64         `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Err                  *CamundaError `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetTaskResp) Reset()         { *m = GetTaskResp{} }
func (m *GetTaskResp) String() string { return proto.CompactTextString(m) }
func (*GetTaskResp) ProtoMessage()    {}
func (*GetTaskResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{4}
}

func (m *GetTaskResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskResp.Unmarshal(m, b)
}
func (m *GetTaskResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskResp.Marshal(b, m, deterministic)
}
func (m *GetTaskResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskResp.Merge(m, src)
}
func (m *GetTaskResp) XXX_Size() int {
	return xxx_messageInfo_GetTaskResp.Size(m)
}
func (m *GetTaskResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskResp proto.InternalMessageInfo

func (m *GetTaskResp) GetTask() *TaskItem {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *GetTaskResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetTaskResp) GetErr() *CamundaError {
	if m != nil {
		return m.Err
	}
	return nil
}

type CompleteTaskReq struct {
	Id                   string               `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Body                 *CompleteTaskReqBody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CompleteTaskReq) Reset()         { *m = CompleteTaskReq{} }
func (m *CompleteTaskReq) String() string { return proto.CompactTextString(m) }
func (*CompleteTaskReq) ProtoMessage()    {}
func (*CompleteTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{5}
}

func (m *CompleteTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteTaskReq.Unmarshal(m, b)
}
func (m *CompleteTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteTaskReq.Marshal(b, m, deterministic)
}
func (m *CompleteTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteTaskReq.Merge(m, src)
}
func (m *CompleteTaskReq) XXX_Size() int {
	return xxx_messageInfo_CompleteTaskReq.Size(m)
}
func (m *CompleteTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteTaskReq proto.InternalMessageInfo

func (m *CompleteTaskReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CompleteTaskReq) GetBody() *CompleteTaskReqBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type CompleteTaskReqBody struct {
	Variables            map[string]*Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CompleteTaskReqBody) Reset()         { *m = CompleteTaskReqBody{} }
func (m *CompleteTaskReqBody) String() string { return proto.CompactTextString(m) }
func (*CompleteTaskReqBody) ProtoMessage()    {}
func (*CompleteTaskReqBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{6}
}

func (m *CompleteTaskReqBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteTaskReqBody.Unmarshal(m, b)
}
func (m *CompleteTaskReqBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteTaskReqBody.Marshal(b, m, deterministic)
}
func (m *CompleteTaskReqBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteTaskReqBody.Merge(m, src)
}
func (m *CompleteTaskReqBody) XXX_Size() int {
	return xxx_messageInfo_CompleteTaskReqBody.Size(m)
}
func (m *CompleteTaskReqBody) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteTaskReqBody.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteTaskReqBody proto.InternalMessageInfo

func (m *CompleteTaskReqBody) GetVariables() map[string]*Variable {
	if m != nil {
		return m.Variables
	}
	return nil
}

type CompleteTaskResp struct {
	Code                 int64         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err                  *CamundaError `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CompleteTaskResp) Reset()         { *m = CompleteTaskResp{} }
func (m *CompleteTaskResp) String() string { return proto.CompactTextString(m) }
func (*CompleteTaskResp) ProtoMessage()    {}
func (*CompleteTaskResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{7}
}

func (m *CompleteTaskResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteTaskResp.Unmarshal(m, b)
}
func (m *CompleteTaskResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteTaskResp.Marshal(b, m, deterministic)
}
func (m *CompleteTaskResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteTaskResp.Merge(m, src)
}
func (m *CompleteTaskResp) XXX_Size() int {
	return xxx_messageInfo_CompleteTaskResp.Size(m)
}
func (m *CompleteTaskResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteTaskResp.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteTaskResp proto.InternalMessageInfo

func (m *CompleteTaskResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CompleteTaskResp) GetErr() *CamundaError {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*GetListTaskReq)(nil), "pb.GetListTaskReq")
	proto.RegisterType((*GetListTaskResp)(nil), "pb.GetListTaskResp")
	proto.RegisterType((*TaskItem)(nil), "pb.TaskItem")
	proto.RegisterType((*GetTaskReq)(nil), "pb.GetTaskReq")
	proto.RegisterType((*GetTaskResp)(nil), "pb.GetTaskResp")
	proto.RegisterType((*CompleteTaskReq)(nil), "pb.CompleteTaskReq")
	proto.RegisterType((*CompleteTaskReqBody)(nil), "pb.CompleteTaskReqBody")
	proto.RegisterMapType((map[string]*Variable)(nil), "pb.CompleteTaskReqBody.VariablesEntry")
	proto.RegisterType((*CompleteTaskResp)(nil), "pb.CompleteTaskResp")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0xdb, 0x38,
	0x10, 0x5d, 0xf9, 0x23, 0xb1, 0xc7, 0x5e, 0xdb, 0xa1, 0xbd, 0x58, 0xc2, 0xc8, 0xc1, 0xd0, 0x21,
	0x30, 0x76, 0x17, 0xc6, 0xc2, 0xbb, 0xc0, 0x2e, 0x7a, 0x6c, 0x12, 0x04, 0x4a, 0x8b, 0x1e, 0xd4,
	0x8f, 0x3b, 0x2d, 0x4e, 0x02, 0x21, 0x96, 0xc8, 0x92, 0x74, 0x52, 0xff, 0x9a, 0x1e, 0xfb, 0x0f,
	0xfa, 0xfb, 0x8a, 0xa1, 0x64, 0xc7, 0xb2, 0xdd, 0x1e, 0x7a, 0xe3, 0xbc, 0xf7, 0x38, 0x1a, 0x3e,
	0x72, 0x46, 0x00, 0x4e, 0xd8, 0x87, 0x99, 0x36, 0xca, 0x29, 0x56, 0xd3, 0x8b, 0x71, 0x37, 0x51,
	0x59, 0xa6, 0xf2, 0x02, 0x09, 0x35, 0xf4, 0x6e, 0xd0, 0xbd, 0x4e, 0xad, 0x7b, 0x27, 0xec, 0x43,
	0x8c, 0x1f, 0xd9, 0x1f, 0x30, 0xd0, 0x46, 0x25, 0x68, 0xed, 0x07, 0x61, 0x52, 0xb1, 0x58, 0xa2,
	0xe5, 0xc1, 0x24, 0x98, 0xb6, 0xe3, 0x03, 0x9c, 0xcd, 0x61, 0x54, 0x62, 0x57, 0x78, 0x97, 0xe6,
	0xa9, 0x4b, 0x55, 0xfe, 0x0a, 0xd7, 0xbc, 0xe6, 0xf5, 0x47, 0xb9, 0x30, 0x83, 0x7e, 0xe5, 0x8b,
	0x56, 0xb3, 0x10, 0x9a, 0x54, 0x24, 0x7d, 0xa7, 0x3e, 0xed, 0xcc, 0xbb, 0x33, 0xbd, 0x98, 0x11,
	0x19, 0x39, 0xcc, 0xe2, 0x82, 0x62, 0x0c, 0x1a, 0x89, 0x92, 0xe8, 0x53, 0xd7, 0x63, 0xbf, 0x66,
	0x21, 0xd4, 0xd1, 0x18, 0x5e, 0x9f, 0x04, 0xd3, 0xce, 0x7c, 0x40, 0xbb, 0x2e, 0x45, 0xb6, 0xca,
	0xa5, 0xb8, 0x36, 0x46, 0x99, 0x98, 0xc8, 0xf0, 0x6b, 0x13, 0x5a, 0x9b, 0x5c, 0xac, 0x07, 0xb5,
	0x54, 0x96, 0xa7, 0xa9, 0xa5, 0x92, 0x92, 0xe6, 0x22, 0xc3, 0xb2, 0x5e, 0xbf, 0x66, 0x63, 0x68,
	0x09, 0x6b, 0xd3, 0xfb, 0x1c, 0xd1, 0x67, 0x6e, 0xc7, 0xdb, 0x98, 0x71, 0x38, 0x4d, 0x0c, 0x0a,
	0x87, 0x92, 0x37, 0x3c, 0xb5, 0x09, 0xd9, 0x00, 0xea, 0x72, 0x85, 0xbc, 0xe9, 0x51, 0x5a, 0x52,
	0x9e, 0x3b, 0xb5, 0x5c, 0xaa, 0xa7, 0xf7, 0x9a, 0x9f, 0x14, 0x79, 0x36, 0x31, 0x9b, 0x42, 0x5f,
	0xe2, 0x12, 0xef, 0x05, 0x99, 0xf2, 0xd6, 0x09, 0x87, 0xfc, 0xd4, 0x4b, 0xf6, 0x61, 0x36, 0x81,
	0x8e, 0x44, 0x9b, 0x98, 0x54, 0x13, 0xc6, 0x5b, 0x5e, 0xb5, 0x0b, 0x91, 0x02, 0x3f, 0x61, 0xb2,
	0xa2, 0x20, 0x92, 0xbc, 0x5d, 0x28, 0x76, 0x20, 0x36, 0x82, 0xa6, 0x7a, 0xca, 0xd1, 0x70, 0xf0,
	0x5c, 0x11, 0xb0, 0x10, 0xba, 0x5a, 0x18, 0xcc, 0xfd, 0x35, 0x44, 0x92, 0x77, 0x3c, 0x59, 0xc1,
	0xe8, 0x0c, 0xda, 0xa4, 0xca, 0xa4, 0x6e, 0xcd, 0xbb, 0xde, 0xf8, 0x6d, 0xcc, 0xfe, 0x86, 0xe1,
	0xc1, 0xfd, 0x46, 0x92, 0xff, 0xea, 0xd3, 0x1c, 0xa3, 0xd8, 0x5f, 0x70, 0x56, 0xc2, 0x51, 0x6e,
	0x9d, 0xc8, 0x13, 0x8c, 0x24, 0xef, 0x79, 0xfd, 0x21, 0x41, 0x6a, 0xba, 0xf9, 0xea, 0xc3, 0xea,
	0x17, 0xea, 0x03, 0x82, 0x1c, 0x4d, 0x84, 0xc5, 0xeb, 0x1d, 0x27, 0x06, 0x85, 0xa3, 0x7b, 0x30,
	0xbb, 0x80, 0x1e, 0x41, 0x3b, 0x25, 0x9c, 0x79, 0xe1, 0x1e, 0x4a, 0x7d, 0x40, 0x48, 0xe5, 0x70,
	0xac, 0xe8, 0x83, 0x7d, 0x9c, 0x9d, 0x43, 0xdb, 0xae, 0xac, 0xc6, 0x5c, 0xa2, 0xe4, 0xc3, 0x49,
	0x30, 0x6d, 0xc5, 0xcf, 0x00, 0xbd, 0x9a, 0x3b, 0x65, 0x32, 0xaa, 0x7f, 0x54, 0xbc, 0x9a, 0x32,
	0x24, 0x7f, 0x1d, 0xe6, 0x22, 0x77, 0x91, 0xe4, 0xbf, 0x15, 0x6f, 0x64, 0x13, 0x87, 0xe7, 0x00,
	0x37, 0xb8, 0xed, 0xca, 0xbd, 0x97, 0x1b, 0xde, 0x43, 0x67, 0xcb, 0x5a, 0xcd, 0x26, 0xd0, 0x20,
	0x4f, 0xbc, 0x60, 0xbf, 0x81, 0x3c, 0xf3, 0xd3, 0xfd, 0xf3, 0x06, 0xfa, 0x97, 0x2a, 0xd3, 0x4b,
	0x74, 0xb8, 0x53, 0x4b, 0xb4, 0xad, 0x25, 0x92, 0xec, 0x4f, 0x68, 0x2c, 0x94, 0x2c, 0xba, 0xbe,
	0x33, 0xff, 0xdd, 0xe7, 0xa9, 0x6e, 0x79, 0xa9, 0xe4, 0x3a, 0xf6, 0xa2, 0xf0, 0x4b, 0x00, 0xc3,
	0x23, 0x2c, 0xbb, 0x82, 0xf6, 0xe3, 0xce, 0xbc, 0xa1, 0x39, 0x70, 0xf1, 0x9d, 0x4c, 0xb3, 0xed,
	0x00, 0xba, 0xce, 0x9d, 0x59, 0xc7, 0xcf, 0x1b, 0xc7, 0xb7, 0xd0, 0xab, 0x92, 0xd4, 0x98, 0x0f,
	0xb8, 0x2e, 0xab, 0xa5, 0x25, 0x4d, 0x9b, 0x47, 0xb1, 0x5c, 0x61, 0x59, 0xaf, 0x37, 0x6b, 0xb3,
	0x29, 0x2e, 0xa8, 0x17, 0xb5, 0xff, 0x83, 0xf0, 0x16, 0x06, 0xd5, 0x8f, 0x5b, 0xbd, 0x75, 0x31,
	0x38, 0x74, 0xb1, 0xf6, 0x03, 0x17, 0xe7, 0x9f, 0x03, 0x68, 0x50, 0x12, 0xf6, 0x2f, 0x9c, 0x96,
	0xd3, 0x8f, 0x31, 0x92, 0x56, 0x87, 0xef, 0x78, 0x78, 0x80, 0x59, 0x1d, 0xfe, 0xc2, 0xfe, 0x83,
	0xd6, 0xa6, 0x14, 0x36, 0x3c, 0xe2, 0xca, 0x78, 0x74, 0x08, 0xfa, 0x8d, 0x53, 0xa8, 0xdf, 0xa0,
	0x63, 0xbd, 0x32, 0xed, 0x46, 0xde, 0xaf, 0xc4, 0xa4, 0x5c, 0x9c, 0xf8, 0xff, 0xc1, 0x3f, 0xdf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x7d, 0xaf, 0x70, 0x2f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	GetList(ctx context.Context, in *GetListTaskReq, opts ...grpc.CallOption) (*GetListTaskResp, error)
	Complete(ctx context.Context, in *CompleteTaskReq, opts ...grpc.CallOption) (*CompleteTaskResp, error)
	Get(ctx context.Context, in *GetTaskReq, opts ...grpc.CallOption) (*GetTaskResp, error)
}

type taskClient struct {
	cc *grpc.ClientConn
}

func NewTaskClient(cc *grpc.ClientConn) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) GetList(ctx context.Context, in *GetListTaskReq, opts ...grpc.CallOption) (*GetListTaskResp, error) {
	out := new(GetListTaskResp)
	err := c.cc.Invoke(ctx, "/pb.Task/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Complete(ctx context.Context, in *CompleteTaskReq, opts ...grpc.CallOption) (*CompleteTaskResp, error) {
	out := new(CompleteTaskResp)
	err := c.cc.Invoke(ctx, "/pb.Task/Complete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Get(ctx context.Context, in *GetTaskReq, opts ...grpc.CallOption) (*GetTaskResp, error) {
	out := new(GetTaskResp)
	err := c.cc.Invoke(ctx, "/pb.Task/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	GetList(context.Context, *GetListTaskReq) (*GetListTaskResp, error)
	Complete(context.Context, *CompleteTaskReq) (*CompleteTaskResp, error)
	Get(context.Context, *GetTaskReq) (*GetTaskResp, error)
}

// UnimplementedTaskServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (*UnimplementedTaskServer) GetList(ctx context.Context, req *GetListTaskReq) (*GetListTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (*UnimplementedTaskServer) Complete(ctx context.Context, req *CompleteTaskReq) (*CompleteTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (*UnimplementedTaskServer) Get(ctx context.Context, req *GetTaskReq) (*GetTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Task/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).GetList(ctx, req.(*GetListTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Task/Complete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Complete(ctx, req.(*CompleteTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Task/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Get(ctx, req.(*GetTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _Task_GetList_Handler,
		},
		{
			MethodName: "Complete",
			Handler:    _Task_Complete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Task_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}
