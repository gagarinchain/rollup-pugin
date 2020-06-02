// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rollup.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

type Rollup struct {
	Accounts             [][]byte       `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Rollup) Reset()         { *m = Rollup{} }
func (m *Rollup) String() string { return proto.CompactTextString(m) }
func (*Rollup) ProtoMessage()    {}
func (*Rollup) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{0}
}

func (m *Rollup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rollup.Unmarshal(m, b)
}
func (m *Rollup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rollup.Marshal(b, m, deterministic)
}
func (m *Rollup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rollup.Merge(m, src)
}
func (m *Rollup) XXX_Size() int {
	return xxx_messageInfo_Rollup.Size(m)
}
func (m *Rollup) XXX_DiscardUnknown() {
	xxx_messageInfo_Rollup.DiscardUnknown(m)
}

var xxx_messageInfo_Rollup proto.InternalMessageInfo

func (m *Rollup) GetAccounts() [][]byte {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Rollup) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Transaction struct {
	From                 int32    `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   int32    `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Value                int64    `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Transaction) GetTo() int32 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Transaction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BlockHeader struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash           []byte   `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	QcHash               []byte   `protobuf:"bytes,3,opt,name=qcHash,proto3" json:"qcHash,omitempty"`
	DataHash             []byte   `protobuf:"bytes,4,opt,name=dataHash,proto3" json:"dataHash,omitempty"`
	TxHash               []byte   `protobuf:"bytes,5,opt,name=txHash,proto3" json:"txHash,omitempty"`
	StateHash            []byte   `protobuf:"bytes,6,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height               int32    `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            int64    `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{2}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetQcHash() []byte {
	if m != nil {
		return m.QcHash
	}
	return nil
}

func (m *BlockHeader) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *BlockHeader) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *BlockHeader) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *BlockHeader) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Rollup)(nil), "Rollup")
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*BlockHeader)(nil), "BlockHeader")
}

func init() { proto.RegisterFile("rollup.proto", fileDescriptor_037f188b50610c79) }

var fileDescriptor_037f188b50610c79 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4e, 0xeb, 0x30,
	0x10, 0x45, 0x15, 0xa7, 0xcd, 0xeb, 0x9b, 0x46, 0x2c, 0x2c, 0x84, 0x4c, 0x85, 0x50, 0xd4, 0x55,
	0x56, 0x2d, 0x82, 0x3f, 0xe8, 0x86, 0xae, 0x2d, 0xc4, 0x82, 0xdd, 0x34, 0x75, 0x9b, 0x88, 0x24,
	0x0e, 0xf6, 0x04, 0xc1, 0xe7, 0xf2, 0x27, 0xc8, 0x93, 0xaa, 0x29, 0x3b, 0x9f, 0x7b, 0xe6, 0x8e,
	0x34, 0x32, 0xa4, 0xce, 0xd6, 0x75, 0xdf, 0xad, 0x3a, 0x67, 0xc9, 0x2e, 0x6e, 0x8f, 0xd6, 0x1e,
	0x6b, 0xb3, 0x66, 0xda, 0xf5, 0x87, 0x35, 0xb6, 0xdf, 0x83, 0x5a, 0xbe, 0x42, 0xa2, 0x79, 0x54,
	0x2e, 0x60, 0x86, 0x45, 0x61, 0xfb, 0x96, 0xbc, 0x8a, 0xb2, 0x38, 0x4f, 0xf5, 0x99, 0xe5, 0x03,
	0xa4, 0xe4, 0xb0, 0xf5, 0x58, 0x50, 0x65, 0x5b, 0xaf, 0x44, 0x16, 0xe7, 0xf3, 0xc7, 0x74, 0xf5,
	0x32, 0x86, 0xfa, 0xcf, 0xc4, 0xf2, 0x19, 0xe6, 0x17, 0x52, 0x4a, 0x98, 0x1c, 0x9c, 0x6d, 0x54,
	0x94, 0x45, 0xf9, 0x54, 0xf3, 0x5b, 0x5e, 0x81, 0x20, 0xab, 0x04, 0x27, 0x82, 0xac, 0xbc, 0x86,
	0xe9, 0x27, 0xd6, 0xbd, 0x51, 0x71, 0x16, 0xe5, 0xb1, 0x1e, 0x60, 0xf9, 0x13, 0xc1, 0x7c, 0x53,
	0xdb, 0xe2, 0x7d, 0x6b, 0x70, 0x6f, 0x5c, 0xd8, 0x54, 0xa2, 0x2f, 0x79, 0x53, 0xaa, 0xf9, 0x2d,
	0xef, 0x01, 0x3a, 0x74, 0xa6, 0xa5, 0x6d, 0x30, 0x82, 0xcd, 0x45, 0x22, 0x6f, 0x20, 0xf9, 0x28,
	0xd8, 0xc5, 0xec, 0x4e, 0x14, 0x4e, 0xde, 0x23, 0x21, 0x9b, 0x09, 0x9b, 0x33, 0x87, 0x0e, 0x7d,
	0xb1, 0x99, 0x0e, 0x9d, 0x81, 0xe4, 0x1d, 0xfc, 0xf7, 0x84, 0x64, 0x58, 0x25, 0xac, 0xc6, 0x20,
	0xb4, 0x4a, 0x53, 0x1d, 0x4b, 0x52, 0xff, 0xf8, 0xae, 0x13, 0x85, 0x16, 0x55, 0x8d, 0xf1, 0x84,
	0x4d, 0xa7, 0x66, 0x7c, 0xdf, 0x18, 0x6c, 0x26, 0x6f, 0xa2, 0xdb, 0xed, 0x12, 0xfe, 0x91, 0xa7,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xb2, 0x12, 0xa6, 0xbc, 0x01, 0x00, 0x00,
}