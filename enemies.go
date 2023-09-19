package main 

type enemies struct {
	Tall int
	Power int
	HP int
	Enemiesweapons int
}

type Tall int
const(
	big Tall = iota
	medium Tall = iota 
	small Tall = iota
)

type Power int 
const(
	verypowerful Power = iota
	powerful = iota 
	weak = iota
)

type HP int
const(
	150 HP = iota 
	100 HP = iota 
	>50 HP = iota
)

type Enemiesweapons int 
const(
	sword Enemiesweapons = iota 
	knife Enemiesweapons = iota 
	hand Enemiesweapons = iota
)