package main

import (
	"ThanosFight/src/app/services"
)


func main() {
	armies := services.GetArmies()
	services.FightAgainstThanos(armies)
}