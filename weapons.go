package weapons

type weapons struct {
	Sworddmg int
	Axedmg int
	Bowdmg int
	Trenchraidingclubdmg int
	Fireball int 
	Magicstick int
}

type Sworddmg int
const(
criticalhit Sworddmg = iota
normalhit Sworddmg = iota
blockedhit Sworddmg = iota 
)

type Axedmg int
const(
	criticalhit Axedmg = iota 
	normalhit Axedmg = iota 
	blockedhit Axedmg = iota
)

type Bowdmg int 
const(
	criticalhit Bowdmg = iota 
	normalhit Bowdmg = iota 
	blockedhit Bowdmg = iota 
)

type Trenchraidingclubdmg int 
const(
	criticalhit Trenchraidingclubdmg = iota 
	normalhit Trenchraidingclubdmg = iota 
	blockedhit Trenchraidingclubdmg = iota 
)

type Fireball int 
const(
	criticalhit Fireball= iota 
	normalhit Fireball = iota 
	blockedhit Fireball = iota 
)

type Magicstick int 
const(
	criticalhit Magicstick = iota 
	normalhit Magicstick = iota 
	blockedhit Fireball = iota 
)

