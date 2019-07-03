// Code generated by protoc-gen-go. DO NOT EDIT.
// source: process_definition.proto

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

type StartProcessDefinitionReq struct {
	Id                   string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 *StartProcessDefinitionReqBody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *StartProcessDefinitionReq) Reset()         { *m = StartProcessDefinitionReq{} }
func (m *StartProcessDefinitionReq) String() string { return proto.CompactTextString(m) }
func (*StartProcessDefinitionReq) ProtoMessage()    {}
func (*StartProcessDefinitionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_88807cd440f06d3d, []int{0}
}

func (m *StartProcessDefinitionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProcessDefinitionReq.Unmarshal(m, b)
}
func (m *StartProcessDefinitionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProcessDefinitionReq.Marshal(b, m, deterministic)
}
func (m *StartProcessDefinitionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProcessDefinitionReq.Merge(m, src)
}
func (m *StartProcessDefinitionReq) XXX_Size() int {
	return xxx_messageInfo_StartProcessDefinitionReq.Size(m)
}
func (m *StartProcessDefinitionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProcessDefinitionReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartProcessDefinitionReq proto.InternalMessageInfo

func (m *StartProcessDefinitionReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StartProcessDefinitionReq) GetBody() *StartProcessDefinitionReqBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type StartProcessDefinitionReqBody struct {
	Variables            map[string]*Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BusinessKey          string               `protobuf:"bytes,2,opt,name=businessKey,proto3" json:"businessKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StartProcessDefinitionReqBody) Reset()         { *m = StartProcessDefinitionReqBody{} }
func (m *StartProcessDefinitionReqBody) String() string { return proto.CompactTextString(m) }
func (*StartProcessDefinitionReqBody) ProtoMessage()    {}
func (*StartProcessDefinitionReqBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_88807cd440f06d3d, []int{1}
}

func (m *StartProcessDefinitionReqBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProcessDefinitionReqBody.Unmarshal(m, b)
}
func (m *StartProcessDefinitionReqBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProcessDefinitionReqBody.Marshal(b, m, deterministic)
}
func (m *StartProcessDefinitionReqBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProcessDefinitionReqBody.Merge(m, src)
}
func (m *StartProcessDefinitionReqBody) XXX_Size() int {
	return xxx_messageInfo_StartProcessDefinitionReqBody.Size(m)
}
func (m *StartProcessDefinitionReqBody) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProcessDefinitionReqBody.DiscardUnknown(m)
}

var xxx_messageInfo_StartProcessDefinitionReqBody proto.InternalMessageInfo

func (m *StartProcessDefinitionReqBody) GetVariables() map[string]*Variable {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *StartProcessDefinitionReqBody) GetBusinessKey() string {
	if m != nil {
		return m.BusinessKey
	}
	return ""
}

type StartProcessDefinitionResp struct {
	Item                 *StartProcessDefinitionRespItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Code                 int64                           `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Err                  *CamundaError                   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *StartProcessDefinitionResp) Reset()         { *m = StartProcessDefinitionResp{} }
func (m *StartProcessDefinitionResp) String() string { return proto.CompactTextString(m) }
func (*StartProcessDefinitionResp) ProtoMessage()    {}
func (*StartProcessDefinitionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_88807cd440f06d3d, []int{2}
}

func (m *StartProcessDefinitionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProcessDefinitionResp.Unmarshal(m, b)
}
func (m *StartProcessDefinitionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProcessDefinitionResp.Marshal(b, m, deterministic)
}
func (m *StartProcessDefinitionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProcessDefinitionResp.Merge(m, src)
}
func (m *StartProcessDefinitionResp) XXX_Size() int {
	return xxx_messageInfo_StartProcessDefinitionResp.Size(m)
}
func (m *StartProcessDefinitionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProcessDefinitionResp.DiscardUnknown(m)
}

var xxx_messageInfo_StartProcessDefinitionResp proto.InternalMessageInfo

func (m *StartProcessDefinitionResp) GetItem() *StartProcessDefinitionRespItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StartProcessDefinitionResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StartProcessDefinitionResp) GetErr() *CamundaError {
	if m != nil {
		return m.Err
	}
	return nil
}

type StartProcessDefinitionRespItem struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DefinitionId         string               `protobuf:"bytes,2,opt,name=definitionId,proto3" json:"definitionId,omitempty"`
	BusinessKey          string               `protobuf:"bytes,3,opt,name=businessKey,proto3" json:"businessKey,omitempty"`
	CaseInstanceId       string               `protobuf:"bytes,4,opt,name=caseInstanceId,proto3" json:"caseInstanceId,omitempty"`
	TenantId             string               `protobuf:"bytes,5,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Ended                bool                 `protobuf:"varint,6,opt,name=ended,proto3" json:"ended,omitempty"`
	Suspended            bool                 `protobuf:"varint,7,opt,name=suspended,proto3" json:"suspended,omitempty"`
	Variable             map[string]*Variable `protobuf:"bytes,8,rep,name=variable,proto3" json:"variable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StartProcessDefinitionRespItem) Reset()         { *m = StartProcessDefinitionRespItem{} }
func (m *StartProcessDefinitionRespItem) String() string { return proto.CompactTextString(m) }
func (*StartProcessDefinitionRespItem) ProtoMessage()    {}
func (*StartProcessDefinitionRespItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_88807cd440f06d3d, []int{3}
}

func (m *StartProcessDefinitionRespItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProcessDefinitionRespItem.Unmarshal(m, b)
}
func (m *StartProcessDefinitionRespItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProcessDefinitionRespItem.Marshal(b, m, deterministic)
}
func (m *StartProcessDefinitionRespItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProcessDefinitionRespItem.Merge(m, src)
}
func (m *StartProcessDefinitionRespItem) XXX_Size() int {
	return xxx_messageInfo_StartProcessDefinitionRespItem.Size(m)
}
func (m *StartProcessDefinitionRespItem) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProcessDefinitionRespItem.DiscardUnknown(m)
}

var xxx_messageInfo_StartProcessDefinitionRespItem proto.InternalMessageInfo

func (m *StartProcessDefinitionRespItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StartProcessDefinitionRespItem) GetDefinitionId() string {
	if m != nil {
		return m.DefinitionId
	}
	return ""
}

func (m *StartProcessDefinitionRespItem) GetBusinessKey() string {
	if m != nil {
		return m.BusinessKey
	}
	return ""
}

func (m *StartProcessDefinitionRespItem) GetCaseInstanceId() string {
	if m != nil {
		return m.CaseInstanceId
	}
	return ""
}

func (m *StartProcessDefinitionRespItem) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *StartProcessDefinitionRespItem) GetEnded() bool {
	if m != nil {
		return m.Ended
	}
	return false
}

func (m *StartProcessDefinitionRespItem) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

func (m *StartProcessDefinitionRespItem) GetVariable() map[string]*Variable {
	if m != nil {
		return m.Variable
	}
	return nil
}

type GetProcessDefinitionReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProcessDefinitionReq) Reset()         { *m = GetProcessDefinitionReq{} }
func (m *GetProcessDefinitionReq) String() string { return proto.CompactTextString(m) }
func (*GetProcessDefinitionReq) ProtoMessage()    {}
func (*GetProcessDefinitionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_88807cd440f06d3d, []int{4}
}

func (m *GetProcessDefinitionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProcessDefinitionReq.Unmarshal(m, b)
}
func (m *GetProcessDefinitionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProcessDefinitionReq.Marshal(b, m, deterministic)
}
func (m *GetProcessDefinitionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProcessDefinitionReq.Merge(m, src)
}
func (m *GetProcessDefinitionReq) XXX_Size() int {
	return xxx_messageInfo_GetProcessDefinitionReq.Size(m)
}
func (m *GetProcessDefinitionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProcessDefinitionReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProcessDefinitionReq proto.InternalMessageInfo

func (m *GetProcessDefinitionReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetProcessDefinitionResp struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string        `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Category             string        `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Description          string        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Name                 string        `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Version              int64         `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	Resource             string        `protobuf:"bytes,7,opt,name=resource,proto3" json:"resource,omitempty"`
	DeploymentId         string        `protobuf:"bytes,8,opt,name=deploymentId,proto3" json:"deploymentId,omitempty"`
	Diagram              string        `protobuf:"bytes,9,opt,name=diagram,proto3" json:"diagram,omitempty"`
	Suspended            bool          `protobuf:"varint,10,opt,name=suspended,proto3" json:"suspended,omitempty"`
	TenantId             string        `protobuf:"bytes,11,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	VersionTag           string        `protobuf:"bytes,12,opt,name=versionTag,proto3" json:"versionTag,omitempty"`
	HistoryTimeToLive    int64         `protobuf:"varint,13,opt,name=historyTimeToLive,proto3" json:"historyTimeToLive,omitempty"`
	StartableInTasklist  bool          `protobuf:"varint,14,opt,name=startableInTasklist,proto3" json:"startableInTasklist,omitempty"`
	Code                 int64         `protobuf:"varint,15,opt,name=code,proto3" json:"code,omitempty"`
	Err                  *CamundaError `protobuf:"bytes,16,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetProcessDefinitionResp) Reset()         { *m = GetProcessDefinitionResp{} }
func (m *GetProcessDefinitionResp) String() string { return proto.CompactTextString(m) }
func (*GetProcessDefinitionResp) ProtoMessage()    {}
func (*GetProcessDefinitionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_88807cd440f06d3d, []int{5}
}

func (m *GetProcessDefinitionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProcessDefinitionResp.Unmarshal(m, b)
}
func (m *GetProcessDefinitionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProcessDefinitionResp.Marshal(b, m, deterministic)
}
func (m *GetProcessDefinitionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProcessDefinitionResp.Merge(m, src)
}
func (m *GetProcessDefinitionResp) XXX_Size() int {
	return xxx_messageInfo_GetProcessDefinitionResp.Size(m)
}
func (m *GetProcessDefinitionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProcessDefinitionResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProcessDefinitionResp proto.InternalMessageInfo

func (m *GetProcessDefinitionResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetProcessDefinitionResp) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetDeploymentId() string {
	if m != nil {
		return m.DeploymentId
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetDiagram() string {
	if m != nil {
		return m.Diagram
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

func (m *GetProcessDefinitionResp) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetVersionTag() string {
	if m != nil {
		return m.VersionTag
	}
	return ""
}

func (m *GetProcessDefinitionResp) GetHistoryTimeToLive() int64 {
	if m != nil {
		return m.HistoryTimeToLive
	}
	return 0
}

func (m *GetProcessDefinitionResp) GetStartableInTasklist() bool {
	if m != nil {
		return m.StartableInTasklist
	}
	return false
}

func (m *GetProcessDefinitionResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetProcessDefinitionResp) GetErr() *CamundaError {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*StartProcessDefinitionReq)(nil), "pb.StartProcessDefinitionReq")
	proto.RegisterType((*StartProcessDefinitionReqBody)(nil), "pb.StartProcessDefinitionReqBody")
	proto.RegisterMapType((map[string]*Variable)(nil), "pb.StartProcessDefinitionReqBody.VariablesEntry")
	proto.RegisterType((*StartProcessDefinitionResp)(nil), "pb.StartProcessDefinitionResp")
	proto.RegisterType((*StartProcessDefinitionRespItem)(nil), "pb.StartProcessDefinitionRespItem")
	proto.RegisterMapType((map[string]*Variable)(nil), "pb.StartProcessDefinitionRespItem.VariableEntry")
	proto.RegisterType((*GetProcessDefinitionReq)(nil), "pb.GetProcessDefinitionReq")
	proto.RegisterType((*GetProcessDefinitionResp)(nil), "pb.GetProcessDefinitionResp")
}

func init() { proto.RegisterFile("process_definition.proto", fileDescriptor_88807cd440f06d3d) }

var fileDescriptor_88807cd440f06d3d = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xfe, 0xa5, 0x69, 0xb7, 0xf6, 0xb4, 0xeb, 0x6f, 0x33, 0x48, 0x98, 0xb0, 0x4d, 0x25, 0x17,
	0xa8, 0x48, 0xa8, 0x9a, 0x8a, 0x40, 0x88, 0xcb, 0xc1, 0x34, 0x02, 0x13, 0x42, 0xa1, 0xe2, 0x16,
	0x39, 0xb1, 0x29, 0xd6, 0x9a, 0x38, 0xd8, 0x6e, 0xa5, 0x3e, 0x01, 0x37, 0xbc, 0x04, 0x0f, 0xc6,
	0x5b, 0xf0, 0x00, 0xc8, 0xce, 0x9f, 0xfe, 0x5b, 0x3b, 0x24, 0xee, 0x7c, 0xce, 0x77, 0xce, 0xe7,
	0x63, 0x7f, 0x9f, 0x13, 0xc0, 0x99, 0x14, 0x31, 0x53, 0xea, 0x33, 0x65, 0x5f, 0x78, 0xca, 0x35,
	0x17, 0xe9, 0x20, 0x93, 0x42, 0x0b, 0x54, 0xcb, 0x22, 0xaf, 0x13, 0x8b, 0x24, 0x29, 0x33, 0x7e,
	0x04, 0xf7, 0x3f, 0x6a, 0x22, 0xf5, 0x87, 0xbc, 0xe5, 0x75, 0xd5, 0x11, 0xb2, 0x6f, 0xa8, 0x0b,
	0x35, 0x4e, 0xb1, 0xd3, 0x73, 0xfa, 0xad, 0xb0, 0xc6, 0x29, 0x7a, 0x06, 0xf5, 0x48, 0xd0, 0x39,
	0xae, 0xf5, 0x9c, 0x7e, 0x7b, 0xf8, 0x70, 0x90, 0x45, 0x83, 0xad, 0xcd, 0xe7, 0x82, 0xce, 0x43,
	0x5b, 0xee, 0xff, 0x72, 0xe0, 0x64, 0x67, 0x1d, 0x7a, 0x0f, 0xad, 0x19, 0x91, 0x9c, 0x44, 0x13,
	0xa6, 0xb0, 0xd3, 0x73, 0xfb, 0xed, 0xe1, 0xd9, 0xad, 0xec, 0x83, 0x4f, 0x65, 0xcb, 0x45, 0xaa,
	0xe5, 0x3c, 0x5c, 0x50, 0xa0, 0x1e, 0xb4, 0xa3, 0xa9, 0xe2, 0x29, 0x53, 0xea, 0x1d, 0xcb, 0xe7,
	0x6d, 0x85, 0xcb, 0x29, 0xef, 0x2d, 0x74, 0x57, 0xdb, 0xd1, 0x21, 0xb8, 0xd7, 0x6c, 0x5e, 0x9c,
	0xd6, 0x2c, 0x91, 0x0f, 0x8d, 0x19, 0x99, 0x4c, 0x59, 0x71, 0xde, 0x8e, 0x99, 0xa8, 0x6c, 0x0a,
	0x73, 0xe8, 0x65, 0xed, 0x85, 0xe3, 0xff, 0x70, 0xc0, 0xdb, 0x36, 0xa9, 0xca, 0xd0, 0x73, 0xa8,
	0x73, 0xcd, 0x12, 0xcb, 0xdc, 0x1e, 0xfa, 0xbb, 0xce, 0xa5, 0xb2, 0x40, 0xb3, 0x24, 0xb4, 0xf5,
	0x08, 0x41, 0x3d, 0x16, 0x34, 0xdf, 0xdd, 0x0d, 0xed, 0x1a, 0xf9, 0xe0, 0x32, 0x29, 0xb1, 0x6b,
	0xa9, 0x0e, 0x0d, 0xd5, 0x2b, 0x92, 0x4c, 0x53, 0x4a, 0x2e, 0xa4, 0x14, 0x32, 0x34, 0xa0, 0xff,
	0xdd, 0x85, 0xd3, 0xdd, 0x1b, 0x6c, 0x08, 0xeb, 0x43, 0x67, 0xe1, 0x95, 0x80, 0x16, 0x17, 0xb6,
	0x92, 0x5b, 0xbf, 0x53, 0x77, 0xe3, 0x4e, 0xd1, 0x23, 0xe8, 0xc6, 0x44, 0xb1, 0x20, 0x55, 0x9a,
	0xa4, 0x31, 0x0b, 0x28, 0xae, 0xdb, 0xa2, 0xb5, 0x2c, 0xf2, 0xa0, 0xa9, 0x59, 0x4a, 0x52, 0x1d,
	0x50, 0xdc, 0xb0, 0x15, 0x55, 0x8c, 0xee, 0x42, 0x83, 0xa5, 0x94, 0x51, 0xbc, 0xd7, 0x73, 0xfa,
	0xcd, 0x30, 0x0f, 0xd0, 0x31, 0xb4, 0xd4, 0x54, 0x65, 0x39, 0xb2, 0x6f, 0x91, 0x45, 0x02, 0x5d,
	0x41, 0xb3, 0x94, 0x1e, 0x37, 0x6f, 0x37, 0x4f, 0x7e, 0x07, 0x95, 0x92, 0xb9, 0x79, 0x2a, 0x06,
	0x2f, 0x80, 0x83, 0x15, 0xe8, 0x1f, 0x8c, 0xf1, 0x18, 0xee, 0x5d, 0xb2, 0xbf, 0x7a, 0x5a, 0xfe,
	0x6f, 0x17, 0xf0, 0xcd, 0xb5, 0x2a, 0xdb, 0x90, 0xab, 0x98, 0xa8, 0xb6, 0x98, 0xc8, 0x83, 0x66,
	0x4c, 0x34, 0x1b, 0x0b, 0x59, 0x2a, 0x53, 0xc5, 0x46, 0x38, 0xca, 0x54, 0x2c, 0x79, 0x66, 0x08,
	0x0b, 0x4d, 0x96, 0x53, 0xc6, 0x69, 0x29, 0x49, 0x58, 0x21, 0x86, 0x5d, 0x23, 0x0c, 0xfb, 0x33,
	0x26, 0x95, 0xe9, 0xd8, 0xb3, 0x06, 0x2c, 0x43, 0xb3, 0x97, 0x64, 0x4a, 0x4c, 0x65, 0xcc, 0xac,
	0x16, 0xad, 0xb0, 0x8a, 0x73, 0x23, 0x65, 0x13, 0x31, 0x4f, 0x98, 0x95, 0xb7, 0x59, 0x1a, 0x69,
	0x91, 0x33, 0xcc, 0x94, 0x93, 0xb1, 0x24, 0x09, 0x6e, 0x59, 0xb8, 0x0c, 0x57, 0x65, 0x86, 0x75,
	0x99, 0x97, 0x6d, 0xd3, 0x5e, 0xb3, 0xcd, 0x29, 0x40, 0x31, 0xde, 0x88, 0x8c, 0x71, 0xc7, 0xa2,
	0x4b, 0x19, 0xf4, 0x04, 0x8e, 0xbe, 0x72, 0xa5, 0x85, 0x9c, 0x8f, 0x78, 0xc2, 0x46, 0xe2, 0x8a,
	0xcf, 0x18, 0x3e, 0xb0, 0xe7, 0xda, 0x04, 0xd0, 0x19, 0xdc, 0x51, 0xc6, 0x3c, 0x46, 0xcf, 0x20,
	0x1d, 0x11, 0x75, 0x3d, 0xe1, 0x4a, 0xe3, 0xae, 0x9d, 0xe8, 0x26, 0xa8, 0x7a, 0xab, 0xff, 0x6f,
	0xbe, 0xd5, 0xc3, 0x1d, 0x6f, 0x75, 0xf8, 0xd3, 0x81, 0xa3, 0x0d, 0xcd, 0xd1, 0x1b, 0x68, 0x58,
	0xf3, 0xa2, 0x93, 0x9d, 0x1f, 0x41, 0xef, 0x74, 0xb7, 0xcd, 0xfd, 0xff, 0xd0, 0x39, 0xb8, 0x97,
	0x4c, 0xa3, 0x07, 0xa6, 0x70, 0x8b, 0x15, 0xbd, 0xe3, 0xed, 0xa0, 0xe1, 0x88, 0xf6, 0xec, 0x9f,
	0xe2, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x37, 0xa2, 0xf3, 0x57, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProcessDefinitionClient is the client API for ProcessDefinition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessDefinitionClient interface {
	Start(ctx context.Context, in *StartProcessDefinitionReq, opts ...grpc.CallOption) (*StartProcessDefinitionResp, error)
	Get(ctx context.Context, in *GetProcessDefinitionReq, opts ...grpc.CallOption) (*GetProcessDefinitionResp, error)
}

type processDefinitionClient struct {
	cc *grpc.ClientConn
}

func NewProcessDefinitionClient(cc *grpc.ClientConn) ProcessDefinitionClient {
	return &processDefinitionClient{cc}
}

func (c *processDefinitionClient) Start(ctx context.Context, in *StartProcessDefinitionReq, opts ...grpc.CallOption) (*StartProcessDefinitionResp, error) {
	out := new(StartProcessDefinitionResp)
	err := c.cc.Invoke(ctx, "/pb.ProcessDefinition/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processDefinitionClient) Get(ctx context.Context, in *GetProcessDefinitionReq, opts ...grpc.CallOption) (*GetProcessDefinitionResp, error) {
	out := new(GetProcessDefinitionResp)
	err := c.cc.Invoke(ctx, "/pb.ProcessDefinition/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessDefinitionServer is the server API for ProcessDefinition service.
type ProcessDefinitionServer interface {
	Start(context.Context, *StartProcessDefinitionReq) (*StartProcessDefinitionResp, error)
	Get(context.Context, *GetProcessDefinitionReq) (*GetProcessDefinitionResp, error)
}

// UnimplementedProcessDefinitionServer can be embedded to have forward compatible implementations.
type UnimplementedProcessDefinitionServer struct {
}

func (*UnimplementedProcessDefinitionServer) Start(ctx context.Context, req *StartProcessDefinitionReq) (*StartProcessDefinitionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedProcessDefinitionServer) Get(ctx context.Context, req *GetProcessDefinitionReq) (*GetProcessDefinitionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterProcessDefinitionServer(s *grpc.Server, srv ProcessDefinitionServer) {
	s.RegisterService(&_ProcessDefinition_serviceDesc, srv)
}

func _ProcessDefinition_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartProcessDefinitionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessDefinitionServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProcessDefinition/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessDefinitionServer).Start(ctx, req.(*StartProcessDefinitionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessDefinition_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessDefinitionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessDefinitionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProcessDefinition/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessDefinitionServer).Get(ctx, req.(*GetProcessDefinitionReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessDefinition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProcessDefinition",
	HandlerType: (*ProcessDefinitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _ProcessDefinition_Start_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProcessDefinition_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "process_definition.proto",
}