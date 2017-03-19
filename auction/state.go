package auction

import (
	"sync"
	"time"
)

const (
	StatusJoining = "JOINING"
	StatusBidding = "BIDDING"
	StatusClosed  = "CLOSED"
	StatusWon     = "WON"
	StatusLost    = "LOST"
	StatusWinning = "WINNING"
)

//This is an implementation of the Rob Pike FSM
// http://rspace.googlecode.com/hg/slide/lex.html#landing-slide

// stateFn represents the state of the auction
// as a function that returns the next state.
type stateFn func(*auction) stateFn

type auction struct {
	descItemNumber int
	endedTime      time.Time
	bidPrice       float32 // current bid price
	maxBid         float32 // user defined maximum bid
	status         string
	mu             sync.RWMutex // guards status
}

// Snipe starts a concurrent state machine process with an auction which is initialised from
// the inputs. A pointer to the auction is then returned
func Snipe(itemNumber int, maxBid float32, bidPrice float32) (*auction) {
	a := &auction{
		descItemNumber: itemNumber,
		bidPrice:       bidPrice,
		maxBid:         maxBid,
		status:         StatusJoining,
		mu:             sync.RWMutex{},
	}

	go a.run()

	return a
}

func (a *auction) run() {
	for state := joining; state != nil; {
		state = state(a)
	}
}

// Status is a concurrency safe way to get the status of
// a given auction
func (a *auction) Status() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.status
}

// AnnounceClosed changes the status of the auction to closed.
func (a *auction) AnnounceClosed() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.status = StatusClosed
}

//////////////////
// Auction States
//////////////////

func joining(a *auction) stateFn {
	a.mu.Lock()
	defer a.mu.Unlock()
	// when the lead time is reached
	a.status = StatusBidding
	return bidding
}

func bidding(a *auction) stateFn {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.status == StatusClosed {
		a.status = StatusLost
		return lost
	}

	if a.bidPrice <= a.maxBid {
		a.status = StatusWon
		return winning
	}

	return bidding
}

func winning(a *auction) stateFn {
	return nil
}

func lost(a *auction) stateFn {
	return nil
}
