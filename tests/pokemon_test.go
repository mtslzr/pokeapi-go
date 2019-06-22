package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestPokemons(t *testing.T) {
	result, err := pokeapi.Pokemons()
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Pokemons{}, result,
		"Expect to receive a Pokemons struct.")
	assert.NotEqual(t, nil, result.Count,
		"Expect to receive a count of pokemon.")
	assert.Equal(t, "bulbasaur", result.Results[0].Name,
		"Expect to receive Bulbasaur as first pokemon.")
}

func TestPokemon(t *testing.T) {
	result, err := pokeapi.Pokemon(1)
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Pokemon{}, result,
		"Expect to receive a Pokemon struct.")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFail(t *testing.T) {
	result, err := pokeapi.Pokemon(9999)
	assert.NotEqual(t, nil, err,
		"Expect to receive an error.")
	assert.IsType(t, &structs.Pokemon{}, result,
		"Expect to receive a Pokemon struct.")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
