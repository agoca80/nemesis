package main

var actions = map[string]*Deck{
	captain: newDeck(Cards{
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
	}),
	explorer: newDeck(Cards{
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
	}),
	scientist: newDeck(Cards{
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
	}),
	soldier: newDeck(Cards{
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
	}),
	mechanic: newDeck(Cards{
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
	}),
	pilot: newDeck(Cards{
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
	}),
}

var areas Areas = Areas{
	newArea(A00, "UNUSED"),
	newArea(A01, room_cockpit),
	newArea(A02, room_1),
	newArea(A03, room_1),
	newArea(A04, room_1),
	newArea(A05, room_1),
	newArea(A06, room_1),
	newArea(A07, room_1),
	newArea(A08, room_2),
	newArea(A09, room_2),
	newArea(A10, room_2),
	newArea(A11, room_hibernatorium),
	newArea(A12, room_2),
	newArea(A13, room_2),
	newArea(A14, room_1),
	newArea(A15, room_1),
	newArea(A16, room_1),
	newArea(A17, room_1),
	newArea(A18, room_1),
	newArea(A19, room_engine1),
	newArea(A20, room_engine2),
	newArea(A21, room_engine3),
	newArea(S01, room_service1),
	newArea(S02, room_service2),
}

var attacks = newDeck(Cards{
	newAttack(6, attack_bite, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(4, attack_bite, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(4, attack_bite, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(0, attack_bite, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(3, attack_recall, intruder_crawler, intruder_queen),
	newAttack(2, attack_tail, intruder_queen),
	newAttack(5, attack_tail, intruder_queen),
	newAttack(3, attack_frenzy, intruder_breeder, intruder_queen),
	newAttack(4, attack_frenzy, intruder_breeder, intruder_queen),
	newAttack(4, attack_transformation, intruder_crawler),
	newAttack(5, attack_transformation, intruder_crawler),
	newAttack(3, attack_scratch, intruder_crawler, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(4, attack_scratch, intruder_crawler, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(5, attack_scratch, intruder_crawler, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(6, attack_scratch, intruder_crawler, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(3, attack_claws, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(4, attack_claws, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(5, attack_claws, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(0, attack_claws, intruder_adult, intruder_breeder, intruder_queen),
	newAttack(5, attack_mucosity, intruder_crawler, intruder_adult, intruder_breeder, intruder_queen),
})

var characters = newDeck(Cards{
	newCharacter(captain),
	newCharacter(explorer),
	newCharacter(mechanic),
	newCharacter(pilot),
	newCharacter(scientist),
	newCharacter(soldier),
})

var contamination = newDeck(Cards{
	newContamination(true),
	newContamination(true),
	newContamination(true),
	newContamination(true),
	newContamination(true),
	newContamination(true),
	newContamination(true),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
	newContamination(false),
})

var coordinates = newDeck(Cards{
	newCoordinates("CBDA"),
	newCoordinates("ABCD"),
	newCoordinates("ACDB"),
	newCoordinates("DABC"),
	newCoordinates("DCAB"),
	newCoordinates("BDCA"),
	newCoordinates("CABD"),
	newCoordinates("BDAC"),
})

var gates = Gates{
	{A01, A02, []int{3}},
	{A01, A03, []int{1, 2}},
	{A01, A04, []int{4}},
	{A02, S01, []int{1, 2}},
	{A02, A06, []int{4}},
	{A03, A07, []int{3, 4}},
	{A04, A08, []int{1}},
	{A04, S01, []int{2, 3}},
	{A05, A10, []int{1, 2}},
	{A05, A06, []int{3}},
	{A05, S01, []int{4}},
	{A06, A07, []int{1}},
	{A06, A11, []int{2}},
	{A07, A08, []int{2}},
	{A08, A09, []int{4}},
	{A08, A11, []int{3}},
	{A09, A12, []int{1, 2}},
	{A09, S01, []int{3}},
	{A10, A13, []int{3, 4}},
	{A11, A14, []int{4}},
	{A11, A15, []int{1}},
	{A12, A16, []int{3, 4}},
	{A13, A14, []int{1}},
	{A13, A19, []int{2}},
	{A14, A17, []int{2}},
	{A14, S01, []int{3}},
	{A15, A16, []int{2}},
	{A15, A18, []int{3}},
	{A15, S01, []int{4}},
	{A16, A21, []int{1}},
	{A17, A19, []int{1}},
	{A17, A20, []int{3, 4}},
	{A18, A20, []int{1, 2}},
	{A18, A21, []int{4}},
	{A19, S01, []int{3, 4}},
	{A21, S01, []int{2, 3}},
}

var events = newDeck(Cards{
	newEvent(1, "Fuga de refrigerante", intruder_adult, intruder_breeder, intruder_queen),
	newEvent(3, "Eclosion", intruder_adult, intruder_breeder),
	newEvent(4, "Ruido en los pasillos de servicio", intruder_adult, intruder_breeder),
	newEvent(2, "Fallo del soporte vital", intruder_adult, intruder_breeder, intruder_queen),
	newEvent(2, "Proteccion de los huevos", intruder_breeder, intruder_queen),
	newEvent(4, "Acechar", intruder_crawler, intruder_adult),
	newEvent(3, "El rastro de la presa", intruder_crawler, intruder_adult),
	newEvent(4, "Compuestos inflamables", intruder_breeder, intruder_queen),
	newEvent(2, "Daño", intruder_crawler, intruder_adult),
	newEvent(3, "Nido", intruder_crawler, intruder_breeder, intruder_queen),
	newEvent(1, "Fuego perjudicial", intruder_crawler, intruder_adult),
	newEvent(1, "Apertura de compuertas", intruder_adult, intruder_breeder),
	newEvent(4, "Fuego devastador", intruder_crawler, intruder_breeder, intruder_queen),
	newEvent(1, "Regeneracion", intruder_breeder, intruder_queen),
	newEvent(2, event_failure, intruder_adult, intruder_breeder),
	newEvent(1, "Nacimiento", intruder_crawler, intruder_breeder, intruder_queen),
	newEvent(3, "Caceria", intruder_crawler, intruder_queen),
	newEvent(4, "cortocircuito", intruder_adult, intruder_breeder, intruder_queen),
	newEvent(2, "caceria", intruder_crawler, intruder_breeder, intruder_queen),
	newEvent(3, "Eyeccion de capsula de evacuacion", intruder_adult, intruder_breeder, intruder_queen),
})

var explorationTokens = newDeck(Cards{
	newExplorationToken(1, ev_damaged),
	newExplorationToken(1, ev_damaged),
	newExplorationToken(2, ev_damaged),
	newExplorationToken(2, ev_damaged),
	newExplorationToken(2, ev_damaged),
	newExplorationToken(2, ev_damaged),
	newExplorationToken(3, ev_damaged),
	newExplorationToken(4, ev_damaged),
	newExplorationToken(2, ev_danger),
	newExplorationToken(3, ev_danger),
	newExplorationToken(1, ev_door),
	newExplorationToken(2, ev_door),
	newExplorationToken(3, ev_door),
	newExplorationToken(4, ev_door),
	newExplorationToken(1, ev_fire),
	newExplorationToken(2, ev_fire),
	newExplorationToken(3, ev_mucus),
	newExplorationToken(4, ev_mucus),
	newExplorationToken(1, ev_silence),
	newExplorationToken(1, ev_silence),
})

var goalsCoop = Cards{
	newGoal(0, "The signal"),                   // Send the signal
	newGoal(0, "Clean team"),                   // Send the signal AND (the nest is destroyed OR the ship is destroyed)
	newGoal(0, "Cut from the root"),            // Send the signal AND (the queen is dead OR the ship is destroyed)
	newGoal(0, "Special delivery"),             // Send the signal AND (the queen is dead OR the nest is destroyed)
	newGoal(0, "We don't leave anyone behind"), // Send the signal AND all compartments in the ship are explored
	newGoal(0, "First contact protocol"),       // Find out two of the intruders' weaknesses
	newGoal(0, "Post-mortem urgency"),          // Bring the blue player corpse to the Operating room and leave it there
	newGoal(0, "Destination: Earth"),           // The ship must arrive to Earth
}

var goalsCorp = Cards{
	newGoal(2, "Big game"),                  // Sendthe signal AND (the queen is dead OR the ship is destroyed)
	newGoal(2, "Restless investigator"),     // Send the signal AND all compartments in the ship are explored
	newGoal(2, "Recover the main resource"), // The ship arrives to Earth OR you are the only one surviving
	newGoal(2, "A proper burial"),           // Send the signal AND end the game in an escape pod or hybernating with the blue player corpse
	newGoal(2, "Quarantine"),                // The ship must arrive to Mars OR the ship arrives to Earth and the nest has been destroyed
	newGoal(2, "Hoarder"),                   // Finish the game in an escape pod with at least 7 items. Mision items only count if they are activated
	newGoal(3, "The oldest friend"),         // The ship must arrive to Earth OR You are the only one surviving
	newGoal(4, "Friends forever"),           // You and another player at least must survive
	newGoal(5, "Aliens in the ship"),        // Send the signal AND (the nest is destroyed OR the ship is destroyed)
}

var goalsPriv = Cards{
	newGoal(2, "Ab Ovo"),                     // Reveal the intruder egg's weakness
	newGoal(2, "Necropsy"),                   // Send the signal AND reveal the intruder corpse's weakness
	newGoal(2, "The best moment for attack"), // Player 1 must not survice OR you are the only one surviving
	newGoal(2, "My treasure"),                // Send the signal AND (end the game in an escape pod or hybernating with an intruder egg)
	newGoal(2, "More exhuberant fields"),     // Player 2 must not survive OR you are the only one surviving
	newGoal(2, "Extreme field biology"),      // Reveal at least two intruder weaknesses
	newGoal(3, "An old enemy"),               // Player 3 must not survive OR you are the only one surviving
	newGoal(4, "Taking by force"),            // Player 4 must not survive OR you are the only one surviving
	newGoal(5, "The ultimate epiphany"),      // Player 5 must not survive OR you are the only one surviving
}

var helpCards = Cards{
	newHelpCard(1),
	newHelpCard(2),
	newHelpCard(3),
	newHelpCard(4),
	newHelpCard(5),
}

var rooms1 = newDeck(Cards{
	newRoom(room_storage, red, false),
	newRoom("Fire control system", yellow, true),
	newRoom("Emergency room", green, false),
	newRoom("Laboratory", intruder_blank, true),
	newRoom("Generator", yellow, true),
	newRoom("Cooms room", yellow, true),
	newRoom("Surgery", green, false),
	newRoom("Nest", intruder_blank, false),
	newRoom("Armory", red, false),
	newRoom("Evacuation section A", all, false),
	newRoom("Evacuation section B", all, false),
})

var rooms2 = newDeck(Cards{
	newRoom("Command deck", intruder_blank, true),
	newRoom("Engine Management", yellow, true),
	newRoom("Showers", all, false),
	newRoom("Decompression control", yellow, false),
	newRoom("Escape pods control", all, false),
	newRoom("Canteen", green, false),
	newRoom("Cabins", all, false),
	newRoom("Mocosous room", intruder_blank, false),
	newRoom("Vigilance", green, true),
})

var weaknesses = newDeck(Cards{
	newWeakness(weakness_fighting),
	newWeakness(weakness_reaction),
	newWeakness(weakness_vitalpoints),
	newWeakness(weakness_phosphates),
	newWeakness(weakness_fire),
	newWeakness(weakness_energy),
	newWeakness(weakness_movement),
	newWeakness(weakness_endangered),
})

var wounds = newDeck(Cards{
	newWound(wound_leg),
	newWound(wound_leg),
	newWound(wound_leg),
	newWound(wound_body),
	newWound(wound_body),
	newWound(wound_body),
	newWound(wound_body),
	newWound(wound_arm),
	newWound(wound_arm),
	newWound(wound_arm),
	newWound(wound_hand),
	newWound(wound_hand),
	newWound(wound_hand),
	newWound(wound_hemorrhage),
	newWound(wound_hemorrhage),
	newWound(wound_hemorrhage),
})
