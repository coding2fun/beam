package main

import (
	"fmt"

	"coding2fun.in/beam/pkg/runtime"
)

func main() {
	fmt.Println("Hello Beam")

	c := runtime.NewIMCache()
	e := runtime.NewDBStorage()

	o := runtime.NewEngine(c, e)

	o.Process()
}
