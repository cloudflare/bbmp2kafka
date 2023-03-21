// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: bbmp/bbmp.proto

package bbmp

import (
	api "github.com/bio-routing/bio-rd/net/api"
	api1 "github.com/bio-routing/bio-rd/route/api"
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

type BBMPMessage_MessageType int32

const (
	BBMPMessage_Unknown                BBMPMessage_MessageType = 0
	BBMPMessage_PeerUpNotification     BBMPMessage_MessageType = 1
	BBMPMessage_PeerDownNotification   BBMPMessage_MessageType = 2
	BBMPMessage_InitiationMessage      BBMPMessage_MessageType = 3
	BBMPMessage_TerminationMessage     BBMPMessage_MessageType = 4
	BBMPMessage_RouteMonitoringMessage BBMPMessage_MessageType = 5
	BBMPMessage_RouteMirroringMessage  BBMPMessage_MessageType = 6
)

// Enum value maps for BBMPMessage_MessageType.
var (
	BBMPMessage_MessageType_name = map[int32]string{
		0: "Unknown",
		1: "PeerUpNotification",
		2: "PeerDownNotification",
		3: "InitiationMessage",
		4: "TerminationMessage",
		5: "RouteMonitoringMessage",
		6: "RouteMirroringMessage",
	}
	BBMPMessage_MessageType_value = map[string]int32{
		"Unknown":                0,
		"PeerUpNotification":     1,
		"PeerDownNotification":   2,
		"InitiationMessage":      3,
		"TerminationMessage":     4,
		"RouteMonitoringMessage": 5,
		"RouteMirroringMessage":  6,
	}
)

func (x BBMPMessage_MessageType) Enum() *BBMPMessage_MessageType {
	p := new(BBMPMessage_MessageType)
	*p = x
	return p
}

func (x BBMPMessage_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BBMPMessage_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_bbmp_bbmp_proto_enumTypes[0].Descriptor()
}

func (BBMPMessage_MessageType) Type() protoreflect.EnumType {
	return &file_bbmp_bbmp_proto_enumTypes[0]
}

func (x BBMPMessage_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BBMPMessage_MessageType.Descriptor instead.
func (BBMPMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_bbmp_bbmp_proto_rawDescGZIP(), []int{0, 0}
}

type BBMPMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType                  BBMPMessage_MessageType       `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=cloudflare.net.bbmp.BBMPMessage_MessageType" json:"message_type,omitempty"`
	BbmpUnicastMonitoringMessage *BBMPUnicastMonitoringMessage `protobuf:"bytes,2,opt,name=bbmp_unicast_monitoring_message,json=bbmpUnicastMonitoringMessage,proto3" json:"bbmp_unicast_monitoring_message,omitempty"`
}

func (x *BBMPMessage) Reset() {
	*x = BBMPMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbmp_bbmp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BBMPMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BBMPMessage) ProtoMessage() {}

func (x *BBMPMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bbmp_bbmp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BBMPMessage.ProtoReflect.Descriptor instead.
func (*BBMPMessage) Descriptor() ([]byte, []int) {
	return file_bbmp_bbmp_proto_rawDescGZIP(), []int{0}
}

func (x *BBMPMessage) GetMessageType() BBMPMessage_MessageType {
	if x != nil {
		return x.MessageType
	}
	return BBMPMessage_Unknown
}

func (x *BBMPMessage) GetBbmpUnicastMonitoringMessage() *BBMPUnicastMonitoringMessage {
	if x != nil {
		return x.BbmpUnicastMonitoringMessage
	}
	return nil
}

type BBMPUnicastMonitoringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterIp      *api.IP       `protobuf:"bytes,1,opt,name=router_ip,json=routerIp,proto3" json:"router_ip,omitempty"`
	LocalBpgIp    *api.IP       `protobuf:"bytes,2,opt,name=local_bpg_ip,json=localBpgIp,proto3" json:"local_bpg_ip,omitempty"`
	NeighborBgpIp *api.IP       `protobuf:"bytes,3,opt,name=neighbor_bgp_ip,json=neighborBgpIp,proto3" json:"neighbor_bgp_ip,omitempty"`
	LocalAs       uint32        `protobuf:"varint,4,opt,name=local_as,json=localAs,proto3" json:"local_as,omitempty"`
	RemoteAs      uint32        `protobuf:"varint,5,opt,name=remote_as,json=remoteAs,proto3" json:"remote_as,omitempty"`
	Announcement  bool          `protobuf:"varint,6,opt,name=announcement,proto3" json:"announcement,omitempty"`
	Pfx           *api.Prefix   `protobuf:"bytes,7,opt,name=pfx,proto3" json:"pfx,omitempty"`
	BgpPath       *api1.BGPPath `protobuf:"bytes,8,opt,name=bgp_path,json=bgpPath,proto3" json:"bgp_path,omitempty"`
	Timestamp     uint32        `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BBMPUnicastMonitoringMessage) Reset() {
	*x = BBMPUnicastMonitoringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbmp_bbmp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BBMPUnicastMonitoringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BBMPUnicastMonitoringMessage) ProtoMessage() {}

func (x *BBMPUnicastMonitoringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bbmp_bbmp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BBMPUnicastMonitoringMessage.ProtoReflect.Descriptor instead.
func (*BBMPUnicastMonitoringMessage) Descriptor() ([]byte, []int) {
	return file_bbmp_bbmp_proto_rawDescGZIP(), []int{1}
}

func (x *BBMPUnicastMonitoringMessage) GetRouterIp() *api.IP {
	if x != nil {
		return x.RouterIp
	}
	return nil
}

func (x *BBMPUnicastMonitoringMessage) GetLocalBpgIp() *api.IP {
	if x != nil {
		return x.LocalBpgIp
	}
	return nil
}

func (x *BBMPUnicastMonitoringMessage) GetNeighborBgpIp() *api.IP {
	if x != nil {
		return x.NeighborBgpIp
	}
	return nil
}

func (x *BBMPUnicastMonitoringMessage) GetLocalAs() uint32 {
	if x != nil {
		return x.LocalAs
	}
	return 0
}

func (x *BBMPUnicastMonitoringMessage) GetRemoteAs() uint32 {
	if x != nil {
		return x.RemoteAs
	}
	return 0
}

func (x *BBMPUnicastMonitoringMessage) GetAnnouncement() bool {
	if x != nil {
		return x.Announcement
	}
	return false
}

func (x *BBMPUnicastMonitoringMessage) GetPfx() *api.Prefix {
	if x != nil {
		return x.Pfx
	}
	return nil
}

func (x *BBMPUnicastMonitoringMessage) GetBgpPath() *api1.BGPPath {
	if x != nil {
		return x.BgpPath
	}
	return nil
}

func (x *BBMPUnicastMonitoringMessage) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_bbmp_bbmp_proto protoreflect.FileDescriptor

var file_bbmp_bbmp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x62, 0x6d, 0x70, 0x2f, 0x62, 0x62, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x65,
	0x74, 0x2e, 0x62, 0x62, 0x6d, 0x70, 0x1a, 0x11, 0x6e, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x03, 0x0a, 0x0b, 0x42, 0x42, 0x4d, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x4f, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c,
	0x61, 0x72, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x62, 0x62, 0x6d, 0x70, 0x2e, 0x42, 0x42, 0x4d,
	0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x78, 0x0a, 0x1f, 0x62, 0x62, 0x6d, 0x70, 0x5f, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x62, 0x62, 0x6d, 0x70,
	0x2e, 0x42, 0x42, 0x4d, 0x50, 0x55, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x1c, 0x62,
	0x62, 0x6d, 0x70, 0x55, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x0b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x65, 0x65, 0x72,
	0x55, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x50, 0x65, 0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x69,
	0x72, 0x72, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x06,
	0x22, 0xf8, 0x02, 0x0a, 0x1c, 0x42, 0x42, 0x4d, 0x50, 0x55, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x28, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49,
	0x50, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x70, 0x12, 0x2d, 0x0a, 0x0c, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x62, 0x70, 0x67, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x52, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x42, 0x70, 0x67, 0x49, 0x70, 0x12, 0x33, 0x0a, 0x0f, 0x6e, 0x65,
	0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x62, 0x67, 0x70, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50,
	0x52, 0x0d, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x42, 0x67, 0x70, 0x49, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x70,
	0x66, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e,
	0x65, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x03, 0x70, 0x66, 0x78, 0x12, 0x2d,
	0x0a, 0x08, 0x62, 0x67, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x42, 0x47, 0x50,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x07, 0x62, 0x67, 0x70, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6c, 0x61, 0x72, 0x65, 0x2f, 0x62, 0x62, 0x6d, 0x70, 0x32, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f,
	0x62, 0x62, 0x6d, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bbmp_bbmp_proto_rawDescOnce sync.Once
	file_bbmp_bbmp_proto_rawDescData = file_bbmp_bbmp_proto_rawDesc
)

func file_bbmp_bbmp_proto_rawDescGZIP() []byte {
	file_bbmp_bbmp_proto_rawDescOnce.Do(func() {
		file_bbmp_bbmp_proto_rawDescData = protoimpl.X.CompressGZIP(file_bbmp_bbmp_proto_rawDescData)
	})
	return file_bbmp_bbmp_proto_rawDescData
}

var file_bbmp_bbmp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bbmp_bbmp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bbmp_bbmp_proto_goTypes = []interface{}{
	(BBMPMessage_MessageType)(0),         // 0: cloudflare.net.bbmp.BBMPMessage.MessageType
	(*BBMPMessage)(nil),                  // 1: cloudflare.net.bbmp.BBMPMessage
	(*BBMPUnicastMonitoringMessage)(nil), // 2: cloudflare.net.bbmp.BBMPUnicastMonitoringMessage
	(*api.IP)(nil),                       // 3: bio.net.IP
	(*api.Prefix)(nil),                   // 4: bio.net.Prefix
	(*api1.BGPPath)(nil),                 // 5: bio.route.BGPPath
}
var file_bbmp_bbmp_proto_depIdxs = []int32{
	0, // 0: cloudflare.net.bbmp.BBMPMessage.message_type:type_name -> cloudflare.net.bbmp.BBMPMessage.MessageType
	2, // 1: cloudflare.net.bbmp.BBMPMessage.bbmp_unicast_monitoring_message:type_name -> cloudflare.net.bbmp.BBMPUnicastMonitoringMessage
	3, // 2: cloudflare.net.bbmp.BBMPUnicastMonitoringMessage.router_ip:type_name -> bio.net.IP
	3, // 3: cloudflare.net.bbmp.BBMPUnicastMonitoringMessage.local_bpg_ip:type_name -> bio.net.IP
	3, // 4: cloudflare.net.bbmp.BBMPUnicastMonitoringMessage.neighbor_bgp_ip:type_name -> bio.net.IP
	4, // 5: cloudflare.net.bbmp.BBMPUnicastMonitoringMessage.pfx:type_name -> bio.net.Prefix
	5, // 6: cloudflare.net.bbmp.BBMPUnicastMonitoringMessage.bgp_path:type_name -> bio.route.BGPPath
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_bbmp_bbmp_proto_init() }
func file_bbmp_bbmp_proto_init() {
	if File_bbmp_bbmp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bbmp_bbmp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBMPMessage); i {
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
		file_bbmp_bbmp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBMPUnicastMonitoringMessage); i {
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
			RawDescriptor: file_bbmp_bbmp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bbmp_bbmp_proto_goTypes,
		DependencyIndexes: file_bbmp_bbmp_proto_depIdxs,
		EnumInfos:         file_bbmp_bbmp_proto_enumTypes,
		MessageInfos:      file_bbmp_bbmp_proto_msgTypes,
	}.Build()
	File_bbmp_bbmp_proto = out.File
	file_bbmp_bbmp_proto_rawDesc = nil
	file_bbmp_bbmp_proto_goTypes = nil
	file_bbmp_bbmp_proto_depIdxs = nil
}