package main

import (
	"os"
	"log"
	"fmt"
)

var cwd string

func init()  {
	var err error
	cwd, err = os.Getwd()
	if err != nil{
		log.Fatalf("Ощибка os.Getwd: %v", err)
	}
}

func main() {

	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
}
	/*
	a := "hello HELLO"
	for _, a := range a {
		a := a + 'a' - 'A'
		fmt.Printf("%c", a)
	}
	*/
/*
	if x := fi(); x == 0{
		fmt.Println(x)
	} else if y := gi(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
//fmt.Println(x, y) // ошибка компиляции: х и у здесь не видимы
*/
//alfavit()
/*
func fi() int {return 3}
func gi(x int) int {return 3}
func initmy() { / ... / }
func alfavit()  {
	for i := 1040; i < 1104; i++ {
		//fmt.Printf("%s:%d\n", i, i)
		fmt.Printf("%c:%d\n", i, i)
	}
	// Вывод символа
	fmt.Println(string(165), "\n")
}
*/