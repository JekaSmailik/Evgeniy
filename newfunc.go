package main

import (
	"fmt"
)

	var global *int

func main() {
	p := new(int)
	*p = 2

	fmt.Println(p) // адрес переменной
	fmt.Println(*p)
	fmt.Println(newInt())
	fmt.Println(delta(5, 10))

	fmt.Println(gcd(8, 2))
}
func f()  {
	var x int
	x = 1
	global = &x
}
func g()  {
	y := new(int)
	*y = 1
}

func gcd(x, y int) int  {
	for y !=0 {
		x, y = y, x%y
	}
	return x
}

func newInt() *int {return new(int)}

func oldInt() *int {
	var dummy int
	return &dummy
}

func delta(old, new int) int {return new - old}
