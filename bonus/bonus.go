package bonus

import "fmt"

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	var damage int
	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, fmt.Errorf("Nível inválido.")
	} else {
		switch {
		case characterLevel < enemyLevel:
			damage = characterLevel * 5
		case characterLevel == enemyLevel:
			damage = characterLevel * 7
		case characterLevel > 100:
			damage = characterLevel * 20
		case enemyLevel > 100:
			damage = characterLevel * 2
		case (characterLevel - enemyLevel) > 20:
			damage = characterLevel * 5
		case (characterLevel - enemyLevel) < 20:
			damage = characterLevel * 2
		}
	} 
	if characterLevel > enemyLevel {
			damage = characterLevel * 10
			if characterLevel > 100 && (characterLevel-enemyLevel) > 20 {
				damage *= 20 * 5
			}
	}
	return damage, nil
}
