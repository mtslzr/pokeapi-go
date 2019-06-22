package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestGeneration(t *testing.T) {
	result, err := pokeapi.Generation("1")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Generation{}, result,
		"Expect to receive a Generation struct.")
	assert.Equal(t, "kanto", result.MainRegion.Name,
		"Expect to receive Kanto.")
}

func TestGenerationByName(t *testing.T) {
	result, err := pokeapi.Generation("generation-i")
	assert.Equal(t, nil, err,
		"Expect to not receive an error.")
	assert.IsType(t, &structs.Generation{}, result,
		"Expect to receive a Generation struct.")
	assert.Equal(t, "kanto", result.MainRegion.Name,
		"Expect to receive Kanto.")
}

func TestGenerationFail(t *testing.T) {
	result, err := pokeapi.Generation("asdf")
	assert.NotEqual(t, nil, err,
		"Expect to receive an error.")
	assert.IsType(t, &structs.Generation{}, result,
		"Expect to receive a Generation struct.")
	assert.Equal(t, "", result.MainRegion.Name,
		"Expect to receive an empty result.")
}
