package main

import (
	"fmt"

	prefixsum "github.com/chinmayb13/schrodinger-cat/internal/array/prefix_sum"
	"github.com/chinmayb13/schrodinger-cat/internal/primer/dsa"
)

func main() {
	v := dsa.GetMinDistance("o..o.xx.ooo.x.o.o")
	fmt.Println(v)
	pa := prefixsum.GetProductArrayAlternate([]int{1, 2, 3, 4, 5})
	fmt.Println(pa)
}
