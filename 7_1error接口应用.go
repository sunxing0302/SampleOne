// 7_1error接口应用.go
package main

import (
	"errors"
	"fmt"
)

func myDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("除数不能为0")

	} else {
		result = a / b
	}
	return

}

func main() {
	result, err := myDiv(10, 3)
	fmt.Println(result, err)

	result2, err2 := myDiv(10, 0)
	fmt.Println(result2, err2)
}
