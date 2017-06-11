package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/Tympanix/artoodetoo/api"
	"github.com/Tympanix/artoodetoo/cli"
	"github.com/Tympanix/artoodetoo/config"
	"github.com/Tympanix/artoodetoo/storage"
	"github.com/Tympanix/artoodetoo/types"
	"github.com/gorilla/mux"
)

const (
	static = "web/dist/"
)

func main() {
	cli.Run(serve)
}

func serve(args types.AppArgs) {
	// Parse application configuration
	config.Parse(args)

	// Set up storage driver
	initStorage()

	router := mux.NewRouter()

	// Set up api handler
	//router.Handle(apiroot, http.StripPrefix(apiroot, api.API))
	router.PathPrefix("/api/").Handler(http.StripPrefix("/api", api.API))

	// Set up file server for static files
	fs := newFileServer(static)
	router.PathPrefix("/").Handler(fs)

	// Serve the web server
	log.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(addr(), router))
}

func addr() string {
	return ":" + strconv.Itoa(config.Port)
}

type fileServer struct {
	http.Handler
}

func newFileServer(path string) *fileServer {
	return &fileServer{http.FileServer(http.Dir(static))}
}

func (f *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file := path.Clean(r.URL.Path)
	if fi, err := os.Stat(path.Join(static, file)); err == nil && !fi.IsDir() {
		f.Handler.ServeHTTP(w, r)
	} else {
		http.ServeFile(w, r, path.Join(static, "index.html"))
	}
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
