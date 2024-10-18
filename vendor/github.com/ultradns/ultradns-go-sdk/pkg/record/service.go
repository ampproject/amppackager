package record

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceName = "Record"

type Service struct {
	c *client.Client
}

func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, errors.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) Create(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s create started", serviceName)

	if err := validatePoolProfile(rrSet); err != nil {
		s.c.Error("%s create failed with error: %v", serviceName, err)
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		s.c.Error("%s create failed with error: %v", serviceName, err)
		return res, errors.CreateError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("%s create completed successfully", serviceName)

	return res, nil
}

func (s *Service) Read(rrSetKey *rrset.RRSetKey) (*http.Response, *rrset.ResponseList, error) {
	target := client.Target(&rrset.ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s read started", serviceName)

	res, err := s.c.Do(http.MethodGet, rrSetKey.RecordURI(), nil, target)

	if err != nil {
		s.c.Error("%s read failed with error: %v", serviceName, err)
		return res, nil, errors.ReadError(serviceName, rrSetKey.RecordID(), err)
	}

	rrsetList := target.Data.(*rrset.ResponseList)

	if len(rrsetList.RRSets) != 1 {
		s.c.Error("%s read failed with error: multiple resource for the filter applied", serviceName)
		return nil, nil, errors.MultipleResourceFoundError(serviceName, rrSetKey.RecordID())
	}

	profile := rrsetList.RRSets[0].Profile

	if profile != nil && getPoolSchema(rrSetKey.PType) != profile.GetContext() {
		s.c.Error("%s read failed with error: queried pool data not available for the owner name", serviceName)
		return nil, nil, errors.ResourceTypeNotFoundError(serviceName, rrSetKey.PType, rrSetKey.RecordID())
	}

	s.c.Trace("%s read completed successfully", serviceName)

	return res, rrsetList, nil
}

func (s *Service) Update(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s update started", serviceName)

	if err := validatePoolProfile(rrSet); err != nil {
		s.c.Error("%s update failed with error: %v", serviceName, err)
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		s.c.Error("%s update failed with error: %v", serviceName, err)
		return res, errors.UpdateError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("%s update completed successfully", serviceName)

	return res, nil
}

func (s *Service) PartialUpdate(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s partial update started", serviceName)

	res, err := s.c.Do(http.MethodPatch, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		s.c.Error("%s partial update failed with error: %v", serviceName, err)
		return res, errors.PartialUpdateError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("%s partial update completed successfully", serviceName)

	return res, nil
}

func (s *Service) Delete(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s delete started", serviceName)

	res, err := s.c.Do(http.MethodDelete, rrSetKey.RecordURI(), nil, target)

	if err != nil {
		s.c.Error("%s delete failed with error: %v", serviceName, err)
		return res, errors.DeleteError(serviceName, rrSetKey.RecordID(), err)
	}

	s.c.Trace("%s delete completed successfully", serviceName)

	return res, nil
}

func (s *Service) List(rrSetKey *rrset.RRSetKey, queryInfo *helper.QueryInfo) (*http.Response, *rrset.ResponseList, error) {
	target := client.Target(&rrset.ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	s.c.Trace("%s list started", serviceName)

	res, err := s.c.Do(http.MethodGet, rrSetKey.RecordURI()+queryInfo.URI(), nil, target)

	if err != nil {
		s.c.Error("%s list failed with error: %v", serviceName, err)
		return res, nil, errors.ListError(serviceName, rrSetKey.RecordID()+queryInfo.URI(), err)
	}

	rrsetList := target.Data.(*rrset.ResponseList)

	s.c.Trace("%s list completed successfully", serviceName)

	return res, rrsetList, nil
}
