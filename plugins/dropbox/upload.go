package dropbox

import (
	"log"

	"github.com/Tympanix/artoodetoo/data"
	"github.com/Tympanix/artoodetoo/unit"
	"github.com/stacktic/dropbox"
)

// UploadFile uploads a file to dropbox
type UploadFile struct {
	dropboxStyle
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
func (u *UploadFile) Execute() error {
	db := dropbox.NewDropbox()

	db.SetAppInfo("rr6d38abezy4l6u", "u99czi59xv4tjwb")
	db.SetAccessToken(string(u.Token))

	reader, err := u.File.NewReader()
	if err != nil {
		return err
	}
	_, err = db.UploadByChunk(reader, 1<<14, u.Destination, u.Overwrite, "")

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
