package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Tympanix/automato/api"
	"github.com/Tympanix/automato/config"
	"github.com/Tympanix/automato/storage"
)

const (
	apiroot  = "/api"
	authroot = "/auth"
)

func main() {
	// Parse application configuration
	config.Parse()

	// Set up storage driver
	initStorage()

	// Set up api handler
	http.Handle(apiroot+"/", http.StripPrefix(apiroot, api.API))

	// Set up file server for static files
	fs := http.FileServer(http.Dir("web/dist"))
	http.Handle("/", fs)

	// Serve the web server
	log.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(addr(), nil))
}

func addr() string {
	return ":" + strconv.Itoa(config.Port)
}

func initStorage() {
	json, err := storage.NewJSONFile("./store.json")
	if err != nil {
		log.Fatal(err)
	}
	storage.Register(json)
	tasks, events := storage.Load()
	log.Printf("Loaded %d task(s)\n", tasks)
	log.Printf("Loaded %d event(s)\n", events)
}
