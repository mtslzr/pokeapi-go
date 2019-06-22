package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Pokemon returns a single Pokemon (by name or ID).
func Pokemon(id string) (*structs.Pokemon, error) {
	endpoint := fmt.Sprintf("pokemon/%s", id)
	var result structs.Pokemon

	err := do(endpoint, &result)
	return &result, err
}
