package rest

import (
	"fmt"
	"net/http"
)

const statsQPSEndpoint = "stats/qps"

// StatsService handles 'stats/qps' endpoint.
type StatsService service

// GetQPS returns current queries per second (QPS) for the account.
// The QPS number is lagged by approximately 30 seconds for statistics collection;
// and the rate is computed over the preceding minute.
func (s *StatsService) GetQPS() (float32, *http.Response, error) {
	return s.getQPS(statsQPSEndpoint)
}

// GetZoneQPS returns current queries per second (QPS) for a specific zone.
// The QPS number is lagged by approximately 30 seconds for statistics collection;
// and the rate is computed over the preceding minute.
func (s *StatsService) GetZoneQPS(zone string) (float32, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", statsQPSEndpoint, zone)
	return s.getQPS(path)
}

// GetRecordQPS returns current queries per second (QPS) for a specific record.
// The QPS number is lagged by approximately 30 seconds for statistics collection;
// and the rate is computed over the preceding minute.
func (s *StatsService) GetRecordQPS(zone, record, t string) (float32, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/%s/%s", statsQPSEndpoint, zone, record, t)
	return s.getQPS(path)
}

func (s *StatsService) getQPS(path string) (float32, *http.Response, error) {
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return 0, nil, err
	}

	var r map[string]float32
	resp, err := s.client.Do(req, &r)

	if err != nil {
		switch err.(type) {
		case *Error:
			switch err.(*Error).Message {
			case "zone not found":
				return 0, nil, ErrZoneMissing
			case "record not found":
				return 0, nil, ErrRecordMissing
			}
		}
		return 0, nil, err
	}

	qps, ok := r["qps"]
	if !ok {
		return 0, nil, fmt.Errorf("could not find 'qps' key in returned data: %v", resp)
	}
	return qps, resp, nil
}
