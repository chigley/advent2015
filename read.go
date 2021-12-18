package advent2015

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadStrings(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("advent2015: opening %s: %w", name, err)
	}
	defer f.Close()

	return readStrings(f)
}

func readStrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2021: scanner: %w", err)
	}

	return lines, nil
}
