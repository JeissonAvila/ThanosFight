package entities

import "fmt"

type HeroArmy struct {
	HeroName       string
	HeroPower      int
	SideKicksPower int
}

func (h HeroArmy) Name() string {
	return fmt.Sprintf("Army with %s", h.HeroName)
}

func (h HeroArmy) Power() int {
	return h.HeroPower + h.SideKicksPower
}
