// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: contrib/envoy/extensions/filters/network/sip_proxy/v3alpha/sip_proxy.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/contrib/envoy/extensions/filters/network/sip_proxy/tra/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type SipProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	RouteConfig *RouteConfiguration `protobuf:"bytes,2,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// A list of individual Sip filters that make up the filter chain for requests made to the
	// Sip proxy. Order matters as the filters are processed sequentially. For backwards
	// compatibility, if no sip_filters are specified, a default Sip router filter
	// (`envoy.filters.sip.router`) is used.
	// [#extension-category: envoy.sip_proxy.filters]
	SipFilters []*SipFilter          `protobuf:"bytes,3,rep,name=sip_filters,json=sipFilters,proto3" json:"sip_filters,omitempty"`
	Settings   *SipProxy_SipSettings `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *SipProxy) Reset() {
	*x = SipProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipProxy) ProtoMessage() {}

func (x *SipProxy) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipProxy.ProtoReflect.Descriptor instead.
func (*SipProxy) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *SipProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *SipProxy) GetRouteConfig() *RouteConfiguration {
	if x != nil {
		return x.RouteConfig
	}
	return nil
}

func (x *SipProxy) GetSipFilters() []*SipFilter {
	if x != nil {
		return x.SipFilters
	}
	return nil
}

func (x *SipProxy) GetSettings() *SipProxy_SipSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

// SipFilter configures a Sip filter.
type SipFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated. See the supported
	// filters for further documentation.
	//
	// Types that are assignable to ConfigType:
	//	*SipFilter_TypedConfig
	ConfigType isSipFilter_ConfigType `protobuf_oneof:"config_type"`
}

func (x *SipFilter) Reset() {
	*x = SipFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipFilter) ProtoMessage() {}

func (x *SipFilter) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipFilter.ProtoReflect.Descriptor instead.
func (*SipFilter) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *SipFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *SipFilter) GetConfigType() isSipFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *SipFilter) GetTypedConfig() *any.Any {
	if x, ok := x.GetConfigType().(*SipFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

type isSipFilter_ConfigType interface {
	isSipFilter_ConfigType()
}

type SipFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*SipFilter_TypedConfig) isSipFilter_ConfigType() {}

// SipProtocolOptions specifies Sip upstream protocol options. This object is used in
// :ref:`typed_extension_protocol_options<envoy_v3_api_field_config.cluster.v3.Cluster.typed_extension_protocol_options>`,
// keyed by the name `envoy.filters.network.sip_proxy`.
type SipProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All sip messages in one dialog should go to the same endpoint.
	SessionAffinity bool `protobuf:"varint,1,opt,name=session_affinity,json=sessionAffinity,proto3" json:"session_affinity,omitempty"`
	// The Register with Authorization header should go to the same endpoint which send out the 401 Unauthorized.
	RegistrationAffinity bool `protobuf:"varint,2,opt,name=registration_affinity,json=registrationAffinity,proto3" json:"registration_affinity,omitempty"`
	// Customized affinity
	CustomizedAffinity []*CustomizedAffinity `protobuf:"bytes,3,rep,name=customized_affinity,json=customizedAffinity,proto3" json:"customized_affinity,omitempty"`
}

func (x *SipProtocolOptions) Reset() {
	*x = SipProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipProtocolOptions) ProtoMessage() {}

func (x *SipProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipProtocolOptions.ProtoReflect.Descriptor instead.
func (*SipProtocolOptions) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *SipProtocolOptions) GetSessionAffinity() bool {
	if x != nil {
		return x.SessionAffinity
	}
	return false
}

func (x *SipProtocolOptions) GetRegistrationAffinity() bool {
	if x != nil {
		return x.RegistrationAffinity
	}
	return false
}

func (x *SipProtocolOptions) GetCustomizedAffinity() []*CustomizedAffinity {
	if x != nil {
		return x.CustomizedAffinity
	}
	return nil
}

// For affinity
type CustomizedAffinity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyName   string `protobuf:"bytes,1,opt,name=key_name,json=keyName,proto3" json:"key_name,omitempty"`
	Query     bool   `protobuf:"varint,2,opt,name=query,proto3" json:"query,omitempty"`
	Subscribe bool   `protobuf:"varint,3,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
}

func (x *CustomizedAffinity) Reset() {
	*x = CustomizedAffinity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomizedAffinity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomizedAffinity) ProtoMessage() {}

func (x *CustomizedAffinity) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomizedAffinity.ProtoReflect.Descriptor instead.
func (*CustomizedAffinity) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *CustomizedAffinity) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *CustomizedAffinity) GetQuery() bool {
	if x != nil {
		return x.Query
	}
	return false
}

func (x *CustomizedAffinity) GetSubscribe() bool {
	if x != nil {
		return x.Subscribe
	}
	return false
}

type SipProxy_SipSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// transaction timeout timer [Timer B] unit is milliseconds, default value 64*T1.
	//
	// Session Initiation Protocol (SIP) timer summary
	//
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer   | Default value           | Section  | Meaning                                                                      |
	// +=========+=========================+==========+==============================================================================+
	// | T1      | 500 ms                  | 17.1.1.1 | Round-trip time (RTT) estimate                                               |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | T2      | 4 sec                   | 17.1.2.2 | Maximum re-transmission interval for non-INVITE requests and INVITE responses|
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | T4      | 5 sec                   | 17.1.2.2 | Maximum duration that a message can remain in the network                    |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer A | initially T1            | 17.1.1.2 | INVITE request re-transmission interval, for UDP only                        |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer B | 64*T1                   | 17.1.1.2 | INVITE transaction timeout timer                                             |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer D | > 32 sec. for UDP       | 17.1.1.2 | Wait time for response re-transmissions                                      |
	// |         | 0 sec. for TCP and SCTP |          |                                                                              |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer E | initially T1            | 17.1.2.2 | Non-INVITE request re-transmission interval, UDP only                        |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer F | 64*T1                   | 17.1.2.2 | Non-INVITE transaction timeout timer                                         |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer G | initially T1            | 17.2.1   | INVITE response re-transmission interval                                     |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer H | 64*T1                   | 17.2.1   | Wait time for ACK receipt                                                    |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer I | T4 for UDP              | 17.2.1   | Wait time for ACK re-transmissions                                           |
	// |         | 0 sec. for TCP and SCTP |          |                                                                              |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer J | 64*T1 for UDP           | 17.2.2   | Wait time for re-transmissions of non-INVITE requests                        |
	// |         | 0 sec. for TCP and SCTP |          |                                                                              |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	// | Timer K | T4 for UDP              | 17.1.2.2 | Wait time for response re-transmissions                                      |
	// |         | 0 sec. for TCP and SCTP |          |                                                                              |
	// +---------+-------------------------+----------+------------------------------------------------------------------------------+
	TransactionTimeout *duration.Duration `protobuf:"bytes,1,opt,name=transaction_timeout,json=transactionTimeout,proto3" json:"transaction_timeout,omitempty"`
	// own domain name
	OwnDomain string `protobuf:"bytes,2,opt,name=own_domain,json=ownDomain,proto3" json:"own_domain,omitempty"`
	// points to domain match with own_domain
	DomainMatchParameterName string                    `protobuf:"bytes,3,opt,name=domain_match_parameter_name,json=domainMatchParameterName,proto3" json:"domain_match_parameter_name,omitempty"`
	TraServiceConfig         *v3alpha.TraServiceConfig `protobuf:"bytes,4,opt,name=tra_service_config,json=traServiceConfig,proto3" json:"tra_service_config,omitempty"`
}

func (x *SipProxy_SipSettings) Reset() {
	*x = SipProxy_SipSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipProxy_SipSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipProxy_SipSettings) ProtoMessage() {}

func (x *SipProxy_SipSettings) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipProxy_SipSettings.ProtoReflect.Descriptor instead.
func (*SipProxy_SipSettings) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SipProxy_SipSettings) GetTransactionTimeout() *duration.Duration {
	if x != nil {
		return x.TransactionTimeout
	}
	return nil
}

func (x *SipProxy_SipSettings) GetOwnDomain() string {
	if x != nil {
		return x.OwnDomain
	}
	return ""
}

func (x *SipProxy_SipSettings) GetDomainMatchParameterName() string {
	if x != nil {
		return x.DomainMatchParameterName
	}
	return ""
}

func (x *SipProxy_SipSettings) GetTraServiceConfig() *v3alpha.TraServiceConfig {
	if x != nil {
		return x.TraServiceConfig
	}
	return nil
}

var File_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x69, 0x70, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x69, 0x70,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73,
	0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x48, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x69, 0x70, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x74, 0x72, 0x61, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76,
	0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x05, 0x0a, 0x08, 0x53, 0x69, 0x70, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x69, 0x0a, 0x0c,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x46, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5e, 0x0a, 0x0b, 0x73, 0x69, 0x70, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x73, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x53, 0x69, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x69, 0x70,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x64, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x69, 0x70,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x69, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x69, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0xaf, 0x02,
	0x0a, 0x0b, 0x53, 0x69, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4a, 0x0a,
	0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e,
	0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x77, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x1b, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x76, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x74, 0x72, 0x61, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x72, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x74,
	0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x72, 0x0a, 0x09, 0x53, 0x69, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x74, 0x79, 0x70,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x12, 0x53, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x66, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x66, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x77, 0x0a, 0x13, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x69, 0x70, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x66, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x52,
	0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x66, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x79, 0x22, 0x63, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65,
	0x64, 0x41, 0x66, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0xb6, 0x01, 0x0a, 0x40, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x69, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0d, 0x53,
	0x69, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescData = file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_goTypes = []interface{}{
	(*SipProxy)(nil),                 // 0: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy
	(*SipFilter)(nil),                // 1: envoy.extensions.filters.network.sip_proxy.v3alpha.SipFilter
	(*SipProtocolOptions)(nil),       // 2: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProtocolOptions
	(*CustomizedAffinity)(nil),       // 3: envoy.extensions.filters.network.sip_proxy.v3alpha.CustomizedAffinity
	(*SipProxy_SipSettings)(nil),     // 4: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy.SipSettings
	(*RouteConfiguration)(nil),       // 5: envoy.extensions.filters.network.sip_proxy.v3alpha.RouteConfiguration
	(*any.Any)(nil),                  // 6: google.protobuf.Any
	(*duration.Duration)(nil),        // 7: google.protobuf.Duration
	(*v3alpha.TraServiceConfig)(nil), // 8: envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraServiceConfig
}
var file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_depIdxs = []int32{
	5, // 0: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy.route_config:type_name -> envoy.extensions.filters.network.sip_proxy.v3alpha.RouteConfiguration
	1, // 1: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy.sip_filters:type_name -> envoy.extensions.filters.network.sip_proxy.v3alpha.SipFilter
	4, // 2: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy.settings:type_name -> envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy.SipSettings
	6, // 3: envoy.extensions.filters.network.sip_proxy.v3alpha.SipFilter.typed_config:type_name -> google.protobuf.Any
	3, // 4: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProtocolOptions.customized_affinity:type_name -> envoy.extensions.filters.network.sip_proxy.v3alpha.CustomizedAffinity
	7, // 5: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy.SipSettings.transaction_timeout:type_name -> google.protobuf.Duration
	8, // 6: envoy.extensions.filters.network.sip_proxy.v3alpha.SipProxy.SipSettings.tra_service_config:type_name -> envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraServiceConfig
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_init() }
func file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_init() {
	if File_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto != nil {
		return
	}
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipProxy); i {
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
		file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipFilter); i {
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
		file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipProtocolOptions); i {
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
		file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomizedAffinity); i {
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
		file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipProxy_SipSettings); i {
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
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SipFilter_TypedConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto = out.File
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_network_sip_proxy_v3alpha_sip_proxy_proto_depIdxs = nil
}
