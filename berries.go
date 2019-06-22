package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Berry returns a single berry (by name or ID).
func Berry(id string) (result structs.Berry, err error) {
	err = do(fmt.Sprintf("berry/%s", id), &result)
	return result, err
}

// BerryFirmness returns a single berry firmness (by name or ID).
func BerryFirmness(id string) (result structs.BerryFirmness, err error) {
	err = do(fmt.Sprintf("berry-firmness/%s", id), &result)
	return result, err
}

// BerryFlavor returns a single berry flavor (by name or ID).
func BerryFlavor(id string) (result structs.BerryFlavor, err error) {
	err = do(fmt.Sprintf("berry-flavor/%s", id), &result)
	return result, err
}
