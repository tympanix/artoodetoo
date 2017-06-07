package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tympanix/artoodetoo/logger"
)

func init() {
	API.Handle("/logs", auth(getLogs)).Methods("GET")
	API.Handle("/logs", auth(clearLogs)).Methods("DELETE")
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	SetJSON(w)
	time := r.URL.Query().Get("t")

	unix, err := strconv.ParseInt(time, 10, 64)

	if err != nil {
		http.Error(w, "Time param required for log view", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(logger.Get(unix))
}

func clearLogs(w http.ResponseWriter, r *http.Request) {
	SetJSON(w)
	logger.Clear()
}
