package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func plusplus(a ...int) int {
	total := 0
	for _,num := range a{
		total += num
	}
	return total
}

func vals(a int) func() (int){
	return func() int{
		a++
		return a
	}
}

func main() {
	res := plus(1,2)
	fmt.Println(res)

	res = plusplus(1,2,3,4,5,6)
	fmt.Println(res)

	a := vals(3)
	fmt.Println(a())
	fmt.Println(a())

}