// Code generated by protoc-gen-goext. DO NOT EDIT.

package apploadbalancer

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetVirtualHostRequest) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *GetVirtualHostRequest) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *ListVirtualHostsRequest) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *ListVirtualHostsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListVirtualHostsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListVirtualHostsResponse) SetVirtualHosts(v []*VirtualHost) {
	m.VirtualHosts = v
}

func (m *ListVirtualHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateVirtualHostRequest) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *CreateVirtualHostRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateVirtualHostRequest) SetAuthority(v []string) {
	m.Authority = v
}

func (m *CreateVirtualHostRequest) SetRoutes(v []*Route) {
	m.Routes = v
}

func (m *CreateVirtualHostRequest) SetModifyRequestHeaders(v []*HeaderModification) {
	m.ModifyRequestHeaders = v
}

func (m *CreateVirtualHostRequest) SetModifyResponseHeaders(v []*HeaderModification) {
	m.ModifyResponseHeaders = v
}

func (m *CreateVirtualHostRequest) SetRouteOptions(v *RouteOptions) {
	m.RouteOptions = v
}

func (m *CreateVirtualHostRequest) SetRateLimit(v *RateLimit) {
	m.RateLimit = v
}

func (m *CreateVirtualHostMetadata) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *CreateVirtualHostMetadata) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *UpdateVirtualHostRequest) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *UpdateVirtualHostRequest) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *UpdateVirtualHostRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateVirtualHostRequest) SetAuthority(v []string) {
	m.Authority = v
}

func (m *UpdateVirtualHostRequest) SetRoutes(v []*Route) {
	m.Routes = v
}

func (m *UpdateVirtualHostRequest) SetModifyRequestHeaders(v []*HeaderModification) {
	m.ModifyRequestHeaders = v
}

func (m *UpdateVirtualHostRequest) SetModifyResponseHeaders(v []*HeaderModification) {
	m.ModifyResponseHeaders = v
}

func (m *UpdateVirtualHostRequest) SetRouteOptions(v *RouteOptions) {
	m.RouteOptions = v
}

func (m *UpdateVirtualHostRequest) SetRateLimit(v *RateLimit) {
	m.RateLimit = v
}

func (m *UpdateVirtualHostMetadata) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *UpdateVirtualHostMetadata) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *DeleteVirtualHostRequest) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *DeleteVirtualHostRequest) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *DeleteVirtualHostMetadata) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *DeleteVirtualHostMetadata) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *RemoveRouteRequest) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *RemoveRouteRequest) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *RemoveRouteRequest) SetRouteName(v string) {
	m.RouteName = v
}

func (m *RemoveRouteMetadata) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *RemoveRouteMetadata) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *RemoveRouteMetadata) SetRouteName(v string) {
	m.RouteName = v
}

type UpdateRouteRequest_Route = isUpdateRouteRequest_Route

func (m *UpdateRouteRequest) SetRoute(v UpdateRouteRequest_Route) {
	m.Route = v
}

func (m *UpdateRouteRequest) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *UpdateRouteRequest) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *UpdateRouteRequest) SetRouteName(v string) {
	m.RouteName = v
}

func (m *UpdateRouteRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateRouteRequest) SetHttp(v *HttpRoute) {
	m.Route = &UpdateRouteRequest_Http{
		Http: v,
	}
}

func (m *UpdateRouteRequest) SetGrpc(v *GrpcRoute) {
	m.Route = &UpdateRouteRequest_Grpc{
		Grpc: v,
	}
}

func (m *UpdateRouteRequest) SetRouteOptions(v *RouteOptions) {
	m.RouteOptions = v
}

func (m *UpdateRouteMetadata) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *UpdateRouteMetadata) SetVirtualHostName(v string) {
	m.VirtualHostName = v
}

func (m *UpdateRouteMetadata) SetRouteName(v string) {
	m.RouteName = v
}
