package main

var actions = map[string]Cards{
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

var attacks = Cards{
	newAttack(6, attack_bite, adult, breeder, queen),
	newAttack(4, attack_bite, adult, breeder, queen),
	newAttack(4, attack_bite, adult, breeder, queen),
	newAttack(0, attack_bite, adult, breeder, queen),
	newAttack(3, attack_recall, crawler, queen),
	newAttack(2, attack_tail, queen),
	newAttack(5, attack_tail, queen),
	newAttack(3, attack_frenzy, breeder, queen),
	newAttack(4, attack_frenzy, breeder, queen),
	newAttack(4, attack_transformation, crawler),
	newAttack(5, attack_transformation, crawler),
	newAttack(3, attack_scratch, crawler, adult, breeder, queen),
	newAttack(4, attack_scratch, crawler, adult, breeder, queen),
	newAttack(5, attack_scratch, crawler, adult, breeder, queen),
	newAttack(6, attack_scratch, crawler, adult, breeder, queen),
	newAttack(3, attack_claws, adult, breeder, queen),
	newAttack(4, attack_claws, adult, breeder, queen),
	newAttack(5, attack_claws, adult, breeder, queen),
	newAttack(0, attack_claws, adult, breeder, queen),
	newAttack(5, attack_mucosity, crawler, adult, breeder, queen),
}

var characters = Cards{
	newCard(captain),
	newCard(explorer),
	newCard(mechanic),
	newCard(pilot),
	newCard(scientist),
	newCard(soldier),
}

var contamination = Cards{
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
}

var coordinates = Cards{
	newCoordinates("CBDA"),
	newCoordinates("ABCD"),
	newCoordinates("ACDB"),
	newCoordinates("DABC"),
	newCoordinates("DCAB"),
	newCoordinates("BDCA"),
	newCoordinates("CABD"),
	newCoordinates("BDAC"),
}

var corridors Corridors = Corridors{
	newCorridor(A01, A02, 3),
	newCorridor(A01, A03, 1, 2),
	newCorridor(A01, A04, 4),
	newCorridor(A02, S01, 1, 2),
	newCorridor(A02, A06, 4),
	newCorridor(A03, A07, 3, 4),
	newCorridor(A04, A08, 1),
	newCorridor(A04, S01, 2, 3),
	newCorridor(A05, A10, 1, 2),
	newCorridor(A05, A06, 3),
	newCorridor(A05, S01, 4),
	newCorridor(A06, A07, 1),
	newCorridor(A06, A11, 2),
	newCorridor(A07, A08, 2),
	newCorridor(A08, A09, 4),
	newCorridor(A08, A11, 3),
	newCorridor(A09, A12, 1, 2),
	newCorridor(A09, S01, 3),
	newCorridor(A10, A13, 3, 4),
	newCorridor(A11, A14, 4),
	newCorridor(A11, A15, 1),
	newCorridor(A12, A16, 3, 4),
	newCorridor(A13, A14, 1),
	newCorridor(A13, A19, 2),
	newCorridor(A14, A17, 2),
	newCorridor(A14, S01, 3),
	newCorridor(A15, A16, 2),
	newCorridor(A15, A18, 3),
	newCorridor(A15, S01, 4),
	newCorridor(A16, A21, 1),
	newCorridor(A17, A19, 1),
	newCorridor(A17, A20, 3, 4),
	newCorridor(A18, A20, 1, 2),
	newCorridor(A18, A21, 4),
	newCorridor(A19, S01, 3, 4),
	newCorridor(A21, S01, 2, 3),
}

var events = Cards{
	newEvent(1, "Fuga de refrigerante", adult, breeder, queen),
	newEvent(3, "Eclosion", adult, breeder),
	newEvent(4, "Ruido en los pasillos de servicio", adult, breeder),
	newEvent(2, "Fallo del soporte vital", adult, breeder, queen),
	newEvent(2, "Proteccion de los huevos", breeder, queen),
	newEvent(4, "Acechar", crawler, adult),
	newEvent(3, "El rastro de la presa", crawler, adult),
	newEvent(4, "Compuestos inflamables", breeder, queen),
	newEvent(2, "Daño", crawler, adult),
	newEvent(3, "Nido", crawler, breeder, queen),
	newEvent(1, "Fuego perjudicial", crawler, adult),
	newEvent(1, "Apertura de compuertas", adult, breeder),
	newEvent(4, "Fuego devastador", crawler, breeder, queen),
	newEvent(1, "Regeneracion", breeder, queen),
	newEvent(2, event_failure, adult, breeder),
	newEvent(1, "Nacimiento", crawler, breeder, queen),
	newEvent(3, "Caceria", crawler, queen),
	newEvent(4, "cortocircuito", adult, breeder, queen),
	newEvent(2, "caceria", crawler, breeder, queen),
	newEvent(3, "Eyeccion de capsula de evacuacion", adult, breeder, queen),
}

var explorationTokens = Cards{
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
}

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

var rooms1 = Cards{
	newRoom(room_storage, red, false),
	newRoom("Fire control system", yellow, true),
	newRoom("Emergency room", green, false),
	newRoom("Laboratory", blank, true),
	newRoom("Generator", yellow, true),
	newRoom("Cooms room", yellow, true),
	newRoom("Surgery", green, false),
	newRoom("Nest", blank, false),
	newRoom("Armory", red, false),
	newRoom("Evacuation section A", all, false),
	newRoom("Evacuation section B", all, false),
}

var rooms2 = Cards{
	newRoom("Command deck", blank, true),
	newRoom("Engine Management", yellow, true),
	newRoom("Showers", all, false),
	newRoom("Decompression control", yellow, false),
	newRoom("Escape pods control", all, false),
	newRoom("Canteen", green, false),
	newRoom("Cabins", all, false),
	newRoom("Mocosous room", blank, false),
	newRoom("Vigilance", green, true),
}

var weaknesses = Cards{
	newWeakness(weakness_fighting),
	newWeakness(weakness_reaction),
	newWeakness(weakness_vitalpoints),
	newWeakness(weakness_phosphates),
	newWeakness(weakness_fire),
	newWeakness(weakness_energy),
	newWeakness(weakness_movement),
	newWeakness(weakness_endangered),
}

var wounds = Cards{
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
}
