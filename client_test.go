package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var endpoint = "machine?offset=0&limit=1"
var mockResource structs.Resource

func TestBuildURL(t *testing.T) {
	scheme := "http"
	host := "test.io"

	err := SetAPIPath(fmt.Sprintf("%s://%s", scheme, host))

	assert.Nil(t, err)

	assert.Equal(t, host, apiHost)
	assert.Equal(t, scheme, apiScheme)
}

var dataTestSetAPIPath = []struct {
	Given        string
	ExpectHost   string
	ExpectScheme string
	ShouldErr    bool
}{
	{"http://foobar.com", "foobar.com", "http", false},
	{"https://zip.co", "zip.co", "https", false},
	{"whatisthis.com/", "whatisthis.com", "", true},
	{(&url.URL{Scheme: defaultAPIScheme, Host: defaultAPIHost}).String(), defaultAPIHost, defaultAPIScheme, false},
}

func TestAPIPath(t *testing.T) {
	for _, tt := range dataTestSetAPIPath {
		t.Run(tt.Given, func(t *testing.T) {
			err := SetAPIPath(tt.Given)

			assert.Equal(t, tt.ShouldErr, err != nil)
			if err != nil {
				return
			}

			assert.Equal(t, tt.ExpectHost, apiHost)
			assert.Equal(t, tt.ExpectScheme, apiScheme)
		})
	}
}

func TestSetCache(t *testing.T) {
	_, found1 := c.Get(endpoint)
	assert.Equal(t, false, found1,
		"Expect to not have cached data before first call.")

	_ = do(endpoint, &mockResource)
	cached2, expires2, found2 := c.GetWithExpiration(endpoint)
	assert.Equal(t, true, found2,
		"Expect to have cached data after first call.")

	_ = do(endpoint, &mockResource)
	var cachedData structs.Resource
	json.Unmarshal(cached2.([]byte), &cachedData)
	cached3, expires3, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, cachedData, mockResource,
		"Expect data to match cached data.")
	assert.Equal(t, cached2, cached3,
		"Expect cached data to match previously-cached data.")
	assert.Equal(t, expires2, expires3,
		"Expect expiration times to match.")
}

func TestClearCache(t *testing.T) {
	_, found := c.Get(endpoint)
	if !found {
		_ = do(endpoint, &mockResource)
		_, found = c.Get(endpoint)
	}
	assert.NotEqual(t, false, found,
		"Expect to start with cached data.")

	ClearCache()
	nocache, nofound := c.Get(endpoint)
	assert.Equal(t, false, nofound,
		"Expect no data found after flushing cache.")
	assert.Equal(t, nil, nocache,
		"Expect no data after flushing cache.")
}

func TestCustomExpiration(t *testing.T) {
	ClearCache()
	defaultExpire := time.Now().Add(defaultCacheSettings.MinExpire).Minute()
	_ = do(endpoint, &mockResource)
	_, expires1, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, defaultExpire, expires1.Minute(),
		"Expect expiration time to match default setting.")

	ClearCache()
	CacheSettings.CustomExpire = 20
	customExpire := time.Now().Add(CacheSettings.CustomExpire * time.Minute).Minute()
	_ = do(endpoint, &mockResource)
	_, expires2, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, customExpire, expires2.Minute(),
		"Expect expiration time to match custom setting.")
}

func TestNoCache(t *testing.T) {
	ClearCache()
	_ = do(endpoint, &mockResource)
	_, expires1, found1 := c.GetWithExpiration(endpoint)
	assert.Equal(t, true, found1,
		"Expect to have cached data after first call.")

	CacheSettings.UseCache = false
	_ = do(endpoint, &mockResource)
	_, expires2, _ := c.GetWithExpiration(endpoint)
	assert.NotEqual(t, expires1, expires2,
		"Expect cache expiration not to match first call.")
}
