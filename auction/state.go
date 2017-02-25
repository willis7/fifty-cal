package auction

import (
	"time"
)

//This is an implenentation of the Rob Pike FSM
// http://rspace.googlecode.com/hg/slide/lex.html#landing-slide

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*auction) stateFn

type auction struct {
	descItemNumber int
	endedTime        time.Time
	bidPrice       float32 // current bid price
	maxBid         float32 // user defined maximum bid
}

func joining(a *auction) stateFn{
	// when the lead time is reached
	return bidding
}

func bidding(a *auction) stateFn {

	return bidding
}
