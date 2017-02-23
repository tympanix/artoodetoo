package example

import (
	"io"
	"log"
)

// DupEvent dusplicates a string
type DupEvent struct {
	Length int
	String string
}

// Trigger start the dubpication producing A characters
func (d DupEvent) Trigger(w *io.PipeWriter) {
	defer w.Close()
	for i := 0; i < d.Length; i++ {
		if _, err := w.Write([]byte(d.String)); err != nil {
			log.Fatal(err)
		}
	}
}
