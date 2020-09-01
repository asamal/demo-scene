// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SensorReading.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SensorReading struct {
	Device               *SensorReading_Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	DateTime             int64                 `protobuf:"varint,2,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	Reading              float64               `protobuf:"fixed64,3,opt,name=reading,proto3" json:"reading,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SensorReading) Reset()         { *m = SensorReading{} }
func (m *SensorReading) String() string { return proto.CompactTextString(m) }
func (*SensorReading) ProtoMessage()    {}
func (*SensorReading) Descriptor() ([]byte, []int) {
	return fileDescriptor_43041f5353350540, []int{0}
}

func (m *SensorReading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorReading.Unmarshal(m, b)
}
func (m *SensorReading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorReading.Marshal(b, m, deterministic)
}
func (m *SensorReading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorReading.Merge(m, src)
}
func (m *SensorReading) XXX_Size() int {
	return xxx_messageInfo_SensorReading.Size(m)
}
func (m *SensorReading) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorReading.DiscardUnknown(m)
}

var xxx_messageInfo_SensorReading proto.InternalMessageInfo

func (m *SensorReading) GetDevice() *SensorReading_Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *SensorReading) GetDateTime() int64 {
	if m != nil {
		return m.DateTime
	}
	return 0
}

func (m *SensorReading) GetReading() float64 {
	if m != nil {
		return m.Reading
	}
	return 0
}

type SensorReading_Device struct {
	DeviceID             string   `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SensorReading_Device) Reset()         { *m = SensorReading_Device{} }
func (m *SensorReading_Device) String() string { return proto.CompactTextString(m) }
func (*SensorReading_Device) ProtoMessage()    {}
func (*SensorReading_Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_43041f5353350540, []int{0, 0}
}

func (m *SensorReading_Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorReading_Device.Unmarshal(m, b)
}
func (m *SensorReading_Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorReading_Device.Marshal(b, m, deterministic)
}
func (m *SensorReading_Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorReading_Device.Merge(m, src)
}
func (m *SensorReading_Device) XXX_Size() int {
	return xxx_messageInfo_SensorReading_Device.Size(m)
}
func (m *SensorReading_Device) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorReading_Device.DiscardUnknown(m)
}

var xxx_messageInfo_SensorReading_Device proto.InternalMessageInfo

func (m *SensorReading_Device) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *SensorReading_Device) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*SensorReading)(nil), "io.confluent.cloud.demo.domain.SensorReading")
	proto.RegisterType((*SensorReading_Device)(nil), "io.confluent.cloud.demo.domain.SensorReading.Device")
}

func init() { proto.RegisterFile("SensorReading.proto", fileDescriptor_43041f5353350540) }

var fileDescriptor_43041f5353350540 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x0f, 0x82, 0x30,
	0x10, 0x46, 0x53, 0x31, 0x15, 0x6b, 0x5c, 0xea, 0xd2, 0x30, 0x18, 0xe2, 0xc4, 0xd4, 0x41, 0x9d,
	0x1d, 0x0c, 0x8b, 0x89, 0x53, 0x75, 0x72, 0x2b, 0xf4, 0x34, 0x4d, 0xa0, 0x67, 0x10, 0xfc, 0xa3,
	0xfe, 0x21, 0x43, 0x15, 0x0c, 0x8b, 0x5b, 0xdf, 0xf0, 0x5e, 0xbf, 0x1c, 0x5b, 0x9c, 0xc0, 0x3d,
	0xb0, 0x52, 0xa0, 0x8d, 0x75, 0x37, 0x79, 0xaf, 0xb0, 0x46, 0xbe, 0xb4, 0x28, 0x73, 0x74, 0xd7,
	0xa2, 0x01, 0x57, 0xcb, 0xbc, 0xc0, 0xc6, 0x48, 0x03, 0x25, 0x4a, 0x83, 0xa5, 0xb6, 0x6e, 0xf5,
	0x22, 0x6c, 0x3e, 0xf0, 0xf8, 0x91, 0x51, 0x03, 0x4f, 0x9b, 0x83, 0x20, 0x31, 0x49, 0x66, 0xeb,
	0xad, 0xfc, 0x9f, 0x90, 0xc3, 0x6f, 0x53, 0xef, 0xaa, 0x6f, 0x83, 0x47, 0x2c, 0x34, 0xba, 0x86,
	0xb3, 0x2d, 0x41, 0x8c, 0x62, 0x92, 0x04, 0xaa, 0x67, 0x2e, 0xd8, 0xa4, 0xfa, 0x58, 0x22, 0x88,
	0x49, 0x42, 0x54, 0x87, 0xd1, 0x8e, 0xd1, 0xf4, 0xe7, 0xfb, 0xd7, 0x21, 0xf5, 0x7b, 0xa6, 0xaa,
	0xe7, 0xd6, 0x07, 0xa7, 0xb3, 0x02, 0x8c, 0x4f, 0x87, 0xaa, 0xc3, 0x3d, 0xbd, 0x8c, 0xdb, 0x69,
	0x19, 0xf5, 0x47, 0xd8, 0xbc, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x3b, 0x80, 0x70, 0x1b, 0x01,
	0x00, 0x00,
}