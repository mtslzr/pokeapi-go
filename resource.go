package pokeapi

import (
	"github.com/mtslzr/pokeapi-go/structs"
)

// Resource returns resource list for an endpoint.
func Resource(endpoint string) (*structs.Resource, error) {
	var result structs.Resource

	err := do(endpoint, &result)
	return &result, err
}
