package main

import (
	"fmt"

	b "github.com/KeiichiHirobe/gomodule-example-b"
	c "github.com/KeiichiHirobe/gomodule-example-c"
)

func main() {
	fmt.Println(c.GetVersionConfirm())
	fmt.Println(b.GetThroughB())
	fmt.Println(b.GetDecolatedThroughB())
	b.SetThroughB("update by module b")
	fmt.Println(b.GetThroughB())
	fmt.Println(c.GetVersionConfirm())
}
