package calc

import (
	"fmt"
)

func Do(op string, a, b float64) (float64, error) {
	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, fmt.Errorf("error, not divisible by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknow operaton")
	}
}
