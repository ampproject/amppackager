package bunny

// EdgeRuleTrigger represents the values of the Trigger field of an EdgeRule.
type EdgeRuleTrigger struct {
	Type                *int     `json:"Type,omitempty"`
	PatternMatches      []string `json:"PatternMatches,omitempty"`
	PatternMatchingType *int     `json:"PatternMatchingType,omitempty"`
	Parameter1          *string  `json:"Parameter1,omitempty"`
}

// Constants for the ActionType fields of an EdgeRule.
const (
	EdgeRuleActionTypeForceSSL int = iota
	EdgeRuleActionTypeRedirect
	EdgeRuleActionTypeOriginURL
	EdgeRuleActionTypeOverrideCacheTime
	EdgeRuleActionTypeBlockRequest
	EdgeRuleActionTypeSetResponseHeader
	EdgeRuleActionTypeSetRequestHeader
	EdgeRuleActionTypeForceDownload
	EdgeRuleActionTypeDisableTokenAuthentication
	EdgeRuleActionTypeEnableTokenAuthentication
	EdgeRuleActionTypeOverrideCacheTimePublic
	EdgeRuleActionTypeIgnoreQueryString
	EdgeRuleActionTypeDisableOptimizer
	EdgeRuleActionTypeForceCompression
	EdgeRuleActionTypeSetStatusCode
	EdgeRuleActionTypeBypassPermaCache
)

// Constants for the Type field of an EdgeRuleTrigger.
const (
	EdgeRuleTriggerTypeURL int = iota
	EdgeRuleTriggerTypeRequestHeader
	EdgeRuleTriggerTypeResponseHeader
	EdgeRuleTriggerTypeURLExtension
	EdgeRuleTriggerTypeCountryCode
	EdgeRuleTriggerTypeRemoteIP
	EdgeRuleTriggerTypeURLQueryString
	EdgeRuleTriggerTypeRandomChance
	EdgeRuleTriggerTypeStatusCode
	EdgeRuleTriggerTypeRequestMethod
)
