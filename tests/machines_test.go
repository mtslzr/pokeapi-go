package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestMachine(t *testing.T) {
	result, _ := pokeapi.Machine("1")
	assert.Equal(t, "tm00", result.Item.Name,
		"Expect to receive TM00.")
}

func TestMachineFail(t *testing.T) {
	result, _ := pokeapi.Machine("asdf")
	assert.Equal(t, "", result.Item.Name,
		"Expect to receive an empty result.")
}
