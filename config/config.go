package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/foomo/htpasswd"
)

// Port is the port which the application listens on
var Port int

// Secret is the crypto secret for the application
var Secret string

// Passwords is used to log into the application
var Passwords htpasswd.HashedPasswords

// Config is the configuration object for the application
type Config struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
}

// Parse parses the application configuration file
func Parse() {
	var path = flag.String("config", "config.json", "the config file for the application")
	var port = flag.Int("port", 2800, "the port for the server to listen on")
	var htpass = flag.String("htpass", ".htpasswd", "set the htpassword file")

	file, err := os.Open(*path)

	if err != nil {
		log.Println("Could not read configuration file")
		log.Fatal(err)
	}

	conf := new(Config)
	dec := json.NewDecoder(file)
	if err = dec.Decode(conf); err != nil {
		log.Fatal(err)
	}

	if *port != 2800 {
		conf.Port = *port
	}

	Port = conf.Port
	Secret = conf.Secret

	pass, err := htpasswd.ParseHtpasswdFile(*htpass)
	if err != nil {
		log.Fatal(err)
	}

	Passwords = pass

}
