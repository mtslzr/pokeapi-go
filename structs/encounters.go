package structs

// EncounterMethod is a single encounter method.
type EncounterMethod struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Order int `json:"order"`
}

// EncounterCondition is a single encounter condition.
type EncounterCondition struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Values []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"values"`
}

// EncounterConditionValue is a single encounter condition value.
type EncounterConditionValue struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Values []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"values"`
}
