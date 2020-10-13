package main

import "fmt"

type celsius float64
type kelvin float64
type farenheit float64

// kelvinToCelcius конвертирует Кельвины в градусы Цельсия
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// celsiusToFarenheit ковертирует градусы Цельсия в градусы Фаренгейта
func celsiusToFarenheit(c celsius) farenheit {
	return farenheit((c * 9.0 / 5.0) + 32.0)
}

// kelvinToFarenheit конвертирует Кельвины в градусы Фаренгейта
func kelvinToFarenheit(k kelvin) farenheit {
	return farenheit((k-273.15)*9.0/5.0 + 32.0)
}

// celsiusToKelvin конвертирует градусы Цельсия в Кельвины
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var t celsius = 127
	fmt.Println(celsiusToKelvin(t))
}
