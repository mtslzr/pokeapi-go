package pokeapi

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var endpoint = "machine?offset=0&limit=1"
var mockResource structs.Resource

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

	CacheSettings.CustomExpire = 20
	customExpire := time.Now().Add(CacheSettings.CustomExpire * time.Minute).Minute()
	ClearCache()
	_ = do(endpoint, &mockResource)
	_, expires2, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, customExpire, expires2.Minute(),
		"Expect expiration time to match custom setting.")
}
