package bitmanipulations1

import (
	"fmt"
	"math"
	"strconv"
)

func GetBitSequence(inp int) {
	shift := 1 << inp
	for i := 0; i < int(math.Pow(2, float64(inp))); i++ {
		b := strconv.FormatInt(int64(i)|int64(shift), 2)
		fmt.Println(b[1:])
	}
}
