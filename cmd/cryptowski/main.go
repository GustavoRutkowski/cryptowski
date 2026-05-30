package main

import (
	"errors"
	"flag"
	"os"
	"fmt"
	"github.com/GustavoRutkowski/cryptowski/internal/cryptowski"
	"github.com/GustavoRutkowski/cryptowski/internal/utils/cli"
)

const HELP = `
Cryptowski - File encryption utility

Usage:
    cryptowski [--v1] <command> <input-file> -o <output-file> -k <key>

Commands:
    encode    Encrypt a file
    decode    Decrypt a file

Flags:
    -o string
        Output filename

    -k string
        Encryption/decryption key

Global Flags:
    --v1
        Use algorithm version 1

Examples:
    cryptowski encode notes.txt -o notes.enc -k secret
    cryptowski decode notes.enc -o notes.txt -k secret
`

const INVALID_COMMAND_ERR = "Invalid Command! Run: cryptowski --help"

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
		fmt.Print(HELP)
	}

	crypto := cli.ParseVersionFlags() // --v1 | --v2 | --v3 | --v4 | ...
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
