package entity

type (
	// Channel data model.
	Channel struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		Order       int32       `json:"int"`
		MedisStatus MediaStatus `json:"mediaStatus"`
	}

	// MediaStatus data model.
	MediaStatus struct {
		DashISOFF bool `json:"dashISOFF,omitempty"`
	}
)
