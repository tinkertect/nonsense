package content

import (
	"fmt"
	"nonsense/arithmetic"
)

type Op = func(int) int

func MakeTable(from int, to int, k int) []string {
	result := []string{}

	var op Op

	switch k {
	case 2:
		op = arithmetic.Times2
	case 3:
		op = arithmetic.Times3
	case 4:
		op = arithmetic.Times4
	default:
		panic(fmt.Sprintf("cant process k = %d", k))
	}

	for i := from; i <= to; i++ {
		product := op(i)
		text := fmt.Sprintf("%d times %d = %d", i, k, product)
		result = append(result, text)
	}

	return result
}
