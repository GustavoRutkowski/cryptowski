package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"github.com/GustavoRutkowski/cryptowski/internal/bruteforce"
	"github.com/GustavoRutkowski/cryptowski/internal/cryptowski"
	"github.com/GustavoRutkowski/cryptowski/internal/utils"
	"github.com/GustavoRutkowski/cryptowski/internal/utils/cli"
)

const HELP = `
Cryptowski Bruteforce - Recover encryption keys by exhaustive search.

Usage:
    bruteforce [--v1] <input-file> [flags]

Flags:
    --size N
        Try only keys with length N

    --minsize N
        Minimum key length

    --maxsize N
        Maximum key length

Global Flags:
    --v1
        Use algorithm version 1

Examples:
    bruteforce secret.enc --size 4

        Try all keys with length 4

    bruteforce secret.enc --maxsize 6

        Try keys from length 1 to 6

    bruteforce secret.enc --minsize 3 --maxsize 8

        Try keys from length 3 to 8
`

const INVALID_COMMAND_ERR = "Invalid arguments. Run: bruteforce --help"

func parseInterval(args []string) (utils.Interval[uint8], error) {
	fs := flag.NewFlagSet("range", flag.ContinueOnError)
	minsize := fs.Uint("minsize", uint(0), "The min size to apply brute force")
	size := fs.Uint("size", uint(0), "A fixed size to apply brute force")
	maxsize := fs.Uint("maxsize", uint(0), "The max size to apply brute force")

	if err := fs.Parse(args); err != nil {
		return utils.Interval[uint8]{}, err
	}

	// ┌───────┬─────────┬─────────┐
	// │ SIZE  │ MINSIZE │ MAXSIZE │
	// ├───────┼─────────┼─────────┤
	// │ FALSE │ FALSE   │ FALSE   │ ➜ ERROR
	// │ FALSE │ FALSE   │ TRUE    │
	// │ FALSE │ TRUE    │ FALSE   │ ➜ ERROR
	// │ FALSE │ TRUE    │ TRUE    │
	// │ TRUE  │ FALSE   │ FALSE   │
	// │ TRUE  │ FALSE   │ TRUE    │ ➜ ERROR
	// │ TRUE  │ TRUE    │ FALSE   │ ➜ ERROR
	// │ TRUE  │ TRUE    │ TRUE    │ ➜ ERROR
	// └───────┴─────────┴─────────┘

	// Nothing
	if *size == 0 && *minsize == 0 && *maxsize == 0 {
		return utils.Interval[uint8]{}, errors.New(INVALID_COMMAND_ERR)
	}

	// Size only
	if *minsize == 0 && *maxsize == 0 {
		return utils.Interval[uint8]{Min: uint8(*size), Max: uint8(*size)}, nil 
	}
	
	// Interval Cases:

	if *size > 0 {
		return utils.Interval[uint8]{}, errors.New(INVALID_COMMAND_ERR)
	}

	if *maxsize == 0 {
		return utils.Interval[uint8]{}, errors.New(INVALID_COMMAND_ERR)
	}

	if *minsize == 0 {
		return utils.Interval[uint8]{Min: 1, Max: uint8(*maxsize)}, nil
	}

	return utils.Interval[uint8]{Min: uint8(*minsize), Max: uint8(*maxsize)}, nil
}

func handleBruteforce(file []byte, interval utils.Interval[uint8], decode cryptowski.DecodeAlg) (string, []byte, error) {
	if interval.Min != interval.Max {
		return bruteforce.Ranged(file, interval, decode)
	}
	return bruteforce.Fixed(file, interval.Min, decode)
}

func main() {
	flag.Usage = func() {
		fmt.Print(HELP)
	}

	crypto := cli.ParseVersionFlags() // --v1 | --v2 | --v3 | --v4 | ...
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println(INVALID_COMMAND_ERR)
		return
	}

	filename := args[0]
	interval, err := parseInterval(args[1:])

	if err != nil {
		fmt.Println(err) // Revisar isso aqui
		return
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Running...")
	fmt.Println()

	key, decoded, err := handleBruteforce(file, interval, crypto.Decode)
	if err != nil {
		fmt.Println(err)
		return
	}

	if key == "" {
		fmt.Println("No key founded with the specified constraints")
		return
	}

	fmt.Println("KEY:")
	fmt.Printf("%s\n\n", key)
	fmt.Println("DECODED:")
	fmt.Printf("%s\n", string(decoded))
}
