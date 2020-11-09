package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	defaultAPIScheme = "https"
	defaultAPIHost   = "pokeapi.co"
	defaultAPIUrl    = "/api/v2/"

	httpRetries = 3
)

var (
	c *cache.Cache

	apiHost   = defaultAPIHost
	apiScheme = defaultAPIScheme
)

func init() {
	c = cache.New(defaultCacheSettings.MinExpire, defaultCacheSettings.MaxExpire)
}

// SetAPIPath configures the Host and Scheme sections used when requesting data.
func SetAPIPath(path string) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}

	if u.Host == "" {
		return fmt.Errorf("failed to parse url host: %s", u.Host)
	}
	apiHost = u.Host

	apiScheme = u.Scheme
	if apiScheme == "" {
		apiScheme = defaultAPIScheme
	}

	return nil
}

// retryableGet performs a HTTP GET on the given URL path. Retries are attempted
// with backoff for server side errors.
func retryableGet(resource string) ([]byte, error) {
	var lasterr error

	for attempt := 0; attempt < httpRetries; attempt++ {
		time.Sleep(time.Second * time.Duration(attempt*attempt)) // exponential backoff

		req, err := http.NewRequest(http.MethodGet, resource, nil)
		if err != nil {
			lasterr = err
			continue
		}
		client := &http.Client{Timeout: 10 * time.Second}

		resp, err := client.Do(req)
		if err != nil {
			lasterr = err
			continue
		}
		if resp.Body != nil { // body can be nil if an err status is returned
			defer resp.Body.Close()
		}

		if resp.StatusCode >= 500 {
			lasterr = fmt.Errorf("server returned error (code: %d)", resp.StatusCode)
			continue
		} else if resp.StatusCode >= 400 {
			// if we got a 4xx code then screwed up, no sense retrying ..
			return nil, err
		}
		// since the go client handles redirects (3xx) codes, we'll leave those out

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			lasterr = err
			continue
		}

		return body, err
	}

	return nil, lasterr
}

// buildURL creates a URL for the given endpoint using the current configuration
func buildURL(endpoint string) string {
	apiurl := &url.URL{Scheme: apiScheme, Host: apiHost, Path: defaultAPIUrl + endpoint}
	return apiurl.String()
}

func do(endpoint string, obj interface{}) error {
	cached, found := c.Get(endpoint)
	if found && CacheSettings.UseCache {
		return json.Unmarshal(cached.([]byte), &obj)
	}

	body, err := retryableGet(buildURL(endpoint))
	if err != nil {
		return err
	}

	setCache(endpoint, body)
	return json.Unmarshal(body, &obj)
}
