package addbinary

import (
	"math/big"
)

func addBinary(a, b string) string {
	x := new(big.Int)
	y := new(big.Int)
	sum := new(big.Int)

	x, _ = x.SetString(a, 2)
	y, _ = y.SetString(b, 2)

	return sum.Add(x, y).Text(2)
}
