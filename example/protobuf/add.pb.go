// Code generated by protoc-gen-go.
// source: add.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
    add.proto

It has these top-level messages:
    AddReq
    AddResp
*/
package protobuf

import proto1 "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type AddReq struct {
	A int32 `protobuf:"varint,1,opt" json:"A,omitempty"`
	B int32 `protobuf:"varint,2,opt" json:"B,omitempty"`
}

func (m *AddReq) Reset()         { *m = AddReq{} }
func (m *AddReq) String() string { return proto1.CompactTextString(m) }
func (*AddReq) ProtoMessage()    {}

type AddResp struct {
	C int32 `protobuf:"varint,1,opt" json:"C,omitempty"`
}

func (m *AddResp) Reset()         { *m = AddResp{} }
func (m *AddResp) String() string { return proto1.CompactTextString(m) }
func (*AddResp) ProtoMessage()    {}

func init() {
}
