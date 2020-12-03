package main

import "fmt"
import "reflect"

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Type :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Value :", reflectValue.Int())
	}
}