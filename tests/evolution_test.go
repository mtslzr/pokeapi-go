package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestEvolutionChain(t *testing.T) {
	result, _ := pokeapi.EvolutionChain("1")
	assert.Equal(t, "bulbasaur", result.Chain.Species.Name,
		"Expect to receive Bulbasaur.")
}

func TestEvolutionChainFail(t *testing.T) {
	result, _ := pokeapi.EvolutionChain("asdf")
	assert.Equal(t, "", result.Chain.Species.Name,
		"Expect to receive an empty result.")
}

func TestEvolutionTrigger(t *testing.T) {
	result, _ := pokeapi.EvolutionTrigger("1")
	assert.Equal(t, "level-up", result.Name,
		"Expect to receive Level Up.")
}

func TestEvolutionTriggerByName(t *testing.T) {
	result, _ := pokeapi.EvolutionTrigger("level-up")
	assert.Equal(t, "level-up", result.Name,
		"Expect to receive Level Up.")
}

func TestEvolutionTriggerFail(t *testing.T) {
	result, _ := pokeapi.EvolutionTrigger("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
