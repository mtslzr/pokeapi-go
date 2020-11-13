package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestAbility(t *testing.T) {
	result, _ := pokeapi.Ability("1")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityByName(t *testing.T) {
	result, _ := pokeapi.Ability("stench")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityFail(t *testing.T) {
	result, _ := pokeapi.Ability("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestCharacteristic(t *testing.T) {
	result, _ := pokeapi.Characteristic("1")
	assert.Equal(t, "Loves to eat", result.Descriptions[2].Description,
		"Expect to receive a description.")
}

func TestCharacteristicFail(t *testing.T) {
	result, _ := pokeapi.Characteristic("asdf")
	assert.Equal(t, 0, len(result.Descriptions),
		"Expect to receive an empty result.")
}

func TestEggGroup(t *testing.T) {
	result, _ := pokeapi.EggGroup("1")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupByName(t *testing.T) {
	result, _ := pokeapi.EggGroup("monster")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupFail(t *testing.T) {
	result, _ := pokeapi.EggGroup("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestGender(t *testing.T) {
	result, _ := pokeapi.Gender("1")
	assert.Equal(t, "female", result.Name,
		"Expect to receive Female.")
}

func TestGenderByName(t *testing.T) {
	result, _ := pokeapi.Gender("female")
	assert.Equal(t, "female", result.Name,
		"Expect to receive Female.")
}

func TestGenderFail(t *testing.T) {
	result, _ := pokeapi.Gender("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestGrowthRate(t *testing.T) {
	result, _ := pokeapi.GrowthRate("1")
	assert.Equal(t, "slow", result.Name,
		"Expect to receive Slow.")
}

func TestGrowthRateByName(t *testing.T) {
	result, _ := pokeapi.GrowthRate("slow")
	assert.Equal(t, "slow", result.Name,
		"Expect to receive Slow.")
}

func TestGrowthRateFail(t *testing.T) {
	result, _ := pokeapi.GrowthRate("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestNature(t *testing.T) {
	result, _ := pokeapi.Nature("1")
	assert.Equal(t, "hardy", result.Name,
		"Expect to receive Hardy.")
}

func TestNatureByName(t *testing.T) {
	result, _ := pokeapi.Nature("hardy")
	assert.Equal(t, "hardy", result.Name,
		"Expect to receive Hardy.")
}

func TestNatureFail(t *testing.T) {
	result, _ := pokeapi.Nature("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokeathlonStat(t *testing.T) {
	result, _ := pokeapi.PokeathlonStat("1")
	assert.Equal(t, "speed", result.Name,
		"Expect to receive Speed.")
}

func TestPokeathlonStatByName(t *testing.T) {
	result, _ := pokeapi.PokeathlonStat("speed")
	assert.Equal(t, "speed", result.Name,
		"Expect to receive Speed.")
}

func TestPokeathlonStatFail(t *testing.T) {
	result, _ := pokeapi.PokeathlonStat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemon(t *testing.T) {
	result, _ := pokeapi.Pokemon("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonByName(t *testing.T) {
	result, _ := pokeapi.Pokemon("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFail(t *testing.T) {
	result, _ := pokeapi.Pokemon("digimon")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonColor(t *testing.T) {
	result, _ := pokeapi.PokemonColor("1")
	assert.Equal(t, "black", result.Name,
		"Expect to receive Black.")
}

func TestPokemonColorByName(t *testing.T) {
	result, _ := pokeapi.PokemonColor("black")
	assert.Equal(t, "black", result.Name,
		"Expect to receive Black.")
}

func TestPokemonColorFail(t *testing.T) {
	result, _ := pokeapi.PokemonColor("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonForm(t *testing.T) {
	result, _ := pokeapi.PokemonForm("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFormByName(t *testing.T) {
	result, _ := pokeapi.PokemonForm("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFormFail(t *testing.T) {
	result, _ := pokeapi.PokemonForm("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonHabitat(t *testing.T) {
	result, _ := pokeapi.PokemonHabitat("1")
	assert.Equal(t, "cave", result.Name,
		"Expect to receive Cave.")
}

func TestPokemonHabitatByName(t *testing.T) {
	result, _ := pokeapi.PokemonHabitat("cave")
	assert.Equal(t, "cave", result.Name,
		"Expect to receive Cave.")
}

func TestPokemonHabitatFail(t *testing.T) {
	result, _ := pokeapi.PokemonHabitat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonShape(t *testing.T) {
	result, _ := pokeapi.PokemonShape("1")
	assert.Equal(t, "ball", result.Name,
		"Expect to receive Ball.")
}

func TestPokemonShapeByName(t *testing.T) {
	result, _ := pokeapi.PokemonShape("ball")
	assert.Equal(t, "ball", result.Name,
		"Expect to receive Ball.")
}

func TestPokemonShapeFail(t *testing.T) {
	result, _ := pokeapi.PokemonShape("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonSpecies(t *testing.T) {
	result, _ := pokeapi.PokemonSpecies("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonSpeciesByName(t *testing.T) {
	result, _ := pokeapi.PokemonSpecies("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonSpeciesFail(t *testing.T) {
	result, _ := pokeapi.PokemonSpecies("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestStat(t *testing.T) {
	result, _ := pokeapi.Stat("1")
	assert.Equal(t, "hp", result.Name,
		"Expect to receive HP.")
}

func TestStatByName(t *testing.T) {
	result, _ := pokeapi.Stat("hp")
	assert.Equal(t, "hp", result.Name,
		"Expect to receive HP.")
}

func TestStatFail(t *testing.T) {
	result, _ := pokeapi.Stat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestType(t *testing.T) {
	result, _ := pokeapi.Type("1")
	assert.Equal(t, "normal", result.Name,
		"Expect to receive Normal.")
}

func TestTypeByName(t *testing.T) {
	result, _ := pokeapi.Type("normal")
	assert.Equal(t, "normal", result.Name,
		"Expect to receive Normal.")
}

func TestTypeFail(t *testing.T) {
	result, _ := pokeapi.Type("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonLocationArea(t *testing.T) {
	result, _ := pokeapi.PokemonLocationAreas("4")
	assert.Equal(t, "pallet-town-area", result[0].LocationArea.Name)
}

// Endpoint returns 404 so request should come back empty
func TestPokemonLocationAreaFail(t *testing.T) {
	result, _ := pokeapi.PokemonLocationAreas("asdf")
	assert.Equal(t, []structs.PokemonLocationArea(nil), result)
}
