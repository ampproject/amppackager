package pool

import (
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

var (
	monitorMethod = map[string]bool{
		"GET":  true,
		"POST": true,
	}
	regionFailureSensitivity = map[string]bool{
		"HIGH": true,
		"LOW":  true,
	}
	poolOrder = map[string]bool{
		"FIXED":       true,
		"RANDOM":      true,
		"ROUND_ROBIN": true,
	}
	poolRecordState = map[string]bool{
		"NORMAL":   true,
		"ACTIVE":   true,
		"INACTIVE": true,
	}
	responseMethod = map[string]bool{
		"PRIORITY_HUNT": true,
		"RANDOM":        true,
		"ROUND_ROBIN":   true,
	}
	servingPreference = map[string]bool{
		"AUTO_SELECT":    true,
		"SERVE_PRIMARY":  true,
		"SERVE_ALL_FAIL": true,
	}
	dirPoolConflict = map[string]bool{
		"GEO": true,
		"IP":  true,
		"":    true,
	}
)

func ValidatePoolOrder(val string) error {
	if isValidField(val, poolOrder) {
		return nil
	}

	list := []string{"FIXED", "RANDOM", "ROUND_ROBIN"}

	return errors.UnknownDataError("Pool order", val, list)
}

func ValidateRegionFailureSensitivity(val string) error {
	if isValidField(val, regionFailureSensitivity) {
		return nil
	}

	list := []string{"HIGH", "LOW"}

	return errors.UnknownDataError("Pool Region Failure Sensitivity", val, list)
}

func ValidateMonitorMethod(monitor *Monitor) error {
	if monitor == nil || isValidField(monitor.Method, monitorMethod) {
		return nil
	}

	list := []string{"GET", "POST"}

	return errors.UnknownDataError("Pool Monitor Method", monitor.Method, list)
}

func ValidateResponseMethod(val string) error {
	if isValidField(val, responseMethod) {
		return nil
	}

	list := []string{"PRIORITY_HUNT", "RANDOM", "ROUND_ROBIN"}

	return errors.UnknownDataError("Pool Response Method", val, list)
}

func ValidateServingPreference(val string) error {
	if isValidField(val, servingPreference) {
		return nil
	}

	list := []string{"AUTO_SELECT", "SERVE_PRIMARY", "SERVE_ALL_FAIL"}

	return errors.UnknownDataError("Pool Serving Preference", val, list)
}

func ValidateConflictResolve(val string) error {
	if isValidField(val, dirPoolConflict) {
		return nil
	}

	list := []string{"GEO", "IP", ""}

	return errors.UnknownDataError("DIR Pool Resolve Conflict", val, list)
}

func ValidatePoolRecordState(rdataInfoData []*RDataInfo) error {
	for _, rdataInfo := range rdataInfoData {
		if !isValidField(rdataInfo.State, poolRecordState) {
			list := []string{"NORMAL", "ACTIVE", "INACTIVE"}

			return errors.UnknownDataError("Pool record state", rdataInfo.State, list)
		}
	}

	return nil
}

func isValidField(val string, dataMap map[string]bool) bool {
	_, ok := dataMap[val]

	return ok
}
