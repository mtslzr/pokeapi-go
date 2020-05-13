package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const apiurl = "https://pokeapi.co/api/v2/"

var c *cache.Cache

func init() {
	c = cache.New(defaultCacheSettings.MinExpire, defaultCacheSettings.MaxExpire)
}

func do(endpoint string, obj interface{}) error {
	cached, found := c.Get(endpoint)
	if found && CacheSettings.UseCache {
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

	setCache(endpoint, body)
	return json.Unmarshal(body, &obj)
}
