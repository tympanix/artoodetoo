package data

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sync"
)

// FromURL return a new stream from a http resource
func FromURL(url string) Stream {
	return &httpStream{
		url:      url,
		mimetype: mime.TypeByExtension(filepath.Ext(url)),
		buffer:   new(bytes.Buffer),
		once:     new(sync.Once),
		lock:     new(sync.Mutex),
	}
}

type httpStream struct {
	url      string
	mimetype string
	buffer   *bytes.Buffer
	once     *sync.Once
	lock     *sync.Mutex
}

func (h *httpStream) getData() {
	resp, err := http.Get(h.url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	h.buffer.ReadFrom(resp.Body)
}

func (h *httpStream) NextReader() io.Reader {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.once.Do(h.getData)

	return bytes.NewReader(h.buffer.Bytes())
}

func (h *httpStream) Mimetype() string {
	return h.mimetype
}
