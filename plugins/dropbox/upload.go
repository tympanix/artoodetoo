package dropbox

import (
	"io/ioutil"
	"log"

	"github.com/Tympanix/automato/data"
	"github.com/Tympanix/automato/unit"
	"github.com/stacktic/dropbox"
)

// UploadFile uploads a file to dropbox
type UploadFile struct {
	File        data.Stream `io:"input"`
	Token       string      `io:"input"`
	Overwrite   bool        `io:"input"`
	Destination string      `io:"input"`
}

func init() {
	unit.Register(new(UploadFile))
}

// Describe what it does
func (UploadFile) Describe() string {
	return "Upload a file to dropbox"
}

// Execute uploads the file
func (u *UploadFile) Execute() {
	db := dropbox.NewDropbox()

	db.SetAppInfo("rr6d38abezy4l6u", "u99czi59xv4tjwb")
	db.SetAccessToken(string(u.Token))

	reader := ioutil.NopCloser(u.File.NextReader())
	_, err := db.UploadByChunk(reader, 1<<14, u.Destination, u.Overwrite, "")

	if err != nil {
		log.Println(err)
	}
}
