package example

import (
	"io"
	"log"
)

// RotConverter rotates the byte data by some offset
type RotConverter struct {
	Rotate uint8
}

// Convert the input and write the rotated output
func (c RotConverter) Convert(r *io.PipeReader, w *io.PipeWriter) {
	defer r.Close()
	defer w.Close()
	buf := make([]byte, 64)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			return
		}
		for i := 0; i < n; i++ {
			buf[i] += c.Rotate
		}
		if _, err := w.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
	}
}
