package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestGeneration(t *testing.T) {
	result, _ := pokeapi.Generation("1")
	assert.Equal(t, "kanto", result.MainRegion.Name,
		"Expect to receive Kanto.")
}

func TestGenerationByName(t *testing.T) {
	result, _ := pokeapi.Generation("generation-i")
	assert.Equal(t, "kanto", result.MainRegion.Name,
		"Expect to receive Kanto.")
}

func TestGenerationFail(t *testing.T) {
	result, _ := pokeapi.Generation("asdf")
	assert.Equal(t, "", result.MainRegion.Name,
		"Expect to receive an empty result.")
}

func TestPokedex(t *testing.T) {
	result, _ := pokeapi.Pokedex("1")
	assert.Equal(t, "national", result.Name,
		"Expect to receive National Dex.")
}

func TestPokedexByName(t *testing.T) {
	result, _ := pokeapi.Pokedex("national")
	assert.Equal(t, "national", result.Name,
		"Expect to receive National Dex.")
}

func TestPokedexFail(t *testing.T) {
	result, _ := pokeapi.Pokedex("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestVersion(t *testing.T) {
	result, _ := pokeapi.Version("1")
	assert.Equal(t, "red", result.Name,
		"Expect to receive Red Version.")
}

func TestVersionByName(t *testing.T) {
	result, _ := pokeapi.Version("red")
	assert.Equal(t, "red", result.Name,
		"Expect to receive Red Version.")
}

func TestVersionFail(t *testing.T) {
	result, _ := pokeapi.Version("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestVersionGroup(t *testing.T) {
	result, _ := pokeapi.VersionGroup("1")
	assert.Equal(t, "red-blue", result.Name,
		"Expect to receive Red-Blue.")
}

func TestVersionGroupByName(t *testing.T) {
	result, _ := pokeapi.VersionGroup("red-blue")
	assert.Equal(t, "red-blue", result.Name,
		"Expect to receive Red-Blue.")
}

func TestVersionGroupFail(t *testing.T) {
	result, _ := pokeapi.VersionGroup("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
