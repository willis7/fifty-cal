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

func TestBidPriceGreaterThanMaxBid(t *testing.T) {
	a := &auction{
		maxBid: 100,
		bidPrice: 100,
	}

	actual := reflect.ValueOf(bidding(a))
	expected := reflect.ValueOf(winning)

	// see: TestJoining
	if actual.Pointer() != expected.Pointer() {
		t.Errorf("Expected; %v, got; %v", actual, expected)
	}
}
