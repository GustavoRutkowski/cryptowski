package bruteforce

import (
	"context"
	"math/big"
	"sync"
	"github.com/GustavoRutkowski/cryptowski/internal/cryptowski"
	"github.com/GustavoRutkowski/cryptowski/internal/utils"
)

// Converts an index to a key in base N.
func keyFrom(index *big.Int, keysize uint8) string {
	bytes := make([]byte, keysize)

	n := new(big.Int).Set(index)
	mod := new(big.Int)

	for i := int(keysize) - 1; i >= 0; i-- {
		mod.Mod(n, bigBase)
		bytes[i] = KEY_CHARSET[mod.Int64()]
		n.Div(n, bigBase)
	}

	return string(bytes)
}

// Tests a key unit (a single index).
// Must be called within a goroutine.
func work(
	data []byte, keysize uint8, decode cryptowski.DecodeAlg,
	i *big.Int,
	ch chan<- result,
	cancel context.CancelFunc,
) {
	key := keyFrom(i, keysize)
	decoded := decode(data, key)

	success, err := verifyResult(decoded)
	if err != nil {
		cancel()
		select {
		case ch <- result{err: err}:
		default:
		}
		return
	}

	if success {
		cancel()
		select {
		case ch <- result{key: key, decoded: decoded}:
		default:
		}
		return
	}
}

// Workers are responsible for processing keys in a specific index range.
// They must be called as goroutines.
func worker(
	data []byte, keysize uint8, decode cryptowski.DecodeAlg,
	interval utils.Interval[*big.Int],
	ch chan<- result,
	ctx context.Context, cancel context.CancelFunc,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	i := new(big.Int).Set(interval.Min)

	for i.Cmp(interval.Max) <= 0 {
		select {
		case <-ctx.Done():
			return
		default:
			work(data, keysize, decode, i, ch, cancel)
			i.Add(i, bigOne)
		}
	}
}
