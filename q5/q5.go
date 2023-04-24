package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	if fromScale != "C" && fromScale != "F" && fromScale != "K" {
		return 0, fmt.Errorf("Inválido")
	}
	if toScale != "C" && toScale != "F" && toScale != "K" {
		return 0, fmt.Errorf("Inválido")
	}
	var expected float64
	if fromScale == "C" && toScale == "F" {
		expected = (temp * 9 / 5) + 32
	}
	if fromScale == "C" && toScale == "K" {
		expected = temp + 273.15
	}
	if fromScale == "F" && toScale == "C" {
		expected = (temp - 32) * 5 / 9
	}
	if fromScale == "F" && toScale == "K" {
		expected = ((temp - 32) * 5 / 9) + 273.15
	}
	if fromScale == "K" && toScale == "C" {
		expected = temp - 273.15
	}
	if fromScale == "K" && toScale == "F" {
		expected = ((temp - 273.15) * 9 / 5) + 32
	}
	return expected, nil
}
