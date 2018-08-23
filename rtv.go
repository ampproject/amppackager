package amppackager

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

// log alias, which can be overridden for testing purposes.
var die = log.Fatalf

// not a const solely for testing purposes (be able to override this value).
var rtvHost = "https://cdn.ampproject.org"

// RTVCache stores the AMP runtime version number and the CSS for that version
type RTVCacheStruct struct {
	RTV, CSS string
}

var c = cron.New()
var RTVCache = new(RTVCacheStruct)

// StartCron starts an hourly cron job to periodically re-fill the RTVCache.
func StartCron() error {
	if err := c.AddFunc("@every 1h", rtvPoll); err != nil {
		return err
	}
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
	// Fetch the runtime version number
	getRtv(&RTVCache.RTV, rtvHost+"/v0/version.txt")
	// Pad to width of 15
	RTVCache.RTV = fmt.Sprintf("%015s", RTVCache.RTV)

	// Fetch the CSS payload
	getRtv(&RTVCache.CSS, rtvHost+"/rtv/"+RTVCache.RTV+"/v0.css")
}

// get retrieves the body contents of the given url, populating the value into
// the given string pointer, fatally dying if there are any errors if this is the first time.
func getRtv(s *string, url string) {
	log.Printf("Fetching URL: %q\n", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	// Decide to die if this is the very first time initializing the value.
	shouldDie := *s == "" || s == nil

	if err != nil {
		if shouldDie {
			die("%+v", err)
		}
		log.Printf("Error getting %s: %+v", url, err)
		return
	}
	if resp.StatusCode != 200 {
		if shouldDie {
			die("Non-200 response for %s, %+v", url, resp)
		}
		log.Printf("Non-200 response for %s, %+v", url, resp)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if shouldDie {
			die("%+v", err)
		}
		log.Printf("Error reading response from %s: %+v", url, err)
		return
	}
	*s = string(body)
}
