package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Equal tests whether two objects are equal using the equality operator
func Equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v got %v", a, b)
	}
}

// NotEqual tests whether two obect are different using the equality operator
func NotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Expected %v to be different from %v", a, b)
	}
}

// True test whether the object passed is true
func True(t *testing.T, a bool) {
	if !a {
		t.Errorf("Expected true got false")
	}
}

// DeepEqual uses reflection to test whether two objects deep equal each other
func DeepEqual(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected deep equal %v got %v", a, b)
	}
}

// NotError test whether the supplied error argument is nil or not
func NotError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// Error tests whether an error occured
func Error(t *testing.T, err error) {
	if err == nil {
		t.Error("Expected error but nil")
	}
}

// ErrorContains tests whether the error message contins all phrases
func ErrorContains(t *testing.T, err error, phrases ...string) {
	if err == nil {
		t.Errorf("Expected error but got nil")
		return
	}
	for _, str := range phrases {
		if !strings.Contains(err.Error(), str) {
			t.Errorf("Expected error to contain '%s' in '%s'", str, err.Error())
			return
		}
	}
}
