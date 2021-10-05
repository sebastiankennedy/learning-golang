// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/maps/routes/v1/compute_custom_routes_response.proto

package routes

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RouteObjective used for the response.
type ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective int32

const (
	// Fallback route objective unspecified.
	ComputeCustomRoutesResponse_FallbackInfo_FALLBACK_ROUTE_OBJECTIVE_UNSPECIFIED ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective = 0
	// If customer requests RateCard and sets include_tolls to true, and
	// Google does not have toll price data for the route, the API falls back
	// to RateCard without considering toll price.
	ComputeCustomRoutesResponse_FallbackInfo_FALLBACK_RATECARD_WITHOUT_TOLL_PRICE_DATA ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective = 1
)

// Enum value maps for ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective.
var (
	ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective_name = map[int32]string{
		0: "FALLBACK_ROUTE_OBJECTIVE_UNSPECIFIED",
		1: "FALLBACK_RATECARD_WITHOUT_TOLL_PRICE_DATA",
	}
	ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective_value = map[string]int32{
		"FALLBACK_ROUTE_OBJECTIVE_UNSPECIFIED":      0,
		"FALLBACK_RATECARD_WITHOUT_TOLL_PRICE_DATA": 1,
	}
)

func (x ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective) Enum() *ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective {
	p := new(ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective)
	*p = x
	return p
}

func (x ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routes_v1_compute_custom_routes_response_proto_enumTypes[0].Descriptor()
}

func (ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective) Type() protoreflect.EnumType {
	return &file_google_maps_routes_v1_compute_custom_routes_response_proto_enumTypes[0]
}

func (x ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective.Descriptor instead.
func (ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescGZIP(), []int{0, 0, 0}
}

// ComputeCustomRoutes response message.
type ComputeCustomRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ‘best’ routes for the input route objective.
	Routes []*CustomRoute `protobuf:"bytes,7,rep,name=routes,proto3" json:"routes,omitempty"`
	// The fastest reference route.
	FastestRoute *CustomRoute `protobuf:"bytes,5,opt,name=fastest_route,json=fastestRoute,proto3" json:"fastest_route,omitempty"`
	// The shortest reference route.
	ShortestRoute *CustomRoute `protobuf:"bytes,6,opt,name=shortest_route,json=shortestRoute,proto3" json:"shortest_route,omitempty"`
	// Fallback info for custom routes.
	FallbackInfo *ComputeCustomRoutesResponse_FallbackInfo `protobuf:"bytes,8,opt,name=fallback_info,json=fallbackInfo,proto3" json:"fallback_info,omitempty"`
}

func (x *ComputeCustomRoutesResponse) Reset() {
	*x = ComputeCustomRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeCustomRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeCustomRoutesResponse) ProtoMessage() {}

func (x *ComputeCustomRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeCustomRoutesResponse.ProtoReflect.Descriptor instead.
func (*ComputeCustomRoutesResponse) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeCustomRoutesResponse) GetRoutes() []*CustomRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *ComputeCustomRoutesResponse) GetFastestRoute() *CustomRoute {
	if x != nil {
		return x.FastestRoute
	}
	return nil
}

func (x *ComputeCustomRoutesResponse) GetShortestRoute() *CustomRoute {
	if x != nil {
		return x.ShortestRoute
	}
	return nil
}

func (x *ComputeCustomRoutesResponse) GetFallbackInfo() *ComputeCustomRoutesResponse_FallbackInfo {
	if x != nil {
		return x.FallbackInfo
	}
	return nil
}

// Encapsulates fallback info for ComputeCustomRoutes. ComputeCustomRoutes
// performs two types of fallbacks:
//
// 1. If it cannot compute the route using the routing_preference requested by
// the customer, it will fallback to another routing mode. In this case
// fallback_routing_mode and routing_mode_fallback_reason are used to
// communicate the fallback routing mode used, as well as the reason for
// fallback. Fallback of routing_preference is not supported in
// ComputeCustomRoutes Alpha.
//
// 2. If it cannot compute a 'best' route for the route objective specified by
// the customer, it might fallback to another objective.
// fallback_route_objective is used to communicate the fallback route
// objective.
type ComputeCustomRoutesResponse_FallbackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Routing mode used for the response. If fallback was triggered, the mode
	// may be different from routing preference set in the original client
	// request.
	RoutingMode FallbackRoutingMode `protobuf:"varint,1,opt,name=routing_mode,json=routingMode,proto3,enum=google.maps.routes.v1.FallbackRoutingMode" json:"routing_mode,omitempty"`
	// The reason why fallback response was used instead of the original
	// response.
	// This field is only populated when the fallback mode is triggered and
	// the fallback response is returned.
	RoutingModeReason FallbackReason `protobuf:"varint,2,opt,name=routing_mode_reason,json=routingModeReason,proto3,enum=google.maps.routes.v1.FallbackReason" json:"routing_mode_reason,omitempty"`
	// The route objective used for the response. If fallback was triggered, the
	// objective may be different from the route objective provided in the
	// original client request.
	RouteObjective ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective `protobuf:"varint,3,opt,name=route_objective,json=routeObjective,proto3,enum=google.maps.routes.v1.ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective" json:"route_objective,omitempty"`
}

func (x *ComputeCustomRoutesResponse_FallbackInfo) Reset() {
	*x = ComputeCustomRoutesResponse_FallbackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeCustomRoutesResponse_FallbackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeCustomRoutesResponse_FallbackInfo) ProtoMessage() {}

func (x *ComputeCustomRoutesResponse_FallbackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeCustomRoutesResponse_FallbackInfo.ProtoReflect.Descriptor instead.
func (*ComputeCustomRoutesResponse_FallbackInfo) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ComputeCustomRoutesResponse_FallbackInfo) GetRoutingMode() FallbackRoutingMode {
	if x != nil {
		return x.RoutingMode
	}
	return FallbackRoutingMode_FALLBACK_ROUTING_MODE_UNSPECIFIED
}

func (x *ComputeCustomRoutesResponse_FallbackInfo) GetRoutingModeReason() FallbackReason {
	if x != nil {
		return x.RoutingModeReason
	}
	return FallbackReason_FALLBACK_REASON_UNSPECIFIED
}

func (x *ComputeCustomRoutesResponse_FallbackInfo) GetRouteObjective() ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective {
	if x != nil {
		return x.RouteObjective
	}
	return ComputeCustomRoutesResponse_FallbackInfo_FALLBACK_ROUTE_OBJECTIVE_UNSPECIFIED
}

var File_google_maps_routes_v1_compute_custom_routes_response_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x05, 0x0a, 0x1b, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0d, 0x66, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x0c, 0x66, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x49, 0x0a,
	0x0e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x66, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0xa8,
	0x03, 0x0a, 0x0c, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x4d, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x55,
	0x0a, 0x13, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x7f, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x56,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x71, 0x0a, 0x16, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x28, 0x0a, 0x24, 0x46, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x52, 0x4f, 0x55,
	0x54, 0x45, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2d, 0x0a, 0x29, 0x46, 0x41,
	0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x43, 0x41, 0x52, 0x44, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x4f, 0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x49,
	0x43, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x42, 0xb6, 0x01, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x4d,
	0x52, 0x53, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescOnce sync.Once
	file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescData = file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDesc
)

func file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescGZIP() []byte {
	file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescOnce.Do(func() {
		file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescData)
	})
	return file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDescData
}

var file_google_maps_routes_v1_compute_custom_routes_response_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_routes_v1_compute_custom_routes_response_proto_goTypes = []interface{}{
	(ComputeCustomRoutesResponse_FallbackInfo_FallbackRouteObjective)(0), // 0: google.maps.routes.v1.ComputeCustomRoutesResponse.FallbackInfo.FallbackRouteObjective
	(*ComputeCustomRoutesResponse)(nil),                                  // 1: google.maps.routes.v1.ComputeCustomRoutesResponse
	(*ComputeCustomRoutesResponse_FallbackInfo)(nil),                     // 2: google.maps.routes.v1.ComputeCustomRoutesResponse.FallbackInfo
	(*CustomRoute)(nil),      // 3: google.maps.routes.v1.CustomRoute
	(FallbackRoutingMode)(0), // 4: google.maps.routes.v1.FallbackRoutingMode
	(FallbackReason)(0),      // 5: google.maps.routes.v1.FallbackReason
}
var file_google_maps_routes_v1_compute_custom_routes_response_proto_depIdxs = []int32{
	3, // 0: google.maps.routes.v1.ComputeCustomRoutesResponse.routes:type_name -> google.maps.routes.v1.CustomRoute
	3, // 1: google.maps.routes.v1.ComputeCustomRoutesResponse.fastest_route:type_name -> google.maps.routes.v1.CustomRoute
	3, // 2: google.maps.routes.v1.ComputeCustomRoutesResponse.shortest_route:type_name -> google.maps.routes.v1.CustomRoute
	2, // 3: google.maps.routes.v1.ComputeCustomRoutesResponse.fallback_info:type_name -> google.maps.routes.v1.ComputeCustomRoutesResponse.FallbackInfo
	4, // 4: google.maps.routes.v1.ComputeCustomRoutesResponse.FallbackInfo.routing_mode:type_name -> google.maps.routes.v1.FallbackRoutingMode
	5, // 5: google.maps.routes.v1.ComputeCustomRoutesResponse.FallbackInfo.routing_mode_reason:type_name -> google.maps.routes.v1.FallbackReason
	0, // 6: google.maps.routes.v1.ComputeCustomRoutesResponse.FallbackInfo.route_objective:type_name -> google.maps.routes.v1.ComputeCustomRoutesResponse.FallbackInfo.FallbackRouteObjective
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_compute_custom_routes_response_proto_init() }
func file_google_maps_routes_v1_compute_custom_routes_response_proto_init() {
	if File_google_maps_routes_v1_compute_custom_routes_response_proto != nil {
		return
	}
	file_google_maps_routes_v1_custom_route_proto_init()
	file_google_maps_routes_v1_fallback_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeCustomRoutesResponse); i {
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
		file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeCustomRoutesResponse_FallbackInfo); i {
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
			RawDescriptor: file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routes_v1_compute_custom_routes_response_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_compute_custom_routes_response_proto_depIdxs,
		EnumInfos:         file_google_maps_routes_v1_compute_custom_routes_response_proto_enumTypes,
		MessageInfos:      file_google_maps_routes_v1_compute_custom_routes_response_proto_msgTypes,
	}.Build()
	File_google_maps_routes_v1_compute_custom_routes_response_proto = out.File
	file_google_maps_routes_v1_compute_custom_routes_response_proto_rawDesc = nil
	file_google_maps_routes_v1_compute_custom_routes_response_proto_goTypes = nil
	file_google_maps_routes_v1_compute_custom_routes_response_proto_depIdxs = nil
}
