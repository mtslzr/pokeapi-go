package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mtslzr/pokeapi-go/structs"
)

// APIURL for PokeAPI.
const APIURL = "https://pokeapi.co/api/v2/"

// AllPokemon returns all pokemon.
func AllPokemon() (*structs.AllPokemon, error) {
	endpoint := "pokemon"
	var result structs.AllPokemon

	err := doCall(endpoint, &result)
	return &result, err
}

// Pokemon returns a single pokemon.
func Pokemon(id int) (*structs.Pokemon, error) {
	endpoint := fmt.Sprintf("pokemon/%d", id)
	var result structs.Pokemon

	err := doCall(endpoint, &result)
	return &result, err
}

func doCall(endpoint string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, APIURL+endpoint, nil)
	if err != nil {
		return err
	}
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}
