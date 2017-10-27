package entity

type (
	// Timetable data model.
	Timetable struct {
		Channels        []Channel         `json:"channels"`
		ChannelSchedule []ChannelSchedule `json:"channelSchedules"`
	}
)
