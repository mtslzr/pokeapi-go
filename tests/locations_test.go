package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	result, _ := pokeapi.Location("1")
	assert.Equal(t, "canalave-city", result.Name,
		"Expect to receive Canalave City.")
}

func TestLocationByName(t *testing.T) {
	result, _ := pokeapi.Location("canalave-city")
	assert.Equal(t, "canalave-city", result.Name,
		"Expect to receive Canalave City.")
}

func TestLocationFail(t *testing.T) {
	result, _ := pokeapi.Location("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestLocationArea(t *testing.T) {
	result, _ := pokeapi.LocationArea("1")
	assert.Equal(t, "canalave-city-area", result.Name,
		"Expect to receive Canalave City area.")
}

func TestLocationAreaByName(t *testing.T) {
	result, _ := pokeapi.LocationArea("canalave-city-area")
	assert.Equal(t, "canalave-city-area", result.Name,
		"Expect to receive Canalave City area.")
}

func TestLocationAreaFail(t *testing.T) {
	result, _ := pokeapi.LocationArea("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPalParkArea(t *testing.T) {
	result, _ := pokeapi.PalParkArea("1")
	assert.Equal(t, "forest", result.Name,
		"Expect to receive Forest.")
}

func TestPalParkAreaByName(t *testing.T) {
	result, _ := pokeapi.PalParkArea("forest")
	assert.Equal(t, "forest", result.Name,
		"Expect to receive Forest.")
}

func TestPalParkAreaFail(t *testing.T) {
	result, _ := pokeapi.PalParkArea("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestRegion(t *testing.T) {
	result, _ := pokeapi.Region("1")
	assert.Equal(t, "kanto", result.Name,
		"Expect to receive Kanto.")
}

func TestRegionByName(t *testing.T) {
	result, _ := pokeapi.Region("kanto")
	assert.Equal(t, "kanto", result.Name,
		"Expect to receive Kanto.")
}

func TestRegionFail(t *testing.T) {
	result, _ := pokeapi.Region("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
