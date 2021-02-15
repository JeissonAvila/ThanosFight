package services

import (
	"ThanosFight/src/app/entities"
	"ThanosFight/src/app/interfaces"
)

func GetArmies () []interfaces.Army {
	army1 := entities.NationalArmy{"Air Forces", 10, 1000}
	army2 := entities.NationalArmy{"Marines", 20, 1000}
	army3 := entities.NationalArmy{"Space Forces", 40, 1000}
	army4 := entities.HeroArmy{"Iron Man", 100000, 50000}
	army5 := entities.HeroArmy{"Black Panther", 50000, 100000}
	army6 := entities.JeiArmy{}
	chuck := entities.ChuckArmy{}

	armies := []interfaces.Army{army1, army2, army3, army4, army5, army6, chuck}

	return armies
}
