package helper

import (
	"fmt"
	"net/url"
	"strings"
)

func appendRootDot(zoneName string) string {
	return fmt.Sprintf("%s.", zoneName)
}

func GetZoneFQDN(zoneName string) string {
	result := zoneName
	if len(zoneName) > 0 {
		if lastChar := zoneName[len(zoneName)-1]; lastChar != '.' {
			result = appendRootDot(zoneName)
		}
	}

	return strings.ToLower(result)
}

func GetOwnerFQDN(owner, zone string) string {
	result := GetZoneFQDN(owner)
	if !strings.Contains(GetZoneFQDN(owner), strings.ToLower(zone)) {
		result = appendRootDot(owner) + GetZoneFQDN(zone)
	}

	return strings.ToLower(result)
}

func GetRecordTypeFullString(key string) string {
	var rrTypes = map[string]string{
		"A":         "A (1)",
		"1":         "A (1)",
		"NS":        "NS (2)",
		"2":         "NS (2)",
		"CNAME":     "CNAME (5)",
		"5":         "CNAME (5)",
		"SOA":       "SOA (6)",
		"6":         "SOA (6)",
		"PTR":       "PTR (12)",
		"12":        "PTR (12)",
		"HINFO":     "HINFO (13)",
		"13":        "HINFO (13)",
		"MX":        "MX (15)",
		"15":        "MX (15)",
		"TXT":       "TXT (16)",
		"16":        "TXT (16)",
		"RP":        "RP (17)",
		"17":        "RP (17)",
		"AAAA":      "AAAA (28)",
		"28":        "AAAA (28)",
		"SRV":       "SRV (33)",
		"33":        "SRV (33)",
		"NAPTR":     "NAPTR (35)",
		"35":        "NAPTR (35)",
		"DS":        "DS (43)",
		"43":        "DS (43)",
		"SSHFP":     "SSHFP (44)",
		"44":        "SSHFP (44)",
		"TLSA":      "TLSA (52)",
		"52":        "TLSA (52)",
		"SPF":       "SPF (99)",
		"99":        "SPF (99)",
		"CAA":       "CAA (257)",
		"257":       "CAA (257)",
		"APEXALIAS": "APEXALIAS (65282)",
		"65282":     "APEXALIAS (65282)",
	}

	return rrTypes[key]
}

func GetRecordTypeString(key string) string {
	var rrTypes = map[string]string{
		"A (1)":             "A",
		"NS (2)":            "NS",
		"CNAME (5)":         "CNAME",
		"SOA (6)":           "SOA",
		"PTR (12)":          "PTR",
		"HINFO (13)":        "HINFO",
		"MX (15)":           "MX",
		"TXT (16)":          "TXT",
		"RP (17)":           "RP",
		"AAAA (28)":         "AAAA",
		"SRV (33)":          "SRV",
		"NAPTR (35)":        "NAPTR",
		"DS (43)":           "DS",
		"SSHFP (44)":        "SSHFP",
		"TLSA (52)":         "TLSA",
		"SPF (99)":          "SPF",
		"CAA (257)":         "CAA",
		"APEXALIAS (65282)": "APEXALIAS",
	}

	return rrTypes[key]
}

func GetRecordTypeNumber(key string) string {
	var rrTypes = map[string]string{
		"A (1)":             "1",
		"NS (2)":            "2",
		"CNAME (5)":         "5",
		"SOA (6)":           "6",
		"PTR (12)":          "12",
		"HINFO (13)":        "13",
		"MX (15)":           "15",
		"TXT (16)":          "16",
		"RP (17)":           "17",
		"AAAA (28)":         "28",
		"SRV (33)":          "33",
		"NAPTR (35)":        "35",
		"DS (43)":           "43",
		"SSHFP (44)":        "44",
		"TLSA (52)":         "52",
		"SPF (99)":          "99",
		"CAA (257)":         "257",
		"APEXALIAS (65282)": "65282",
	}

	return rrTypes[key]
}

func GetAccountName(id string) string {
	geoAccount := strings.Split(id, ":")
	return geoAccount[1]
}

func GetAccountNameFromURI(uri string) string {
	geoAccount := strings.Split(uri, "/")
	return geoAccount[1]
}

func GetDirGroupURI(groupID, groupType string) string {
	groupID = url.PathEscape(groupID)
	groupData := strings.Split(groupID, ":")

	return fmt.Sprintf("accounts/%s/dirgroups/%s/%s", groupData[1], groupType, groupData[0])
}

func GetDirGroupListURI(accountName, groupType string) string {
	accountName = url.PathEscape(accountName)

	return fmt.Sprintf("accounts/%s/dirgroups/%s", accountName, groupType)
}
