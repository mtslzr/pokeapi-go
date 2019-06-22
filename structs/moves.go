package structs

// Move is a single move.
type Move struct {
	Accuracy      int `json:"accuracy"`
	ContestCombos struct {
		Normal struct {
			UseAfter  interface{} `json:"use_after"`
			UseBefore []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"use_before"`
		} `json:"normal"`
		Super struct {
			UseAfter  interface{} `json:"use_after"`
			UseBefore interface{} `json:"use_before"`
		} `json:"super"`
	} `json:"contest_combos"`
	ContestEffect struct {
		URL string `json:"url"`
	} `json:"contest_effect"`
	ContestType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contest_type"`
	DamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"damage_class"`
	EffectChance  interface{}   `json:"effect_chance"`
	EffectChanges []interface{} `json:"effect_changes"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID       int           `json:"id"`
	Machines []interface{} `json:"machines"`
	Meta     struct {
		Ailment struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ailment"`
		AilmentChance int `json:"ailment_chance"`
		Category      struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"category"`
		CritRate     int         `json:"crit_rate"`
		Drain        int         `json:"drain"`
		FlinchChance int         `json:"flinch_chance"`
		Healing      int         `json:"healing"`
		MaxHits      interface{} `json:"max_hits"`
		MaxTurns     interface{} `json:"max_turns"`
		MinHits      interface{} `json:"min_hits"`
		MinTurns     interface{} `json:"min_turns"`
		StatChance   int         `json:"stat_chance"`
	} `json:"meta"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PastValues         []interface{} `json:"past_values"`
	Power              int           `json:"power"`
	Pp                 int           `json:"pp"`
	Priority           int           `json:"priority"`
	StatChanges        []interface{} `json:"stat_changes"`
	SuperContestEffect struct {
		URL string `json:"url"`
	} `json:"super_contest_effect"`
	Target struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"target"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

// MoveAilment is a single move ailment.
type MoveAilment struct {
	ID    int `json:"id"`
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
}

// MoveBattleStyle is a single move battle style.
type MoveBattleStyle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

// MoveCategory is a single move category.
type MoveCategory struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID    int `json:"id"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
	Name string `json:"name"`
}

// MoveDamageClass is a single move damage class.
type MoveDamageClass struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID    int `json:"id"`
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
}

// MoveLearnMethod is a single move learn method.
type MoveLearnMethod struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	VersionGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_groups"`
}

// MoveTarget is a single move target.
type MoveTarget struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID    int `json:"id"`
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
}
