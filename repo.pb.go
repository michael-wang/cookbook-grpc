// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: repo.proto

package cookbook_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Repo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FullName        string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Private         bool                   `protobuf:"varint,4,opt,name=private,proto3" json:"private,omitempty"`
	Owner           *User                  `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Description     string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Fork            bool                   `protobuf:"varint,7,opt,name=fork,proto3" json:"fork,omitempty"`
	Url             string                 `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	StargazersCount uint32                 `protobuf:"varint,11,opt,name=stargazers_count,json=stargazersCount,proto3" json:"stargazers_count,omitempty"`
	ForksCount      uint32                 `protobuf:"varint,12,opt,name=forks_count,json=forksCount,proto3" json:"forks_count,omitempty"`
}

func (x *Repo) Reset() {
	*x = Repo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repo) ProtoMessage() {}

func (x *Repo) ProtoReflect() protoreflect.Message {
	mi := &file_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repo.ProtoReflect.Descriptor instead.
func (*Repo) Descriptor() ([]byte, []int) {
	return file_repo_proto_rawDescGZIP(), []int{0}
}

func (x *Repo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Repo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repo) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Repo) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *Repo) GetOwner() *User {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Repo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Repo) GetFork() bool {
	if x != nil {
		return x.Fork
	}
	return false
}

func (x *Repo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Repo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Repo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Repo) GetStargazersCount() uint32 {
	if x != nil {
		return x.StargazersCount
	}
	return 0
}

func (x *Repo) GetForksCount() uint32 {
	if x != nil {
		return x.ForksCount
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login      string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	AvatarUrl  string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Url        string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	StarredUrl string `protobuf:"bytes,5,opt,name=starred_url,json=starredUrl,proto3" json:"starred_url,omitempty"`
	ReposUrl   string `protobuf:"bytes,6,opt,name=repos_url,json=reposUrl,proto3" json:"repos_url,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_repo_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *User) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *User) GetStarredUrl() string {
	if x != nil {
		return x.StarredUrl
	}
	return ""
}

func (x *User) GetReposUrl() string {
	if x != nil {
		return x.ReposUrl
	}
	return ""
}

var File_repo_proto protoreflect.FileDescriptor

var file_repo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x04, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x6f, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x67,
	0x61, 0x7a, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x67, 0x61, 0x7a, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6b, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6b, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x72, 0x65,
	0x64, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x55, 0x72,
	0x6c, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x2e, 0x77, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6f,
	0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x63, 0x6f, 0x6f, 0x6b, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repo_proto_rawDescOnce sync.Once
	file_repo_proto_rawDescData = file_repo_proto_rawDesc
)

func file_repo_proto_rawDescGZIP() []byte {
	file_repo_proto_rawDescOnce.Do(func() {
		file_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_repo_proto_rawDescData)
	})
	return file_repo_proto_rawDescData
}

var file_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_repo_proto_goTypes = []interface{}{
	(*Repo)(nil),                  // 0: repo.Repo
	(*User)(nil),                  // 1: repo.User
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_repo_proto_depIdxs = []int32{
	1, // 0: repo.Repo.owner:type_name -> repo.User
	2, // 1: repo.Repo.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: repo.Repo.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_repo_proto_init() }
func file_repo_proto_init() {
	if File_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repo); i {
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
		file_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_repo_proto_goTypes,
		DependencyIndexes: file_repo_proto_depIdxs,
		MessageInfos:      file_repo_proto_msgTypes,
	}.Build()
	File_repo_proto = out.File
	file_repo_proto_rawDesc = nil
	file_repo_proto_goTypes = nil
	file_repo_proto_depIdxs = nil
}
