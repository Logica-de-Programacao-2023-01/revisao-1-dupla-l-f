package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {

	numerodeletras := 0


	text1 := strings.ReplaceAll(text, "!", "")
	text1 = strings.ReplaceAll(text1, ",", "")
	text1 = strings.ReplaceAll(text1, ".", "")
	text1 = strings.ReplaceAll(text1, "(", "")
	text1 = strings.ReplaceAll(text1, ")", "")
	text1 = strings.ReplaceAll(text1, "[", "")
	text1 = strings.ReplaceAll(text1, "]", "")
	text1 = strings.ReplaceAll(text1, "{", "")
	text1 = strings.ReplaceAll(text1, "}", "")
	text1 = strings.ReplaceAll(text1, "?", "")
	text1 = strings.ReplaceAll(text1, "0", "")
	text1 = strings.ReplaceAll(text1, "1", "")
	text1 = strings.ReplaceAll(text1, "2", "")
	text1 = strings.ReplaceAll(text1, "3", "")
	text1 = strings.ReplaceAll(text1, "4", "")
	text1 = strings.ReplaceAll(text1, "5", "")
	text1 = strings.ReplaceAll(text1, "6", "")
	text1 = strings.ReplaceAll(text1, "7", "")
	text1 = strings.ReplaceAll(text1, "8", "")
	text1 = strings.ReplaceAll(text1, "9", "")

for _, letra := range text1 {
	if unicode.IsLetter(letra) {
	numerodeletras++}
}
if numerodeletras == 0 {
	return 0, fmt.Errorf("Erro ")
}
numerodepalavras := len(strings.Fields(text1))

media := float64(numerodeletras)/float64(numerodepalavras)

	return media, nil
}

