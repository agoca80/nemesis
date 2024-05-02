package main

import (
	"math/rand"
)

func NewBoard() (b *Board) {
	b = &Board{
		Area: Areas{
			area(A00, "UNUSED"),
			area(A01, room_cockpit),
			area(A02, room_1),
			area(A03, room_1),
			area(A04, room_1),
			area(A05, room_1),
			area(A06, room_1),
			area(A07, room_1),
			area(A08, room_2),
			area(A09, room_2),
			area(A10, room_2),
			area(A11, room_hibernatorium),
			area(A12, room_2),
			area(A13, room_2),
			area(A14, room_1),
			area(A15, room_1),
			area(A16, room_1),
			area(A17, room_1),
			area(A18, room_1),
			area(A19, room_engine1),
			area(A20, room_engine2),
			area(A21, room_engine3),
			area(S01, room_service1),
			area(S02, room_service2),
		},
	}

	tunnelId := 0
	tunnel := func(x, y int, numbers ...int) (t *Tunnel) {
		tunnelId++
		t = &Tunnel{
			Id:      tunnelId,
			Door:    door_open,
			Numbers: numbers,
		}
		b.Area[x].Corridors = append(b.Area[x].Corridors, corridor(b.Area[y], t))
		b.Area[y].Corridors = append(b.Area[y].Corridors, corridor(b.Area[x], t))
		return
	}

	b.Tunnel = Tunnels{
		tunnel(A01, A02, 3),
		tunnel(A01, A03, 1, 2),
		tunnel(A01, A04, 4),
		tunnel(A02, S01, 1, 2),
		tunnel(A02, A06, 4),
		tunnel(A03, A07, 3, 4),
		tunnel(A04, A08, 1),
		tunnel(A04, S01, 2, 3),
		tunnel(A05, A10, 1, 2),
		tunnel(A05, A06, 3),
		tunnel(A05, S01, 4),
		tunnel(A06, A07, 1),
		tunnel(A06, A11, 2),
		tunnel(A07, A08, 2),
		tunnel(A08, A09, 4),
		tunnel(A08, A11, 3),
		tunnel(A09, A12, 1, 2),
		tunnel(A09, S01, 3),
		tunnel(A10, A13, 3, 4),
		tunnel(A11, A14, 4),
		tunnel(A11, A15, 1),
		tunnel(A12, A16, 3, 4),
		tunnel(A13, A14, 1),
		tunnel(A13, A19, 2),
		tunnel(A14, A17, 2),
		tunnel(A14, S01, 3),
		tunnel(A15, A16, 2),
		tunnel(A15, A18, 3),
		tunnel(A15, S01, 4),
		tunnel(A16, A21, 1),
		tunnel(A17, A19, 1),
		tunnel(A17, A20, 3, 4),
		tunnel(A18, A20, 1, 2),
		tunnel(A18, A21, 4),
		tunnel(A19, S01, 3, 4),
		tunnel(A21, S01, 2, 3),
	}

	for _, area := range b.Area {
		if area.Class != room_1 && area.Class != room_2 {
			area.Room = &Room{
				card: newCard(area.Class),
			}
		}
	}

	return
}

func NewGame(players int) (g *Game) {
	g = &Game{
		Board:   NewBoard(),
		Players: Players{},
		EscapePods: []Cards{
			{},
			{},
		},

		Attacks: NewDeck(Cards{
			NewAttackCard(6, attack_bite, adult, breeder, queen),
			NewAttackCard(4, attack_bite, adult, breeder, queen),
			NewAttackCard(4, attack_bite, adult, breeder, queen),
			NewAttackCard(0, attack_bite, adult, breeder, queen),
			NewAttackCard(3, attack_recall, crawler, queen),
			NewAttackCard(2, attack_tail, queen),
			NewAttackCard(5, attack_tail, queen),
			NewAttackCard(3, attack_frenzy, breeder, queen),
			NewAttackCard(4, attack_frenzy, breeder, queen),
			NewAttackCard(4, attack_transformation, crawler),
			NewAttackCard(5, attack_transformation, crawler),
			NewAttackCard(3, attack_scratch, crawler, adult, breeder, queen),
			NewAttackCard(4, attack_scratch, crawler, adult, breeder, queen),
			NewAttackCard(5, attack_scratch, crawler, adult, breeder, queen),
			NewAttackCard(6, attack_scratch, crawler, adult, breeder, queen),
			NewAttackCard(3, attack_claws, adult, breeder, queen),
			NewAttackCard(4, attack_claws, adult, breeder, queen),
			NewAttackCard(5, attack_claws, adult, breeder, queen),
			NewAttackCard(0, attack_claws, adult, breeder, queen),
			NewAttackCard(5, attack_mucosity, crawler, adult, breeder, queen),
		}),

		Contamination: NewDeck(Cards{
			contaminationCard(true),
			contaminationCard(true),
			contaminationCard(true),
			contaminationCard(true),
			contaminationCard(true),
			contaminationCard(true),
			contaminationCard(true),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
			contaminationCard(false),
		}),

		Events: NewDeck(Cards{
			NewEventCard(1, "Fuga de refrigerante", adult, breeder, queen),
			NewEventCard(3, "Eclosion", adult, breeder),
			NewEventCard(4, "Ruido en los pasillos de servicio", adult, breeder),
			NewEventCard(2, "Fallo del soporte vital", adult, breeder, queen),
			NewEventCard(2, "Proteccion de los huevos", breeder, queen),
			NewEventCard(4, "Acechar", crawler, adult),
			NewEventCard(3, "El rastro de la presa", crawler, adult),
			NewEventCard(4, "Compuestos inflamables", breeder, queen),
			NewEventCard(2, "Daño", crawler, adult),
			NewEventCard(3, "Nido", crawler, breeder, queen),
			NewEventCard(1, "Fuego perjudicial", crawler, adult),
			NewEventCard(1, "Apertura de compuertas", adult, breeder),
			NewEventCard(4, "Fuego devastador", crawler, breeder, queen),
			NewEventCard(1, "Regeneracion", breeder, queen),
			NewEventCard(2, event_failure, adult, breeder),
			NewEventCard(1, "Nacimiento", crawler, breeder, queen),
			NewEventCard(3, "Caceria", crawler, queen),
			NewEventCard(4, "cortocircuito", adult, breeder, queen),
			NewEventCard(2, "caceria", crawler, breeder, queen),
			NewEventCard(3, "Eyeccion de capsula de evacuacion", adult, breeder, queen),
		}),

		Weaknesses: NewDeck(Cards{
			NewWeakness(weakness_fighting),
			NewWeakness(weakness_reaction),
			NewWeakness(weakness_vitalpoints),
			NewWeakness(weakness_phosphates),
			NewWeakness(weakness_fire),
			NewWeakness(weakness_energy),
			NewWeakness(weakness_movement),
			NewWeakness(weakness_endangered),
		}),

		Wounds: NewDeck(Cards{
			NewSeriousWound(wound_leg),
			NewSeriousWound(wound_leg),
			NewSeriousWound(wound_leg),
			NewSeriousWound(wound_body),
			NewSeriousWound(wound_body),
			NewSeriousWound(wound_body),
			NewSeriousWound(wound_body),
			NewSeriousWound(wound_arm),
			NewSeriousWound(wound_arm),
			NewSeriousWound(wound_arm),
			NewSeriousWound(wound_hand),
			NewSeriousWound(wound_hand),
			NewSeriousWound(wound_hand),
			NewSeriousWound(wound_hemorrhage),
			NewSeriousWound(wound_hemorrhage),
			NewSeriousWound(wound_hemorrhage),
		}),
	}

	for range players {
		g.Players = append(g.Players, NewPlayer(g))
	}

	return
}

func (g *Game) Prepare(coop bool) {
	// Prepare 1
	g.Board = NewBoard()

	// Prepare 2
	room2 := NewDeck(Cards{
		NewRoom("Command deck", blank, true),
		NewRoom("Engine Management", yellow, true),
		NewRoom("Showers", all, false),
		NewRoom("Decompression control", yellow, false),
		NewRoom("Escape pods control", all, false),
		NewRoom("Canteen", green, false),
		NewRoom("Cabins", all, false),
		NewRoom("Mocosous room", blank, false),
		NewRoom("Vigilance", green, true),
	})
	for _, a := range g.Area {
		if a.Class == room_2 {
			a.Room = room2.Draw().(*Room)
		}
	}

	// Prepare 3
	room1 := NewDeck(Cards{
		NewRoom(room_storage, red, false),
		NewRoom("Fire control system", yellow, true),
		NewRoom("Emergency room", green, false),
		NewRoom("Laboratory", blank, true),
		NewRoom("Generator", yellow, true),
		NewRoom("Cooms room", yellow, true),
		NewRoom("Surgery", green, false),
		NewRoom("Nest", blank, false),
		NewRoom("Armory", red, false),
		NewRoom("Evacuation section A", all, false),
		NewRoom("Evacuation section B", all, false),
	})
	for _, a := range g.Area {
		if a.Class == room_1 {
			a.Room = room1.Draw().(*Room)
		}
	}

	// Prepare 4
	explorationTokens := NewDeck(Cards{
		NewExplorationToken(1, ev_damaged),
		NewExplorationToken(1, ev_damaged),
		NewExplorationToken(2, ev_damaged),
		NewExplorationToken(2, ev_damaged),
		NewExplorationToken(2, ev_damaged),
		NewExplorationToken(2, ev_damaged),
		NewExplorationToken(3, ev_damaged),
		NewExplorationToken(4, ev_damaged),
		NewExplorationToken(2, ev_danger),
		NewExplorationToken(3, ev_danger),
		NewExplorationToken(1, ev_door),
		NewExplorationToken(2, ev_door),
		NewExplorationToken(3, ev_door),
		NewExplorationToken(4, ev_door),
		NewExplorationToken(1, ev_fire),
		NewExplorationToken(2, ev_fire),
		NewExplorationToken(3, ev_mucus),
		NewExplorationToken(4, ev_mucus),
		NewExplorationToken(1, ev_silence),
		NewExplorationToken(1, ev_silence),
	})
	for _, a := range g.Area {
		if a.Class == room_1 || a.Class == room_2 {
			a.ExplorationToken = explorationTokens.Draw().(*ExplorationToken)
		}
	}

	// Initialize coordinates card
	coordinates := NewDeck(Cards{
		CoordinatesCard("CBDA"),
		CoordinatesCard("ABCD"),
		CoordinatesCard("ACDB"),
		CoordinatesCard("DABC"),
		CoordinatesCard("DCAB"),
		CoordinatesCard("BDCA"),
		CoordinatesCard("CABD"),
		CoordinatesCard("BDAC"),
	})
	g.Coordinate = coordinates.Draw().(*Coordinates)

	// Initialize destination
	g.Destination = "B"

	// Initialize escape pods
	// order := rand.Perm(2)
	// alternate := []int{0, 1, order[0], order[1]}
	// numOfEscapePods := []int{2, 2, 3, 3, 4}[len(g.Players)-1]
	// for n := range numOfEscapePods {
	// 	search := func(c Card) bool {
	// 		ep := c.(*EscapePod)
	// 		return ep.number == n
	// 	}
	// 	g.EscapePods[alternate[n]] = append(g.EscapePods[alternate[n]], g.Take(escapePods, search))
	// }

	// initialize engines states
	g.EngineStatus = [3]bool{
		rand.Intn(2) == 0,
		rand.Intn(2) == 0,
		rand.Intn(2) == 0,
	}

	// Initialize intruder board
	g.Eggs = 5
	g.Weakness = []*Weakness{
		g.Weaknesses.Draw().(*Weakness),
		g.Weaknesses.Draw().(*Weakness),
		g.Weaknesses.Draw().(*Weakness),
	}

	// Initialize intruder bag
	g.IntruderBag = NewIntruderBag(len(g.Players))

	// Initialize hyperdrive countdown
	g.hyperdriveCountdown = 15

	// Crew preparation step 18 A,B,C
	for _, p := range g.Players {
		p.Area, g.Area[A11].Players = g.Area[A11], append(g.Area[A11].Players, p)
	}

	// Crew preparation step 14
	helpCards := Cards{
		NewHelpCard(1),
		NewHelpCard(2),
		NewHelpCard(3),
		NewHelpCard(4),
		NewHelpCard(5),
	}
	helpDeck := NewDeck(helpCards[:len(g.Players)])
	for _, p := range g.Players {
		p.HelpCard = helpDeck.Draw().(*HelpCard)
	}

	// Crew preparation step 16
	goals := NewDeck(Cards{
		// Solo/Cooperative
		NewGoal(0, "The signal"),                   // Send the signal
		NewGoal(0, "Clean team"),                   // Send the signal AND (the nest is destroyed OR the ship is destroyed)
		NewGoal(0, "Cut from the root"),            // Send the signal AND (the queen is dead OR the ship is destroyed)
		NewGoal(0, "Special delivery"),             // Send the signal AND (the queen is dead OR the nest is destroyed)
		NewGoal(0, "We don't leave anyone behind"), // Send the signal AND all compartments in the ship are explored
		NewGoal(0, "First contact protocol"),       // Find out two of the intruders' weaknesses
		NewGoal(0, "Post-mortem urgency"),          // Bring the blue player corpse to the Operating room and leave it there
		NewGoal(0, "Destination: Earth"),           // The ship must arrive to Earth
		// Coorporative
		NewGoal(5, "Aliens in the ship"),        // Send the signal AND (the nest is destroyed OR the ship is destroyed)
		NewGoal(4, "Friends forever"),           // You and another player at least must survive
		NewGoal(3, "The oldest friend"),         // The ship must arrive to Earth OR You are the only one surviving
		NewGoal(2, "Big game"),                  // Sendthe signal AND (the queen is dead OR the ship is destroyed)
		NewGoal(2, "Restless investigator"),     // Send the signal AND all compartments in the ship are explored
		NewGoal(2, "Recover the main resource"), // The ship arrives to Earth OR you are the only one surviving
		NewGoal(2, "A proper burial"),           // Send the signal AND end the game in an escape pod or hybernating with the blue player corpse
		NewGoal(2, "Quarantine"),                // The ship must arrive to Mars OR the ship arrives to Earth and the nest has been destroyed
		NewGoal(2, "Hoarder"),                   // Finish the game in an escape pod with at least 7 items. Mision items only count if they are activated
		// Personal
		NewGoal(5, "The ultimate epiphany"),      // Player 5 must not survive OR you are the only one surviving
		NewGoal(4, "Taking by force"),            // Player 4 must not survive OR you are the only one surviving
		NewGoal(3, "An old enemy"),               // Player 3 must not survive OR you are the only one surviving
		NewGoal(2, "Ab Ovo"),                     // Reveal the intruder egg's weakness
		NewGoal(2, "Necropsy"),                   // Send the signal AND reveal the intruder corpse's weakness
		NewGoal(2, "The best moment for attack"), // Player 1 must not survice OR you are the only one surviving
		NewGoal(2, "My treasure"),                // Send the signal AND (end the game in an escape pod or hybernating with an intruder egg)
		NewGoal(2, "More exhuberant fields"),     // Player 2 must not survive OR you are the only one surviving
		NewGoal(2, "Extreme field biology"),      // Reveal at least two intruder weaknesses
	})
	for _, p := range g.Players {
		if coop {
			p.Goals = append(p.Goals, goals.Draw())
		} else {
			p.Goals = append(p.Goals, goals.Draw(), goals.Draw())
		}
	}

	// Crew preparation step 17
	characters := NewDeck(Cards{
		newCard(captain),
		newCard(explorer),
		newCard(mechanic),
		newCard(pilot),
		newCard(scientist),
		newCard(soldier),
	})
	for _, p := range g.Players {
		p.chooseCharacter(characters)
	}

	// Crew preparation step 18
	actions := map[string]Cards{
		captain: {
			NewActionCard(0, action_rest),
			NewActionCard(0, action_order),
			NewActionCard(0, action_demolition),
			NewActionCard(0, action_reload),
			NewActionCard(0, action_suppresion),
			NewActionCard(0, action_motivation),
			NewActionCard(0, action_register),
			NewActionCard(0, action_basic_repairs),
			NewActionCard(0, action_interruption),
			NewActionCard(0, action_register),
		},
		explorer: {
			NewActionCard(0, action_scavenger),
			NewActionCard(0, action_rest),
			NewActionCard(0, action_scout),
			NewActionCard(0, action_demolition),
			NewActionCard(0, action_register),
			NewActionCard(0, action_suppresion),
			NewActionCard(0, action_register),
			NewActionCard(0, action_basic_repairs),
			NewActionCard(0, action_adrenaline),
			NewActionCard(0, action_interruption),
		},
		scientist: {
			NewActionCard(0, action_register),
			NewActionCard(0, action_intranet),
			NewActionCard(0, action_computer_skills),
			NewActionCard(0, action_repair),
			NewActionCard(0, action_demolition),
			NewActionCard(0, action_risk_evaluation),
			NewActionCard(0, action_block_access),
			NewActionCard(0, action_rest),
			NewActionCard(0, action_interruption),
			NewActionCard(0, action_register),
		},
		soldier: {
			NewActionCard(0, action_basic_repairs),
			NewActionCard(0, action_automatic_fire),
			NewActionCard(0, action_covering_fire),
			NewActionCard(0, action_register),
			NewActionCard(0, action_rest),
			NewActionCard(0, action_aimed_fire),
			NewActionCard(0, action_still_nerves),
			NewActionCard(0, action_interruption),
			NewActionCard(0, action_demolition),
			NewActionCard(0, action_register),
		},
		mechanic: {
			NewActionCard(0, action_interruption),
			NewActionCard(0, action_register),
			NewActionCard(0, action_computer_skills),
			NewActionCard(0, action_register),
			NewActionCard(0, action_improvisation), // Inventiva???
			NewActionCard(0, action_quick_repairs),
			NewActionCard(0, action_demolition),
			NewActionCard(0, action_fireworks),
			NewActionCard(0, action_service_corridors),
			NewActionCard(0, action_rest),
		},
		pilot: {
			NewActionCard(0, action_interruption),
			NewActionCard(0, action_rest),
			NewActionCard(0, action_computer_skills),
			NewActionCard(0, action_demolition),
			NewActionCard(0, action_register),
			NewActionCard(0, action_ship_knownledge),
			NewActionCard(0, action_old_friend),
			NewActionCard(0, action_piloting),
			NewActionCard(0, action_repair),
			NewActionCard(0, action_register),
		},
	}
	for _, p := range g.Players {
		p.Deck = NewDeck(actions[p.Character])
	}

	// Step 19
	for _, p := range g.Players {
		if p.Number == 1 {
			p.Jonesy = true
		}
	}

	// Preparation 20
	// hybernarium.Objects = append(hybernarium.Objects, g.BlueCorpseToken)
}
