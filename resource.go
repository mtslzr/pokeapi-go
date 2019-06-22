package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Resource returns resource list for an endpoint.
func Resource(endpoint string, params ...int) (result structs.Resource,
	err error) {
	offset, limit := parseParams(params)
	err = do(fmt.Sprintf("%s?offset=%d&limit=%d", endpoint, offset, limit),
		&result)
	return result, err
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
