package bonus

import "fmt"

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	var damage int
	if characterLevel <= 0 || enemyLevel <= 0 {
		fmt.Errorf("Nível inválido.")
	} else {
		switch {
		case characterLevel > enemyLevel:
			damage = characterLevel * 10
		case characterLevel < enemyLevel:
			damage = characterLevel * 5
		case characterLevel == enemyLevel:
			damage = characterLevel * 7
		case enemyLevel > 100:
			damage = characterLevel * 20
		case (characterLevel - enemyLevel) > 20:
			damage = characterLevel * 5
		case (characterLevel - enemyLevel) < 20:
			damage = characterLevel * 2
		}
	}
	return damage, nil
}
