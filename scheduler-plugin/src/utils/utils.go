package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func LoadTupleFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, os.ErrNotExist
		} else {
			log.Fatal(err)
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// read first line of file
	line := lines[0]

	// remove tuple brackets
	line = line[1 : len(line)-1]
	lineContent := strings.Split(line, ",")

	return lineContent, nil
}

func MaxOf(vars ...int64) int64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func MinOf(vars ...int64) int64 {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
