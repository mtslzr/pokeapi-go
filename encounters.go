package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// EncounterMethod returns a single encounter method (by name or ID).
func EncounterMethod(id string) (result structs.EncounterMethod, err error) {
	err = do(fmt.Sprintf("encounter-method/%s", id), &result)
	return result, err
}

// EncounterCondition returns a single encounter condition (by name or ID).
func EncounterCondition(id string) (result structs.EncounterCondition,
	err error) {
	err = do(fmt.Sprintf("encounter-condition/%s", id), &result)
	return result, err
}

// EncounterConditionValue returns a single encounter condition value
//   (by name or ID).
func EncounterConditionValue(id string) (result structs.EncounterConditionValue,
	err error) {
	err = do(fmt.Sprintf("encounter-condition-value/%s", id), &result)
	return result, err
}
