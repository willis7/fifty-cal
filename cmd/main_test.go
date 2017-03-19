package main

import (
	"runtime"
	"testing"

	"github.com/willis7/fifty-cal/auction"
)

func TestJoinAndLoseWithoutBidding(t *testing.T) {
	// Join an auction by creating a new instance
	a := auction.Snipe(1234, 1.00, 4.00)

	runtime.Gosched()

	status := a.Status()
	expected := auction.StatusBidding
	if status != expected {
		t.Errorf("Expected: %s, got: %s", expected, status)
	}

	// Announce a closing auction
	a.AnnounceClosed()

	runtime.Gosched()

	status = a.Status()
	expected = auction.StatusLost
	if status != expected {
		t.Errorf("Expected: %s, got: %s", expected, status)
	}
}
