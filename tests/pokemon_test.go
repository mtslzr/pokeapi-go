package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestPokemon(t *testing.T) {
	result, err := pokeapi.Pokemon("1")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Pokemon{}, result,
		"Expect to receive a Pokemon struct.")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonByName(t *testing.T) {
	result, err := pokeapi.Pokemon("bulbasaur")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Pokemon{}, result,
		"Expect to receive a Pokemon struct.")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFail(t *testing.T) {
	result, err := pokeapi.Pokemon("digimon")
	assert.NotEqual(t, nil, err,
		"Expect to receive an error.")
	assert.IsType(t, &structs.Pokemon{}, result,
		"Expect to receive a Pokemon struct.")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
