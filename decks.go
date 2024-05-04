package main

import (
	"fmt"
	"math/rand"
)

var cardId = 0
var deckId = 'a'

type card struct {
	id   string
	name string
}

type Deck struct {
	cards    Cards
	discards Cards
}

type Card interface {
	Id() string
	Name() string
	String() string
}

type Cards []Card

func newCard(name string) *card {
	cardId++
	return &card{
		id:   fmt.Sprintf("%c%02d", deckId, cardId),
		name: name,
	}
}

func (c *card) Id() string {
	return c.id
}

func (c *card) Name() string {
	return c.name
}

func (c *card) String() string {
	return c.id
}

func (c Cards) shuffle() (shuffled Cards) {
	length := len(c)
	shuffled = make(Cards, length)
	perm := rand.Perm(length)
	for i, j := range perm {
		shuffled[i] = c[j]
	}
	return
}

func newDeck(cards Cards) (d *Deck) {
	deckId, cardId = deckId+1, 0
	return &Deck{
		cards:    cards.shuffle(),
		discards: Cards{},
	}
}

func (d *Deck) Shuffle() {
	cards := append(d.cards, d.discards...)
	d.cards, d.discards = cards.shuffle(), Cards{}
}

// Drawing from a deck is the same as popping from a stack
func (d *Deck) Draw() (c Card) {
	if len(d.cards) == 0 {
		d.Shuffle()
	}

	c, d.cards = d.cards[0], d.cards[1:]
	return
}

func (d *Deck) Discard(c Card) Card {
	d.discards = append(d.discards, c)
	return c
}

func (d *Deck) Return(c Card) {
	d.cards = append(d.cards, c)
}

func (d *Deck) Next() Card {
	return d.Discard(d.Draw())
}

func (d *Deck) Random() (card Card) {
	index := rand.Intn(len(d.cards))
	card, d.cards[index], d.cards = d.cards[index], d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]
	return
}
