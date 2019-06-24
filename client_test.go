package pokeapi

import (
	"encoding/json"
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var endpoint = "machine/1"
var mockMachine structs.Machine

func TestSetCache(t *testing.T) {
	_, found1 := c.Get(endpoint)
	assert.Equal(t, false, found1,
		"Expect to not have cached data before first call.")

	_ = do(endpoint, &mockMachine)
	cached2, expires2, found2 := c.GetWithExpiration(endpoint)
	assert.Equal(t, true, found2,
		"Expect to have cached data after first call.")

	_ = do(endpoint, &mockMachine)
	var cachedData structs.Machine
	json.Unmarshal(cached2.([]byte), &cachedData)
	cached3, expires3, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, cachedData, mockMachine,
		"Expect data to match cached data.")
	assert.Equal(t, cached2, cached3,
		"Expect cached data to match previously-cached data.")
	assert.Equal(t, expires2, expires3,
		"Expect expiration times to match.")
}

func TestClearCache(t *testing.T) {
	_, found := c.Get(endpoint)
	if !found {
		_ = do(endpoint, &mockMachine)
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
