package main

import (
	"fmt"
	"time"
)

func getType(typeToCheck interface{}) string {

	switch typeToCheck.(type) {
	case int:
		return "Integer"
	case float32:
		return "Float 32 bits"
	case float64:
		return "Float 64 bits"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Unknown Type"
	}

}

func main() {
	fmt.Println(getType(2.3))
	fmt.Println(getType("João"))
	fmt.Println(getType(2))
	fmt.Println(getType(func() {}))
	fmt.Println(getType(time.Now()))
}
