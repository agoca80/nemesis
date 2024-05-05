package main

import (
	"bufio"
	"fmt"
	"os"
)

type InputReader struct {
	*bufio.Reader
}

func newHuman() *InputReader {
	return &InputReader{
		Reader: bufio.NewReader(os.Stdin),
	}
}

func (input *InputReader) Choose(cards Cards) (selected, rejected Card) {
	var choice string
	for _, c := range cards {
		Show(c.Id(), c.Name())
	}
	for {
		Prompt("Choose a card")
		input, _ := fmt.Fscanln(input.Reader, &choice)
		switch {
		case input == 0:
			fallthrough
		case choice == cards[0].Id():
			return cards[0], cards[1]
		case choice == cards[1].Id():
			return cards[1], cards[0]
		}
	}
}
