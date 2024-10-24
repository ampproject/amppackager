// Code generated by protoc-gen-goext. DO NOT EDIT.

package apploadbalancer

import (
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

func (m *VirtualHost) SetName(v string) {
	m.Name = v
}

func (m *VirtualHost) SetAuthority(v []string) {
	m.Authority = v
}

func (m *VirtualHost) SetRoutes(v []*Route) {
	m.Routes = v
}

func (m *VirtualHost) SetModifyRequestHeaders(v []*HeaderModification) {
	m.ModifyRequestHeaders = v
}

func (m *VirtualHost) SetModifyResponseHeaders(v []*HeaderModification) {
	m.ModifyResponseHeaders = v
}

func (m *VirtualHost) SetRouteOptions(v *RouteOptions) {
	m.RouteOptions = v
}

func (m *RouteOptions) SetModifyRequestHeaders(v []*HeaderModification) {
	m.ModifyRequestHeaders = v
}

func (m *RouteOptions) SetModifyResponseHeaders(v []*HeaderModification) {
	m.ModifyResponseHeaders = v
}

func (m *RouteOptions) SetRbac(v *RBAC) {
	m.Rbac = v
}

func (m *RouteOptions) SetSecurityProfileId(v string) {
	m.SecurityProfileId = v
}

func (m *RBAC) SetAction(v RBAC_Action) {
	m.Action = v
}

func (m *RBAC) SetPrincipals(v []*Principals) {
	m.Principals = v
}

func (m *Principals) SetAndPrincipals(v []*Principal) {
	m.AndPrincipals = v
}

type Principal_Identifier = isPrincipal_Identifier

func (m *Principal) SetIdentifier(v Principal_Identifier) {
	m.Identifier = v
}

func (m *Principal) SetHeader(v *Principal_HeaderMatcher) {
	m.Identifier = &Principal_Header{
		Header: v,
	}
}

func (m *Principal) SetRemoteIp(v string) {
	m.Identifier = &Principal_RemoteIp{
		RemoteIp: v,
	}
}

func (m *Principal) SetAny(v bool) {
	m.Identifier = &Principal_Any{
		Any: v,
	}
}

func (m *Principal_HeaderMatcher) SetName(v string) {
	m.Name = v
}

func (m *Principal_HeaderMatcher) SetValue(v *StringMatch) {
	m.Value = v
}

type HeaderModification_Operation = isHeaderModification_Operation

func (m *HeaderModification) SetOperation(v HeaderModification_Operation) {
	m.Operation = v
}

func (m *HeaderModification) SetName(v string) {
	m.Name = v
}

func (m *HeaderModification) SetAppend(v string) {
	m.Operation = &HeaderModification_Append{
		Append: v,
	}
}

func (m *HeaderModification) SetReplace(v string) {
	m.Operation = &HeaderModification_Replace{
		Replace: v,
	}
}

func (m *HeaderModification) SetRemove(v bool) {
	m.Operation = &HeaderModification_Remove{
		Remove: v,
	}
}

func (m *HeaderModification) SetRename(v string) {
	m.Operation = &HeaderModification_Rename{
		Rename: v,
	}
}

type Route_Route = isRoute_Route

func (m *Route) SetRoute(v Route_Route) {
	m.Route = v
}

func (m *Route) SetName(v string) {
	m.Name = v
}

func (m *Route) SetHttp(v *HttpRoute) {
	m.Route = &Route_Http{
		Http: v,
	}
}

func (m *Route) SetGrpc(v *GrpcRoute) {
	m.Route = &Route_Grpc{
		Grpc: v,
	}
}

func (m *Route) SetRouteOptions(v *RouteOptions) {
	m.RouteOptions = v
}

type HttpRoute_Action = isHttpRoute_Action

func (m *HttpRoute) SetAction(v HttpRoute_Action) {
	m.Action = v
}

func (m *HttpRoute) SetMatch(v *HttpRouteMatch) {
	m.Match = v
}

func (m *HttpRoute) SetRoute(v *HttpRouteAction) {
	m.Action = &HttpRoute_Route{
		Route: v,
	}
}

func (m *HttpRoute) SetRedirect(v *RedirectAction) {
	m.Action = &HttpRoute_Redirect{
		Redirect: v,
	}
}

func (m *HttpRoute) SetDirectResponse(v *DirectResponseAction) {
	m.Action = &HttpRoute_DirectResponse{
		DirectResponse: v,
	}
}

type GrpcRoute_Action = isGrpcRoute_Action

func (m *GrpcRoute) SetAction(v GrpcRoute_Action) {
	m.Action = v
}

func (m *GrpcRoute) SetMatch(v *GrpcRouteMatch) {
	m.Match = v
}

func (m *GrpcRoute) SetRoute(v *GrpcRouteAction) {
	m.Action = &GrpcRoute_Route{
		Route: v,
	}
}

func (m *GrpcRoute) SetStatusResponse(v *GrpcStatusResponseAction) {
	m.Action = &GrpcRoute_StatusResponse{
		StatusResponse: v,
	}
}

func (m *HttpRouteMatch) SetHttpMethod(v []string) {
	m.HttpMethod = v
}

func (m *HttpRouteMatch) SetPath(v *StringMatch) {
	m.Path = v
}

func (m *GrpcRouteMatch) SetFqmn(v *StringMatch) {
	m.Fqmn = v
}

type StringMatch_Match = isStringMatch_Match

func (m *StringMatch) SetMatch(v StringMatch_Match) {
	m.Match = v
}

func (m *StringMatch) SetExactMatch(v string) {
	m.Match = &StringMatch_ExactMatch{
		ExactMatch: v,
	}
}

func (m *StringMatch) SetPrefixMatch(v string) {
	m.Match = &StringMatch_PrefixMatch{
		PrefixMatch: v,
	}
}

func (m *StringMatch) SetRegexMatch(v string) {
	m.Match = &StringMatch_RegexMatch{
		RegexMatch: v,
	}
}

type RedirectAction_Path = isRedirectAction_Path

func (m *RedirectAction) SetPath(v RedirectAction_Path) {
	m.Path = v
}

func (m *RedirectAction) SetReplaceScheme(v string) {
	m.ReplaceScheme = v
}

func (m *RedirectAction) SetReplaceHost(v string) {
	m.ReplaceHost = v
}

func (m *RedirectAction) SetReplacePort(v int64) {
	m.ReplacePort = v
}

func (m *RedirectAction) SetReplacePath(v string) {
	m.Path = &RedirectAction_ReplacePath{
		ReplacePath: v,
	}
}

func (m *RedirectAction) SetReplacePrefix(v string) {
	m.Path = &RedirectAction_ReplacePrefix{
		ReplacePrefix: v,
	}
}

func (m *RedirectAction) SetRemoveQuery(v bool) {
	m.RemoveQuery = v
}

func (m *RedirectAction) SetResponseCode(v RedirectAction_RedirectResponseCode) {
	m.ResponseCode = v
}

func (m *DirectResponseAction) SetStatus(v int64) {
	m.Status = v
}

func (m *DirectResponseAction) SetBody(v *Payload) {
	m.Body = v
}

func (m *GrpcStatusResponseAction) SetStatus(v GrpcStatusResponseAction_Status) {
	m.Status = v
}

type HttpRouteAction_HostRewriteSpecifier = isHttpRouteAction_HostRewriteSpecifier

func (m *HttpRouteAction) SetHostRewriteSpecifier(v HttpRouteAction_HostRewriteSpecifier) {
	m.HostRewriteSpecifier = v
}

func (m *HttpRouteAction) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *HttpRouteAction) SetTimeout(v *durationpb.Duration) {
	m.Timeout = v
}

func (m *HttpRouteAction) SetIdleTimeout(v *durationpb.Duration) {
	m.IdleTimeout = v
}

func (m *HttpRouteAction) SetHostRewrite(v string) {
	m.HostRewriteSpecifier = &HttpRouteAction_HostRewrite{
		HostRewrite: v,
	}
}

func (m *HttpRouteAction) SetAutoHostRewrite(v bool) {
	m.HostRewriteSpecifier = &HttpRouteAction_AutoHostRewrite{
		AutoHostRewrite: v,
	}
}

func (m *HttpRouteAction) SetPrefixRewrite(v string) {
	m.PrefixRewrite = v
}

func (m *HttpRouteAction) SetUpgradeTypes(v []string) {
	m.UpgradeTypes = v
}

type GrpcRouteAction_HostRewriteSpecifier = isGrpcRouteAction_HostRewriteSpecifier

func (m *GrpcRouteAction) SetHostRewriteSpecifier(v GrpcRouteAction_HostRewriteSpecifier) {
	m.HostRewriteSpecifier = v
}

func (m *GrpcRouteAction) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *GrpcRouteAction) SetMaxTimeout(v *durationpb.Duration) {
	m.MaxTimeout = v
}

func (m *GrpcRouteAction) SetIdleTimeout(v *durationpb.Duration) {
	m.IdleTimeout = v
}

func (m *GrpcRouteAction) SetHostRewrite(v string) {
	m.HostRewriteSpecifier = &GrpcRouteAction_HostRewrite{
		HostRewrite: v,
	}
}

func (m *GrpcRouteAction) SetAutoHostRewrite(v bool) {
	m.HostRewriteSpecifier = &GrpcRouteAction_AutoHostRewrite{
		AutoHostRewrite: v,
	}
}
