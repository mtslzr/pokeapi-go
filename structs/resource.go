package structs

// Resource is a resource list for an endpoint.
type Resource struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

// Result is a resource list result.
type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
