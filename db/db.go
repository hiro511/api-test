package db

import (
	"encoding/json"
	"os"

	"github.com/hiro511/api-test/config"
	"github.com/hiro511/api-test/entity"
)

const (
	timetablePath = "data/timetable.json"
)

// GetTimetable from database
func GetTimetable() (*entity.Timetable, error) {
	file, err := os.Open(config.DBBasePath() + timetablePath)
	if err != nil {
		return nil, err
	}

	timetable := entity.Timetable{}
	err = json.NewDecoder(file).Decode(&timetable)
	if err != nil {
		return nil, err
	}

	return &timetable, nil
}
