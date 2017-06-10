package data

import (
	"bytes"
	"io"
	"io/ioutil"
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

func (b *byteStream) NewReader() (io.ReadCloser, error) {
	r := bytes.NewReader(b.data)
	return ioutil.NopCloser(r), nil
}

func (b *byteStream) Mimetype() string {
	return b.mimetype
}
