package auction

import (
	"reflect"
	"testing"
)

func TestJoining(t *testing.T) {
	a := &auction{}

	actual := reflect.ValueOf(joining(a))
	expected := reflect.ValueOf(bidding)

	// validate the identity of the function being returned is what we expect
	// using Pointer on the reflected value
	if actual.Pointer() != expected.Pointer() {
		t.Errorf("Expected; %v, got; %v", actual, expected)
	}

}
