package data_test

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/Tympanix/artoodetoo/data"
	"github.com/stretchr/testify/assert"
)

const stream = `Lorem ipsum dolor sit amet, consectetur adipiscing elit,
    sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi
    ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit
    in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur
    sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
    mollit anim id est laborum.`

func TestStreamBuffer(t *testing.T) {
	const READERS = 3
	sb, err := data.NewStreamBuffer("test001")
	assert.Nil(t, err)

	begin := new(sync.WaitGroup)
	begin.Add(READERS)
	wg := new(sync.WaitGroup)
	wg.Add(READERS)

	for i := 0; i < 3; i++ {
		go func() {
			r, err := sb.NewReader()
			begin.Done()
			defer wg.Done()
			defer r.Close()
			assert.Nil(t, err)
			data, err := ioutil.ReadAll(r)
			assert.Nil(t, err)
			assert.Equal(t, string(data), stream)
		}()
	}

	r := strings.NewReader(stream)
	io.Copy(sb, r)
	begin.Wait()
	err = sb.Close()
	assert.Nil(t, err)
	wg.Wait()

	sb.End()
	_, err = os.Stat(sb.File().Name())
	assert.True(t, os.IsNotExist(err))

}

func TestSlowStream(t *testing.T) {

}
