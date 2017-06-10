package data

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sync"
)

// FromURL return a new stream from a http resource
func FromURL(url string) (s Stream, err error) {
	stream, err := NewStreamBuffer("http")
	if err != nil {
		return
	}
	return &httpStream{
		StreamBuffer: stream,
		url:          url,
		mimetype:     mime.TypeByExtension(filepath.Ext(url)),
		once:         new(sync.Once),
	}, nil
}

type httpStream struct {
	*StreamBuffer
	url      string
	mimetype string
	once     *sync.Once
}

func (h *httpStream) getData() {
	resp, err := http.Get(h.url)
	if err != nil {
		return
	}
	defer h.Close()
	defer resp.Body.Close()
	io.Copy(h, resp.Body)
}

func (h *httpStream) String() string {
	return fmt.Sprintf("HTTP data stream: %s", h.url)
}

func (h *httpStream) NewReader() (io.ReadCloser, error) {
	go h.once.Do(h.getData)
	return h.StreamBuffer.NewReader()
}

func (h *httpStream) Mimetype() string {
	return h.mimetype
}
