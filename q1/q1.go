package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	// Implemente sua solução aqui
	var SomaHistory float64
	expectedDiscount := 0.0
	for _, valor := range purchaseHistory {
		SomaHistory += valor
	}
	media := (SomaHistory) / float64(len(purchaseHistory))
	switch {
	case currentPurchase <= 0:
		return 0, fmt.Errorf("Valor de compra inválido")
	case SomaHistory > 1000:
		expectedDiscount = currentPurchase * 0.1
	case SomaHistory <= 1000:
		expectedDiscount = currentPurchase * 0.05
	case SomaHistory <= 500:
		expectedDiscount = currentPurchase * 0.02
	case SomaHistory == 0:
		expectedDiscount = currentPurchase * 0.1
	case media > 1000:
		expectedDiscount = currentPurchase * 0.2
	}
	return expectedDiscount, nil
}

