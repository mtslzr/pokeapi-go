package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// AllPokemon returns all pokemon.
func AllPokemon() (*structs.AllPokemon, error) {
	endpoint := "pokemon"
	var result structs.AllPokemon

	err := do(endpoint, &result)
	return &result, err
}

// Pokemon returns a single pokemon.
func Pokemon(id int) (*structs.Pokemon, error) {
	endpoint := fmt.Sprintf("pokemon/%d", id)
	var result structs.Pokemon

	err := do(endpoint, &result)
	return &result, err
}
