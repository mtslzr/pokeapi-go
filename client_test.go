package pokeapi

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

type CachedData struct {
	Data    interface{}
	Expires time.Time
	Found   bool
}

func TestCache(t *testing.T) {
	endpoint := "machine"
	fullEndpoint := endpoint + "?offset=0&limit=0"

	_, found1 := c.Get(endpoint)
	assert.Equal(t, false, found1,
		"Expect to not have cached data before first call.")

	result, _ := Resource(endpoint)
	cached2, expires2, found2 := c.GetWithExpiration(fullEndpoint)
	var cachedData structs.Resource
	json.Unmarshal(cached2.([]byte), &cachedData)
	assert.Equal(t, true, found2,
		"Expect to have cached data after first call.")

	result, _ = Resource(endpoint)
	cached3, expires3, _ := c.GetWithExpiration(fullEndpoint)
	assert.Equal(t, cachedData, result,
		"Expect data to match cached data.")
	assert.Equal(t, cached2, cached3,
		"Expect cached data to match previously-cached data.")
	assert.Equal(t, expires2, expires3,
		"Expect expiration times to match.")
}
