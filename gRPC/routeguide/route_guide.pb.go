// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: routeguide/route_guide.proto

package gRPC

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeacherID int32      `protobuf:"varint,3,opt,name=teacherID,proto3" json:"teacherID,omitempty"`
	Ects      int32      `protobuf:"varint,4,opt,name=ects,proto3" json:"ects,omitempty"`
	Rating    *int32     `protobuf:"varint,5,opt,name=rating,proto3,oneof" json:"rating,omitempty"`
	Students  []*Student `protobuf:"bytes,6,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{1}
}

func (x *Course) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetTeacherID() int32 {
	if x != nil {
		return x.TeacherID
	}
	return 0
}

func (x *Course) GetEcts() int32 {
	if x != nil {
		return x.Ects
	}
	return 0
}

func (x *Course) GetRating() int32 {
	if x != nil && x.Rating != nil {
		return *x.Rating
	}
	return 0
}

func (x *Course) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type CourseID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CourseID) Reset() {
	*x = CourseID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseID) ProtoMessage() {}

func (x *CourseID) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseID.ProtoReflect.Descriptor instead.
func (*CourseID) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{2}
}

func (x *CourseID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CourseRatingChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Rating int32 `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *CourseRatingChangeRequest) Reset() {
	*x = CourseRatingChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseRatingChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseRatingChangeRequest) ProtoMessage() {}

func (x *CourseRatingChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseRatingChangeRequest.ProtoReflect.Descriptor instead.
func (*CourseRatingChangeRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{3}
}

func (x *CourseRatingChangeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CourseRatingChangeRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type CourseEnrollStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Student int32 `protobuf:"varint,2,opt,name=Student,proto3" json:"Student,omitempty"`
}

func (x *CourseEnrollStudentRequest) Reset() {
	*x = CourseEnrollStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseEnrollStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseEnrollStudentRequest) ProtoMessage() {}

func (x *CourseEnrollStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseEnrollStudentRequest.ProtoReflect.Descriptor instead.
func (*CourseEnrollStudentRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{4}
}

func (x *CourseEnrollStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CourseEnrollStudentRequest) GetStudent() int32 {
	if x != nil {
		return x.Student
	}
	return 0
}

type CourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Messsage *string `protobuf:"bytes,2,opt,name=messsage,proto3,oneof" json:"messsage,omitempty"`
}

func (x *CourseResponse) Reset() {
	*x = CourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseResponse) ProtoMessage() {}

func (x *CourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseResponse.ProtoReflect.Descriptor instead.
func (*CourseResponse) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{5}
}

func (x *CourseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CourseResponse) GetMesssage() string {
	if x != nil && x.Messsage != nil {
		return *x.Messsage
	}
	return ""
}

type CourseRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Course  *Course `protobuf:"bytes,2,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *CourseRequestResponse) Reset() {
	*x = CourseRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseRequestResponse) ProtoMessage() {}

func (x *CourseRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseRequestResponse.ProtoReflect.Descriptor instead.
func (*CourseRequestResponse) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{6}
}

func (x *CourseRequestResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CourseRequestResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

var File_routeguide_route_guide_proto protoreflect.FileDescriptor

var file_routeguide_route_guide_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d,
	0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x01,
	0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1b,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x08, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x19, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x46, 0x0a,
	0x1a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x52, 0x0a, 0x15, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x32, 0xce, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x12, 0x07, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0f, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x07,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x09, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x44, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x09, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1a, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x15,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1a, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x45, 0x58, 0x4b, 0x72, 0x61, 0x73, 0x68, 0x2f, 0x44, 0x49, 0x53, 0x59,
	0x53, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x52, 0x45, 0x53, 0x54, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routeguide_route_guide_proto_rawDescOnce sync.Once
	file_routeguide_route_guide_proto_rawDescData = file_routeguide_route_guide_proto_rawDesc
)

func file_routeguide_route_guide_proto_rawDescGZIP() []byte {
	file_routeguide_route_guide_proto_rawDescOnce.Do(func() {
		file_routeguide_route_guide_proto_rawDescData = protoimpl.X.CompressGZIP(file_routeguide_route_guide_proto_rawDescData)
	})
	return file_routeguide_route_guide_proto_rawDescData
}

var file_routeguide_route_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_routeguide_route_guide_proto_goTypes = []interface{}{
	(*Student)(nil),                    // 0: Student
	(*Course)(nil),                     // 1: Course
	(*CourseID)(nil),                   // 2: CourseID
	(*CourseRatingChangeRequest)(nil),  // 3: CourseRatingChangeRequest
	(*CourseEnrollStudentRequest)(nil), // 4: CourseEnrollStudentRequest
	(*CourseResponse)(nil),             // 5: CourseResponse
	(*CourseRequestResponse)(nil),      // 6: CourseRequestResponse
}
var file_routeguide_route_guide_proto_depIdxs = []int32{
	0, // 0: Course.students:type_name -> Student
	1, // 1: CourseRequestResponse.course:type_name -> Course
	1, // 2: CourseService.SetCourse:input_type -> Course
	1, // 3: CourseService.CreateCourse:input_type -> Course
	2, // 4: CourseService.GetCourse:input_type -> CourseID
	2, // 5: CourseService.DeleteCourse:input_type -> CourseID
	3, // 6: CourseService.SetRatingCourse:input_type -> CourseRatingChangeRequest
	3, // 7: CourseService.EnrollStudentInCourse:input_type -> CourseRatingChangeRequest
	5, // 8: CourseService.SetCourse:output_type -> CourseResponse
	5, // 9: CourseService.CreateCourse:output_type -> CourseResponse
	6, // 10: CourseService.GetCourse:output_type -> CourseRequestResponse
	5, // 11: CourseService.DeleteCourse:output_type -> CourseResponse
	5, // 12: CourseService.SetRatingCourse:output_type -> CourseResponse
	5, // 13: CourseService.EnrollStudentInCourse:output_type -> CourseResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_routeguide_route_guide_proto_init() }
func file_routeguide_route_guide_proto_init() {
	if File_routeguide_route_guide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routeguide_route_guide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routeguide_route_guide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routeguide_route_guide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routeguide_route_guide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseRatingChangeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routeguide_route_guide_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseEnrollStudentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routeguide_route_guide_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routeguide_route_guide_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseRequestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_routeguide_route_guide_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_routeguide_route_guide_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_routeguide_route_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routeguide_route_guide_proto_goTypes,
		DependencyIndexes: file_routeguide_route_guide_proto_depIdxs,
		MessageInfos:      file_routeguide_route_guide_proto_msgTypes,
	}.Build()
	File_routeguide_route_guide_proto = out.File
	file_routeguide_route_guide_proto_rawDesc = nil
	file_routeguide_route_guide_proto_goTypes = nil
	file_routeguide_route_guide_proto_depIdxs = nil
}
