package main

type Creature struct {
	Sts Status
}

type Status struct {
	MaxHp int
	Hp    int
	MaxMp int
	Mp    int
}

func InitCreature(mhp, mmp int) *Creature {
	return &Creature{
		Sts: Status{
			MaxHp: mhp,
			Hp:    mhp,
			MaxMp: mmp,
			Mp:    mmp,
		},
	}
}
