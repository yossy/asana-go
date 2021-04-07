package main

import (
	"flag"
	"fmt"

	asana "github.com/yossy/asana-go"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	asana.UpdateTask()
}
