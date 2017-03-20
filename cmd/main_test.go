package main

import (
	"runtime"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/willis7/fifty-cal/auction"
)

func TestJoinAndLoseWithoutBidding(t *testing.T) {
	gomega.RegisterTestingT(t)

	// Join an auction by creating a new instance
	a := auction.Snipe(1234, 1.00, 4.00)

	runtime.Gosched()

	gomega.Eventually(a.Status(), time.Duration(time.Second*2)).Should(gomega.Equal(auction.StatusBidding))

	// Announce a closing auction
	a.AnnounceClosed()

	runtime.Gosched()

	gomega.Eventually(a.Status(), time.Duration(time.Second*2)).Should(gomega.Equal(auction.StatusLost))
}
