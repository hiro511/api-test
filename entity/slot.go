package entity

type (
	// Slot data model.
	Slot struct {
		ID       string    `json:"id"`
		Title    string    `json:"title"`
		StartAt  int64     `json:"startAt"`
		EndAt    int64     `json:"endAt"`
		Programs []Program `json:"programs"`
	}

	// SlotGroup data model.
	SlotGroup struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		LastSlotID string `json:"lastSlotId"`
	}
)
