package main

import "slices"

type Symbols []string

func (s Symbols) Includes(symbol string) bool {
	return slices.Contains(s, symbol)
}

// Actions
const (
	// Basic actions
	basic_move     = "Move"
	basic_fire     = "Fire"
	basic_fight    = "Fight"
	basic_pickup   = "Pick up"
	basic_exchange = "Exchange"
	basic_prepare  = "Prepare"
	basic_sneak    = "Sneak"

	// Card actions
	action_adrenaline        = "Adrenaline"
	action_aimed_fire        = "Aimed fire"
	action_automatic_fire    = "Automatic fire"
	action_basic_repairs     = "Basic repairs"
	action_block_access      = "Block access"
	action_computer_skills   = "Computer skills"
	action_covering_fire     = "Covering fire"
	action_demolition        = "Demolition"
	action_fireworks         = "Fireworks"
	action_improvisation     = "Improvisation"
	action_interruption      = "Interruption"
	action_intranet          = "Intranet"
	action_motivation        = "Motivation"
	action_old_friend        = "Old friend"
	action_order             = "Order"
	action_piloting          = "Piloting"
	action_quick_repairs     = "Quick repairs"
	action_register          = "Register"
	action_reload            = "Reload"
	action_repair            = "Repair"
	action_rest              = "Rest"
	action_risk_evaluation   = "Risk evaluation"
	action_scavenger         = "Scavenger"
	action_scout             = "Scout"
	action_service_corridors = "Service corridors"
	action_ship_knownledge   = "Ship knownledge"
	action_still_nerves      = "Still nerves"
	action_suppresion        = "Suppresion fire"
)

// area symbols
const (
	A00 = iota
	A01
	A02
	A03
	A04
	A05
	A06
	A07
	A08
	A09
	A10
	A11
	A12
	A13
	A14
	A15
	A16
	A17
	A18
	A19
	A20
	A21
	S01
	S02
)

// Intruder symbols
const (
	intruder_blank   = "blank"
	intruder_egg     = "egg"
	intruder_larva   = "larva"
	intruder_crawler = "crawler"
	intruder_adult   = "adult"
	intruder_breeder = "breeder"
	intruder_queen   = "queen"
)

// Color symbols
const (
	all    = "all"
	blue   = "blue"
	green  = "green"
	red    = "red"
	yellow = "yellow"
)

// Character symbols
const (
	captain   = "captain"
	explorer  = "explorer"
	scientist = "scientist"
	soldier   = "soldier"
	mechanic  = "mechanic"
	pilot     = "pilot"
)

// Door states
const (
	door_open   = "O"
	door_closed = "|"
	door_broken = "X"
)

// event symbols
const (
	event_failure = "failure"
)

// Exploration token events
const (
	ev_silence = "silence"
	ev_danger  = "danger"
	ev_fire    = "fire"
	ev_damaged = "damaged"
	ev_mucus   = "mucus"
	ev_door    = "door"
)

// game steps
const (
	step_draw      = "draw action cards"
	step_token     = "pass Token"
	step_turn      = "turn"
	step_counters  = "counters"
	step_attack    = "attack"
	step_fire      = "fire Damage"
	step_event     = "event"
	step_evolution = "evolution"
)

// noise events
const (
	silence = "silence"
	danger  = "danger"
	n1      = "1"
	n2      = "2"
	n3      = "3"
	n4      = "4"
)

// player states
const (
	player_alive    = "alive"
	player_dead     = "dead"
	player_sleeping = "sleeping"
	player_escaped  = "escaped"
)

// room symbols
const (
	room_1             = "class1"
	room_2             = "class2"
	room_cockpit       = "Cockpit"
	room_engine1       = "Engine #1"
	room_engine2       = "Engine #2"
	room_engine3       = "Engine #3"
	room_hibernatorium = "Hibernatorium"
	room_nest          = "Nest"
	room_service1      = "Service Tunnels #1"
	room_service2      = "Service Tunnels #1"
	room_slime         = "Room covered in Slime"
	room_storage       = "Storage"
)

// weaknesses
const (
	weakness_fighting    = "The fighting tactics"
	weakness_reaction    = "Reaction to danger"
	weakness_vitalpoints = "Vital points"
	weakness_phosphates  = "Weak to phosphates"
	weakness_fire        = "Weak to fire"
	weakness_energy      = "Weak to energy"
	weakness_movement    = "The way of moving"
	weakness_endangered  = "Endangered species"
)

// Wound cards
const (
	wound_arm        = "Arm wound"
	wound_body       = "Body wound"
	wound_hand       = "Hand wound"
	wound_hemorrhage = "Hemorrhage"
	wound_leg        = "Leg wound"
)
