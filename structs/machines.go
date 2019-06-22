package structs

// Machine is a single machine.
type Machine struct {
	ID   int `json:"id"`
	Item struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"item"`
	Move struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move"`
	VersionGroup struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_group"`
}
