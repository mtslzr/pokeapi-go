package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Ability returns a single ability (by name or ID).
func Ability(id string) (result structs.Ability, err error) {
	err = do(fmt.Sprintf("ability/%s", id), &result)
	return result, err
}

// Characteristic returns a single characteristic (by ID).
func Characteristic(id string) (result structs.Characteristic, err error) {
	err = do(fmt.Sprintf("characteristic/%s", id), &result)
	return result, err
}

// EggGroup returns a single egg group (by name or ID).
func EggGroup(id string) (result structs.EggGroup, err error) {
	err = do(fmt.Sprintf("egg-group/%s", id), &result)
	return result, err
}

// Gender returns a single gender (by name or ID).
func Gender(id string) (result structs.Gender, err error) {
	err = do(fmt.Sprintf("gender/%s", id), &result)
	return result, err
}

// GrowthRate returns a single growth rate (by name or ID).
func GrowthRate(id string) (result structs.GrowthRate, err error) {
	err = do(fmt.Sprintf("growth-rate/%s", id), &result)
	return result, err
}

// Nature returns a single nature (by name or ID).
func Nature(id string) (result structs.Nature, err error) {
	err = do(fmt.Sprintf("nature/%s", id), &result)
	return result, err
}

// PokeathlonStat returns a single Pokeathlon state (by name or ID).
func PokeathlonStat(id string) (result structs.PokeathlonStat, err error) {
	err = do(fmt.Sprintf("pokeathlon-stat/%s", id), &result)
	return result, err
}

// Pokemon returns a single Pokemon (by name or ID).
func Pokemon(id string) (result structs.Pokemon, err error) {
	err = do(fmt.Sprintf("pokemon/%s", id), &result)
	return result, err
}

// PokemonColor returns a single Pokemon color (by name or ID).
func PokemonColor(id string) (result structs.PokemonColor, err error) {
	err = do(fmt.Sprintf("pokemon-color/%s", id), &result)
	return result, err
}

// PokemonForm returns a single Pokemon form (by name or ID).
func PokemonForm(id string) (result structs.PokemonForm, err error) {
	err = do(fmt.Sprintf("pokemon-form/%s", id), &result)
	return result, err
}

// PokemonHabitat returns a single Pokemon habitat (by name or ID).
func PokemonHabitat(id string) (result structs.PokemonHabitat, err error) {
	err = do(fmt.Sprintf("pokemon-habitat/%s", id), &result)
	return result, err
}

// PokemonShape returns a single Pokemon shape (by name or ID).
func PokemonShape(id string) (result structs.PokemonShape, err error) {
	err = do(fmt.Sprintf("pokemon-shape/%s", id), &result)
	return result, err
}

// PokemonSpecies returns a single Pokemon species (by name or ID).
func PokemonSpecies(id string) (result structs.PokemonSpecies, err error) {
	err = do(fmt.Sprintf("pokemon-species/%s", id), &result)
	return result, err
}

// Stat returns a single stat (by name or ID).
func Stat(id string) (result structs.Stat, err error) {
	err = do(fmt.Sprintf("stat/%s", id), &result)
	return result, err
}

// Type returns a single type (by name or ID).
func Type(id string) (result structs.Type, err error) {
	err = do(fmt.Sprintf("type/%s", id), &result)
	return result, err
}
