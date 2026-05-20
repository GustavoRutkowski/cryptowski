package main

import (
	"errors"
	"flag"
	"os"
	"fmt"
	"github.com/GustavoRutkowsik/cryptowski/internal/cryptowski"
	"github.com/GustavoRutkowsik/cryptowski/internal/cryptowski/v1"
)

// Version Flag: --v1 | --v2 | --v3 | --v4 | ...
// Default: Latest Version

// encode foobar.txt -o foobar.enc
// decode foobar.enc -o foobar.dec

const INVALID_COMMAND_ERR = "Invalid Command! Run: cryptowski --help"

func parseVersionFlags() cryptowski.ICrypto {
	isV1 := flag.Bool("v1", false, "Use version 1")
	// isV2 := flag.Bool("v2", false, "Use version 2")
	// ...
	flag.Parse()
	
	// Se passar muitas flags, pega sempre a versão mais recente informada
	switch {
	case *isV1:
		return v1.Crypto{}
	default:
		latest := v1.Crypto{}
		return latest
	}
}

func resolveAction(crypto cryptowski.ICrypto, command string) (cryptowski.CryptoAlg, error) {
	switch command {
	case "encode":
		fmt.Println("Encoding...")
		return crypto.Encode, nil
	case "decode":
		fmt.Println("Decoding...")
		return crypto.Decode, nil
	default:
		return nil, errors.New(INVALID_COMMAND_ERR)
	}
}

func parseCommandFlags(args []string) (string, string, error) {
	fs := flag.NewFlagSet("command", flag.ContinueOnError)
	output := fs.String("o", "", "Filename for output file. e.g.: encoded.enc")
	key := fs.String("k", "", "Key for encoding/decoding")

	fs.Usage = func() {
		fmt.Println(`
Usage:
	cryptowski encode <input-file> -o <output-file> -k <key>

Flags:
		`)
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		return "", "", err
	}

	if *output == "" {
		return "", "", errors.New("Error: The flag \"-o\" is required")
	}

	if *key == "" {
		return "", "", errors.New("Error: The flag \"-k\" is required")
	}
	
	return *output, *key, nil
}

func main() {
	flag.Usage = func() {
		fmt.Println(`
Usage:
	cryptowski [--v1] <command> <input-file> [flags]
Commands:
	encode    Encode a file
	decode    Decode a file

Global Flags:
		`)

		flag.PrintDefaults()
	}

	crypto := parseVersionFlags() // --v1 | --v2 | --v3 | --v4 | ...
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println(INVALID_COMMAND_ERR)
		return
	}
	
	command := args[0]
	action, err := resolveAction(crypto, command)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	input := args[1]
	output, key, err := parseCommandFlags(args[2:])
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Call Encode/Decode method
	computedBytes := action(file, key)

	// Owner: read & write | Group: read | Others: read
	const PERMS os.FileMode = 0644

	err = os.WriteFile(output, computedBytes, PERMS)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s completed successfully!\n", command)
}
