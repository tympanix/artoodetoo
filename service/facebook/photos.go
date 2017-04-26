package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Tympanix/automato/event"
)

// Photos listen for new photo uploads or tags on facebook
type Photos struct {
	event.Base
	last     time.Time
	Token    string  `io:"input"`
	Interval float64 `io:"input"`

	ImgURL string `io:"output"`
}

func init() {
	event.Register(new(Photos))
}

// Describe describes what the facebook photo event does
func (p *Photos) Describe() string {
	return "Triggers whenever you are tagged or uploads a photo on facebook"
}

// Listen starts listening for new photo uploads and tags on facebook
func (p *Photos) Listen() error {
	if p.Interval < 1000 {
		return errors.New("Interval must be above 1 second")
	}

	err := p.updateLast()

	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(p.Interval) * time.Millisecond)
	for range ticker.C {
		photos, err := p.getPhotos()

		if err != nil {
			return nil
		}

		log.Printf("Found %d photos", len(photos))

		for _, photo := range photos {
			if photo.Created.After(p.last) {
				details, err := p.getPhotoDeails(photo.ID)
				if err != nil {
					return err
				}
				if len(details.Images) > 0 {
					p.ImgURL = details.Images[0].Source
					fmt.Println(details.Images[0].Source)
					p.Trigger()
				} else {
					return errors.New("No images found")
				}
			}
		}
	}

	return nil
}

type photoInfo struct {
	CreatedTime string    `json:"created_time"`
	ID          string    `json:"id"`
	Created     time.Time `json:"-"`
}

func (p *photoInfo) UnmarshalJSON(data []byte) error {
	type alias photoInfo
	var photo alias
	if err := json.Unmarshal(data, &photo); err != nil {
		return err
	}
	time, err := time.Parse(TIME, p.CreatedTime)
	if err != nil {
		return err
	}
	p.Created = time
	*p = photoInfo(photo)
	return nil
}

func (p *Photos) updateLast() error {
	photos, err := p.getPhotos()
	if err != nil {
		return err
	}
	if len(photos) > 0 {
		p.last = photos[0].Created
	} else {
		p.last = time.Now()
	}
	return nil
}

func (p *Photos) getPhotos() (photos []*photoInfo, err error) {
	if _, err = getData("/me/photos", p.Token, &photos); err != nil {
		return
	}

	return
}

type photoDetails struct {
	ID     string   `json:"id"`
	Images []*image `json:"images"`
}

type image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Source string `json:"source"`
}

func (p *Photos) getPhotoDeails(id string) (details photoDetails, err error) {
	if err = getNode(id, p.Token, &details); err != nil {
		return
	}
	return
}
