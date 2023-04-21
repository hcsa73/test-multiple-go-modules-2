package core

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const CONST = "CORE-V2.0.0"

func getConstReversed() string {
	return stringutil.Reverse(CONST)
}

func GetInfo() string {
	out := "Hello from core!\n"
	out += fmt.Sprintf("Const : %v\n", CONST)
	out += fmt.Sprintf("Const reversed: %v\n", getConstReversed())
	return out
}
