package q4

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {

	var taxCodeValor float64
	var stateValor float64
	var finalPrice float64

	if taxCode == 1 {
		taxCodeValor = 0.05
	}
	if taxCode == 2 {
		taxCodeValor = 0.1
	}
	if taxCode == 3 {
		taxCodeValor = 0.15
	}
	if taxCode != 1 || taxCode != 2 || taxCode != 3 {
		return 0, fmt.Errorf("Este não é um código de imposto válido ")
	}
	if state == "SP" {
		stateValor = 0.1
	}
	if state == "RJ" {
		stateValor = 0.15
}
	if state == "MG" {
		stateValor = 0.2
	}
	if state == "ES" {
		stateValor = 0.25
	}
	if state != "SP" || state != "RJ" || state != "MG" || state != "ES" {
		stateValor = 0.3
	}
	



	finalPrice = basePrice + (basePrice * float64(stateValor)) + (basePrice * float64(taxCodeValor))

	return finalPrice, nil
}
