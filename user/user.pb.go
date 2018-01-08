// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
	Response
	ID
*/
package user

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string         `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string         `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Email    string         `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Fname    string         `protobuf:"bytes,5,opt,name=fname" json:"fname,omitempty"`
	Mname    string         `protobuf:"bytes,6,opt,name=mname" json:"mname,omitempty"`
	Lname    string         `protobuf:"bytes,7,opt,name=lname" json:"lname,omitempty"`
	Location *User_Location `protobuf:"bytes,8,opt,name=location" json:"location,omitempty"`
}

// Value implements Valueer interface to marshal object into []byte type before
// storing into DB.
func (u User) Value() (driver.Value, error) {
	j, err := json.Marshal(u)
	return j, err
}

// Scan implements Scanner interface to Unmarshal return []byte array from DB into User
func (u *User) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	err := json.Unmarshal(source, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFname() string {
	if m != nil {
		return m.Fname
	}
	return ""
}

func (m *User) GetMname() string {
	if m != nil {
		return m.Mname
	}
	return ""
}

func (m *User) GetLname() string {
	if m != nil {
		return m.Lname
	}
	return ""
}

func (m *User) GetLocation() *User_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type User_Location struct {
	State       string  `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	City        string  `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	Zipcode     int32   `protobuf:"varint,3,opt,name=zipcode" json:"zipcode,omitempty"`
	Phonenumber string  `protobuf:"bytes,4,opt,name=phonenumber" json:"phonenumber,omitempty"`
	Lat         float32 `protobuf:"fixed32,5,opt,name=lat" json:"lat,omitempty"`
	Long        float32 `protobuf:"fixed32,6,opt,name=long" json:"long,omitempty"`
}

func (m *User_Location) Reset()                    { *m = User_Location{} }
func (m *User_Location) String() string            { return proto.CompactTextString(m) }
func (*User_Location) ProtoMessage()               {}
func (*User_Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *User_Location) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *User_Location) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *User_Location) GetZipcode() int32 {
	if m != nil {
		return m.Zipcode
	}
	return 0
}

func (m *User_Location) GetPhonenumber() string {
	if m != nil {
		return m.Phonenumber
	}
	return ""
}

func (m *User_Location) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *User_Location) GetLong() float32 {
	if m != nil {
		return m.Long
	}
	return 0
}

type Response struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*User_Location)(nil), "user.User.Location")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*ID)(nil), "user.ID")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6a, 0xe3, 0x30,
	0x14, 0x86, 0x63, 0xc5, 0x89, 0x35, 0x2f, 0x10, 0x06, 0x4d, 0x16, 0x22, 0xab, 0xe0, 0x61, 0x20,
	0x8b, 0x21, 0x85, 0xb4, 0x17, 0x68, 0x6b, 0x08, 0x81, 0x2e, 0x8a, 0xa0, 0x07, 0x70, 0xec, 0xd7,
	0xd6, 0x60, 0x5b, 0xc6, 0x52, 0x28, 0xe9, 0xba, 0x67, 0xe8, 0x39, 0x7b, 0x84, 0xa2, 0x27, 0x3b,
	0x0d, 0x69, 0x77, 0xff, 0xff, 0xbd, 0x1f, 0x59, 0xff, 0xb3, 0x00, 0xf6, 0x06, 0xdb, 0x55, 0xd3,
	0x6a, 0xab, 0x45, 0xe8, 0x74, 0xfc, 0xc1, 0x20, 0x7c, 0x30, 0xd8, 0x8a, 0x29, 0xb0, 0x22, 0x97,
	0xc1, 0x22, 0x58, 0xfe, 0x52, 0xac, 0xc8, 0xc5, 0x1c, 0xb8, 0x0b, 0xd4, 0x69, 0x85, 0x92, 0x11,
	0x3d, 0x7a, 0x37, 0x6b, 0x52, 0x63, 0x5e, 0x74, 0x9b, 0xcb, 0xa1, 0x9f, 0xf5, 0x5e, 0xcc, 0x60,
	0x84, 0x55, 0x5a, 0x94, 0x32, 0xa4, 0x81, 0x37, 0x8e, 0x3e, 0xd2, 0x51, 0x23, 0x4f, 0xc9, 0x38,
	0x5a, 0x11, 0x1d, 0x7b, 0x5a, 0xf5, 0xb4, 0x24, 0x1a, 0x79, 0x4a, 0x46, 0x5c, 0x00, 0x2f, 0x75,
	0x96, 0xda, 0x42, 0xd7, 0x92, 0x2f, 0x82, 0xe5, 0x64, 0xfd, 0x67, 0x45, 0x6d, 0xdc, 0xed, 0x57,
	0x77, 0xdd, 0x48, 0x1d, 0x43, 0xf3, 0xf7, 0x00, 0x78, 0x8f, 0xdd, 0x99, 0xc6, 0xa6, 0x16, 0xbb,
	0x82, 0xde, 0x08, 0x01, 0x61, 0x56, 0xd8, 0x43, 0xd7, 0x8f, 0xb4, 0x90, 0x10, 0xbd, 0x16, 0x4d,
	0xa6, 0x73, 0xa4, 0x6a, 0x23, 0xd5, 0x5b, 0xb1, 0x80, 0x49, 0xf3, 0xac, 0x6b, 0xac, 0xf7, 0xd5,
	0x0e, 0xdb, 0xae, 0xdf, 0x29, 0x12, 0xbf, 0x61, 0x58, 0xa6, 0x96, 0x3a, 0x32, 0xe5, 0xa4, 0xfb,
	0x42, 0xa9, 0xeb, 0x27, 0x2a, 0xc8, 0x14, 0xe9, 0xf8, 0x0a, 0xb8, 0x42, 0xd3, 0xe8, 0xda, 0xe0,
	0xb7, 0xad, 0x4b, 0x88, 0xcc, 0x3e, 0xcb, 0xd0, 0x18, 0xba, 0x14, 0x57, 0xbd, 0x8d, 0x67, 0xc0,
	0xb6, 0xc9, 0x79, 0x7e, 0xfd, 0x16, 0x00, 0x77, 0x0b, 0xd8, 0xa8, 0xfb, 0x5b, 0xf1, 0x17, 0xa2,
	0xeb, 0x3c, 0xa7, 0xbf, 0x09, 0x5f, 0xbb, 0x99, 0x9f, 0xe8, 0x78, 0x20, 0xfe, 0xc1, 0x64, 0x83,
	0xd6, 0x99, 0x9b, 0xc3, 0x36, 0x11, 0xdc, 0x0f, 0xb7, 0xc9, 0x59, 0xec, 0x3f, 0x4c, 0x13, 0x2c,
	0xd1, 0xe2, 0x0f, 0xc9, 0xa9, 0x57, 0x7d, 0x89, 0x78, 0xb0, 0x1b, 0xd3, 0x93, 0xba, 0xfc, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0x8f, 0xbd, 0x06, 0x60, 0x02, 0x00, 0x00,
}
