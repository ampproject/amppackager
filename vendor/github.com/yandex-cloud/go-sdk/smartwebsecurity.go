package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/smartwebsecurity"
	advanced_rate_limiter "github.com/yandex-cloud/go-sdk/gen/smartwebsecurity/arl"
	waf "github.com/yandex-cloud/go-sdk/gen/smartwebsecurity/waf"
)

const (
	SmartWebSecurityServiceID Endpoint = "smart-web-security"
)

// SmartWebSecurity returns SmartWebSecurity object that is used to operate with security profiles
func (sdk *SDK) SmartWebSecurity() *smartwebsecurity.SmartWebSecurity {
	return smartwebsecurity.NewSmartWebSecurity(sdk.getConn(SmartWebSecurityServiceID))
}

// SmartWebSecurityWaf returns SmartWebSecurityWaf object that is used to operate with WAF profiles
func (sdk *SDK) SmartWebSecurityWaf() *waf.SmartWebSecurityWaf {
	return waf.NewSmartWebSecurityWaf(sdk.getConn(SmartWebSecurityServiceID))
}

// SmartWebSecurityArl returns SmartWebSecurityArl object that is used to operate with AdvancedRateLimiter profiles
func (sdk *SDK) SmartWebSecurityArl() *advanced_rate_limiter.SmartWebSecurityArl {
	return advanced_rate_limiter.NewSmartWebSecurityArl(sdk.getConn(SmartWebSecurityServiceID))
}
