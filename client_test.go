package pokeapi

import (
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestAllPokemon(t *testing.T) {
	result, err := AllPokemon()
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.AllPokemon{}, result,
		"Expect to receive an AllPokemon struct.")
	assert.NotEqual(t, nil, result.Count,
		"Expect to receive a count of pokemon.")
	assert.Equal(t, "bulbasaur", result.Results[0].Name,
		"Expect to receive Bulbasaur as first pokemon.")
}

func TestPokemon(t *testing.T) {
	result, err := Pokemon(1)
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Pokemon{}, result,
		"Expect to receive a Pokemon struct.")
	assert.NotEqual(t, nil, result.Name,
		"Expect to receive Bulbasaur.")
}
