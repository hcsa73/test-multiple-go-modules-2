package main

import (
	"fmt"

	"github.com/hcsa73/test-multiple-go-modules-2/core"
	"github.com/hcsa73/test-multiple-go-modules-2/services/argus"
)

func main() {
	fmt.Println("Hello from main")
	fmt.Println("==== core info ====")
	fmt.Print(core.GetInfo())
	fmt.Println("==== argus info ====")
	fmt.Print(argus.GetInfo())
}
