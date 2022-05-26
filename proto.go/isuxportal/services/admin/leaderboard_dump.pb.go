// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: isuxportal/services/admin/leaderboard_dump.proto

package admin

import (
	_ "github.com/golang/protobuf/ptypes/timestamp"
	resources "github.com/isucon/isucon10-portal/proto.go/isuxportal/resources"
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

type GetLeaderboardDumpQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	When string `protobuf:"bytes,1,opt,name=when,proto3" json:"when,omitempty"` // ISO8601 or "qualify-end"
}

func (x *GetLeaderboardDumpQuery) Reset() {
	*x = GetLeaderboardDumpQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaderboardDumpQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaderboardDumpQuery) ProtoMessage() {}

func (x *GetLeaderboardDumpQuery) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaderboardDumpQuery.ProtoReflect.Descriptor instead.
func (*GetLeaderboardDumpQuery) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_leaderboard_dump_proto_rawDescGZIP(), []int{0}
}

func (x *GetLeaderboardDumpQuery) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

type GetLeaderboardDumpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetLeaderboardDumpResponse_LeaderboardDumpItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetLeaderboardDumpResponse) Reset() {
	*x = GetLeaderboardDumpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaderboardDumpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaderboardDumpResponse) ProtoMessage() {}

func (x *GetLeaderboardDumpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaderboardDumpResponse.ProtoReflect.Descriptor instead.
func (*GetLeaderboardDumpResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_leaderboard_dump_proto_rawDescGZIP(), []int{1}
}

func (x *GetLeaderboardDumpResponse) GetItems() []*GetLeaderboardDumpResponse_LeaderboardDumpItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetLeaderboardDumpResponse_LeaderboardDumpItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position    int64                                                   `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Team        *resources.Team                                         `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	BestScore   *resources.Leaderboard_LeaderboardItem_LeaderboardScore `protobuf:"bytes,3,opt,name=best_score,json=bestScore,proto3" json:"best_score,omitempty"`
	LatestScore *resources.Leaderboard_LeaderboardItem_LeaderboardScore `protobuf:"bytes,4,opt,name=latest_score,json=latestScore,proto3" json:"latest_score,omitempty"`
	Target      *resources.ContestantInstance                           `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) Reset() {
	*x = GetLeaderboardDumpResponse_LeaderboardDumpItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaderboardDumpResponse_LeaderboardDumpItem) ProtoMessage() {}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaderboardDumpResponse_LeaderboardDumpItem.ProtoReflect.Descriptor instead.
func (*GetLeaderboardDumpResponse_LeaderboardDumpItem) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_leaderboard_dump_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) GetTeam() *resources.Team {
	if x != nil {
		return x.Team
	}
	return nil
}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) GetBestScore() *resources.Leaderboard_LeaderboardItem_LeaderboardScore {
	if x != nil {
		return x.BestScore
	}
	return nil
}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) GetLatestScore() *resources.Leaderboard_LeaderboardItem_LeaderboardScore {
	if x != nil {
		return x.LatestScore
	}
	return nil
}

func (x *GetLeaderboardDumpResponse_LeaderboardDumpItem) GetTarget() *resources.ContestantInstance {
	if x != nil {
		return x.Target
	}
	return nil
}

var File_isuxportal_services_admin_leaderboard_dump_proto protoreflect.FileDescriptor

var file_isuxportal_services_admin_leaderboard_dump_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x1a, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x69, 0x73, 0x75,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44,
	0x75, 0x6d, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x22, 0x8b, 0x04, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44,
	0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x69, 0x73, 0x75,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0x85, 0x03, 0x0a, 0x13, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x67, 0x0a, 0x0a,
	0x62, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x48, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x09, 0x62, 0x65, 0x73, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x69, 0x73,
	0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f,
	0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x2f, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_services_admin_leaderboard_dump_proto_rawDescOnce sync.Once
	file_isuxportal_services_admin_leaderboard_dump_proto_rawDescData = file_isuxportal_services_admin_leaderboard_dump_proto_rawDesc
)

func file_isuxportal_services_admin_leaderboard_dump_proto_rawDescGZIP() []byte {
	file_isuxportal_services_admin_leaderboard_dump_proto_rawDescOnce.Do(func() {
		file_isuxportal_services_admin_leaderboard_dump_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_services_admin_leaderboard_dump_proto_rawDescData)
	})
	return file_isuxportal_services_admin_leaderboard_dump_proto_rawDescData
}

var file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_isuxportal_services_admin_leaderboard_dump_proto_goTypes = []interface{}{
	(*GetLeaderboardDumpQuery)(nil),                                // 0: isuxportal.proto.services.admin.GetLeaderboardDumpQuery
	(*GetLeaderboardDumpResponse)(nil),                             // 1: isuxportal.proto.services.admin.GetLeaderboardDumpResponse
	(*GetLeaderboardDumpResponse_LeaderboardDumpItem)(nil),         // 2: isuxportal.proto.services.admin.GetLeaderboardDumpResponse.LeaderboardDumpItem
	(*resources.Team)(nil),                                         // 3: isuxportal.proto.resources.Team
	(*resources.Leaderboard_LeaderboardItem_LeaderboardScore)(nil), // 4: isuxportal.proto.resources.Leaderboard.LeaderboardItem.LeaderboardScore
	(*resources.ContestantInstance)(nil),                           // 5: isuxportal.proto.resources.ContestantInstance
}
var file_isuxportal_services_admin_leaderboard_dump_proto_depIdxs = []int32{
	2, // 0: isuxportal.proto.services.admin.GetLeaderboardDumpResponse.items:type_name -> isuxportal.proto.services.admin.GetLeaderboardDumpResponse.LeaderboardDumpItem
	3, // 1: isuxportal.proto.services.admin.GetLeaderboardDumpResponse.LeaderboardDumpItem.team:type_name -> isuxportal.proto.resources.Team
	4, // 2: isuxportal.proto.services.admin.GetLeaderboardDumpResponse.LeaderboardDumpItem.best_score:type_name -> isuxportal.proto.resources.Leaderboard.LeaderboardItem.LeaderboardScore
	4, // 3: isuxportal.proto.services.admin.GetLeaderboardDumpResponse.LeaderboardDumpItem.latest_score:type_name -> isuxportal.proto.resources.Leaderboard.LeaderboardItem.LeaderboardScore
	5, // 4: isuxportal.proto.services.admin.GetLeaderboardDumpResponse.LeaderboardDumpItem.target:type_name -> isuxportal.proto.resources.ContestantInstance
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_isuxportal_services_admin_leaderboard_dump_proto_init() }
func file_isuxportal_services_admin_leaderboard_dump_proto_init() {
	if File_isuxportal_services_admin_leaderboard_dump_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaderboardDumpQuery); i {
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
		file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaderboardDumpResponse); i {
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
		file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaderboardDumpResponse_LeaderboardDumpItem); i {
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
			RawDescriptor: file_isuxportal_services_admin_leaderboard_dump_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_services_admin_leaderboard_dump_proto_goTypes,
		DependencyIndexes: file_isuxportal_services_admin_leaderboard_dump_proto_depIdxs,
		MessageInfos:      file_isuxportal_services_admin_leaderboard_dump_proto_msgTypes,
	}.Build()
	File_isuxportal_services_admin_leaderboard_dump_proto = out.File
	file_isuxportal_services_admin_leaderboard_dump_proto_rawDesc = nil
	file_isuxportal_services_admin_leaderboard_dump_proto_goTypes = nil
	file_isuxportal_services_admin_leaderboard_dump_proto_depIdxs = nil
}
