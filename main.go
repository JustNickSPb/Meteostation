package main

import "fmt"

type celsius float64
type kelvin float64
type farenheit float64

func (c celsius) toKelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) toFarenheit() farenheit {
	return farenheit((c * 9.0 / 5.0) + 32.0)
}

func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) toFarenheit() farenheit {
	return farenheit((k-273.15)*9.0/5.0 + 32.0)
}

func (f farenheit) toCelsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f farenheit) toKelvin() kelvin {
	return kelvin((f-32.0)*5.0/9.0 + 273.15)
}

func main() {
	var t celsius = 127
	fmt.Println(t.toFarenheit())
}
