//Package temptypes holds types of temperature scales
package temptypes

//Celsius is a type of temperature to hold it in Celsius scale
type Celsius float64

func (c Celsius) toKelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c Celsius) ToFarenheit() Farenheit {
	return Farenheit((c * 9.0 / 5.0) + 32.0)
}

type kelvin float64

func (k kelvin) ToCelsius() Celsius {
	return Celsius(k - 273.15)
}

func (k kelvin) ToFarenheit() Farenheit {
	return Farenheit((k-273.15)*9.0/5.0 + 32.0)
}

//Farenheit is type of temperature to hold it in farenheit scale
type Farenheit float64

func (f Farenheit) ToCelsius() Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

func (f Farenheit) toKelvin() kelvin {
	return kelvin((f-32.0)*5.0/9.0 + 273.15)
}
