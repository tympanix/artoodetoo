package example

import (
	"bufio"
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
	bw := bufio.NewWriter(w)

	for i := 0; i < d.Length; i++ {
		if _, err := bw.Write([]byte(d.String)); err != nil {
			log.Fatal(err)
		}
	}
	bw.Flush()
}
