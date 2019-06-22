package structs

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

// Pokedex is a single Pokedex.
type Pokedex struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID           int    `json:"id"`
	IsMainSeries bool   `json:"is_main_series"`
	Name         string `json:"name"`
	Names        []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEntries []struct {
		EntryNumber    int `json:"entry_number"`
		PokemonSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"pokemon_entries"`
	Region        interface{}   `json:"region"`
	VersionGroups []interface{} `json:"version_groups"`
}

// Version is a single version.
type Version struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	VersionGroup struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_group"`
}

// VersionGroup is a single version group.
type VersionGroup struct {
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID               int `json:"id"`
	MoveLearnMethods []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_learn_methods"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Pokedexes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokedexes"`
	Regions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"regions"`
	Versions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"versions"`
}
