package main

type Inventory struct {
	Potion int 
	Shield int 
	Weapons string
}

type Weapons int
	const (
		sword Weapons = iota
		bow Weapons = iota
		axe Weapons = iota
	)