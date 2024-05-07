package main

func attackByte(i *Intruder, p *Player) {
	Show(i, "bytes", p, "!")
	if len(p.Wounds) == 2 {
		p.Dies()
	} else {
		p.SufferSeriousWound()
	}
}

func attackClaws(i *Intruder, p *Player) {
	Show(i, "attacks", p, "with its claws!")
	p.Discard(contamination.Draw())
	p.SuffersLightWound()
	if p.Alive() {
		p.SuffersLightWound()
	}
}

func attackFrenzy(i *Intruder, p *Player) {
	Show(i, "in area", i.Area, "goes into a frenzy!!!")
	for _, player := range i.Area.Players.Alive() {
		if len(player.Wounds) < 2 {
			player.SufferSeriousWound()
		} else {
			player.Dies()
		}
	}
}

func attackMucosity(i *Intruder, p *Player) {
	Show(i, "covers", p, " in mucus!")
	p.Discard(contamination.Draw())
	p.IsDrenched = true
}

func attackRecall(i *Intruder, p *Player) {
	token := game.FetchToken()
	if token.Kind == token_blank {
		Show(i, "in area calls for friends, but it has no effect.")
	} else {
		Show(i, "in area calls for friends!")
		spawnIntruder(token, p.Area)
	}
}

func attackScratch(i *Intruder, p *Player) {
	Show(i, "scratches", p, "!")
	p.Discard(contamination.Draw())
	p.SuffersLightWound()
}

func attackTail(i *Intruder, p *Player) {
	Show(i, "atacks", p, "with its tail!")
	if len(p.Wounds) == 1 {
		p.Dies()
	} else {
		p.SufferSeriousWound()
	}
}

func attackTransformation(i *Intruder, p *Player) {
	breeders := 0
	for _, intruder := range game.Intruders() {
		if intruder.Kind == intruder_breeder {
			breeders++
		}
	}

	if breeders == 2 {
		Show(i, "convulses and tries to transform!")
		return
	}

	Show(i, "convulses and transforms into an effing breeder!")
	RemIntruder(i)
	breeder := newIntruder(intruder_breeder, p.Area)
	if p.HandSize() == 0 {
		breeder.ResolveAttack(p)
	}
}

func (a *Attack) Retreats() bool {
	return a.Wounds == 0
}
