package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Location returns a single location (by name or ID).
func Location(id string) (result structs.Location, err error) {
	err = do(fmt.Sprintf("location/%s", id), &result)
	return result, err
}

// LocationArea returns a single location area (by name or ID).
func LocationArea(id string) (result structs.LocationArea, err error) {
	err = do(fmt.Sprintf("location-area/%s", id), &result)
	return result, err
}

// PalParkArea returns a single Pal Park area (by name or ID).
func PalParkArea(id string) (result structs.PalParkArea, err error) {
	err = do(fmt.Sprintf("pal-park-area/%s", id), &result)
	return result, err
}

// Region returns a single region (by name or ID).
func Region(id string) (result structs.Region, err error) {
	err = do(fmt.Sprintf("region/%s", id), &result)
	return result, err
}
