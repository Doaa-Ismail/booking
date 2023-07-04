package functions

import (
	"errors"
	"fmt"
)

func DivideNum(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	return x / y, nil
}
func Sumfunc(x, y int) int {
	fmt.Printf("Sum of %d + %d is equal : %d\n", x, y, x+y)
	return x + y
}
