package main

import (
	"fmt"

	"github.com/samber/lo"
)

func Test() {
	a := lo.Shuffle([]int{0, 1, 2, 3, 4, 5})
	b := lo.Filter(a, func(i int, x int) bool {
		return x%2 == 0
	})

	fmt.Printf("b: %v\n", b)
}
