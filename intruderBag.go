package main

import (
	"math/rand"
	"slices"
)

const (
	token_blank   = "blank"
	token_larva   = "larva"
	token_crawler = "crawler"
	token_adult   = "adult"
	token_breeder = "breeder"
	token_queen   = "queen"
)

type IntruderBag struct {
	inside  IntruderTokens
	outside map[string]IntruderTokens
}

func (ib *IntruderBag) FetchToken() (token *IntruderToken) {
	index := rand.Intn(len(ib.inside))
	token = ib.inside[index]
	return
}

func (ib *IntruderBag) Retire(token *IntruderToken) {
	index := slices.Index(ib.inside, token)
	ib.inside = slices.Delete(ib.inside, index, index+1)
}

func (ib *IntruderBag) Return(kind string) {
	if len(ib.outside[kind]) == 0 {
		return
	}

	index := rand.Intn(len(ib.outside[kind]))
	token := ib.outside[kind][index]
	ib.inside = append(ib.inside, token)
	ib.outside[kind] = slices.Delete(ib.outside[kind], index, index+1)
}

func newIntruderBag(players int) (ib *IntruderBag) {
	ib = &IntruderBag{
		inside: IntruderTokens{},
		outside: map[string]IntruderTokens{
			token_larva:   {},
			token_crawler: {},
			token_adult:   {},
			token_breeder: {},
			token_queen:   {},
		},
	}

	tokens := []*IntruderToken{
		NewIntruderToken(0, token_blank),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_larva),
		NewIntruderToken(1, token_crawler),
		NewIntruderToken(1, token_crawler),
		NewIntruderToken(1, token_crawler),
		NewIntruderToken(2, token_adult),
		NewIntruderToken(2, token_adult),
		NewIntruderToken(2, token_adult),
		NewIntruderToken(2, token_adult),
		NewIntruderToken(3, token_adult),
		NewIntruderToken(3, token_adult),
		NewIntruderToken(3, token_adult),
		NewIntruderToken(3, token_adult),
		NewIntruderToken(3, token_adult),
		NewIntruderToken(4, token_adult),
		NewIntruderToken(4, token_adult),
		NewIntruderToken(4, token_adult),
		NewIntruderToken(3, token_breeder),
		NewIntruderToken(4, token_breeder),
		NewIntruderToken(4, token_queen),
	}
	perm := rand.Perm(len(tokens))
	for _, j := range perm {
		token := tokens[j]
		kind := token.Kind
		ib.outside[kind] = append(ib.outside[kind], token)
	}

	ib.Return(token_blank)
	ib.Return(token_larva)
	ib.Return(token_larva)
	ib.Return(token_larva)
	ib.Return(token_larva)
	ib.Return(token_crawler)
	ib.Return(token_adult)
	ib.Return(token_adult)
	ib.Return(token_adult)
	ib.Return(token_queen)
	for range players {
		ib.Return(token_adult)
	}

	return
}
