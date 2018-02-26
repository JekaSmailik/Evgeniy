package main
// 2.3 Переиенные
import (
	"fmt"
	"os"
)

//	var name type = expression

func main() {
	var s string
	fmt.Println(s)

	var i, j ,k int // int, int, int
	var b, f, a = true, 2.3, "four" // bool, float64, string

	fmt.Println(i, j ,k)
	fmt.Println(b, f, a)

//	var f, err = os.Open(name) // os.Open возврощает файл и ошибку
}