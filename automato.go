package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Tympanix/automato/api"
	"github.com/Tympanix/automato/storage"

	_ "github.com/Tympanix/automato/example"
	_ "github.com/Tympanix/automato/service/cron"
)

const (
	apiroot = "/api"
	port    = 2800
)

func main() {
	// Set up storage driver
	initStorage()

	// Set up api handler
	http.Handle(apiroot+"/", http.StripPrefix(apiroot, api.API))

	// Set up file server for static files
	fs := http.FileServer(http.Dir("web/dist"))
	http.Handle("/", fs)

	// Serve the web server
	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(addr(), nil))
}

func addr() string {
	return ":" + strconv.Itoa(port)
}

func initStorage() {
	json, err := storage.NewJSONFile("./store.json")
	if err != nil {
		log.Fatal(err)
	}
	storage.Register(json)
	loaded := storage.Load()
	log.Printf("Loaded %d tasks\n", loaded)
}
