package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	// Implemente sua solução aqui
	expectedMax := -9999999999
	expectedMin := 99999999999
	var soma, embaixo, expectedAverage float64
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("Lista vazia")
	}
	for i := 0; i < len(numbers); i++ {
		soma += float64(numbers[i])
		embaixo++
		if numbers[i] > expectedMax {
			expectedMax = numbers[i]
		}
		if numbers[i] < expectedMin {
			expectedMin = numbers[i]
		}
	}
	expectedAverage = soma / embaixo
	return expectedMax, expectedMin, expectedAverage, nil
}
