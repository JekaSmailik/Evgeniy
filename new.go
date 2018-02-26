package main

import (
	"fmt"
	"github.com/derekparker/delve/pkg/dwarf/reader"
)

type MyReader struct{}

func (MyReader) Read(b [] byte) (int,  error)  {
	b[0] = byte('A')
	fmt.Println(b)
	return 1, nil
}
// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	fmt.Println(MyReader.Read([]byte))
	reader.Validate(MyReader{})
}