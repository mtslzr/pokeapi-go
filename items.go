package pokeapi

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// Item returns a single item (by name or ID).
func Item(id string) (result structs.Item, err error) {
	err = do(fmt.Sprintf("item/%s", id), &result)
	return result, err
}

// ItemAttribute returns a single item attribute (by name or ID).
func ItemAttribute(id string) (result structs.ItemAttribute, err error) {
	err = do(fmt.Sprintf("item-attribute/%s", id), &result)
	return result, err
}

// ItemCategory returns a single item category (by name or ID).
func ItemCategory(id string) (result structs.ItemCategory, err error) {
	err = do(fmt.Sprintf("item-category/%s", id), &result)
	return result, err
}

// ItemFlingEffect returns a single item fling effect (by name or ID).
func ItemFlingEffect(id string) (result structs.ItemFlingEffect, err error) {
	err = do(fmt.Sprintf("item-fling-effect/%s", id), &result)
	return result, err
}

// ItemPocket returns a single item pocket (by name or ID).
func ItemPocket(id string) (result structs.ItemPocket, err error) {
	err = do(fmt.Sprintf("item-pocket/%s", id), &result)
	return result, err
}
