package entities

type NationalArmy struct {
	ArmyName  string
	UnitPower int
	Soldiers  int
}

func (n NationalArmy) Name() string {
	return n.ArmyName
}

func (n NationalArmy) Power() int {
	return n.UnitPower * n.Soldiers
}
