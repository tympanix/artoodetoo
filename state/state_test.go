package state_test

import (
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/state"
)

func timeTest(t *testing.T, fn func(), mili int) {
	done := make(chan bool)

	go func() {
		fn()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(time.Duration(mili) * time.Millisecond):
		t.FailNow()
	}
}

func timeFail(t *testing.T, fn func(), mili int) {
	done := make(chan bool)

	go func() {
		fn()
		done <- true
	}()

	select {
	case <-done:
		t.FailNow()
	case <-time.After(time.Duration(mili) * time.Millisecond):
	}
}

func TestStateStaticPutGet(t *testing.T) {
	s := state.New()

	s.Put("test", 10)

	var i int

	timeTest(t, func() {
		if err := s.Query("test", &i); err != nil {
			t.Error(err)
		}
	}, 100)

	assert.Equal(t, i, 10)

}

func TestStateMultiType(t *testing.T) {
	s := state.New()

	s.Put("test", "hej", 10, 3.1415)

	var l string
	var i int
	var f float64

	timeTest(t, func() {
		if err := s.Query("test", &l, &i, &f); err != nil {
			t.Error(err)
		}
	}, 100)

	assert.Equal(t, l, "hej")
	assert.Equal(t, i, 10)
	assert.Equal(t, f, 3.1415)
}

func TestStateSameKey(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
		127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193}

	s := state.New()

	go func() {
		for _, number := range primes {
			s.Put("primes", number)
		}
	}()

	var found int32

	for range primes {
		timeTest(t, func() {
			var prime int
			if err := s.Get("primes", &prime); err != nil {
				t.Error(err)
			}
			atomic.AddInt32(&found, 1)
		}, 100)
	}

	assert.Equal(t, int(found), len(primes))

	timeFail(t, func() {
		var prime int
		if err := s.Get("primes", &prime); err != nil {
			t.Error(err)
		}
	}, 32)

}

func TestStateBlock(t *testing.T) {
	s := state.New()

	s.Put("block", true)

	timeTest(t, func() {
		var b bool
		if err := s.Get("block", &b); err != nil {
			t.Error(err)
		}
	}, 100)

	timeFail(t, func() {
		var b bool
		if err := s.Get("block", &b); err != nil {
			t.Error(err)
		}
	}, 32)
}

func TestStatePredicate(t *testing.T) {
	s := state.New()

	s.Put("string", "bbb")
	s.Put("string", "aaa")

	p := func(s string) bool {
		return strings.Contains(s, "a")
	}

	timeTest(t, func() {
		if err := s.Get("string", p); err != nil {
			t.Error(err)
		}
	}, 100)
}

func TestNilValues(t *testing.T) {
	s := state.New()

	err := s.Put("nil", nil)
	assert.Error(t, err)
}
