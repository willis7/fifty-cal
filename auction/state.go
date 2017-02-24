package auction

import "fmt"

type Updater interface {
	Update()
}

type State struct {
	Updater
	counter int
}

type Joining State
type Bidding State
type Winning State
type Lost State
type Won State

func (self *Joining) Update() {
	fmt.Println("Joining")

	// receive price
	// self.Updater = (*Bidding)(self)

	// auction closed
	// self.Updater = (*Lost)(self)
}


func (self *Bidding) Update() {
	fmt.Println("Bidding")

	// price <= bid
	// self.Updater = (*Winning)(self)

	// auction closed
	// self.Updater = (*Lost)(self)

	// price > bid
	// stay in state
}

func (self *Winning) Update() {
	fmt.Println("Winning")

	// auction finished
	// self.Updater = (*Won)(self)

	// price > bid
	// self.Update = (*Bidding)(self)

	// price <= bid
	// stay in state
}

func (self *Lost) Update() {
	fmt.Println("Lost")

	// report & break
}

func (self *Won) Update() {
	fmt.Println("Won")

	// report & break
}
