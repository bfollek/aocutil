package aocutil

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// MustBitSliceToInt converts `bits`,  an int slice of 0's and 1's, to an integer
// value. It treats each `bits` element as a bit, e.g. [0,1,1,0] => 6. It panics
// if any `bits` element is neither 0 nor 1.
func MustBitSliceToInt(bits []int) int {
	sum := float64(0)
	j := float64(0)
	for i := len(bits) - 1; i >= 0; i-- {
		switch bits[i] {
		case 0:
		case 1:
			sum += math.Pow(2, j)
		default:
			panic(fmt.Sprintf("Found %d at index %d in %v. Should be 0 or 1\n", bits[i], i, bits))
		}
		j++
	}
	return int(sum)
}

// MustFileToOneString loads all lines in a text file into one string and panics on error.
func MustFileToOneString(fileName string) (string, error) {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		log.Panic(err)
	}
	bytes, err := os.ReadFile(absPath)
	if err != nil {
		log.Panic(err)
	}
	s := string(bytes)
	// Trim any leading/trailing whitespace.
	return strings.TrimSpace(s), nil
}

// MustFileToStrings reads a text file into a string slice and panics on error.
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
	/*
	 * https://pkg.go.dev/bufio#ScanLines
	 *
	 * "NewScanner returns a new Scanner to read from r. The split function defaults to ScanLines."
	 *
	 * "ScanLines is a split function for a Scanner that returns each line of text, stripped of any
	 * trailing end-of-line marker."
	 */
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

// MustAtoi wraps `strconv.Atoi` and panics on error.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Panic(err)
	}
	return i
}

// MustAtoi wraps `strconv.ParseInt` and panics on error.
func MustParseInt(s string, base int, bitSize int) int64 {
	i, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func SortString(s string) string {
	// https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
	r := []rune(s)
	sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
	return string(r)
}
