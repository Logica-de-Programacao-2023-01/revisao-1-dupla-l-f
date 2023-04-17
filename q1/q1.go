package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	var somaHistory float64
	for i := 0; i < len(purchaseHistory); i++ {
		somaHistory += purchaseHistory[i]
	}
	somaFull := somaHistory + currentPurchase
	discount := 0.0
	switch {
	case currentPurchase <= 0:
		return 0, fmt.Errorf("Valor inválido")
	case somaFull == float64(0):
		discount = currentPurchase * 0.1
	case somaFull/float64(len(purchaseHistory)) > float64(1000):
		discount = currentPurchase * 0.2
	case somaFull > float64(1000) && somaHistory == float64(0):
		discount = currentPurchase * 0.1
	case somaFull <= float64(1000):
		discount = currentPurchase * 0.05
	case somaFull <= float64(500):
		discount = currentPurchase * 0.02
	}
	return discount, nil
}
