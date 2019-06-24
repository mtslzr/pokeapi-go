package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const apiurl = "https://pokeapi.co/api/v2/"
const cachemin = 5
const cachemax = 10

var c *cache.Cache

func init() {
	c = cache.New(cachemin*time.Minute, cachemax*time.Minute)
}

// ClearCache clears all cached data.
func ClearCache() {
	c.Flush()
}

func do(endpoint string, obj interface{}) error {
	cached, found := c.Get(endpoint)
	if found {
		return json.Unmarshal(cached.([]byte), &obj)
	}

	req, err := http.NewRequest(http.MethodGet, apiurl+endpoint, nil)
	if err != nil {
		return err
	}
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.Set(endpoint, body, cache.DefaultExpiration)
	return json.Unmarshal(body, &obj)
}
