package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Generation returns a single generation (by name or ID).
func Generation(id string) (result structs.Generation, err error) {
	err = do(fmt.Sprintf("generation/%s", id), &result)
	return result, err
}

// Pokedex returns a single Pokedex (by name or ID).
func Pokedex(id string) (result structs.Pokedex, err error) {
	err = do(fmt.Sprintf("pokedex/%s", id), &result)
	return result, err
}

// Version returns a single version (by name or ID).
func Version(id string) (result structs.Version, err error) {
	err = do(fmt.Sprintf("version/%s", id), &result)
	return result, err
}

// VersionGroup returns a single version group (by name or ID).
func VersionGroup(id string) (result structs.VersionGroup, err error) {
	err = do(fmt.Sprintf("version-group/%s", id), &result)
	return result, err
}
