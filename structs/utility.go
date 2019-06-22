package structs

// Language is a single language.
type Language struct {
	ID      int    `json:"id"`
	Iso3166 string `json:"iso3166"`
	Iso639  string `json:"iso639"`
	Name    string `json:"name"`
	Names   []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Official bool `json:"official"`
}
