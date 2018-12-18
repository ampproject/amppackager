package rtv

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	defaultHTTPTimeout  = 1 * time.Minute
	defaultPollInterval = 1 * time.Hour
)

// not a const for testing purposes
var rtvHost = "https://cdn.ampproject.org"

// rtvData stores the AMP runtime version number and the CSS for that version
// Note: fields must be exported for json unmarshaling.
type rtvData struct {
	RTV                   string `json:"ampRuntimeVersion"`
	CSSURL                string `json:"ampCssUrl"`
	CanaryPercentage, CSS string
}

type RTVCache struct {
	d  *rtvData
	c  http.Client
	lk sync.Mutex
	stop chan struct{}
}

// New returns a new cache for storing AMP runtime values, or an
// error if there was a problem initializing. To have it auto-refresh,
// call StartCron().
func New() (*RTVCache, error) {
	r := &RTVCache{c: http.Client{Timeout: defaultHTTPTimeout}, d: &rtvData{}, stop: make(chan struct{})}
	if err := r.poll(); err != nil {
		return nil, err
	}
	return r, nil
}

// StartCron starts a cron job to re-fill the RTVCache hourly.
func (r *RTVCache) StartCron() {
	go func() {
		ticker := time.NewTicker(defaultPollInterval)

		for {
			select {
			case <-ticker.C:
				r.poll()  // Ignores error return.
			case <-r.stop:
				ticker.Stop()
				return
			}
		}
	}()
}

// StopCron stops the cron job.
func (r *RTVCache) StopCron() {
	r.stop <- struct{}{}
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
	// Minimal validation of expected values.
	if d.RTV == "" {
		return nil, errors.Errorf("Could not unmarshal RTV value from %s", b)
	}
	if d.CSSURL == "" {
		return nil, errors.Errorf("Could not unmarshal CSS URL value from %s", b)
	}
	if _, err := url.Parse(d.CSSURL); err != nil {
		return nil, errors.Wrapf(err, "Error parsing CSS URL %s", d.CSSURL)
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
