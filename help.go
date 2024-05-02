package main

type HelpCard struct {
	Card
	Number  int
	Flipped bool
}

func NewHelpCard(number int) *HelpCard {
	return &HelpCard{
		Card:    newCard("Help card"),
		Number:  number,
		Flipped: false,
	}
}
func (hc *HelpCard) Unflips() {
	hc.Flipped = false
}

func (hc *HelpCard) Flips() {
	hc.Flipped = true
}
