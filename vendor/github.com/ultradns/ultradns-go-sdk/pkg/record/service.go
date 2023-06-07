package record

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
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

	if err := validatePoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		return nil, errors.CreateError(serviceName, rrSetKey.RecordID(), err)
	}

	return res, nil
}

func (s *Service) Read(rrSetKey *rrset.RRSetKey) (*http.Response, *rrset.ResponseList, error) {
	rrSetTarget := &rrset.RRSet{}

	setPoolProfile(rrSetKey.PType, rrSetTarget)

	rrSetResList := &rrset.ResponseList{}
	rrSetResList.RRSets = make([]*rrset.RRSet, 1)
	rrSetResList.RRSets[0] = rrSetTarget
	target := client.Target(rrSetResList)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, rrSetKey.RecordURI(), nil, target)

	if err != nil {
		return nil, nil, errors.ReadError(serviceName, rrSetKey.RecordID(), err)
	}

	rrsetList := target.Data.(*rrset.ResponseList)

	profile := rrsetList.RRSets[0].Profile

	if profile != nil && getPoolSchema(rrSetKey.PType) != profile.GetContext() {
		return nil, nil, errors.ResourceTypeNotFoundError(serviceName, rrSetKey.PType, rrSetKey.RecordID())
	}

	return res, rrsetList, nil
}

func (s *Service) Update(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	if err := validatePoolProfile(rrSet); err != nil {
		return nil, err
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		return nil, errors.UpdateError(serviceName, rrSetKey.RecordID(), err)
	}

	return res, nil
}

func (s *Service) PartialUpdate(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, rrSetKey.RecordURI(), rrSet, target)

	if err != nil {
		return nil, errors.PartialUpdateError(serviceName, rrSetKey.RecordID(), err)
	}

	return res, nil
}

func (s *Service) Delete(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, rrSetKey.RecordURI(), nil, target)

	if err != nil {
		return nil, errors.DeleteError(serviceName, rrSetKey.RecordID(), err)
	}

	return res, nil
}
