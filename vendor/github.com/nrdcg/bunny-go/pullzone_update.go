package bunny

import (
	"context"
	"fmt"
)

// PullZoneUpdateOptions represents the request parameters for the Update Pull
// Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_updatepullzone
type PullZoneUpdateOptions struct {
	AWSSigningEnabled                     *bool    `json:"AWSSigningEnabled,omitempty"`
	AWSSigningKey                         *string  `json:"AWSSigningKey,omitempty"`
	AWSSigningRegionName                  *string  `json:"AWSSigningRegionName,omitempty"`
	AWSSigningSecret                      *string  `json:"AWSSigningSecret,omitempty"`
	AccessControlOriginHeaderExtensions   []string `json:"AccessControlOriginHeaderExtensions,omitempty"`
	AddCanonicalHeader                    *bool    `json:"AddCanonicalHeader,omitempty"`
	AddHostHeader                         *bool    `json:"AddHostHeader,omitempty"`
	AllowedReferrers                      []string `json:"AllowedReferrers,omitempty"`
	BlockPostRequests                     *bool    `json:"BlockPostRequests,omitempty"`
	BlockRootPathAccess                   *bool    `json:"BlockRootPathAccess,omitempty"`
	BlockedCountries                      []string `json:"BlockedCountries,omitempty"`
	BlockedIPs                            []string `json:"BlockedIps,omitempty"`
	BudgetRedirectedCountries             []string `json:"BudgetRedirectedCountries,omitempty"`
	CacheControlBrowserMaxAgeOverride     *int64   `json:"CacheControlBrowserMaxAgeOverride,omitempty"`
	CacheControlMaxAgeOverride            *int64   `json:"CacheControlMaxAgeOverride,omitempty"`
	CacheErrorResponses                   *bool    `json:"CacheErrorResponses,omitempty"`
	ConnectionLimitPerIPCount             *int32   `json:"ConnectionLimitPerIPCount,omitempty"`
	CookieVaryParameters                  []string `json:"CookieVaryParameters,omitempty"`
	DisableCookies                        *bool    `json:"DisableCookies,omitempty"`
	EnableAccessControlOriginHeader       *bool    `json:"EnableAccessControlOriginHeader,omitempty"`
	EnableAvifVary                        *bool    `json:"EnableAvifVary,omitempty"`
	EnableCacheSlice                      *bool    `json:"EnableCacheSlice,omitempty"`
	EnableCookieVary                      *bool    `json:"EnableCookieVary,omitempty"`
	EnableCountryCodeVary                 *bool    `json:"EnableCountryCodeVary,omitempty"`
	EnableGeoZoneAF                       *bool    `json:"EnableGeoZoneAF,omitempty"`
	EnableGeoZoneAsia                     *bool    `json:"EnableGeoZoneASIA,omitempty"`
	EnableGeoZoneEU                       *bool    `json:"EnableGeoZoneEU,omitempty"`
	EnableGeoZoneSA                       *bool    `json:"EnableGeoZoneSA,omitempty"`
	EnableGeoZoneUS                       *bool    `json:"EnableGeoZoneUS,omitempty"`
	EnableHostnameVary                    *bool    `json:"EnableHostnameVary,omitempty"`
	EnableLogging                         *bool    `json:"EnableLogging,omitempty"`
	EnableMobileVary                      *bool    `json:"EnableMobileVary,omitempty"`
	EnableOriginShield                    *bool    `json:"EnableOriginShield,omitempty"`
	EnableQueryStringOrdering             *bool    `json:"EnableQueryStringOrdering,omitempty"`
	EnableSafeHop                         *bool    `json:"EnableSafeHop,omitempty"`
	EnableTLS1                            *bool    `json:"EnableTLS1,omitempty"`
	EnableTLS11                           *bool    `json:"EnableTLS1_1,omitempty"`
	EnableWebPVary                        *bool    `json:"EnableWebPVary,omitempty"`
	ErrorPageCustomCode                   *string  `json:"ErrorPageCustomCode,omitempty"`
	ErrorPageEnableCustomCode             *bool    `json:"ErrorPageEnableCustomCode,omitempty"`
	ErrorPageEnableStatuspageWidget       *bool    `json:"ErrorPageEnableStatuspageWidget,omitempty"`
	ErrorPageStatuspageCode               *string  `json:"ErrorPageStatuspageCode,omitempty"`
	ErrorPageWhitelabel                   *bool    `json:"ErrorPageWhitelabel,omitempty"`
	FollowRedirects                       *bool    `json:"FollowRedirects,omitempty"`
	IgnoreQueryStrings                    *bool    `json:"IgnoreQueryStrings,omitempty"`
	LogForwardingEnabled                  *bool    `json:"LogForwardingEnabled,omitempty"`
	LogForwardingHostname                 *string  `json:"LogForwardingHostname,omitempty"`
	LogForwardingPort                     *int32   `json:"LogForwardingPort,omitempty"`
	LogForwardingToken                    *string  `json:"LogForwardingToken,omitempty"`
	LoggingIPAnonymizationEnabled         *bool    `json:"LoggingIPAnonymizationEnabled,omitempty"`
	LoggingSaveToStorage                  *bool    `json:"LoggingSaveToStorage,omitempty"`
	LoggingStorageZoneID                  *int64   `json:"LoggingStorageZoneId,omitempty"`
	MonthlyBandwidthLimit                 *int64   `json:"MonthlyBandwidthLimit,omitempty"`
	OptimizerAutomaticOptimizationEnabled *bool    `json:"OptimizerAutomaticOptimizationEnabled,omitempty"`
	OptimizerDesktopMaxWidth              *int32   `json:"OptimizerDesktopMaxWidth,omitempty"`
	OptimizerEnableManipulationEngine     *bool    `json:"OptimizerEnableManipulationEngine,omitempty"`
	OptimizerEnableWebP                   *bool    `json:"OptimizerEnableWebP,omitempty"`
	OptimizerEnabled                      *bool    `json:"OptimizerEnabled,omitempty"`
	OptimizerImageQuality                 *int32   `json:"OptimizerImageQuality,omitempty"`
	OptimizerMinifyCSS                    *bool    `json:"OptimizerMinifyCSS,omitempty"`
	OptimizerMinifyJavaScript             *bool    `json:"OptimizerMinifyJavaScript,omitempty"`
	OptimizerMobileImageQuality           *int32   `json:"OptimizerMobileImageQuality,omitempty"`
	OptimizerMobileMaxWidth               *int32   `json:"OptimizerMobileMaxWidth,omitempty"`
	OptimizerWatermarkEnabled             *bool    `json:"OptimizerWatermarkEnabled,omitempty"`
	OptimizerWatermarkMinImageSize        *int32   `json:"OptimizerWatermarkMinImageSize,omitempty"`
	OptimizerWatermarkOffset              *float64 `json:"OptimizerWatermarkOffset,omitempty"`
	OptimizerWatermarkPosition            *int     `json:"OptimizerWatermarkPosition,omitempty"`
	OptimizerWatermarkURL                 *string  `json:"OptimizerWatermarkUrl,omitempty"`
	OriginConnectTimeout                  *int32   `json:"OriginConnectTimeout,omitempty"`
	OriginResponseTimeout                 *int32   `json:"OriginResponseTimeout,omitempty"`
	OriginRetries                         *int32   `json:"OriginRetries,omitempty"`
	OriginRetry5xxResponses               *bool    `json:"OriginRetry5xxResponses,omitempty"`
	OriginRetryConnectionTimeout          *bool    `json:"OriginRetryConnectionTimeout,omitempty"`
	OriginRetryDelay                      *int32   `json:"OriginRetryDelay,omitempty"`
	OriginRetryResponseTimeout            *bool    `json:"OriginRetryResponseTimeout,omitempty"`
	OriginShieldEnableConcurrencyLimit    *bool    `json:"OriginShieldEnableConcurrencyLimit,omitempty"`
	OriginShieldMaxConcurrentRequests     *int32   `json:"OriginShieldMaxConcurrentRequests,omitempty"`
	OriginShieldMaxQueuedRequests         *int32   `json:"OriginShieldMaxQueuedRequests,omitempty"`
	OriginShieldQueueMaxWaitTime          *int32   `json:"OriginShieldQueueMaxWaitTime,omitempty"`
	OriginShieldZoneCode                  *string  `json:"OriginShieldZoneCode,omitempty"`
	OriginURL                             *string  `json:"OriginUrl,omitempty"`
	PermaCacheStorageZoneID               *int64   `json:"PermaCacheStorageZoneId,omitempty"`
	QueryStringVaryParameters             []string `json:"QueryStringVaryParameters,omitempty"`
	RequestLimit                          *int32   `json:"RequestLimit,omitempty"`
	Type                                  *int     `json:"Type,omitempty"`
	UseStaleWhileOffline                  *bool    `json:"UseStaleWhileOffline,omitempty"`
	UseStaleWhileUpdating                 *bool    `json:"UseStaleWhileUpdating,omitempty"`
	VerifyOriginSSL                       *bool    `json:"VerifyOriginSSL,omitempty"`
	WAFEnabled                            *bool    `json:"WAFEnabled,omitempty"`
	WAFEnabledRules                       []int32  `json:"WAFEnabledRules,omitempty"`
	ZoneSecurityEnabled                   *bool    `json:"ZoneSecurityEnabled,omitempty"`
	ZoneSecurityIncludeHashRemoteIP       *bool    `json:"ZoneSecurityIncludeHashRemoteIP,omitempty"`
}

// Update changes the configuration the Pull-Zone with the given ID.
// The updated Pull Zone is returned.
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_updatepullzone
func (s *PullZoneService) Update(ctx context.Context, id int64, opts *PullZoneUpdateOptions) (*PullZone, error) {
	path := fmt.Sprintf("pullzone/%d", id)
	return resourcePostWithResponse[PullZone](
		ctx,
		s.client,
		path,
		opts,
	)
}
