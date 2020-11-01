//Package sensor is to get temperature from somewhere to metheostation
package sensor

import (
	"fmt"
	"math/rand"
	"time"
)

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
