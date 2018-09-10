package amppackager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	// TODO(twifkak): Replace this with time.NewTicker.
	"github.com/robfig/cron"
)

const (
	defaultHTTPTimeout  = 1 * time.Minute
	defaultPollInterval = "@every 1h"
)

// not a const for testing purposes
var rtvHost = "https://cdn.ampproject.org"

// rtvData stores the AMP runtime version number and the CSS for that version
// Note: fields must be exported for json unmarshaling.
type rtvData struct {
	RTV              string `json:"ampRuntimeVersion"`
	CSSURL           string `json:"ampCssUrl"`
	CanaryPercentage string `json:"canaryPercentage"`
	CSS              string
}

type RTVCache struct {
	d  *rtvData
	c  http.Client
	lk sync.Mutex
	cr *cron.Cron
}

// NewRTV returns a new cache for storing AMP runtime values, or an
// error if there was a problem initializing. To have it auto-refresh,
// call StartCron().
func NewRTV() (*RTVCache, error) {
	r := &RTVCache{c: http.Client{Timeout: defaultHTTPTimeout}, d: &rtvData{}, cr: cron.New()}
	if err := r.poll(); err != nil {
		return nil, err
	}
	return r, nil
}

// StartCron starts a cron job to periodically re-fill the RTVCache,
// based on the given cron expression format. If empty, defaults to hourly.
func (r *RTVCache) StartCron(spec string) error {
	if spec == "" {
		spec = defaultPollInterval
	}
	if err := r.cr.AddFunc(spec, func() { r.poll() }); err != nil {
		return err
	}
	r.cr.Start()
	return nil
}

// StopCron stops the cron job.
func (r *RTVCache) StopCron() {
	r.cr.Stop()
}

// getRTVData returns the cached rtvData.
func (r *RTVCache) getRTVData() *rtvData {
	r.lk.Lock()
	defer r.lk.Unlock()
	return r.d
}

// GetRTV returns the cached value for the runtime version.
func (r *RTVCache) GetRTV() string {
	return r.getRTVData().RTV
}

// GetCSS returns the cached value for the inline CSS.
func (r *RTVCache) GetCSS() string {
	return r.getRTVData().CSS
}

// poll attempts to re-populate the RTVCache, returning an error if there
// were any problems.
func (r *RTVCache) poll() error {
	// Fetch the runtime metadata
	d, err := getMetadata(r)
	if err != nil {
		return err
	}

	// If the value is unchanged, skip CSS call
	if d.RTV == r.GetRTV() {
		return nil
	}

	// Fetch the CSS payload
	var b []byte
	b, err = getRTVBody(r.c, d.CSSURL)
	if err != nil {
		return err
	}
	d.CSS = string(b)

	// No errors, update cache.
	r.lk.Lock()
	defer r.lk.Unlock()
	r.d = d
	return nil
}

// getMetadata fetches the JSON from the metadata endpoint and returns the
// data parsed as a struct.
func getMetadata(r *RTVCache) (*rtvData, error) {
	// Fetch the runtime metadata json
	b, err := getRTVBody(r.c, rtvHost+"/rtv/metadata")
	if err != nil {
		return nil, err
	}
	var d rtvData
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// getRTVBody returns the body contents of the given url, or an error
// if there was problem.
func getRTVBody(c http.Client, url string) ([]byte, error) {
	log.Printf("Fetching URL: %q\n", url)
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("Non-200 response fetching %s, %+v", url, resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
