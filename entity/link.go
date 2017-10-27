package entity

type (
	// SharedLink data model.
	SharedLink struct {
		Twitter  string `json:"twitter"`
		Facebook string `json:"facebook"`
		Google   string `json:"google"`
		Line     string `json:"line"`
		Copy     string `json:"copy"`
		Screen   string `json:"screen"`
	}
)
