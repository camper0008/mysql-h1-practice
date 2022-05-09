package fs

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadCSVHeader(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func ReadCSVEverythingElse(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // advance once to skip header
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}
