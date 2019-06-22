package pokeapi

import (
	"fmt"
	"strings"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Resource returns resource list for an endpoint.
func Resource(endpoint string, params ...int) (result structs.Resource,
	err error) {
	offset, limit := parseParams(params)
	err = do(fmt.Sprintf("%s?offset=%d&limit=%d", endpoint, offset, limit),
		&result)
	return
}

// Search returns resource list, filtered by search term.
func Search(endpoint string, search string) (result structs.Resource,
	err error) {
	err = do(fmt.Sprintf("%s?offset=0&limit=9999", endpoint), &result)
	result.Results = parseSearch(result.Results, search)
	return
}

func parseParams(params []int) (offset int, limit int) {
	switch x := len(params); {
	case x == 2:
		limit = params[1]
		fallthrough
	case x >= 1:
		offset = params[0]
	}
	return
}

func parseSearch(results []structs.Result, search string) []structs.Result {
	var x int
	for _, result := range results {
		if !strings.Contains(result.Name, search) {
			continue
		}
		results[x] = result
		x++
	}
	return results[:x]
}
