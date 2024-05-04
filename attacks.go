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

func attackByte(i *Intruder, p *player) {
	Show(i, "bytes", p, "!")
	if len(p.Wounds) == 2 {
		p.Dies()
	} else {
		p.SufferSeriousWound()
	}
}

func attackClaws(i *Intruder, p *player) {
	Show(i, "attacks", p, "with its claws!")
	p.Discard(contamination.Draw())
	p.SuffersLightWound()
	if p.Alive() {
		p.SuffersLightWound()
	}
}

func attackFrenzy(i *Intruder, p *player) {
	Show(i, "in area", i.Area, "goes into a frenzy!!!")
	for _, player := range i.Area.Players.Alive() {
		if len(player.Wounds) < 2 {
			player.SufferSeriousWound()
		} else {
			player.Dies()
		}
	}
}

func attackMucosity(i *Intruder, p *player) {
	Show(i, "covers", p, " in mucus!")
	p.Discard(contamination.Draw())
	p.IsDrenched = true
}

func attackRecall(i *Intruder, p *player) {
	token := game.FetchToken()
	if token.Kind == token_blank {
		Show(i, "in area calls for friends, but it has no effect.")
	} else {
		Show(i, "in area calls for friends!")
		NewIntruder(token, p.Area)
	}
}

func attackScratch(i *Intruder, p *player) {
	Show(i, "scratches", p, "!")
	p.Discard(contamination.Draw())
	p.SuffersLightWound()
}

func attackTail(i *Intruder, p *player) {
	Show(i, "atacks", p, "with its tail!")
	if len(p.Wounds) == 1 {
		p.Dies()
	} else {
		p.SufferSeriousWound()
	}
}

func attackTransformation(i *Intruder, p *player) {
	breeders := 0
	for _, intruder := range game.Intruders {
		if intruder.Kind == breeder {
			breeders++
		}
	}

	if breeders == 2 {
		Show(i, "convulses and tries to transform!")
		return
	}

	Show(i, "convulses and transforms into an effing breeder!")
	RemIntruder(i)
	transformed := SpawnIntruder(breeder, p.Area)
	if p.HandSize() == 0 {
		ResolveIntruderAttack(transformed, p)
	}
}

func ResolveIntruderAttack(i *Intruder, p *player) {
	effect := map[string]func(*Intruder, *player){
		attack_claws:          attackClaws,
		attack_frenzy:         attackFrenzy,
		attack_tail:           attackTail,
		attack_transformation: attackTransformation,
		attack_bite:           attackByte,
		attack_mucosity:       attackMucosity,
		attack_recall:         attackRecall,
		attack_scratch:        attackScratch,
	}

	if !p.Alive() {
		// Show(i, "was going to attack", p, "but they are already dead!")
		return
	}

	if i.Kind == larva {
		Show(i, "infestes", p, "!")
		p.IsInfected = true
		p.SuffersContamination()
		RemIntruder(i)
		return
	}

	attack := attacks.Next().(*AttackCard)
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

func (p *player) RollDamage() (result string) {
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
