package main

import (
	"fmt"
	"math/rand"
	"time"
)

type celsius float64

func (c celsius) toKelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) toFarenheit() farenheit {
	return farenheit((c * 9.0 / 5.0) + 32.0)
}

type kelvin float64

func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) toFarenheit() farenheit {
	return farenheit((k-273.15)*9.0/5.0 + 32.0)
}

type farenheit float64

func (f farenheit) toCelsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f farenheit) toKelvin() kelvin {
	return kelvin((f-32.0)*5.0/9.0 + 273.15)
}

type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0 // TODO: Внедрить реальный сенсор
}

func measureTemperature(samples int, sensor sensor) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := fakeSensor
	measureTemperature(5, sensor)
}
