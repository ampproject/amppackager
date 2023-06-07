package record

import (
	"github.com/ultradns/ultradns-go-sdk/pkg/record/dirpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sfpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/slbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/tcpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func getPoolSchema(poolType string) string {
	var poolSchema = map[string]string{
		pool.RD:  rdpool.Schema,
		pool.SF:  sfpool.Schema,
		pool.SLB: slbpool.Schema,
		pool.SB:  sbpool.Schema,
		pool.TC:  tcpool.Schema,
		pool.DIR: dirpool.Schema,
	}

	return poolSchema[poolType]
}

func setPoolProfile(profileType string, rrSet *rrset.RRSet) {
	switch profileType {
	case pool.RD:
		rrSet.Profile = &rdpool.Profile{}
	case pool.SF:
		rrSet.Profile = &sfpool.Profile{}
	case pool.SLB:
		rrSet.Profile = &slbpool.Profile{}
	case pool.SB:
		rrSet.Profile = &sbpool.Profile{}
	case pool.TC:
		rrSet.Profile = &tcpool.Profile{}
	case pool.DIR:
		rrSet.Profile = &dirpool.Profile{}
	}
}

func validatePoolProfile(rrSet *rrset.RRSet) error {
	if rrSet.Profile == nil {
		return nil
	}

	rrSet.Profile.SetContext()

	switch rrSet.Profile.GetContext() {
	case rdpool.Schema:
		return pool.ValidatePoolOrder(rrSet.Profile.(*rdpool.Profile).Order)
	case sfpool.Schema:
		return validateSFPoolProfile(rrSet.Profile.(*sfpool.Profile))
	case slbpool.Schema:
		return validateSLBPoolProfile(rrSet.Profile.(*slbpool.Profile))
	case sbpool.Schema:
		return validateSBPoolProfile(rrSet.Profile.(*sbpool.Profile))
	case tcpool.Schema:
		return pool.ValidatePoolRecordState(rrSet.Profile.(*tcpool.Profile).RDataInfo)
	case dirpool.Schema:
		return pool.ValidateConflictResolve(rrSet.Profile.(*dirpool.Profile).ConflictResolve)
	}

	return nil
}

func validateSFPoolProfile(profile *sfpool.Profile) error {
	if err := pool.ValidateMonitorMethod(profile.Monitor); err != nil {
		return err
	}

	if err := pool.ValidateRegionFailureSensitivity(profile.RegionFailureSensitivity); err != nil {
		return err
	}

	return nil
}

func validateSLBPoolProfile(profile *slbpool.Profile) error {
	if err := pool.ValidateMonitorMethod(profile.Monitor); err != nil {
		return err
	}

	if err := pool.ValidateRegionFailureSensitivity(profile.RegionFailureSensitivity); err != nil {
		return err
	}

	if err := pool.ValidateResponseMethod(profile.ResponseMethod); err != nil {
		return err
	}

	if err := pool.ValidateServingPreference(profile.ServingPreference); err != nil {
		return err
	}

	return nil
}

func validateSBPoolProfile(profile *sbpool.Profile) error {
	if err := pool.ValidatePoolOrder(profile.Order); err != nil {
		return err
	}

	if err := pool.ValidatePoolRecordState(profile.RDataInfo); err != nil {
		return err
	}

	return nil
}
