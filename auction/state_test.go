package auction

import (
	"reflect"
	"testing"
)

func TestAuction_Joining(t *testing.T) {
	a := &auction{}

	actual := reflect.ValueOf(joining(a))
	expected := reflect.ValueOf(bidding)

	// validate the identity of the function being returned is what we expect
	// using Pointer on the reflected value
	if actual.Pointer() != expected.Pointer() {
		t.Errorf("Expected; %v, got; %v", actual, expected)
	}

}

func TestAuction_BidPriceGreaterThanMaxBid(t *testing.T) {
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

func TestAuction_AnnounceClosed(t *testing.T) {
	a := &auction{}

	a.AnnounceClosed()

	expected := "CLOSED"
	if actual := a.GetStatus(); actual != expected {
		t.Errorf("Expected; %v, got; %v", actual, expected)
	}

	actual := reflect.ValueOf(bidding(a))
	expectedFn := reflect.ValueOf(lost)

	// see: TestJoining
	if actual.Pointer() != expectedFn.Pointer() {
		t.Errorf("Expected; %v, got; %v", actual, expected)
	}
}
