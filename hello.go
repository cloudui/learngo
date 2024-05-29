package main

import (
	"errors"
	"fmt"
)

func main() {
	var mystr = "Try me"
	var i uint32 = 0

	var j, k int = 5, 6

	m := j * k

	fmt.Println("Hello, World!")
	fmt.Println(mystr + "Testing things out")
	fmt.Println(i*3 - 4)
	fmt.Println(m)
	fmt.Println(getArea(3, 4))
	fmt.Println(getPerimeter(3, 4))
	fmt.Println(getAreaPerimeter(3, 4))

	val, err := divide(3, 0)
	if err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}

}

func getArea(x int, y int) int {
	return x * y
}

func getPerimeter(x int, y int) int {
	return 2*x + 2*y
}

func getAreaPerimeter(x int, y int) (int, int) {
	return getArea(x, y), getPerimeter(x, y)
}

func divide(x float64, y float64) (float64, error) {
	var err error
	if y == 0. {
		err = errors.New("cannot divide by zero")
		return 0., err
	}

	return x / y, err
}
