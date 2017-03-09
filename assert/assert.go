package assert

import (
	"reflect"
	"testing"
)

// Equal tests whether two objects are equal using the equality operator
func Equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v got %v", a, b)
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

func Value(t *testing.T, a reflect.Value, b interface{}) {
	bVal := reflect.ValueOf(b)
	if a != b {
		t.Errorf("Expected %v got %v", a, bVal)
	}
}
