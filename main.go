package main

import (
	"fmt"

	"github.com/sourcegraph/conc/pool"
)

func main() {
	fmt.Println("Hello, world!!")
	p := pool.NewWithResults[int]().WithMaxGoroutines(8)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, result := range nums {

		p.Go(func() int {
			return result
		})
	}
	res := p.Wait()
	fmt.Printf("res: %v\n", res)
}
