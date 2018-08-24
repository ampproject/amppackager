package amppackager

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron"
)

// die is an alias for log.Fatalf, which can be overridden for testing purposes.
var die = log.Fatalf

// not a const solely for testing purposes (be able to override this value).
var rtvHost = "https://cdn.ampproject.org"

// rtvCacheStruct stores the AMP runtime version number and the CSS for that version
type rtvCacheStruct struct {
	RTV, CSS string
}

var RTVCache = new(rtvCacheStruct)
var c = cron.New()
var rtvClient = http.Client{Timeout: 60 * time.Second}

// StartCron starts an hourly cron job to periodically re-fill the RTVCache.
func StartCron() error {
	if err := c.AddFunc("@every 1h", rtvPoll); err != nil {
		return err
	}
	// Initialize the cache. Then cron will trigger after every interval.
	rtvPoll()
	c.Start()
	return nil
}

// StopCron stops the cron job.
func StopCron() {
	c.Stop()
}

// rtvPoll attempts to re-populate the RTVCache. If this is the very first time,
// and there are any errors, this will die fatally.
func rtvPoll() {
	// Make a copy for transactional state.
	newCache := *RTVCache

	// Decide to die if this is the very first time initializing the value.
	shouldDie := RTVCache.RTV == ""
	maybeDie := func(err interface{}) {
		if shouldDie {
			die("%+v", err)
		}
		log.Print(err)
	}

	// Fetch the runtime version number
	// TODO(angielin): This is a temporary endpoint. Migrate to metadata
	// endpoint when ready.
	var err error
	if newCache.RTV, err = getRTVBody(rtvHost + "/v0/version.txt"); err != nil {
		// If there is a problem getting the RTV value, there is no need to
		// get the CSS
		maybeDie(err)
		return
	}
	// Pad to width of 15
	newCache.RTV = fmt.Sprintf("%015s", newCache.RTV)

	// If the value is the same, skip CSS call
	if newCache.RTV == RTVCache.RTV {
		return
	}

	// Fetch the CSS payload
	if newCache.CSS, err = getRTVBody(rtvHost + "/rtv/" + newCache.RTV + "/v0.css"); err != nil {
		// If there was a problem getting CSS, abort and don't write new cache value.
		maybeDie(err)
		return
	}
	RTVCache = &newCache
}

// getRTVBody returns the body contents of the given url, or an error if there was problem.
func getRTVBody(url string) (string, error) {
	log.Printf("Fetching URL: %q\n", url)
	resp, err := rtvClient.Get(url)
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
