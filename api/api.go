package api

import (
	"encoding/json"
	"net/http"

	"github.com/hiro511/api-test/db"
)

// Register api
func Register(mux *http.ServeMux) {
	mux.HandleFunc("/timetable", timetableHandler)
}

func timetableHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTimetable(w, r)
	default:
		http.Error(w, "unexpected method", http.StatusNotFound)
	}
}

func getTimetable(w http.ResponseWriter, r *http.Request) {
	timetable, err := db.GetTimetable()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Cache-Control", "public,max-age=100")

	sendResponse(w, r, timetable)
}

func sendResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("content-type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
