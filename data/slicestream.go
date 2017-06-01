package data

import (
	"bytes"
	"io"
)

// FromByteArray return a new stream from a byte object
func FromByteArray(data []byte, mimetype string) Stream {
	return &byteStream{
		mimetype: mimetype,
		data:     data,
	}
}

type byteStream struct {
	mimetype string
	data     []byte
}

func (b *byteStream) NextReader() io.Reader {
	return bytes.NewReader(b.data)
}

func (b *byteStream) Mimetype() string {
	return b.mimetype
}
