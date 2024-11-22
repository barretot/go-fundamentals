package challengers

import "fmt"

func Double(value *int, valueToDouble int) int {
	doubleValue := *value * valueToDouble

	fmt.Println(doubleValue)

	return doubleValue
}

func FloatValue(value *float64, divider int) float64 {
	divideValue := *value / float64(divider)

	fmt.Println(divideValue)

	return divideValue
}
