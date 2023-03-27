package main

import (
	_ "embed"
	"fmt"
)

//go:embed greeting
var greeting string

func main() {
	fmt.Printf("%d\n%s", foobar(), greeting)
}
