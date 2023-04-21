package core

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const CONST = "CORE-V1.0.0"

func GetConst() string {
	return CONST
}

func getConstReversed() string {
	return stringutil.Reverse(CONST)
}

func GetInfo() string {
	out := "Hello from core!\n"
	out += fmt.Sprintf("Const reversed: %v\n", getConstReversed())
	return out
}
