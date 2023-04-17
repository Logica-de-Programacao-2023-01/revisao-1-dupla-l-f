package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	var letras float64
	var espaços float64
	if text == "" || text == " " {
		return 0.0, fmt.Errorf("Texto inválido. ")
	}
	for i := 0; i < len(text); i++ {
		if strings.ContainsAny(text, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz") {
			letras++
			if letras == 0 {
				return 0.0, fmt.Errorf("Texto inválido")
			}
		}
	}
	espaços = float64(len(strings.Split(text, " ")))
	media := letras / espaços
	return media, nil
}

