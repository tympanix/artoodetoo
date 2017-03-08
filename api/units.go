package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tympanix/automato/util"
)

func init() {
	API.HandleFunc("/units", listUnits).Methods("GET")
}

func listUnits(w http.ResponseWriter, r *http.Request) {
	SetJSON(w)
	json.NewEncoder(w).Encode(util.AllMetas())
}
