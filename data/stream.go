package data

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const tmpDir = "tmp"

// Stream is an object which produces readers from which to obtain data
type Stream interface {
	Mimetype() string
	NewReader() (io.ReadCloser, error)
}

// StreamBuffer is a buffer with a file as its underlying buffer
type StreamBuffer struct {
	*sync.Cond
	file    *os.File
	readers int
	closed  bool
	ended   bool
}

// StreamReader is a thread safe reader which reads from a stream buffer
type StreamReader struct {
	file *os.File
	buf  *StreamBuffer
}

// StreamEndError is returned when the stream has ended
type StreamEndError struct{}

func (*StreamEndError) Error() string {
	return "The data stream has ended"
}

// Close closes the underlying file
func (sr *StreamReader) Close() error {
	if err := sr.file.Close(); err != nil {
		return err
	}
	sr.buf.Lock()
	defer sr.buf.Unlock()
	sr.buf.readers--
	if sr.buf.readers == 0 && sr.buf.ended {
		return sr.buf.Cleanup()
	}
	return nil
}

func (sr *StreamReader) Read(p []byte) (n int, err error) {

	n, err = sr.file.Read(p)

	for err == io.EOF {
		sr.buf.Lock()
		if sr.buf.closed {
			sr.buf.Unlock()
			return
		}
		sr.buf.Wait()
		sr.buf.Unlock()
		n, err = sr.file.Read(p)
	}

	return
}

// NewReader returns a new reader to the buffer
func (mb *StreamBuffer) NewReader() (io.ReadCloser, error) {
	mb.Lock()
	defer mb.Unlock()
	if mb.ended {
		return nil, new(StreamEndError)
	}
	f, err := os.Open(mb.file.Name())
	if err != nil {
		return nil, err
	}
	mb.readers++
	return &StreamReader{f, mb}, nil
}

// Lock locks the data stream for manipulation
func (mb *StreamBuffer) Lock() {
	mb.L.Lock()
}

// Unlock unlocks the data stream
func (mb *StreamBuffer) Unlock() {
	mb.L.Unlock()
}

// Cleanup removes the temporary file
func (mb *StreamBuffer) Cleanup() error {
	_, err := os.Stat(mb.file.Name())
	if err == nil {
		err := os.Remove(mb.file.Name())
		log.Println(err)
		return err
	}
	return nil
}

// File returns the underlying data file
func (mb *StreamBuffer) File() *os.File {
	return mb.file
}

// Write write data to the buffer
func (mb *StreamBuffer) Write(b []byte) (n int, err error) {
	n, err = mb.file.Write(b)
	mb.Broadcast()
	return
}

// Close closes the stream. No new readers allowed
func (mb *StreamBuffer) Close() (err error) {
	mb.Lock()
	defer mb.Unlock()
	if err = mb.file.Close(); err != nil {
		return
	}
	mb.closed = true
	mb.Broadcast()
	return
}

// End ends the stream and allows no more readers
func (mb *StreamBuffer) End() (err error) {
	mb.Lock()
	defer mb.Unlock()
	mb.ended = true
	if mb.readers == 0 {
		return mb.Cleanup()
	}
	return nil
}

// NewStreamBuffer returns a new stream buffer
func NewStreamBuffer(prefix string) (*StreamBuffer, error) {
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		os.Mkdir(tmpDir, os.ModeDir)
	}
	f, err := ioutil.TempFile(tmpDir, prefix)
	if err != nil {
		return nil, err
	}
	return &StreamBuffer{
		file:    f,
		Cond:    sync.NewCond(new(sync.Mutex)),
		readers: 0,
	}, nil
}
