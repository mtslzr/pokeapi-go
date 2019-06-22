package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	result, _ := pokeapi.Move("1")
	assert.Equal(t, "pound", result.Name,
		"Expect to receive Pound.")
}

func TestMoveByName(t *testing.T) {
	result, _ := pokeapi.Move("pound")
	assert.Equal(t, "pound", result.Name,
		"Expect to receive Pound.")
}

func TestMoveFail(t *testing.T) {
	result, _ := pokeapi.Move("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestMoveAilment(t *testing.T) {
	result, _ := pokeapi.MoveAilment("1")
	assert.Equal(t, "paralysis", result.Name,
		"Expect to receive Paralysis.")
}

func TestMoveAilmentByName(t *testing.T) {
	result, _ := pokeapi.MoveAilment("paralysis")
	assert.Equal(t, "paralysis", result.Name,
		"Expect to receive Paralysis.")
}

func TestMoveAilmentFail(t *testing.T) {
	result, _ := pokeapi.MoveAilment("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestMoveBattleStyle(t *testing.T) {
	result, _ := pokeapi.MoveBattleStyle("1")
	assert.Equal(t, "attack", result.Name,
		"Expect to receive Attack.")
}

func TestMoveBattleStyleByName(t *testing.T) {
	result, _ := pokeapi.MoveBattleStyle("attack")
	assert.Equal(t, "attack", result.Name,
		"Expect to receive Attack.")
}

func TestMoveBattleStyleFail(t *testing.T) {
	result, _ := pokeapi.MoveBattleStyle("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestMoveCategory(t *testing.T) {
	result, _ := pokeapi.MoveCategory("1")
	assert.Equal(t, "ailment", result.Name,
		"Expect to receive Ailment.")
}

func TestMoveCategoryByName(t *testing.T) {
	result, _ := pokeapi.MoveCategory("ailment")
	assert.Equal(t, "ailment", result.Name,
		"Expect to receive Ailment.")
}

func TestMoveCategoryFail(t *testing.T) {
	result, _ := pokeapi.MoveCategory("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestMoveDamageClass(t *testing.T) {
	result, _ := pokeapi.MoveDamageClass("1")
	assert.Equal(t, "status", result.Name,
		"Expect to receive Status.")
}

func TestMoveDamageClassByName(t *testing.T) {
	result, _ := pokeapi.MoveDamageClass("status")
	assert.Equal(t, "status", result.Name,
		"Expect to receive Status.")
}

func TestMoveDamageClassFail(t *testing.T) {
	result, _ := pokeapi.MoveDamageClass("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestMoveLearnMethod(t *testing.T) {
	result, _ := pokeapi.MoveLearnMethod("1")
	assert.Equal(t, "level-up", result.Name,
		"Expect to receive Level Up.")
}

func TestMoveLearnMethodByName(t *testing.T) {
	result, _ := pokeapi.MoveLearnMethod("level-up")
	assert.Equal(t, "level-up", result.Name,
		"Expect to receive Level Up.")
}

func TestMoveLearnMethodFail(t *testing.T) {
	result, _ := pokeapi.MoveLearnMethod("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestMoveTarget(t *testing.T) {
	result, _ := pokeapi.MoveTarget("1")
	assert.Equal(t, "specific-move", result.Name,
		"Expect to receive Specific Move.")
}

func TestMoveTargetByName(t *testing.T) {
	result, _ := pokeapi.MoveTarget("specific-move")
	assert.Equal(t, "specific-move", result.Name,
		"Expect to receive Specific Move.")
}

func TestMoveTargetFail(t *testing.T) {
	result, _ := pokeapi.MoveTarget("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
