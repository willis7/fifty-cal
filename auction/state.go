package auction

import (
	"sync"
	"time"
)

//This is an implenentation of the Rob Pike FSM
// http://rspace.googlecode.com/hg/slide/lex.html#landing-slide

// stateFn represents the state of the auction
// as a function that returns the next state.
type stateFn func(*auction) stateFn

type auction struct {
	descItemNumber int
	endedTime      time.Time
	bidPrice       float32 // current bid price
	maxBid         float32 // user defined maximum bid
	Status         string
	mu             sync.Mutex // guards status
}

func joining(a *auction) stateFn {
	// when the lead time is reached
	return bidding
}

func bidding(a *auction) stateFn {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Status == "CLOSED" {
		return lost
	}
	a.Status = "BIDDING"

	if a.bidPrice <= a.maxBid {
		return winning
	}

	return bidding
}

func winning(a *auction) stateFn {
	return nil
}

func lost(a *auction) stateFn {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Status = "LOST"

	return nil
}

func Snipe(itemNumber int, maxBid float32) (*auction) {
	a := &auction{
		descItemNumber: itemNumber,
		maxBid:         maxBid,
		Status:         "JOINING",
		mu:             sync.Mutex{},
	}

	go a.run()

	return a
}

func (a *auction) run() {
	for state := joining; state != nil; {
		state = state(a)
	}
}

// GetStatus is a concurrency safe way to get the status of
// a given auction
func (a *auction)GetStatus() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Status
}

// AnnounceClosed changes the status of the auction to closed.
func (a *auction) AnnounceClosed() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Status = "CLOSED"
}
