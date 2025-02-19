// Code generated by protoc-gen-go. DO NOT EDIT.
// source: segcore.proto

package segcorepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	schemapb "github.com/milvus-io/milvus/api/schemapb"
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

type RetrieveResults struct {
	Ids                  *schemapb.IDs         `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	Offset               []int64               `protobuf:"varint,2,rep,packed,name=offset,proto3" json:"offset,omitempty"`
	FieldsData           []*schemapb.FieldData `protobuf:"bytes,3,rep,name=fields_data,json=fieldsData,proto3" json:"fields_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RetrieveResults) Reset()         { *m = RetrieveResults{} }
func (m *RetrieveResults) String() string { return proto.CompactTextString(m) }
func (*RetrieveResults) ProtoMessage()    {}
func (*RetrieveResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{0}
}

func (m *RetrieveResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveResults.Unmarshal(m, b)
}
func (m *RetrieveResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveResults.Marshal(b, m, deterministic)
}
func (m *RetrieveResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveResults.Merge(m, src)
}
func (m *RetrieveResults) XXX_Size() int {
	return xxx_messageInfo_RetrieveResults.Size(m)
}
func (m *RetrieveResults) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveResults.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveResults proto.InternalMessageInfo

func (m *RetrieveResults) GetIds() *schemapb.IDs {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *RetrieveResults) GetOffset() []int64 {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *RetrieveResults) GetFieldsData() []*schemapb.FieldData {
	if m != nil {
		return m.FieldsData
	}
	return nil
}

type LoadFieldMeta struct {
	MinTimestamp         int64    `protobuf:"varint,1,opt,name=min_timestamp,json=minTimestamp,proto3" json:"min_timestamp,omitempty"`
	MaxTimestamp         int64    `protobuf:"varint,2,opt,name=max_timestamp,json=maxTimestamp,proto3" json:"max_timestamp,omitempty"`
	RowCount             int64    `protobuf:"varint,3,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadFieldMeta) Reset()         { *m = LoadFieldMeta{} }
func (m *LoadFieldMeta) String() string { return proto.CompactTextString(m) }
func (*LoadFieldMeta) ProtoMessage()    {}
func (*LoadFieldMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{1}
}

func (m *LoadFieldMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadFieldMeta.Unmarshal(m, b)
}
func (m *LoadFieldMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadFieldMeta.Marshal(b, m, deterministic)
}
func (m *LoadFieldMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadFieldMeta.Merge(m, src)
}
func (m *LoadFieldMeta) XXX_Size() int {
	return xxx_messageInfo_LoadFieldMeta.Size(m)
}
func (m *LoadFieldMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadFieldMeta.DiscardUnknown(m)
}

var xxx_messageInfo_LoadFieldMeta proto.InternalMessageInfo

func (m *LoadFieldMeta) GetMinTimestamp() int64 {
	if m != nil {
		return m.MinTimestamp
	}
	return 0
}

func (m *LoadFieldMeta) GetMaxTimestamp() int64 {
	if m != nil {
		return m.MaxTimestamp
	}
	return 0
}

func (m *LoadFieldMeta) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

type LoadSegmentMeta struct {
	// TODOs
	Metas                []*LoadFieldMeta `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
	TotalSize            int64            `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LoadSegmentMeta) Reset()         { *m = LoadSegmentMeta{} }
func (m *LoadSegmentMeta) String() string { return proto.CompactTextString(m) }
func (*LoadSegmentMeta) ProtoMessage()    {}
func (*LoadSegmentMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{2}
}

func (m *LoadSegmentMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadSegmentMeta.Unmarshal(m, b)
}
func (m *LoadSegmentMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadSegmentMeta.Marshal(b, m, deterministic)
}
func (m *LoadSegmentMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadSegmentMeta.Merge(m, src)
}
func (m *LoadSegmentMeta) XXX_Size() int {
	return xxx_messageInfo_LoadSegmentMeta.Size(m)
}
func (m *LoadSegmentMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadSegmentMeta.DiscardUnknown(m)
}

var xxx_messageInfo_LoadSegmentMeta proto.InternalMessageInfo

func (m *LoadSegmentMeta) GetMetas() []*LoadFieldMeta {
	if m != nil {
		return m.Metas
	}
	return nil
}

func (m *LoadSegmentMeta) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

type InsertRecord struct {
	FieldsData           []*schemapb.FieldData `protobuf:"bytes,1,rep,name=fields_data,json=fieldsData,proto3" json:"fields_data,omitempty"`
	NumRows              int64                 `protobuf:"varint,2,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InsertRecord) Reset()         { *m = InsertRecord{} }
func (m *InsertRecord) String() string { return proto.CompactTextString(m) }
func (*InsertRecord) ProtoMessage()    {}
func (*InsertRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d79fce784797357, []int{3}
}

func (m *InsertRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertRecord.Unmarshal(m, b)
}
func (m *InsertRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertRecord.Marshal(b, m, deterministic)
}
func (m *InsertRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertRecord.Merge(m, src)
}
func (m *InsertRecord) XXX_Size() int {
	return xxx_messageInfo_InsertRecord.Size(m)
}
func (m *InsertRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertRecord.DiscardUnknown(m)
}

var xxx_messageInfo_InsertRecord proto.InternalMessageInfo

func (m *InsertRecord) GetFieldsData() []*schemapb.FieldData {
	if m != nil {
		return m.FieldsData
	}
	return nil
}

func (m *InsertRecord) GetNumRows() int64 {
	if m != nil {
		return m.NumRows
	}
	return 0
}

func init() {
	proto.RegisterType((*RetrieveResults)(nil), "milvus.proto.segcore.RetrieveResults")
	proto.RegisterType((*LoadFieldMeta)(nil), "milvus.proto.segcore.LoadFieldMeta")
	proto.RegisterType((*LoadSegmentMeta)(nil), "milvus.proto.segcore.LoadSegmentMeta")
	proto.RegisterType((*InsertRecord)(nil), "milvus.proto.segcore.InsertRecord")
}

func init() { proto.RegisterFile("segcore.proto", fileDescriptor_1d79fce784797357) }

var fileDescriptor_1d79fce784797357 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4b, 0xeb, 0x40,
	0x10, 0xc6, 0x49, 0xc3, 0xeb, 0x6b, 0xb7, 0x2d, 0x85, 0xf0, 0x78, 0x44, 0x45, 0x29, 0xe9, 0xa5,
	0x08, 0x26, 0x50, 0x45, 0xf0, 0x24, 0x68, 0x11, 0x0a, 0x7a, 0xd9, 0x7a, 0xf2, 0x12, 0xb6, 0xc9,
	0xb4, 0x5d, 0xcd, 0xee, 0x96, 0xdd, 0x49, 0x53, 0xfa, 0x87, 0xf8, 0xf7, 0x4a, 0x36, 0x11, 0xad,
	0xf4, 0xe2, 0x6d, 0xe7, 0xdb, 0xdf, 0xcc, 0x37, 0xf3, 0x91, 0x9e, 0x81, 0x65, 0xa2, 0x34, 0x84,
	0x6b, 0xad, 0x50, 0x79, 0xff, 0x04, 0xcf, 0x36, 0xb9, 0xa9, 0xaa, 0xb0, 0xfe, 0x3b, 0xee, 0x9a,
	0x64, 0x05, 0x82, 0x55, 0x6a, 0xf0, 0xee, 0x90, 0x3e, 0x05, 0xd4, 0x1c, 0x36, 0x40, 0xc1, 0xe4,
	0x19, 0x1a, 0xef, 0x9c, 0xb8, 0x3c, 0x35, 0xbe, 0x33, 0x70, 0x46, 0x9d, 0xb1, 0x1f, 0xee, 0x4f,
	0xa9, 0x9a, 0xa7, 0x13, 0x43, 0x4b, 0xc8, 0xfb, 0x4f, 0x9a, 0x6a, 0xb1, 0x30, 0x80, 0x7e, 0x63,
	0xe0, 0x8e, 0x5c, 0x5a, 0x57, 0xde, 0x2d, 0xe9, 0x2c, 0x38, 0x64, 0xa9, 0x89, 0x53, 0x86, 0xcc,
	0x77, 0x07, 0xee, 0xa8, 0x33, 0x3e, 0x3b, 0x38, 0xeb, 0xa1, 0xe4, 0x26, 0x0c, 0x19, 0x25, 0x55,
	0x4b, 0xf9, 0x0e, 0x36, 0xa4, 0xf7, 0xa8, 0x58, 0x6a, 0x3f, 0x9f, 0x00, 0x99, 0x37, 0x24, 0x3d,
	0xc1, 0x65, 0x8c, 0x5c, 0x80, 0x41, 0x26, 0xd6, 0x76, 0x3f, 0x97, 0x76, 0x05, 0x97, 0xcf, 0x9f,
	0x9a, 0x85, 0xd8, 0xf6, 0x1b, 0xd4, 0xa8, 0x21, 0xb6, 0xfd, 0x82, 0x4e, 0x48, 0x5b, 0xab, 0x22,
	0x4e, 0x54, 0x2e, 0xd1, 0x77, 0x2d, 0xd0, 0xd2, 0xaa, 0xb8, 0x2f, 0xeb, 0xe0, 0x8d, 0xf4, 0x4b,
	0xdf, 0x19, 0x2c, 0x05, 0x48, 0xb4, 0xce, 0x37, 0xe4, 0x8f, 0x00, 0x64, 0x65, 0x22, 0xe5, 0x15,
	0xc3, 0xf0, 0x50, 0xae, 0xe1, 0xde, 0xb6, 0xb4, 0xea, 0xf0, 0x4e, 0x09, 0x41, 0x85, 0x2c, 0x8b,
	0x0d, 0xdf, 0x41, 0xbd, 0x4c, 0xdb, 0x2a, 0x33, 0xbe, 0x83, 0xe0, 0x95, 0x74, 0xa7, 0xd2, 0x80,
	0x46, 0x0a, 0x89, 0xd2, 0xe9, 0xcf, 0xd4, 0x9c, 0xdf, 0xa6, 0xe6, 0x1d, 0x91, 0x96, 0xcc, 0x45,
	0xac, 0x55, 0x61, 0x6a, 0xb7, 0xbf, 0x32, 0x17, 0x54, 0x15, 0xe6, 0xee, 0xfa, 0xe5, 0x6a, 0xc9,
	0x71, 0x95, 0xcf, 0xc3, 0x44, 0x89, 0xa8, 0x1a, 0x79, 0xc1, 0x55, 0xfd, 0x8a, 0xb8, 0x44, 0xd0,
	0x92, 0x65, 0x91, 0x75, 0x89, 0xea, 0xab, 0xd6, 0xf3, 0x79, 0xd3, 0x0a, 0x97, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd6, 0x48, 0x04, 0x17, 0x5d, 0x02, 0x00, 0x00,
}
