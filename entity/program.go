package entity

type (
	// Program data model.
	Program struct {
		ID           string       `json:"id"`
		Episode      Episode      `json:"episode"`
		Credit       Credit       `json:"credit"`
		Series       Series       `json:"series"`
		ProvidedInfo ProvidedInfo `json:"providedInfo"`
	}

	// Episode data model.
	Episode struct {
		Sequence int32 `json:"sequence"`
	}

	// Credit data model.
	Credit struct {
		Casts      []string `json:"casts"`
		Crews      []string `json:"crews"`
		Copyrights []string `json:"copyrights"`
	}

	// Series data model.
	Series struct {
		ID         string     `json:"id"`
		ThemeColor ThemeColor `json:"themeColor"`
		UpdatedAt  int64      `json:"updatedAt"`
	}

	// ThemeColor data model.
	ThemeColor struct {
		Primary    string `json:"primary"`
		Secondary  string `json:"secondary"`
		Detail     string `json:"detail"`
		Background string `json:"background"`
	}

	// ProvidedInfo data model.
	ProvidedInfo struct {
		ThumbImg  string `json:"thumbImg"`
		UpdatedAt int64  `json:"updatedAt"`
	}
)
