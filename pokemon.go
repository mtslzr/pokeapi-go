package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Ability returns a single ability (by name or ID).
func Ability(id string) (*structs.Ability, error) {
	endpoint := fmt.Sprintf("ability/%s", id)
	var result structs.Ability

	err := do(endpoint, &result)
	return &result, err
}

// EggGroup returns a single egg group (by name or ID).
func EggGroup(id string) (*structs.EggGroup, error) {
	endpoint := fmt.Sprintf("egg-group/%s", id)
	var result structs.EggGroup

	err := do(endpoint, &result)
	return &result, err
}

// Gender returns a single gender (by name or ID).
func Gender(id string) (*structs.Gender, error) {
	endpoint := fmt.Sprintf("gender/%s", id)
	var result structs.Gender

	err := do(endpoint, &result)
	return &result, err
}

// GrowthRate returns a single growth rate (by name or ID).
func GrowthRate(id string) (*structs.GrowthRate, error) {
	endpoint := fmt.Sprintf("growth-rate/%s", id)
	var result structs.GrowthRate

	err := do(endpoint, &result)
	return &result, err
}

// Nature returns a single nature (by name or ID).
func Nature(id string) (*structs.Nature, error) {
	endpoint := fmt.Sprintf("nature/%s", id)
	var result structs.Nature

	err := do(endpoint, &result)
	return &result, err
}

// PokeathlonStat returns a single Pokeathlon stat (by name or ID).
func PokeathlonStat(id string) (*structs.PokeathlonStat, error) {
	endpoint := fmt.Sprintf("pokeathlon-stat/%s", id)
	var result structs.PokeathlonStat

	err := do(endpoint, &result)
	return &result, err
}

// Pokemon returns a single Pokemon (by name or ID).
func Pokemon(id string) (*structs.Pokemon, error) {
	endpoint := fmt.Sprintf("pokemon/%s", id)
	var result structs.Pokemon

	err := do(endpoint, &result)
	return &result, err
}

// PokemonColor returns a single Pokemon color (by name or ID).
func PokemonColor(id string) (*structs.PokemonColor, error) {
	endpoint := fmt.Sprintf("pokemon-color/%s", id)
	var result structs.PokemonColor

	err := do(endpoint, &result)
	return &result, err
}

// PokemonForm returns a single Pokemon form (by name or ID).
func PokemonForm(id string) (*structs.PokemonForm, error) {
	endpoint := fmt.Sprintf("pokemon-form/%s", id)
	var result structs.PokemonForm

	err := do(endpoint, &result)
	return &result, err
}

// PokemonHabitat returns a single Pokemon habitat (by name or ID).
func PokemonHabitat(id string) (*structs.PokemonHabitat, error) {
	endpoint := fmt.Sprintf("pokemon-habitat/%s", id)
	var result structs.PokemonHabitat

	err := do(endpoint, &result)
	return &result, err
}

// PokemonSpecies returns a single Pokemon species (by name or ID).
func PokemonSpecies(id string) (*structs.PokemonSpecies, error) {
	endpoint := fmt.Sprintf("pokemon-species/%s", id)
	var result structs.PokemonSpecies

	err := do(endpoint, &result)
	return &result, err
}

// Stat returns a single stat (by name or ID).
func Stat(id string) (*structs.Stat, error) {
	endpoint := fmt.Sprintf("stat/%s", id)
	var result structs.Stat

	err := do(endpoint, &result)
	return &result, err
}

// Type returns a single type (by name or ID).
func Type(id string) (*structs.Type, error) {
	endpoint := fmt.Sprintf("type/%s", id)
	var result structs.Type

	err := do(endpoint, &result)
	return &result, err
}
