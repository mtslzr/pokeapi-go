package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestItem(t *testing.T) {
	result, _ := pokeapi.Item("1")
	assert.Equal(t, "master-ball", result.Name,
		"Expect to receive Master Ball.")
}

func TestItemByName(t *testing.T) {
	result, _ := pokeapi.Item("master-ball")
	assert.Equal(t, "master-ball", result.Name,
		"Expect to receive Master Ball.")
}

func TestItemFail(t *testing.T) {
	result, _ := pokeapi.Item("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestItemAttribute(t *testing.T) {
	result, _ := pokeapi.ItemAttribute("1")
	assert.Equal(t, "countable", result.Name,
		"Expect to receive Countable.")
}

func TestItemAttributeByName(t *testing.T) {
	result, _ := pokeapi.ItemAttribute("countable")
	assert.Equal(t, "countable", result.Name,
		"Expect to receive Countable.")
}

func TestItemAttributeFail(t *testing.T) {
	result, _ := pokeapi.ItemAttribute("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestItemCategory(t *testing.T) {
	result, _ := pokeapi.ItemCategory("1")
	assert.Equal(t, "stat-boosts", result.Name,
		"Expect to receive Stat Boosts.")
}

func TestItemCategoryByName(t *testing.T) {
	result, _ := pokeapi.ItemCategory("stat-boosts")
	assert.Equal(t, "stat-boosts", result.Name,
		"Expect to receive Stat Boosts.")
}

func TestItemCategoryFail(t *testing.T) {
	result, _ := pokeapi.ItemCategory("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestItemFlingEffect(t *testing.T) {
	result, _ := pokeapi.ItemFlingEffect("1")
	assert.Equal(t, "badly-poison", result.Name,
		"Expect to receive Badly Poison.")
}

func TestItemFlingEffectByName(t *testing.T) {
	result, _ := pokeapi.ItemFlingEffect("badly-poison")
	assert.Equal(t, "badly-poison", result.Name,
		"Expect to receive Badly Poison.")
}

func TestItemFlingEffectFail(t *testing.T) {
	result, _ := pokeapi.ItemFlingEffect("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestItemPocket(t *testing.T) {
	result, _ := pokeapi.ItemPocket("1")
	assert.Equal(t, "misc", result.Name,
		"Expect to receive Miscellaneous.")
}

func TestItemPocketByName(t *testing.T) {
	result, _ := pokeapi.ItemPocket("misc")
	assert.Equal(t, "misc", result.Name,
		"Expect to receive Miscellaneous.")
}

func TestItemPocketFail(t *testing.T) {
	result, _ := pokeapi.ItemPocket("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
