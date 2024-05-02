package main

// attack symbols
const (
	attack_bite           = "bite"
	attack_claws          = "claws"
	attack_frenzy         = "frenzy"
	attack_mucosity       = "mucosity"
	attack_recall         = "recall"
	attack_scratch        = "scratch"
	attack_tail           = "tail"
	attack_transformation = "transformation"
)

func (g *Game) attackByte(i *Intruder, p *Player) {
	Show(i, "bytes", p, "!")
	if len(p.SeriousWounds) == 2 {
		p.Dies()
	} else {
		p.SufferSeriousWound()
	}
}

func (g *Game) attackClaws(i *Intruder, p *Player) {
	Show(i, "attacks", p, "with its claws!")
	p.Discard(g.Contamination.Draw())
	p.SuffersLightWound()
	if p.Alive() {
		p.SuffersLightWound()
	}
}

func (g *Game) attackFrenzy(i *Intruder, p *Player) {
	Show(i, "in area", i.Area, "goes into a frenzy!!!")
	for _, player := range i.Area.Players.Alive() {
		if len(player.SeriousWounds) < 2 {
			player.SufferSeriousWound()
		} else {
			player.Dies()
		}
	}
}

func (g *Game) attackMucosity(i *Intruder, p *Player) {
	Show(i, "covers", p, " in mucus!")
	p.Discard(g.Contamination.Draw())
	p.IsDrenched = true
}

func (g *Game) attackRecall(i *Intruder, p *Player) {
	token := g.FetchToken()
	if token.Kind == token_blank {
		Show(i, "in area calls for friends, but it has no effect.")
	} else {
		Show(i, "in area calls for friends!")
		g.NewIntruder(token, p.Area)
	}
}

func (g *Game) attackScratch(i *Intruder, p *Player) {
	Show(i, "scratches", p, "!")
	p.Discard(g.Contamination.Draw())
	p.SuffersLightWound()
}

func (g *Game) attackTail(i *Intruder, p *Player) {
	Show(i, "atacks", p, "with its tail!")
	if len(p.SeriousWounds) == 1 {
		p.Dies()
	} else {
		p.SufferSeriousWound()
	}
}

func (g *Game) attackTransformation(i *Intruder, p *Player) {
	breeders := 0
	for _, intruder := range g.Intruders {
		if intruder.Kind == breeder {
			breeders++
		}
	}

	if breeders == 2 {
		Show(i, "convulses and tries to transform!")
		return
	}

	Show(i, "convulses and transforms into an effing breeder!")
	g.RemIntruder(i)
	transformed := g.SpawnIntruder(breeder, p.Area)
	if p.HandSize() == 0 {
		g.ResolveIntruderAttack(transformed, p)
	}
}

func (g *Game) ResolveIntruderAttack(i *Intruder, p *Player) {
	effect := map[string]func(*Intruder, *Player){
		attack_claws:          g.attackClaws,
		attack_frenzy:         g.attackFrenzy,
		attack_tail:           g.attackTail,
		attack_transformation: g.attackTransformation,
		attack_bite:           g.attackByte,
		attack_mucosity:       g.attackMucosity,
		attack_recall:         g.attackRecall,
		attack_scratch:        g.attackScratch,
	}

	if !p.Alive() {
		// Show(i, "was going to attack", p, "but they are already dead!")
		return
	}

	if i.Kind == larva {
		Show(i, "infestes", p, "!")
		p.IsInfested = true
		p.SuffersContamination()
		g.RemIntruder(i)
		return
	}

	attack := i.Attacks.Next().(*AttackCard)
	if !attack.Includes(i.Kind) {
		Show(i, "attacks", p, "but fails!")
		return
	}

	if fn, ok := effect[attack.name]; !ok {
		Show("PENDING", attack, "function not implemented.")
	} else {
		fn(i, p)
	}
}

func (a *AttackCard) Retreats() bool {
	return a.Wounds == 0
}

func (a *AttackCard) String() string {
	return a.card.name
}

// Damage dice
const (
	damage_blank   = "none"
	damage_crawler = "crawler"
	damage_adult   = "adult"
	damage_single  = "single"
	damage_double  = "double"
)

func (p *Player) RollDamage() (result string) {
	damageDice := Symbols{
		damage_blank,
		damage_crawler,
		damage_crawler,
		damage_adult,
		damage_single,
		damage_double,
	}

	return Roll(damageDice)
}
