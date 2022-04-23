package aocutil

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// MustFileToOneString loads all lines in a text file into one string.
func MustFileToOneString(fileName string) (string, error) {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		log.Panic(err)
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Panic(err)
	}
	s := string(bytes)
	// Trim any leading/trailing whitespace.
	return strings.TrimSpace(s), nil
}

// MustFileToStrings reads a text file into a string slice.
func MustFileToStrings(fileName string) []string {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		log.Panic(err)
	}
	file, err := os.Open(absPath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}
	return lines
}

// MustFileToInts reads a text file into an int slice.
func MustFileToInts(fileName string) []int {
	is := []int{}
	ss := MustFileToStrings(fileName)
	for _, s := range ss {
		is = append(is, MustAtoi(s))
	}
	return is
}

// MustAtoi converts a string to an integer.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Panic(err)
	}
	return i
}
