package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	// Implemente sua solução aqui
	return 0, 0, 0, nil
	var Max, Min int 
	var soma, embaixo, media float64
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("Lista vazia")
	}
	for i := 0; i < len(numbers); i++ {
		soma += float64(numbers[i])
		embaixo++
		if numbers[i] > Max{
			Max = numbers[i]
		} else if numbers[i] < Min {
			Min = numbers[i]
		}
	}
media = soma/embaixo
	return Max, Min, media, nil
}
