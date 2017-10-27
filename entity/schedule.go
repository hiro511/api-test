package entity

type (
	// ChannelSchedule data model.
	ChannelSchedule struct {
		ID               string     `json:"id"`
		Date             string     `json:"date"`
		Slots            []Slot     `json:"slots"`
		TableStartAt     int64      `json:"tableStartAt"`
		TableEndAt       int64      `json:"tableEndAt"`
		Highlight        string     `json:"highlight"`
		TableHighlight   string     `json:"tableHighlight"`
		DetailHighlight  string     `json:"detailHighlight"`
		Content          string     `json:"content"`
		DisplayProgramID string     `json:"displayProgramId"`
		Mark             struct{}   `json:"mark"`
		Flags            struct{}   `json:"flags"`
		ChannelID        string     `json:"channelID"`
		SlotGroup        SlotGroup  `json:"slotGroup"`
		Links            []struct{} `json:"links"`
		SharedLink       SharedLink `json:"sharedLink"`
	}
)
