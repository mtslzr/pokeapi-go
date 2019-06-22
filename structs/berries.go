package structs

// Berry is a single berry.
type Berry struct {
	Firmness struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"firmness"`
	Flavors []struct {
		Flavor struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"flavor"`
		Potency int `json:"potency"`
	} `json:"flavors"`
	GrowthTime int `json:"growth_time"`
	ID         int `json:"id"`
	Item       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"item"`
	MaxHarvest       int    `json:"max_harvest"`
	Name             string `json:"name"`
	NaturalGiftPower int    `json:"natural_gift_power"`
	NaturalGiftType  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"natural_gift_type"`
	Size        int `json:"size"`
	Smoothness  int `json:"smoothness"`
	SoilDryness int `json:"soil_dryness"`
}

// BerryFirmness is a single berry firmness.
type BerryFirmness struct {
	Berries []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"berries"`
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

// BerryFlavor is a single berry flavor.
type BerryFlavor struct {
	Berries []struct {
		Berry struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"berry"`
		Potency int `json:"potency"`
	} `json:"berries"`
	ContestType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contest_type"`
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
