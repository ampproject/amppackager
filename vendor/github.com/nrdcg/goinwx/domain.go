package goinwx

import (
	"errors"
	"strconv"
	"time"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

const (
	methodDomainCheck       = "domain.check"
	methodDomainCreate      = "domain.create"
	methodDomainDelete      = "domain.delete"
	methodDomainGetPrices   = "domain.getPrices"
	methodDomainGetRules    = "domain.getRules"
	methodDomainInfo        = "domain.info"
	methodDomainList        = "domain.list"
	methodDomainLog         = "domain.log"
	methodDomainPush        = "domain.push"
	methodDomainRenew       = "domain.renew"
	methodDomainRestore     = "domain.restore"
	methodDomainStats       = "domain.stats"
	methodDomainTrade       = "domain.trade"
	methodDomainTransfer    = "domain.transfer"
	methodDomainTransferOut = "domain.transferOut"
	methodDomainUpdate      = "domain.update"
	methodDomainWhois       = "domain.whois"
)

// DomainService API access to Domain.
type DomainService service

// Check checks domains.
func (s *DomainService) Check(domains []string) ([]DomainCheckResponse, error) {
	req := s.client.NewRequest(methodDomainCheck, map[string]interface{}{
		"domain": domains,
		"wide":   "2",
	})

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	root := new(domainCheckResponseRoot)
	err = mapstructure.Decode(resp, &root)
	if err != nil {
		return nil, err
	}

	return root.Domains, nil
}

// GetPrices gets TLDS prices.
func (s *DomainService) GetPrices(tlds []string) ([]DomainPriceResponse, error) {
	req := s.client.NewRequest(methodDomainGetPrices, map[string]interface{}{
		"tld": tlds,
		"vat": false,
	})

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	root := new(domainPriceResponseRoot)
	err = mapstructure.Decode(resp, &root)
	if err != nil {
		return nil, err
	}

	return root.Prices, nil
}

// Register registers a domain.
func (s *DomainService) Register(request *DomainRegisterRequest) (*DomainRegisterResponse, error) {
	req := s.client.NewRequest(methodDomainCreate, structs.Map(request))

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result DomainRegisterResponse
	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Delete deletes a domain.
func (s *DomainService) Delete(domain string, scheduledDate time.Time) error {
	req := s.client.NewRequest(methodDomainDelete, map[string]interface{}{
		"domain": domain,
		"scDate": scheduledDate.Format(time.RFC3339),
	})

	_, err := s.client.Do(req)
	return err
}

// Info gets information about a domain.
func (s *DomainService) Info(domain string, roID int) (*DomainInfoResponse, error) {
	req := s.client.NewRequest(methodDomainInfo, map[string]interface{}{
		"domain": domain,
		"wide":   "2",
	})
	if roID != 0 {
		req.Args["roId"] = roID
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result DomainInfoResponse
	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// List lists domains.
func (s *DomainService) List(request *DomainListRequest) (*DomainList, error) {
	if request == nil {
		return nil, errors.New("request can't be nil")
	}

	requestMap := structs.Map(request)
	requestMap["wide"] = "2"

	req := s.client.NewRequest(methodDomainList, requestMap)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	// This is a (temporary) workaround to convert the API response from string to int.
	// The code only applies when string is being return, otherwise it's being skipped.
	// As per docs at https://www.inwx.com/en/help/apidoc/f/ch02s08.html#domain.list we should get int here, but apparently it's not the case.
	// Ticket 10026265 with INWX was raised.
	if countStr, ok := resp["count"].(string); ok {
		// If we land here, we got string back, but we expect int.
		// Converting value to int and writing it to the response.
		if num, ok := strconv.Atoi(countStr); ok == nil {
			resp["count"] = num
		}
	}

	var result DomainList
	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Whois information about a domains.
func (s *DomainService) Whois(domain string) (string, error) {
	req := s.client.NewRequest(methodDomainWhois, map[string]interface{}{
		"domain": domain,
	})

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}

	var result map[string]string
	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return "", err
	}

	return result["whois"], nil
}

// Update updates domain information.
func (s *DomainService) Update(request *DomainUpdateRequest) (float32, error) {
	req := s.client.NewRequest(methodDomainUpdate, structs.Map(request))

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result DomainUpdateResponse
	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return 0, err
	}

	return result.Price, nil
}

type domainCheckResponseRoot struct {
	Domains []DomainCheckResponse `mapstructure:"domain"`
}

// DomainCheckResponse API model.
type DomainCheckResponse struct {
	Available   int     `mapstructure:"avail"`
	Status      string  `mapstructure:"status"`
	Name        string  `mapstructure:"name"`
	Domain      string  `mapstructure:"domain"`
	TLD         string  `mapstructure:"tld"`
	CheckMethod string  `mapstructure:"checkmethod"`
	Price       float32 `mapstructure:"price"`
	CheckTime   float32 `mapstructure:"checktime"`
}

type domainPriceResponseRoot struct {
	Prices []DomainPriceResponse `mapstructure:"price"`
}

// DomainPriceResponse API model.
type DomainPriceResponse struct {
	Tld                 string  `mapstructure:"tld"`
	Currency            string  `mapstructure:"currency"`
	CreatePrice         float32 `mapstructure:"createPrice"`
	MonthlyCreatePrice  float32 `mapstructure:"monthlyCreatePrice"`
	TransferPrice       float32 `mapstructure:"transferPrice"`
	RenewalPrice        float32 `mapstructure:"renewalPrice"`
	MonthlyRenewalPrice float32 `mapstructure:"monthlyRenewalPrice"`
	UpdatePrice         float32 `mapstructure:"updatePrice"`
	TradePrice          float32 `mapstructure:"tradePrice"`
	TrusteePrice        float32 `mapstructure:"trusteePrice"`
	MonthlyTrusteePrice float32 `mapstructure:"monthlyTrusteePrice"`
	CreatePeriod        int     `mapstructure:"createPeriod"`
	TransferPeriod      int     `mapstructure:"transferPeriod"`
	RenewalPeriod       int     `mapstructure:"renewalPeriod"`
	TradePeriod         int     `mapstructure:"tradePeriod"`
}

// DomainRegisterRequest API model.
type DomainRegisterRequest struct {
	Domain        string   `structs:"domain"`
	Period        string   `structs:"period,omitempty"`
	Registrant    int      `structs:"registrant"`
	Admin         int      `structs:"admin"`
	Tech          int      `structs:"tech"`
	Billing       int      `structs:"billing"`
	Nameservers   []string `structs:"ns,omitempty"`
	TransferLock  string   `structs:"transferLock,omitempty"`
	RenewalMode   string   `structs:"renewalMode,omitempty"`
	WhoisProvider string   `structs:"whoisProvider,omitempty"`
	WhoisURL      string   `structs:"whoisUrl,omitempty"`
	ScDate        string   `structs:"scDate,omitempty"`
	ExtDate       string   `structs:"extDate,omitempty"`
	Asynchron     string   `structs:"asynchron,omitempty"`
	Voucher       string   `structs:"voucher,omitempty"`
	Testing       string   `structs:"testing,omitempty"`
}

// DomainRegisterResponse API model.
type DomainRegisterResponse struct {
	RoID     int     `mapstructure:"roId"`
	Price    float32 `mapstructure:"price"`
	Currency string  `mapstructure:"currency"`
}

// DomainInfoResponse API model.
type DomainInfoResponse struct {
	RoID         int                `mapstructure:"roId"`
	Domain       string             `mapstructure:"domain"`
	DomainAce    string             `mapstructure:"domainAce"`
	Period       string             `mapstructure:"period"`
	CrDate       time.Time          `mapstructure:"crDate"`
	ExDate       time.Time          `mapstructure:"exDate"`
	UpDate       time.Time          `mapstructure:"upDate"`
	ReDate       time.Time          `mapstructure:"reDate"`
	ScDate       time.Time          `mapstructure:"scDate"`
	TransferLock int                `mapstructure:"transferLock"`
	Status       string             `mapstructure:"status"`
	AuthCode     string             `mapstructure:"authCode"`
	RenewalMode  string             `mapstructure:"renewalMode"`
	TransferMode string             `mapstructure:"transferMode"`
	Registrant   int                `mapstructure:"registrant"`
	Admin        int                `mapstructure:"admin"`
	Tech         int                `mapstructure:"tech"`
	Billing      int                `mapstructure:"billing"`
	Nameservers  []string           `mapstructure:"ns"`
	NoDelegation int                `mapstructure:"noDelegation"`
	Contacts     map[string]Contact `mapstructure:"contact"`
}

// Contact API model.
type Contact struct {
	RoID          int    `mapstructure:"roId"`
	ID            string `mapstructure:"id"`
	Type          string `mapstructure:"type"`
	Name          string `mapstructure:"name"`
	Org           string `mapstructure:"org"`
	Street        string `mapstructure:"street"`
	City          string `mapstructure:"city"`
	PostalCode    string `mapstructure:"pc"`
	StateProvince string `mapstructure:"sp"`
	Country       string `mapstructure:"cc"`
	Phone         string `mapstructure:"voice"`
	Fax           string `mapstructure:"fax"`
	Email         string `mapstructure:"email"`
	Remarks       string `mapstructure:"remarks"`
	Protection    string `mapstructure:"protection"`
}

// DomainListRequest API model.
type DomainListRequest struct {
	Domain       string `structs:"domain,omitempty"`
	RoID         int    `structs:"roId,omitempty"`
	Status       int    `structs:"status,omitempty"`
	Registrant   int    `structs:"registrant,omitempty"`
	Admin        int    `structs:"admin,omitempty"`
	Tech         int    `structs:"tech,omitempty"`
	Billing      int    `structs:"billing,omitempty"`
	RenewalMode  int    `structs:"renewalMode,omitempty"`
	TransferLock int    `structs:"transferLock,omitempty"`
	NoDelegation int    `structs:"noDelegation,omitempty"`
	Tag          int    `structs:"tag,omitempty"`
	Order        int    `structs:"order,omitempty"`
	Page         int    `structs:"page,omitempty"`
	PageLimit    int    `structs:"pagelimit,omitempty"`
}

// DomainList API model.
type DomainList struct {
	Count   int
	Domains []DomainInfoResponse `mapstructure:"domain"`
}

// DomainUpdateRequest API model.
type DomainUpdateRequest struct {
	Domain       string   `structs:"domain"`
	Nameservers  []string `structs:"ns,omitempty"`
	TransferLock int      `structs:"transferLock,omitempty"`
	RenewalMode  string   `structs:"renewalMode,omitempty"`
	TransferMode string   `structs:"transferMode,omitempty"`
	// unsupported fields:
	// registrant:	New owner contact handle id	(int, false)
	// admin:	New administrative contact handle id (int, false)
	// tech:	New technical contact handle id	(int, false)
	// billing:	New billing contact handle id	(int, false)
	// authCode:	Authorization code (if supported)	(text64, false)
	// scDate:	Time of scheduled execution	(timestamp, false)
	// whoisProvider:	Whois provider	(token0255, false)
	// whoisUrl:	Whois url	(token0255, false)
	// extData:	Domain extra data	(extData, false)
	// asynchron:	Asynchron domain update	boolean	(false, false)
	// testing:	Execute command in testing mode	boolean	(false, false)
}

// DomainUpdateResponse API model.
type DomainUpdateResponse struct {
	Price float32 `mapstructure:"price"`
}
