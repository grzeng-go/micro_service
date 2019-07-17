// Code generated by protoc-gen-go. DO NOT EDIT.
// source: term.proto

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

type ListTermInfoRequest struct {
	Page                 int32          `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32          `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Item                 *TermInfoField `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListTermInfoRequest) Reset()         { *m = ListTermInfoRequest{} }
func (m *ListTermInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ListTermInfoRequest) ProtoMessage()    {}
func (*ListTermInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{0}
}

func (m *ListTermInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTermInfoRequest.Unmarshal(m, b)
}
func (m *ListTermInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTermInfoRequest.Marshal(b, m, deterministic)
}
func (m *ListTermInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTermInfoRequest.Merge(m, src)
}
func (m *ListTermInfoRequest) XXX_Size() int {
	return xxx_messageInfo_ListTermInfoRequest.Size(m)
}
func (m *ListTermInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTermInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTermInfoRequest proto.InternalMessageInfo

func (m *ListTermInfoRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTermInfoRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListTermInfoRequest) GetItem() *TermInfoField {
	if m != nil {
		return m.Item
	}
	return nil
}

type ListTermInfoReply struct {
	Page                 int32            `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32            `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Count                int32            `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Items                []*TermInfoField `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	Err                  *Error           `protobuf:"bytes,5,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListTermInfoReply) Reset()         { *m = ListTermInfoReply{} }
func (m *ListTermInfoReply) String() string { return proto.CompactTextString(m) }
func (*ListTermInfoReply) ProtoMessage()    {}
func (*ListTermInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{1}
}

func (m *ListTermInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTermInfoReply.Unmarshal(m, b)
}
func (m *ListTermInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTermInfoReply.Marshal(b, m, deterministic)
}
func (m *ListTermInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTermInfoReply.Merge(m, src)
}
func (m *ListTermInfoReply) XXX_Size() int {
	return xxx_messageInfo_ListTermInfoReply.Size(m)
}
func (m *ListTermInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTermInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListTermInfoReply proto.InternalMessageInfo

func (m *ListTermInfoReply) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTermInfoReply) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListTermInfoReply) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListTermInfoReply) GetItems() []*TermInfoField {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ListTermInfoReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type TermInfoField struct {
	MchtCd               string   `protobuf:"bytes,1,opt,name=mchtCd,proto3" json:"mchtCd,omitempty"`
	TermId               string   `protobuf:"bytes,2,opt,name=termId,proto3" json:"termId,omitempty"`
	TermTp               string   `protobuf:"bytes,3,opt,name=termTp,proto3" json:"termTp,omitempty"`
	Belong               string   `protobuf:"bytes,4,opt,name=belong,proto3" json:"belong,omitempty"`
	BelongSub            string   `protobuf:"bytes,5,opt,name=belongSub,proto3" json:"belongSub,omitempty"`
	TmnlMoneyIntype      string   `protobuf:"bytes,6,opt,name=tmnlMoneyIntype,proto3" json:"tmnlMoneyIntype,omitempty"`
	TmnlMoney            float64  `protobuf:"fixed64,7,opt,name=tmnlMoney,proto3" json:"tmnlMoney,omitempty"`
	TmnlBrand            string   `protobuf:"bytes,8,opt,name=tmnlBrand,proto3" json:"tmnlBrand,omitempty"`
	TmnlModelNo          string   `protobuf:"bytes,9,opt,name=tmnlModelNo,proto3" json:"tmnlModelNo,omitempty"`
	TmnlBarcode          string   `protobuf:"bytes,10,opt,name=tmnlBarcode,proto3" json:"tmnlBarcode,omitempty"`
	DeviceCd             string   `protobuf:"bytes,11,opt,name=deviceCd,proto3" json:"deviceCd,omitempty"`
	InstallLocation      string   `protobuf:"bytes,12,opt,name=installLocation,proto3" json:"installLocation,omitempty"`
	TmnlIntype           string   `protobuf:"bytes,13,opt,name=tmnlIntype,proto3" json:"tmnlIntype,omitempty"`
	DialOut              string   `protobuf:"bytes,14,opt,name=dialOut,proto3" json:"dialOut,omitempty"`
	DealTypes            string   `protobuf:"bytes,15,opt,name=dealTypes,proto3" json:"dealTypes,omitempty"`
	RecOprId             string   `protobuf:"bytes,16,opt,name=recOprId,proto3" json:"recOprId,omitempty"`
	RecUpdOpr            string   `protobuf:"bytes,17,opt,name=recUpdOpr,proto3" json:"recUpdOpr,omitempty"`
	CreatedAt            string   `protobuf:"bytes,18,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,19,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	AppCd                string   `protobuf:"bytes,20,opt,name=appCd,proto3" json:"appCd,omitempty"`
	SystemFlag           string   `protobuf:"bytes,21,opt,name=systemFlag,proto3" json:"systemFlag,omitempty"`
	Status               string   `protobuf:"bytes,22,opt,name=status,proto3" json:"status,omitempty"`
	ActiveCode           string   `protobuf:"bytes,23,opt,name=activeCode,proto3" json:"activeCode,omitempty"`
	NoFlag               string   `protobuf:"bytes,24,opt,name=noFlag,proto3" json:"noFlag,omitempty"`
	MsgResvFld1          string   `protobuf:"bytes,25,opt,name=msgResvFld1,proto3" json:"msgResvFld1,omitempty"`
	MsgResvFld2          string   `protobuf:"bytes,26,opt,name=msgResvFld2,proto3" json:"msgResvFld2,omitempty"`
	MsgResvFld3          string   `protobuf:"bytes,27,opt,name=msgResvFld3,proto3" json:"msgResvFld3,omitempty"`
	MsgResvFld4          string   `protobuf:"bytes,28,opt,name=msgResvFld4,proto3" json:"msgResvFld4,omitempty"`
	MsgResvFld5          string   `protobuf:"bytes,29,opt,name=msgResvFld5,proto3" json:"msgResvFld5,omitempty"`
	MsgResvFld6          string   `protobuf:"bytes,30,opt,name=msgResvFld6,proto3" json:"msgResvFld6,omitempty"`
	MsgResvFld7          string   `protobuf:"bytes,31,opt,name=msgResvFld7,proto3" json:"msgResvFld7,omitempty"`
	MsgResvFld8          string   `protobuf:"bytes,32,opt,name=msgResvFld8,proto3" json:"msgResvFld8,omitempty"`
	MsgResvFld9          string   `protobuf:"bytes,33,opt,name=msgResvFld9,proto3" json:"msgResvFld9,omitempty"`
	MsgResvFld10         string   `protobuf:"bytes,34,opt,name=msgResvFld10,proto3" json:"msgResvFld10,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TermInfoField) Reset()         { *m = TermInfoField{} }
func (m *TermInfoField) String() string { return proto.CompactTextString(m) }
func (*TermInfoField) ProtoMessage()    {}
func (*TermInfoField) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{2}
}

func (m *TermInfoField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TermInfoField.Unmarshal(m, b)
}
func (m *TermInfoField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TermInfoField.Marshal(b, m, deterministic)
}
func (m *TermInfoField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TermInfoField.Merge(m, src)
}
func (m *TermInfoField) XXX_Size() int {
	return xxx_messageInfo_TermInfoField.Size(m)
}
func (m *TermInfoField) XXX_DiscardUnknown() {
	xxx_messageInfo_TermInfoField.DiscardUnknown(m)
}

var xxx_messageInfo_TermInfoField proto.InternalMessageInfo

func (m *TermInfoField) GetMchtCd() string {
	if m != nil {
		return m.MchtCd
	}
	return ""
}

func (m *TermInfoField) GetTermId() string {
	if m != nil {
		return m.TermId
	}
	return ""
}

func (m *TermInfoField) GetTermTp() string {
	if m != nil {
		return m.TermTp
	}
	return ""
}

func (m *TermInfoField) GetBelong() string {
	if m != nil {
		return m.Belong
	}
	return ""
}

func (m *TermInfoField) GetBelongSub() string {
	if m != nil {
		return m.BelongSub
	}
	return ""
}

func (m *TermInfoField) GetTmnlMoneyIntype() string {
	if m != nil {
		return m.TmnlMoneyIntype
	}
	return ""
}

func (m *TermInfoField) GetTmnlMoney() float64 {
	if m != nil {
		return m.TmnlMoney
	}
	return 0
}

func (m *TermInfoField) GetTmnlBrand() string {
	if m != nil {
		return m.TmnlBrand
	}
	return ""
}

func (m *TermInfoField) GetTmnlModelNo() string {
	if m != nil {
		return m.TmnlModelNo
	}
	return ""
}

func (m *TermInfoField) GetTmnlBarcode() string {
	if m != nil {
		return m.TmnlBarcode
	}
	return ""
}

func (m *TermInfoField) GetDeviceCd() string {
	if m != nil {
		return m.DeviceCd
	}
	return ""
}

func (m *TermInfoField) GetInstallLocation() string {
	if m != nil {
		return m.InstallLocation
	}
	return ""
}

func (m *TermInfoField) GetTmnlIntype() string {
	if m != nil {
		return m.TmnlIntype
	}
	return ""
}

func (m *TermInfoField) GetDialOut() string {
	if m != nil {
		return m.DialOut
	}
	return ""
}

func (m *TermInfoField) GetDealTypes() string {
	if m != nil {
		return m.DealTypes
	}
	return ""
}

func (m *TermInfoField) GetRecOprId() string {
	if m != nil {
		return m.RecOprId
	}
	return ""
}

func (m *TermInfoField) GetRecUpdOpr() string {
	if m != nil {
		return m.RecUpdOpr
	}
	return ""
}

func (m *TermInfoField) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *TermInfoField) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *TermInfoField) GetAppCd() string {
	if m != nil {
		return m.AppCd
	}
	return ""
}

func (m *TermInfoField) GetSystemFlag() string {
	if m != nil {
		return m.SystemFlag
	}
	return ""
}

func (m *TermInfoField) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TermInfoField) GetActiveCode() string {
	if m != nil {
		return m.ActiveCode
	}
	return ""
}

func (m *TermInfoField) GetNoFlag() string {
	if m != nil {
		return m.NoFlag
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld1() string {
	if m != nil {
		return m.MsgResvFld1
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld2() string {
	if m != nil {
		return m.MsgResvFld2
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld3() string {
	if m != nil {
		return m.MsgResvFld3
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld4() string {
	if m != nil {
		return m.MsgResvFld4
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld5() string {
	if m != nil {
		return m.MsgResvFld5
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld6() string {
	if m != nil {
		return m.MsgResvFld6
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld7() string {
	if m != nil {
		return m.MsgResvFld7
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld8() string {
	if m != nil {
		return m.MsgResvFld8
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld9() string {
	if m != nil {
		return m.MsgResvFld9
	}
	return ""
}

func (m *TermInfoField) GetMsgResvFld10() string {
	if m != nil {
		return m.MsgResvFld10
	}
	return ""
}

type SaveTermRequest struct {
	Item                 *TermInfoField `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SaveTermRequest) Reset()         { *m = SaveTermRequest{} }
func (m *SaveTermRequest) String() string { return proto.CompactTextString(m) }
func (*SaveTermRequest) ProtoMessage()    {}
func (*SaveTermRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{3}
}

func (m *SaveTermRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTermRequest.Unmarshal(m, b)
}
func (m *SaveTermRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTermRequest.Marshal(b, m, deterministic)
}
func (m *SaveTermRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTermRequest.Merge(m, src)
}
func (m *SaveTermRequest) XXX_Size() int {
	return xxx_messageInfo_SaveTermRequest.Size(m)
}
func (m *SaveTermRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTermRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTermRequest proto.InternalMessageInfo

func (m *SaveTermRequest) GetItem() *TermInfoField {
	if m != nil {
		return m.Item
	}
	return nil
}

type SaveTermReply struct {
	Err                  *Error   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveTermReply) Reset()         { *m = SaveTermReply{} }
func (m *SaveTermReply) String() string { return proto.CompactTextString(m) }
func (*SaveTermReply) ProtoMessage()    {}
func (*SaveTermReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{4}
}

func (m *SaveTermReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTermReply.Unmarshal(m, b)
}
func (m *SaveTermReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTermReply.Marshal(b, m, deterministic)
}
func (m *SaveTermReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTermReply.Merge(m, src)
}
func (m *SaveTermReply) XXX_Size() int {
	return xxx_messageInfo_SaveTermReply.Size(m)
}
func (m *SaveTermReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTermReply.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTermReply proto.InternalMessageInfo

func (m *SaveTermReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type TermRiskField struct {
	MchtCd               string   `protobuf:"bytes,1,opt,name=MchtCd,proto3" json:"MchtCd,omitempty"`
	TermId               string   `protobuf:"bytes,2,opt,name=TermId,proto3" json:"TermId,omitempty"`
	CardType             string   `protobuf:"bytes,3,opt,name=CardType,proto3" json:"CardType,omitempty"`
	TotalLimitMoney      float64  `protobuf:"fixed64,4,opt,name=TotalLimitMoney,proto3" json:"TotalLimitMoney,omitempty"`
	AccpetStartTime      string   `protobuf:"bytes,5,opt,name=AccpetStartTime,proto3" json:"AccpetStartTime,omitempty"`
	AccpetStartDate      string   `protobuf:"bytes,6,opt,name=AccpetStartDate,proto3" json:"AccpetStartDate,omitempty"`
	AccpetEndTime        string   `protobuf:"bytes,7,opt,name=AccpetEndTime,proto3" json:"AccpetEndTime,omitempty"`
	AccpetEndDate        string   `protobuf:"bytes,8,opt,name=AccpetEndDate,proto3" json:"AccpetEndDate,omitempty"`
	SingleLimitMoney     float64  `protobuf:"fixed64,9,opt,name=SingleLimitMoney,proto3" json:"SingleLimitMoney,omitempty"`
	ControlWay           string   `protobuf:"bytes,10,opt,name=ControlWay,proto3" json:"ControlWay,omitempty"`
	SingleMinMoney       float64  `protobuf:"fixed64,11,opt,name=SingleMinMoney,proto3" json:"SingleMinMoney,omitempty"`
	TotalPeriod          string   `protobuf:"bytes,12,opt,name=TotalPeriod,proto3" json:"TotalPeriod,omitempty"`
	RecOprId             string   `protobuf:"bytes,13,opt,name=RecOprId,proto3" json:"RecOprId,omitempty"`
	RecUpdOpr            string   `protobuf:"bytes,14,opt,name=RecUpdOpr,proto3" json:"RecUpdOpr,omitempty"`
	OperIn               string   `protobuf:"bytes,15,opt,name=OperIn,proto3" json:"OperIn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TermRiskField) Reset()         { *m = TermRiskField{} }
func (m *TermRiskField) String() string { return proto.CompactTextString(m) }
func (*TermRiskField) ProtoMessage()    {}
func (*TermRiskField) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{5}
}

func (m *TermRiskField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TermRiskField.Unmarshal(m, b)
}
func (m *TermRiskField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TermRiskField.Marshal(b, m, deterministic)
}
func (m *TermRiskField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TermRiskField.Merge(m, src)
}
func (m *TermRiskField) XXX_Size() int {
	return xxx_messageInfo_TermRiskField.Size(m)
}
func (m *TermRiskField) XXX_DiscardUnknown() {
	xxx_messageInfo_TermRiskField.DiscardUnknown(m)
}

var xxx_messageInfo_TermRiskField proto.InternalMessageInfo

func (m *TermRiskField) GetMchtCd() string {
	if m != nil {
		return m.MchtCd
	}
	return ""
}

func (m *TermRiskField) GetTermId() string {
	if m != nil {
		return m.TermId
	}
	return ""
}

func (m *TermRiskField) GetCardType() string {
	if m != nil {
		return m.CardType
	}
	return ""
}

func (m *TermRiskField) GetTotalLimitMoney() float64 {
	if m != nil {
		return m.TotalLimitMoney
	}
	return 0
}

func (m *TermRiskField) GetAccpetStartTime() string {
	if m != nil {
		return m.AccpetStartTime
	}
	return ""
}

func (m *TermRiskField) GetAccpetStartDate() string {
	if m != nil {
		return m.AccpetStartDate
	}
	return ""
}

func (m *TermRiskField) GetAccpetEndTime() string {
	if m != nil {
		return m.AccpetEndTime
	}
	return ""
}

func (m *TermRiskField) GetAccpetEndDate() string {
	if m != nil {
		return m.AccpetEndDate
	}
	return ""
}

func (m *TermRiskField) GetSingleLimitMoney() float64 {
	if m != nil {
		return m.SingleLimitMoney
	}
	return 0
}

func (m *TermRiskField) GetControlWay() string {
	if m != nil {
		return m.ControlWay
	}
	return ""
}

func (m *TermRiskField) GetSingleMinMoney() float64 {
	if m != nil {
		return m.SingleMinMoney
	}
	return 0
}

func (m *TermRiskField) GetTotalPeriod() string {
	if m != nil {
		return m.TotalPeriod
	}
	return ""
}

func (m *TermRiskField) GetRecOprId() string {
	if m != nil {
		return m.RecOprId
	}
	return ""
}

func (m *TermRiskField) GetRecUpdOpr() string {
	if m != nil {
		return m.RecUpdOpr
	}
	return ""
}

func (m *TermRiskField) GetOperIn() string {
	if m != nil {
		return m.OperIn
	}
	return ""
}

type SaveTermRiskRequest struct {
	Item                 *TermRiskField `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SaveTermRiskRequest) Reset()         { *m = SaveTermRiskRequest{} }
func (m *SaveTermRiskRequest) String() string { return proto.CompactTextString(m) }
func (*SaveTermRiskRequest) ProtoMessage()    {}
func (*SaveTermRiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{6}
}

func (m *SaveTermRiskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTermRiskRequest.Unmarshal(m, b)
}
func (m *SaveTermRiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTermRiskRequest.Marshal(b, m, deterministic)
}
func (m *SaveTermRiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTermRiskRequest.Merge(m, src)
}
func (m *SaveTermRiskRequest) XXX_Size() int {
	return xxx_messageInfo_SaveTermRiskRequest.Size(m)
}
func (m *SaveTermRiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTermRiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTermRiskRequest proto.InternalMessageInfo

func (m *SaveTermRiskRequest) GetItem() *TermRiskField {
	if m != nil {
		return m.Item
	}
	return nil
}

type SaveTermRiskReply struct {
	Err                  *Error   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveTermRiskReply) Reset()         { *m = SaveTermRiskReply{} }
func (m *SaveTermRiskReply) String() string { return proto.CompactTextString(m) }
func (*SaveTermRiskReply) ProtoMessage()    {}
func (*SaveTermRiskReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f9b522c29056d7, []int{7}
}

func (m *SaveTermRiskReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTermRiskReply.Unmarshal(m, b)
}
func (m *SaveTermRiskReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTermRiskReply.Marshal(b, m, deterministic)
}
func (m *SaveTermRiskReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTermRiskReply.Merge(m, src)
}
func (m *SaveTermRiskReply) XXX_Size() int {
	return xxx_messageInfo_SaveTermRiskReply.Size(m)
}
func (m *SaveTermRiskReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTermRiskReply.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTermRiskReply proto.InternalMessageInfo

func (m *SaveTermRiskReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*ListTermInfoRequest)(nil), "pb.ListTermInfoRequest")
	proto.RegisterType((*ListTermInfoReply)(nil), "pb.ListTermInfoReply")
	proto.RegisterType((*TermInfoField)(nil), "pb.TermInfoField")
	proto.RegisterType((*SaveTermRequest)(nil), "pb.SaveTermRequest")
	proto.RegisterType((*SaveTermReply)(nil), "pb.SaveTermReply")
	proto.RegisterType((*TermRiskField)(nil), "pb.TermRiskField")
	proto.RegisterType((*SaveTermRiskRequest)(nil), "pb.SaveTermRiskRequest")
	proto.RegisterType((*SaveTermRiskReply)(nil), "pb.SaveTermRiskReply")
}

func init() { proto.RegisterFile("term.proto", fileDescriptor_d8f9b522c29056d7) }

var fileDescriptor_d8f9b522c29056d7 = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0xab, 0xc6, 0x4e, 0x6c, 0x3a, 0x69, 0x6a, 0xa6, 0x1f, 0x9c, 0xdb, 0x75, 0x9e, 0xb0,
	0x8f, 0x60, 0x18, 0x82, 0x2e, 0x49, 0xdb, 0x14, 0xd8, 0x45, 0x5b, 0x6f, 0x01, 0x0c, 0x24, 0xcb,
	0xa0, 0x78, 0xd8, 0x35, 0x2d, 0x72, 0x9e, 0x50, 0x7d, 0x70, 0x14, 0x1d, 0xc0, 0x7b, 0x87, 0x5d,
	0xef, 0x11, 0xf6, 0x1a, 0x7b, 0xb4, 0xe1, 0xf0, 0x48, 0xb2, 0xc8, 0x18, 0x6b, 0xef, 0x74, 0x7e,
	0xe7, 0x7f, 0x48, 0x1e, 0xea, 0xe8, 0x6f, 0x13, 0x62, 0xa4, 0xce, 0x8e, 0x94, 0x2e, 0x4c, 0x41,
	0xef, 0xaa, 0xf9, 0x68, 0x20, 0xb5, 0x2e, 0x34, 0x82, 0x50, 0x90, 0x83, 0x8b, 0xa4, 0x34, 0x33,
	0xa9, 0xb3, 0x69, 0xfe, 0x5b, 0x11, 0xc9, 0x3f, 0x96, 0xb2, 0x34, 0x94, 0x92, 0x8e, 0xe2, 0x0b,
	0xc9, 0x82, 0x71, 0x70, 0xd8, 0x8d, 0xec, 0x33, 0xb0, 0x32, 0xf9, 0x53, 0xb2, 0xbb, 0xc8, 0xe0,
	0x99, 0x7e, 0x49, 0x3a, 0x89, 0x91, 0x19, 0xdb, 0x1a, 0x07, 0x87, 0x83, 0xe3, 0xe1, 0x91, 0x9a,
	0x1f, 0xd5, 0x4b, 0x9d, 0x27, 0x32, 0x15, 0x91, 0x4d, 0x87, 0x7f, 0x07, 0x64, 0xe8, 0x6e, 0xa3,
	0xd2, 0xd5, 0x47, 0x6f, 0xf2, 0x80, 0x74, 0xe3, 0x62, 0x99, 0x1b, 0xbb, 0x4b, 0x37, 0xc2, 0x80,
	0x7e, 0x4d, 0xba, 0xb0, 0x76, 0xc9, 0x3a, 0xe3, 0xad, 0xcd, 0x7b, 0x63, 0x9e, 0x3e, 0x21, 0x5b,
	0x52, 0x6b, 0xd6, 0xb5, 0x47, 0xec, 0x83, 0xcc, 0x5e, 0x40, 0x04, 0x34, 0xfc, 0xa7, 0x47, 0xf6,
	0x9c, 0x2a, 0xfa, 0x88, 0x6c, 0x67, 0xf1, 0xef, 0x66, 0x22, 0xec, 0xb9, 0xfa, 0x51, 0x15, 0x01,
	0x87, 0x8b, 0x9c, 0x0a, 0x7b, 0xb6, 0x7e, 0x54, 0x45, 0x35, 0x9f, 0x29, 0x7b, 0xbc, 0x8a, 0xcf,
	0x14, 0xf0, 0xb9, 0x4c, 0x8b, 0x7c, 0xc1, 0x3a, 0xc8, 0x31, 0xa2, 0x4f, 0x49, 0x1f, 0x9f, 0xae,
	0x97, 0x73, 0x7b, 0xa8, 0x7e, 0xb4, 0x06, 0xf4, 0x90, 0xec, 0x9b, 0x2c, 0x4f, 0x2f, 0x8b, 0x5c,
	0xae, 0xa6, 0xb9, 0x59, 0x29, 0xc9, 0xb6, 0xad, 0xc6, 0xc7, 0xb0, 0x4e, 0x83, 0xd8, 0xce, 0x38,
	0x38, 0x0c, 0xa2, 0x35, 0xa8, 0xb3, 0xef, 0x34, 0xcf, 0x05, 0xeb, 0xe1, 0x2e, 0x0d, 0xa0, 0x63,
	0x32, 0x40, 0xa9, 0x90, 0xe9, 0x4f, 0x05, 0xeb, 0xdb, 0x7c, 0x1b, 0xd5, 0x8a, 0x77, 0x5c, 0xc7,
	0x85, 0x90, 0x8c, 0xac, 0x15, 0x15, 0xa2, 0x23, 0xd2, 0x13, 0xf2, 0x26, 0x89, 0xe5, 0x44, 0xb0,
	0x81, 0x4d, 0x37, 0x31, 0x74, 0x91, 0xe4, 0xa5, 0xe1, 0x69, 0x7a, 0x51, 0xc4, 0xdc, 0x24, 0x45,
	0xce, 0x76, 0xb1, 0x0b, 0x0f, 0xd3, 0x67, 0x84, 0xc0, 0xa2, 0x55, 0xab, 0x7b, 0x56, 0xd4, 0x22,
	0x94, 0x91, 0x1d, 0x91, 0xf0, 0xf4, 0x6a, 0x69, 0xd8, 0x3d, 0x9b, 0xac, 0x43, 0xe8, 0x50, 0x48,
	0x9e, 0xce, 0x56, 0x4a, 0x96, 0x6c, 0x1f, 0x3b, 0x6c, 0x00, 0x9c, 0x4e, 0xcb, 0xf8, 0x4a, 0xe9,
	0xa9, 0x60, 0xf7, 0xf1, 0x74, 0x75, 0x0c, 0x95, 0x5a, 0xc6, 0xbf, 0x28, 0x71, 0xa5, 0x34, 0x1b,
	0x62, 0x65, 0x03, 0x20, 0x1b, 0x6b, 0xc9, 0x8d, 0x14, 0x6f, 0x0d, 0xa3, 0x98, 0x6d, 0x00, 0x64,
	0x97, 0x4a, 0x54, 0xd9, 0x03, 0xcc, 0x36, 0x00, 0x26, 0x95, 0x2b, 0x35, 0x11, 0xec, 0x81, 0xcd,
	0x60, 0x00, 0x3d, 0x96, 0xab, 0xd2, 0xc8, 0xec, 0x3c, 0xe5, 0x0b, 0xf6, 0x10, 0x7b, 0x5c, 0x13,
	0x98, 0x94, 0xd2, 0x70, 0xb3, 0x2c, 0xd9, 0x23, 0x9c, 0x14, 0x8c, 0xa0, 0x8e, 0xc7, 0x26, 0xb9,
	0x91, 0x13, 0x78, 0x05, 0x8f, 0xb1, 0x6e, 0x4d, 0xa0, 0x2e, 0x2f, 0xec, 0x9a, 0x0c, 0xeb, 0x30,
	0x82, 0x77, 0x97, 0x95, 0x8b, 0x48, 0x96, 0x37, 0xe7, 0xa9, 0xf8, 0x8e, 0x7d, 0x82, 0xef, 0xae,
	0x85, 0x5c, 0xc5, 0x31, 0x1b, 0xf9, 0x8a, 0x63, 0x57, 0x71, 0xc2, 0x9e, 0xf8, 0x8a, 0x13, 0x57,
	0x71, 0xca, 0x9e, 0xfa, 0x8a, 0x53, 0x57, 0xf1, 0x82, 0x7d, 0xea, 0x2b, 0x5e, 0xb8, 0x8a, 0x97,
	0xec, 0x99, 0xaf, 0x78, 0xe9, 0x2a, 0x5e, 0xb1, 0xcf, 0x7c, 0xc5, 0x2b, 0x57, 0x71, 0xc6, 0xc6,
	0xbe, 0xe2, 0xcc, 0x55, 0xbc, 0x66, 0x9f, 0xfb, 0x8a, 0xd7, 0x34, 0x24, 0xbb, 0xad, 0xeb, 0x79,
	0xce, 0x42, 0x2b, 0x71, 0x58, 0x78, 0x46, 0xf6, 0xaf, 0xf9, 0x8d, 0x04, 0xb3, 0xa8, 0x5d, 0xb2,
	0x76, 0xbf, 0xe0, 0xff, 0xdd, 0xef, 0x5b, 0xb2, 0xb7, 0xae, 0x04, 0xe3, 0xab, 0x1c, 0x29, 0xd8,
	0xe8, 0x48, 0x7f, 0x75, 0xd0, 0x91, 0xa2, 0xa4, 0x7c, 0xdf, 0x38, 0xd2, 0xa5, 0xe3, 0x48, 0x97,
	0x8d, 0x23, 0xcd, 0x1c, 0x47, 0xc2, 0x08, 0x66, 0x7f, 0xc2, 0xb5, 0x80, 0x0f, 0xa1, 0xf2, 0xa4,
	0x26, 0x86, 0x2f, 0x73, 0x56, 0x18, 0x9e, 0x5e, 0x24, 0x59, 0x62, 0xd0, 0x3b, 0x3a, 0xd6, 0x3b,
	0x7c, 0x0c, 0xca, 0xb7, 0x71, 0xac, 0xa4, 0xb9, 0x36, 0x5c, 0x9b, 0x59, 0x92, 0xc9, 0xca, 0xad,
	0x7c, 0xec, 0x29, 0x7f, 0xe0, 0xa6, 0xf1, 0x2c, 0x0f, 0xd3, 0x2f, 0xc8, 0x1e, 0xa2, 0x1f, 0x73,
	0x61, 0x57, 0xdc, 0xb1, 0x3a, 0x17, 0x3a, 0x2a, 0xbb, 0x5a, 0xcf, 0x53, 0xd9, 0xb5, 0xbe, 0x21,
	0xf7, 0xaf, 0x93, 0x7c, 0x91, 0xca, 0x56, 0x2b, 0x7d, 0xdb, 0xca, 0x2d, 0x0e, 0x5f, 0xd2, 0xa4,
	0xc8, 0x8d, 0x2e, 0xd2, 0x5f, 0xf9, 0xaa, 0x32, 0xb3, 0x16, 0xa1, 0x5f, 0x91, 0x7b, 0x58, 0x73,
	0x99, 0xe4, 0xb8, 0xd2, 0xc0, 0xae, 0xe4, 0x51, 0x98, 0x24, 0x7b, 0x4d, 0x3f, 0x4b, 0x9d, 0x14,
	0xa2, 0xf2, 0xb4, 0x36, 0x82, 0xbb, 0x8f, 0x6a, 0xdf, 0x41, 0x37, 0x6b, 0x62, 0xf0, 0x8e, 0xa8,
	0xf1, 0x1d, 0x74, 0xb3, 0x35, 0x80, 0xb7, 0x79, 0xa5, 0xa4, 0x9e, 0xe6, 0x95, 0x99, 0x55, 0x51,
	0xf8, 0x3d, 0x39, 0x68, 0xa6, 0x27, 0x29, 0xdf, 0x7f, 0x60, 0xf6, 0x9a, 0xa9, 0xa9, 0x66, 0xef,
	0x39, 0x19, 0xba, 0xd5, 0x1f, 0x9a, 0xbf, 0xe3, 0x7f, 0x03, 0xd2, 0x01, 0x39, 0x7d, 0x43, 0x76,
	0xdb, 0xbf, 0xd9, 0xf4, 0x31, 0x08, 0x37, 0xfc, 0x59, 0x18, 0x3d, 0xbc, 0x9d, 0x50, 0xe9, 0x2a,
	0xbc, 0x43, 0x4f, 0x49, 0xaf, 0xde, 0x9c, 0x1e, 0x80, 0xc8, 0xfb, 0x80, 0x46, 0x43, 0x17, 0x62,
	0xd5, 0x1b, 0xb2, 0xdb, 0x3e, 0x32, 0xee, 0xbb, 0xe1, 0x0a, 0x70, 0xdf, 0x5b, 0xdd, 0x85, 0x77,
	0xe6, 0xdb, 0xf6, 0xbf, 0xcd, 0xc9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0x5f, 0xb7, 0xa0,
	0xfa, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TermClient is the client API for Term service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TermClient interface {
	ListTermInfo(ctx context.Context, in *ListTermInfoRequest, opts ...grpc.CallOption) (*ListTermInfoReply, error)
	SaveTerm(ctx context.Context, in *SaveTermRequest, opts ...grpc.CallOption) (*SaveTermReply, error)
	SaveTermRisk(ctx context.Context, in *SaveTermRiskRequest, opts ...grpc.CallOption) (*SaveTermRiskReply, error)
}

type termClient struct {
	cc *grpc.ClientConn
}

func NewTermClient(cc *grpc.ClientConn) TermClient {
	return &termClient{cc}
}

func (c *termClient) ListTermInfo(ctx context.Context, in *ListTermInfoRequest, opts ...grpc.CallOption) (*ListTermInfoReply, error) {
	out := new(ListTermInfoReply)
	err := c.cc.Invoke(ctx, "/pb.Term/ListTermInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termClient) SaveTerm(ctx context.Context, in *SaveTermRequest, opts ...grpc.CallOption) (*SaveTermReply, error) {
	out := new(SaveTermReply)
	err := c.cc.Invoke(ctx, "/pb.Term/SaveTerm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termClient) SaveTermRisk(ctx context.Context, in *SaveTermRiskRequest, opts ...grpc.CallOption) (*SaveTermRiskReply, error) {
	out := new(SaveTermRiskReply)
	err := c.cc.Invoke(ctx, "/pb.Term/SaveTermRisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TermServer is the server API for Term service.
type TermServer interface {
	ListTermInfo(context.Context, *ListTermInfoRequest) (*ListTermInfoReply, error)
	SaveTerm(context.Context, *SaveTermRequest) (*SaveTermReply, error)
	SaveTermRisk(context.Context, *SaveTermRiskRequest) (*SaveTermRiskReply, error)
}

// UnimplementedTermServer can be embedded to have forward compatible implementations.
type UnimplementedTermServer struct {
}

func (*UnimplementedTermServer) ListTermInfo(ctx context.Context, req *ListTermInfoRequest) (*ListTermInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTermInfo not implemented")
}
func (*UnimplementedTermServer) SaveTerm(ctx context.Context, req *SaveTermRequest) (*SaveTermReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTerm not implemented")
}
func (*UnimplementedTermServer) SaveTermRisk(ctx context.Context, req *SaveTermRiskRequest) (*SaveTermRiskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTermRisk not implemented")
}

func RegisterTermServer(s *grpc.Server, srv TermServer) {
	s.RegisterService(&_Term_serviceDesc, srv)
}

func _Term_ListTermInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTermInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServer).ListTermInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Term/ListTermInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServer).ListTermInfo(ctx, req.(*ListTermInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Term_SaveTerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServer).SaveTerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Term/SaveTerm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServer).SaveTerm(ctx, req.(*SaveTermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Term_SaveTermRisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTermRiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServer).SaveTermRisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Term/SaveTermRisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServer).SaveTermRisk(ctx, req.(*SaveTermRiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Term_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Term",
	HandlerType: (*TermServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTermInfo",
			Handler:    _Term_ListTermInfo_Handler,
		},
		{
			MethodName: "SaveTerm",
			Handler:    _Term_SaveTerm_Handler,
		},
		{
			MethodName: "SaveTermRisk",
			Handler:    _Term_SaveTermRisk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "term.proto",
}
