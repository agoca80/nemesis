package main

type EscapePod struct {
	*card
	blocked  bool
	launched bool
	number   int
	Player   Players
}

func NewEscapePod(number int) Card {
	return &EscapePod{
		card:     newCard("Escape Pod"),
		blocked:  true,
		launched: false,
		number:   number,
	}
}

func (ep *EscapePod) IsFull() bool {
	return len(ep.Player) == 2
}

// var escapePodCards = Cards{
// 	NewEscapePod(1),
// 	NewEscapePod(2),
// 	NewEscapePod(3),
// 	NewEscapePod(4),
// }
