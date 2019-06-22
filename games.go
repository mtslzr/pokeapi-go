package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Generation returns a single generation (by name or ID).
func Generation(id string) (*structs.Generation, error) {
	endpoint := fmt.Sprintf("generation/%s", id)
	var result structs.Generation

	err := do(endpoint, &result)
	return &result, err
}
