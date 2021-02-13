package services

import (
	"fmt"
	"ThanosFight/src/app/interfaces"
)

func FightAgainstThanos(armies []interfaces.Army) {
	thanosLife := 9999999
	for _, army := range armies {
		fmt.Printf("%s fight against Thanos with a force of %d\n", army.Name(), army.Power())
		thanosLife -= army.Power()
	}
	if thanosLife <= 0 {
		fmt.Println("The terrible Thanos has been defeated")
	} else {
		fmt.Printf("Thanos destroy entire universe, and his life is %d yet\n", thanosLife)
	}
}
