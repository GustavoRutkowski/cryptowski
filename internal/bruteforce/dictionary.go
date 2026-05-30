package bruteforce

import (
	"bufio"
	"embed"
	"strings"
)

//go:embed google-10000-english.txt
var fsys embed.FS

const WORDS_LIST_FILENAME = "google-10000-english.txt"

type dictionary map[string]struct{}

func (dict dictionary) Has(value string) bool {
	_, ok := dict[value]
	return ok
}

func DownloadDictionary() (dictionary, error) {
	file, err := fsys.Open(WORDS_LIST_FILENAME)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dict := make(dictionary, 10000)
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
