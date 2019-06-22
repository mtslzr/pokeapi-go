package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestContestType(t *testing.T) {
	result, _ := pokeapi.ContestType("1")
	assert.Equal(t, "cool", result.Name,
		"Expect to receive Cool.")
}

func TestContestTypeByName(t *testing.T) {
	result, _ := pokeapi.ContestType("cool")
	assert.Equal(t, "cool", result.Name,
		"Expect to receive Cool.")
}

func TestContestTypeFail(t *testing.T) {
	result, _ := pokeapi.ContestType("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestContestEffect(t *testing.T) {
	result, _ := pokeapi.ContestEffect("1")
	assert.Equal(t, "A highly appealing move.",
		result.FlavorTextEntries[0].FlavorText, "Expect to receive flavor text.")
}

func TestContestEffectFail(t *testing.T) {
	result, _ := pokeapi.ContestEffect("asdf")
	assert.Equal(t, 0, len(result.FlavorTextEntries),
		"Expect to receive an empty result.")
}

func TestSuperContestEffect(t *testing.T) {
	result, _ := pokeapi.SuperContestEffect("1")
	assert.Equal(t, "Enables the user to perform first in the next turn.",
		result.FlavorTextEntries[0].FlavorText, "Expect to receive flavor text.")
}

func TestSuperContestEffectFail(t *testing.T) {
	result, _ := pokeapi.SuperContestEffect("asdf")
	assert.Equal(t, 0, len(result.FlavorTextEntries),
		"Expect to receive an empty result.")
}
