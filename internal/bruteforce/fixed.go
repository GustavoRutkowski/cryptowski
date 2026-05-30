package bruteforce

import (
	"context"
	"math/big"
	"runtime"
	"sync"
	"github.com/GustavoRutkowski/cryptowski/internal/cryptowski"
	"github.com/GustavoRutkowski/cryptowski/internal/utils"
)

var (
	bigOne  = big.NewInt(1)
	bigBase = big.NewInt(int64(KEY_CHARSET_LEN))
)

type result struct {
	key     string
	decoded []byte
	err     error
}

func Fixed(data []byte, keysize uint8, decode cryptowski.DecodeAlg) (string, []byte, error) {
	cores := runtime.NumCPU()
	bigCores := big.NewInt(int64(cores))

	// KEY_CHARSET_LEN^keysize --> 62^keysize
	combinations := new(big.Int).Exp(
		big.NewInt(int64(KEY_CHARSET_LEN)),
		big.NewInt(int64(keysize)),
		nil,
	)
	// Divide the keyspace into equal parts for each worker
	keysPerWorker := new(big.Int).Div(
		combinations,
		bigCores,
	)

	var wg sync.WaitGroup
	ch := make(chan result, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range cores {
		iBig := big.NewInt(int64(i))

		// min = i * keysPerWorker
		min := new(big.Int).Mul(
			iBig,
			keysPerWorker,
		)

		// max = (i + 1) * keysPerWorker - 1
		max := new(big.Int).Mul(
			new(big.Int).Add(iBig, bigOne),
			keysPerWorker,
		)

		max.Sub(max, bigOne)

		// Add the remaining keys for the last worker
		if i == cores-1 {
			rest := new(big.Int).Mod(
				combinations,
				bigCores,
			)
			max.Add(max, rest)
		}

		interval := utils.Interval[*big.Int]{Min: min, Max: max}
		wg.Add(1)

		go worker(data, keysize, decode, interval, ch, ctx, cancel, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	res, ok := <-ch

	if !ok {
		return "", nil, ErrKeyNotFound
	}

	return res.key, res.decoded, res.err
}
