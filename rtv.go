package amppackager

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron"
)

const (
	defaultHTTPTimeout = 1 * time.Minute
)

// not a const for testing purposes
var rtvHost = "https://cdn.ampproject.org"

// rtvData stores the AMP runtime version number and the CSS for that version
type rtvData struct {
	rtv, css string
}

type RTVCache struct {
	d  *rtvData
	c  http.Client
	lk sync.Mutex
	// TODO(alin04): Switch to NewTicker
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

// StartCron starts an hourly cron job to periodically re-fill the RTVCache.
func (r *RTVCache) StartCron() error {
	if err := r.cr.AddFunc("@every 1h", func() { r.poll() }); err != nil {
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
	return r.d.rtv
}

// GetRTV returns the cached value for the runtime version.
func (r *RTVCache) GetRTV() string {
	return getRTVData().rtv
}

// GetCSS returns the cached value for the inline CSS.
func (r *RTVCache) GetCSS() string {
	return getRTVData().css
}

// poll attempts to re-populate the RTVCache, returning an error if there
// were any problems.
func (r *RTVCache) poll() error {
	// Make a copy for atomicity.
	d := *getRTVData()

	// Fetch the runtime version number
	// TODO(alin04): This is a temporary endpoint. Migrate to metadata
	// endpoint when ready.
	var err error
	if d.rtv, err = r.getRTVBody(rtvHost + "/v0/version.txt"); err != nil {
		// If there is a problem getting the RTV value, skip getting CSS.
		return err
	}
	// Pad to width of 15
	d.rtv = fmt.Sprintf("%015s", d.rtv)

	// If the value is unchanged, skip CSS call
	if d.rtv == r.GetRTV() {
		return nil
	}

	// Fetch the CSS payload
	if d.css, err = r.getRTVBody(rtvHost + "/rtv/" + d.rtv + "/v0.css"); err != nil {
		// If there was a problem getting CSS, abort and don't write
		// new cache value.
		return err
	}
	r.lk.Lock()
	defer r.lk.Unlock()
	r.d = &d
	return nil
}

// getRTVBody returns the body contents of the given url, or an error
// if there was problem.
func (r *RTVCache) getRTVBody(url string) (string, error) {
	log.Printf("Fetching URL: %q\n", url)
	resp, err := r.c.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.Errorf("Non-200 response fetching %s, %+v", url, resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
