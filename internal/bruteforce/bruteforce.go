package bruteforce

import (
	"bufio"
	"os"
	"strings"
)

type Rate float32
const CHARSET_RATE Rate = 0.8
const DICTIONARY_RATE Rate = 0.5
// 62 characters in charset
// Complexity: 62^keysize
const KEY_CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Dictionary map[string]struct{}

func (dict Dictionary) Has(value string) bool {
	_, ok := dict[value]
	return ok
}

func DownloadDictionary() (Dictionary, error) {
	const WORDS_LIST_FILENAME = "google-10000-english.txt"

	file, err := os.Open(WORDS_LIST_FILENAME)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dict := make(Dictionary, 10000)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word == "" {
			continue
		}
		dict[word] = struct{}{}
	}

	return dict, scanner.Err()
}
