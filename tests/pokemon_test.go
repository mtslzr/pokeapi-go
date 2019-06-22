package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestAbility(t *testing.T) {
	result, err := pokeapi.Ability("1")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Ability{}, result,
		"Expect to receive an Ability struct.")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityByName(t *testing.T) {
	result, err := pokeapi.Ability("stench")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Ability{}, result,
		"Expect to receive an Ability struct.")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityFail(t *testing.T) {
	result, err := pokeapi.Ability("asdf")
	assert.NotEqual(t, nil, err,
		"Expect to receive an error.")
	assert.IsType(t, &structs.Ability{}, result,
		"Expect to receive an Ability struct.")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestEggGroup(t *testing.T) {
	result, err := pokeapi.EggGroup("1")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.EggGroup{}, result,
		"Expect to receive an EggGroup struct.")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupByName(t *testing.T) {
	result, err := pokeapi.EggGroup("monster")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.EggGroup{}, result,
		"Expect to receive an EggGroup struct.")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupFail(t *testing.T) {
	result, err := pokeapi.EggGroup("asdf")
	assert.NotEqual(t, nil, err,
		"Expect to receive an error.")
	assert.IsType(t, &structs.EggGroup{}, result,
		"Expect to receive am EggGroup struct.")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

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
