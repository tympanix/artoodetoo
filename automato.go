package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Tympanix/automato/api"

	_ "github.com/Tympanix/automato/example"
)

const (
	apiroot = "/api"
	port    = 2800
)

func main() {
	// Set up api handler
	http.Handle(apiroot+"/", http.StripPrefix(apiroot, api.API))

	// Set up file server for static files
	fs := http.FileServer(http.Dir("web/dist"))
	http.Handle("/", fs)

	// Serve the web server
	log.Printf("Listeing on port %d\n", port)
	log.Fatal(http.ListenAndServe(addr(), nil))
}

func addr() string {
	return ":" + strconv.Itoa(port)
}
