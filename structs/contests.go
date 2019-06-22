package structs

// ContestType is a single contest type.
type ContestType struct {
	BerryFlavor struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"berry_flavor"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Color    string `json:"color"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

// ContestEffect is a single contest effect.
type ContestEffect struct {
	Appeal        int `json:"appeal"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	ID  int `json:"id"`
	Jam int `json:"jam"`
}

// SuperContestEffect is a single super contest effect.
type SuperContestEffect struct {
	Appeal            int `json:"appeal"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	ID    int `json:"id"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
}
