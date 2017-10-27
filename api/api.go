package api

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"strings"

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sendResponse(w, r, timetable)
}

func sendResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	var writer io.Writer
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Add("Content-Encoding", "gzip")
		gzw := gzip.NewWriter(w)
		defer gzw.Close()
		writer = gzw
	} else {
		writer = w
	}

	w.Header().Add("content-type", "application/json")
	err := json.NewEncoder(writer).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
