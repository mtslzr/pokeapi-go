package structs

// Encounter is a single encounter
type Encounter struct {
	MinLevel       int                     `json:"min_level"`
	MaxLevel       int                     `json:"max_level"`
	ConditionValue EncounterConditionValue `json:"condition_value"`
	Chance         int                     `json:"chance"`
	Method         EncounterMethod         `json:"method"`
}

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

// VersionEncounterDetail is a single encounter detail.
type VersionEncounterDetail struct {
	Version         Version     `json:"version"`
	MaxChance       int         `json:"max_chance"`
	EcounterDetails []Encounter `json:"encounter_details"`
}
