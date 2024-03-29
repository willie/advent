package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// SessionCookie returns a session cookie for AoC
func SessionCookie() string {
	value, exists := os.LookupEnv("AOC_SESSION_COOKIE")
	if !exists {
		log.Fatal("no AOC_SESSION_COOKIE defined")
	}
	return value
}

// Input caches and returns input data from a given URL
func Input(url string) (input []byte) {
	if strings.Contains(url, "http") {
		url += "/input"
	}

	_, filename := path.Split(url)
	if path.Ext(url) == "" {
		filename += ".txt"
	}

	// local?
	input, err := os.ReadFile(filename)
	if err == nil {
		// println("file:", filename)
		return
	}

	// remote?
	input = bodyFromURL(url, SessionCookie())

	// cache it.
	err = os.WriteFile(filename, input, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url, "saved")

	return
}

// String from url
func String(url string) string {
	return string(Input(url))
}

// Strings returns each line in input as a string array
func Strings(url string) (strings []string) {
	scanner := bufio.NewScanner(bytes.NewReader(Input(url)))
	// scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s := scanner.Text()
		strings = append(strings, s)
	}
	return
}

// LoadInts returns each line in input as a int array
func LoadInts(url string) (ints Ints) {
	for _, s := range Strings(url) {
		ints = append(ints, AtoI(s))
	}
	return
}

// StringsSplit returns each line in input separated by a delimiter as an [][]string
func StringsSplit(url string, delimiter string) (out [][]string) {
	for _, s := range Strings(url) {
		t := strings.Split(s, delimiter)
		out = append(out, t)
	}
	return
}

// LoadGrid returns the input as a Grid
func LoadGrid(url string) (grid Grid) { return NewGrid(Strings(url)) }

// Test prints output and compares to expected
func Test[T comparable](label string, result T, expected T) {
	extra := "PASS"

	if result != expected {
		extra = fmt.Sprint("FAIL, expected: ", expected)
	}

	fmt.Println(label+":\t", result, extra)
}

// Test64 prints output and compares to expected
func Test64(label string, result int64, expected int64) {
	extra := "PASS"

	if result != expected {
		extra = fmt.Sprint("FAIL, expected: ", expected)
	}

	fmt.Println(label+":\t", result, extra)
}

// Run prints output
func Run[T any](label string, result T) { fmt.Println(label+":\t", result) }

// TestX prints output and compares results to expected (results..., expected...)
func TestX[T comparable](label string, resultExpected ...T) {
	if len(resultExpected)%2 != 0 {
		log.Fatalln(len(resultExpected), "resultedExpected is results, expected")
	}

	half := len(resultExpected) / 2
	for i := 0; i < half; i++ {
		Test(fmt.Sprint(label, i+1), resultExpected[i], resultExpected[i+half])
	}
}

// RunX prints output
func RunX[T any](label string, results ...T) {
	for i, result := range results {
		Run(fmt.Sprint(label, i+1), result)
	}
}
