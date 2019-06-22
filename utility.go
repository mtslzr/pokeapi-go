package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Language returns a single language (by name or ID).
func Language(id string) (result structs.Language, err error) {
	err = do(fmt.Sprintf("language/%s", id), &result)
	return result, err
}
