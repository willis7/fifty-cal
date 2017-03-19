package main

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/willis7/fifty-cal/auction"
)

func TestJoinAndLoseWithoutBidding(t *testing.T) {
	// Join an auction by creating a new instance
	a := auction.Snipe(1234, 1.00, 4.00)

	gomega.Eventually(a.Status()).Should(gomega.Equal(auction.StatusBidding))

	// Announce a closing auction
	a.AnnounceClosed()

	gomega.Eventually(a.Status()).Should(gomega.Equal(auction.StatusLost))
}
