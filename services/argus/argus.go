package argus

import (
	"fmt"

	"github.com/hcsa73/test-multiple-go-modules-2/core"
)

const CONST = "SERVICES/ARGUS-V1.0.1"

func GetConst() string {
	return CONST
}

func GetInfo() string {
	out := "Hello from services/argus!\n"
	out += fmt.Sprintf("Const: %v\n", GetConst())
	out += fmt.Sprintf("Const from core: %v\n", core.GetConst())
	return out
}
