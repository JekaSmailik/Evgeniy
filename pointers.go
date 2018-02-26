package main
//2.3.2 Указатели/
import "fmt"

var newfile = "../newfile.txt"

func main() {
	var v int = 1
	incr(&v)
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}