package utils

import "math/big"

func ByteToBigint(b byte) *big.Int {
	return new(big.Int).SetUint64(uint64(b))
}
