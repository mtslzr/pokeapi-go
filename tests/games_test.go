package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestGenerations(t *testing.T) {
	result, err := pokeapi.Generations()
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Generations{}, result,
		"Expect to receive a Generations struct.")
	assert.NotEqual(t, nil, result.Count,
		"Expect to receive a count of generations.")
	assert.Equal(t, "generation-i", result.Results[0].Name,
		"Expect to receive Generation I as first generation.")
}

func TestGeneration(t *testing.T) {
	result, err := pokeapi.Generation(1)
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Generation{}, result,
		"Expect to receive a Generation struct.")
	assert.Equal(t, "kanto", result.MainRegion.Name,
		"Expect to receive Kanto.")
}

func TestGenerationFail(t *testing.T) {
	result, err := pokeapi.Generation(9999)
	assert.NotEqual(t, nil, err,
		"Expect to receive an error.")
	assert.IsType(t, &structs.Generation{}, result,
		"Expect to receive a Generation struct.")
	assert.Equal(t, "", result.MainRegion.Name,
		"Expect to receive an empty result.")
}
