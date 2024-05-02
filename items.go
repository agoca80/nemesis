package main

type item struct {
	*card
}

type Items Cards

func NewItem(name string) *item {
	return &item{
		card: newCard(name),
	}
}

type weapon struct {
	*item
	ammo int
}

func (w *weapon) Shoot(ammo int) int {
	w.ammo -= ammo
	return w.ammo
}

func NewWeapon() *weapon {
	return &weapon{
		item: NewItem("weapon"),
	}
}
