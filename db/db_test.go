package db

import (
	"testing"

	"github.com/hiro511/api-test/config"
)

func TestGetTimetable(t *testing.T) {
	config.SetDBBasePath("../")
	timetable, err := GetTimetable()
	if err != nil {
		t.Fatalf("err should be nil, err = %s", err.Error())
	}

	if timetable == nil {
		t.Fatalf("timetable should not be nil")
	}
}
