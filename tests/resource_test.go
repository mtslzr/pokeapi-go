package tests

import (
	"testing"

	"github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {
	tests := []string{
		"berry", "berry-firmness", "berry-flavor", "contest-type",
		"contest-effect", "super-contest-effect", "encounter-method",
		"encounter-condition", "encounter-condition-value", "evolution-chain",
		"evolution-trigger", "generation", "pokedex", "version", "version-group",
		"item", "item-attribute", "item-category", "item-fling-effect",
		"item-pocket", "location", "location-area", "pal-park-area", "region",
		"machine", "move", "move-ailment", "move-battle-style", "move-category",
		"move-damage-class", "move-learn-method", "move-target", "ability",
		"egg-group", "gender", "growth-rate", "nature", "pokeathlon-stat",
		"pokemon", "pokemon-color", "pokemon-form", "pokemon-habitat",
		"pokemon-species", "stat", "type", "language",
	}
	for _, test := range tests {
		result, _ := pokeapi.Resource(test)
		assert.NotEqual(t, 0, len(result.Results),
			"Expect to receive more than zero results.")
		assert.NotEqual(t, nil, result.Results[0].URL,
			"Expect to receive results with URLs.")
	}
}

func TestResourceOffset(t *testing.T) {
	result, _ := pokeapi.Resource("pokemon", 3)
	assert.Equal(t, "charmander", result.Results[0].Name,
		"Expect to receive Charmander.")
}

func TestResourceOffsetLimit(t *testing.T) {
	result, _ := pokeapi.Resource("pokemon", 3, 3)
	assert.Equal(t, 3, len(result.Results),
		"Expect to receive exactly three results.")
	assert.Equal(t, "charizard", result.Results[2].Name,
		"Expect to receive Charizard last.")
}

func TestSearch(t *testing.T) {
	result, _ := pokeapi.Search("pokemon", "saur")
	assert.Equal(t, 4, len(result.Results),
		"Expect to receive four results.")
	assert.Equal(t, "venusaur", result.Results[2].Name,
		"Expect to receive Venusaur.")
}

func TestSearchFail(t *testing.T) {
	result, _ := pokeapi.Search("pokemon", "asdf")
	assert.Equal(t, 0, len(result.Results),
		"Expect to receive zero results.")
}

func TestSearchStartsWith(t *testing.T) {
	result, _ := pokeapi.Search("pokemon", "^a")
	assert.Equal(t, 50, len(result.Results),
		"Expect to receive four results.")
	assert.Equal(t, "arbok", result.Results[0].Name,
		"Expect to receive Arbok.")

	result, _ = pokeapi.Search("pokemon", "^bla")
	assert.Equal(t, 5, len(result.Results),
		"Expect to receive four results.")
	assert.Equal(t, "blastoise", result.Results[0].Name,
		"Expect to receive Blastoise.")
}

func TestSearchStartsWithFail(t *testing.T) {
	result, _ := pokeapi.Search("pokemon", "^zzz")
	assert.Equal(t, 0, len(result.Results),
		"Expect to receive zero results.")

	result, _ = pokeapi.Search("pokemon", "^asdfasdfasdfasdfasdfasdf")
	assert.Equal(t, 0, len(result.Results),
		"Expect to receive zero results.")
}
