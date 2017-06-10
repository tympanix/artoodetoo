package drive

import (
	"log"

	"golang.org/x/oauth2"

	"github.com/Tympanix/artoodetoo/data"
	"github.com/Tympanix/artoodetoo/plugins/google"
	"github.com/Tympanix/artoodetoo/unit"
	"google.golang.org/api/drive/v3"
)

// UploadFile uploads a file to dropbox
type UploadFile struct {
	driveStyle
	File     data.Stream  `io:"input"`
	Filename string       `io:"input"`
	Token    google.Token `io:"input"`
}

func init() {
	unit.Register(new(UploadFile))
}

// Describe what it does
func (UploadFile) Describe() string {
	return "Upload a file to google drive"
}

// Execute uploads the file
func (u *UploadFile) Execute() error {

	token := &oauth2.Token{
		AccessToken: string(u.Token),
	}

	client := oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(token))

	driveService, err := drive.New(client)

	if err != nil {
		log.Println(err)
		return err
	}

	file := &drive.File{
		Name: u.Filename,
	}

	reader, err := u.File.NewReader()
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = driveService.Files.Create(file).Media(reader).Do()

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
