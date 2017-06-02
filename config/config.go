package config

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Tympanix/artoodetoo/types"
	"github.com/foomo/htpasswd"
)

// Port is the port which the application listens on
var Port int

// Secret is the crypto secret for the application
var Secret []byte

// Passwords is used to log into the application
var Passwords htpasswd.HashedPasswords

// Config is the configuration object for the application
type Config struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
}

// Parse parses the application configuration file
func Parse(args types.AppArgs) {
	Port = args.Port()

	file, err := os.Open(args.SecretPath())

	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(file); err != nil {
		log.Fatal(err)
	}
	Secret = buf.Bytes()

	pass, err := htpasswd.ParseHtpasswdFile(args.HtpasswdPath())
	if err != nil {
		log.Fatal(err)
	}

	Passwords = pass

}

// AddUser adds a new user to the system
func AddUser(path string, username string, password string) error {
	return htpasswd.SetPassword(path, username, password, htpasswd.HashBCrypt)
}

// DeleteUser removes a user from the application
func DeleteUser(path string, username string) error {
	return htpasswd.RemoveUser(path, username)
}

// GetUsers returns the users of the application
func GetUsers(path string) (htpasswd.HashedPasswords, error) {
	return htpasswd.ParseHtpasswdFile(path)
}

// GenerateSecret generates a new application secret
func GenerateSecret(path string, length int, force bool) (err error) {

	min := 64
	max := 1 << 14

	if length < min || length > max {
		return fmt.Errorf("Length must be between %v and %v bytes", min, max)
	}

	if _, err = os.Stat(path); err == nil {
		if !force {
			return errors.New("File already exists. Use force argument to overwrite")
		}
	}

	file, err := os.Create(path)

	if err != nil {
		return
	}

	defer file.Close()
	b := make([]byte, length)
	_, err = rand.Read(b)

	if err != nil {
		return
	}

	if err = binary.Write(file, binary.LittleEndian, b); err != nil {
		return
	}

	return nil
}
