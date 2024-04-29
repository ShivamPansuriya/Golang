package advance

import (
	"errors"
	"fmt"
)

func isEligible(age int) (bool, error) {
	if age == 30 {
		return true, errors.New("your age is 30")
	}
	return false, nil
}

func Error() {

	if a, err := isEligible(32); err != nil {
		fmt.Println("there is an error", err)
	} else {
		fmt.Println(a)
	}
}
