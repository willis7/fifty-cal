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

//func TestAuction_Bidding_BidPriceLowerThanMaxBid(t *testing.T) {
//	a := &auction{
//		maxBid:   100,
//		bidPrice: 50,
//	}
//
//	actual := reflect.ValueOf(bidding(a))
//	expected := reflect.ValueOf(bidding)
//
//	// see: TestJoining
//	if actual.Pointer() != expected.Pointer() {
//		t.Errorf("Expected; %v, got; %v", expected, actual)
//	}
//}

func TestAuction_Bidding_AnnounceClosed(t *testing.T) {
	a := &auction{}

	a.AnnounceClosed()

	expected := "CLOSED"
	if actual := a.Status(); actual != expected {
		t.Errorf("Expected; %v, got; %v", expected, actual)
	}

	actual := reflect.ValueOf(bidding(a))
	expectedFn := reflect.ValueOf(lost)

	// see: TestJoining
	if actual.Pointer() != expectedFn.Pointer() {
		t.Errorf("Expected; %v, got; %v", expectedFn, actual)
	}
}

func TestStateMachine(t *testing.T) {
	tests := []struct {
		description    string
		bidPrice       float32
		maxBid         float32
		currentStatus  string
		expectedStatus string
		startStateFn   stateFn
		nextStateFn    stateFn
	}{
		// #1
		{
			description:    "bidding: bid price greater than max bid",
			bidPrice:       100,
			maxBid:         100,
			currentStatus:  "BIDDING",
			expectedStatus: "WON",
			startStateFn:   bidding,
			nextStateFn:    winning,
		},
	}

	for _, tc := range tests {
		a := &auction{
			status:   tc.currentStatus,
			maxBid:   tc.maxBid,
			bidPrice: tc.bidPrice,
		}

		actualFn := reflect.ValueOf(tc.startStateFn(a))
		expectedFn := reflect.ValueOf(tc.nextStateFn)

		// see: TestJoining
		if actualFn.Pointer() != expectedFn.Pointer() {
			t.Errorf("Expected; %v, got; %v", expectedFn, actualFn)
		}


		if actual := a.Status(); actual != tc.expectedStatus {
			t.Errorf("Expected; %v, got; %v", tc.expectedStatus, actual)
		}
	}
}
