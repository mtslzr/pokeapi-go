package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// EvolutionChain returns a single evolution chain (by ID).
func EvolutionChain(id string) (result structs.EvolutionChain, err error) {
	err = do(fmt.Sprintf("evolution-chain/%s", id), &result)
	return result, err
}

// EvolutionTrigger returns a single evolution trigger (by ID or name).
func EvolutionTrigger(id string) (result structs.EvolutionTrigger, err error) {
	err = do(fmt.Sprintf("evolution-trigger/%s", id), &result)
	return result, err
}
