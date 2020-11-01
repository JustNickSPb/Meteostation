//Package tables handles all code needed to show metheo data as a table
package tables

import (
	temptypes "Meteostation/pkg/temptypes"
	"fmt"
)

const tableGrid = "======================"

func celsToFar(c temptypes.Celsius) temptypes.Farenheit {
	return c.ToFarenheit()
}

func farToCels(f temptypes.Farenheit) temptypes.Celsius {
	return f.ToCelsius()
}

//DrawTable is to draw a table, "C" is for Celsius, "F" - for Farenheit
/*USAGE:
DrawTable("C","F"):
|C|F|
======================
|40|104|
DrawTable("F", "C"):
|F|C|
======================
|40|4.444444444444445|
*/
func DrawTable(sourceDataCaption string, outputDataCaption string) {

	fmt.Println(tableGrid)
	fmt.Printf("|%8s|%8s|\n", sourceDataCaption, outputDataCaption)
	fmt.Println(tableGrid)
	if sourceDataCaption == "C" {
		var cels temptypes.Celsius
		for cels = 40; cels <= 100; cels += 5 {
			fmt.Printf("|%8v|%8v|\n", cels, celsToFar(cels))
			//| %8s | %8s |\n
		}
	}
	if sourceDataCaption == "F" {
		var far temptypes.Farenheit
		for far = 40; far <= 100; far += 5 {
			fmt.Printf("|%8v|%8v|\n", far, farToCels(far))
		}
	}

	fmt.Println(tableGrid)
}
