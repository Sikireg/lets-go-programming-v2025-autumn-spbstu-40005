package main

import (
	"fmt"

	"alexandra.karnauhova/task-2-1/internal/temperature"
)

func main() {
	var numberDepartment int
	lol := temperature.RangeTemp{Range: ">=", Temperature: 20}

	_, err := fmt.Scan(&numberDepartment)
	if err != nil {
		fmt.Println("Invalid number of departments")
		return
	}
	//пару значений, знак и число
	//массив этих пар +метод в массиве для подсчета норм знач
	//вывести все:) ( Т_Т )

}
