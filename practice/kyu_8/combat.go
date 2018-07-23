package kyu_8

func Combat(health, damage float64) float64 {
	if health < damage {
		return 0.0
	} else {
		return health - damage
	}
}
