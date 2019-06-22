package structs

// Generations is a list of all generations.
type Generations struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Generation is a single generation.
type Generation struct {
	Abilities  []interface{} `json:"abilities"`
	ID         int           `json:"id"`
	MainRegion struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"main_region"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
	Types []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"types"`
	VersionGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_groups"`
}
