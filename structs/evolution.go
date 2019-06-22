package structs

// EvolutionChain is a single evolution chain.
type EvolutionChain struct {
	BabyTriggerItem interface{} `json:"baby_trigger_item"`
	Chain           struct {
		EvolutionDetails []interface{} `json:"evolution_details"`
		EvolvesTo        []struct {
			EvolutionDetails []struct {
				Gender                interface{} `json:"gender"`
				HeldItem              interface{} `json:"held_item"`
				Item                  interface{} `json:"item"`
				KnownMove             interface{} `json:"known_move"`
				KnownMoveType         interface{} `json:"known_move_type"`
				Location              interface{} `json:"location"`
				MinAffection          interface{} `json:"min_affection"`
				MinBeauty             interface{} `json:"min_beauty"`
				MinHappiness          interface{} `json:"min_happiness"`
				MinLevel              int         `json:"min_level"`
				NeedsOverworldRain    bool        `json:"needs_overworld_rain"`
				PartySpecies          interface{} `json:"party_species"`
				PartyType             interface{} `json:"party_type"`
				RelativePhysicalStats interface{} `json:"relative_physical_stats"`
				TimeOfDay             string      `json:"time_of_day"`
				TradeSpecies          interface{} `json:"trade_species"`
				Trigger               struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"trigger"`
				TurnUpsideDown bool `json:"turn_upside_down"`
			} `json:"evolution_details"`
			EvolvesTo []struct {
				EvolutionDetails []struct {
					Gender                interface{} `json:"gender"`
					HeldItem              interface{} `json:"held_item"`
					Item                  interface{} `json:"item"`
					KnownMove             interface{} `json:"known_move"`
					KnownMoveType         interface{} `json:"known_move_type"`
					Location              interface{} `json:"location"`
					MinAffection          interface{} `json:"min_affection"`
					MinBeauty             interface{} `json:"min_beauty"`
					MinHappiness          interface{} `json:"min_happiness"`
					MinLevel              int         `json:"min_level"`
					NeedsOverworldRain    bool        `json:"needs_overworld_rain"`
					PartySpecies          interface{} `json:"party_species"`
					PartyType             interface{} `json:"party_type"`
					RelativePhysicalStats interface{} `json:"relative_physical_stats"`
					TimeOfDay             string      `json:"time_of_day"`
					TradeSpecies          interface{} `json:"trade_species"`
					Trigger               struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					} `json:"trigger"`
					TurnUpsideDown bool `json:"turn_upside_down"`
				} `json:"evolution_details"`
				EvolvesTo []interface{} `json:"evolves_to"`
				IsBaby    bool          `json:"is_baby"`
				Species   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"species"`
			} `json:"evolves_to"`
			IsBaby  bool `json:"is_baby"`
			Species struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"species"`
		} `json:"evolves_to"`
		IsBaby  bool `json:"is_baby"`
		Species struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"species"`
	} `json:"chain"`
	ID int `json:"id"`
}

// EvolutionTrigger is a single evolution trigger.
type EvolutionTrigger struct {
	ID    int    `json:"id"`
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
}
