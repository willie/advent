package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

// SessionCookie returns a session cookie for AoC
func SessionCookie() string {
	return "53616c7465645f5f6b02c508a13810bbd6286c82b3c85792aec35e78df61c9975af13d1164f104652d177e51c7d9236c"
}

// Input caches and returns input data from a given URL
func Input(url string) (input []byte) {
	if strings.Contains(url, "http") {
		url += "/input"
	}

	_, filename := path.Split(url)
	filename += ".txt"

	// local?
	input, err := ioutil.ReadFile(filename)
	if err == nil {
		// println("file:", filename)
		return
	}

	// remote?
	println("url:", url)
	input = bodyFromURL(url, SessionCookie())

	// cache it.
	println("saving:", filename)
	err = ioutil.WriteFile(filename, input, 0644)
	if err != nil {
		log.Fatal(err)
	}
	println()

	return
}

// String from url
func String(url string) string {
	return string(Input(url))
}

// Strings returns each line in input as a string array
func Strings(url string) (strings []string) {
	scanner := bufio.NewScanner(bytes.NewReader(Input(url)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s := scanner.Text()
		strings = append(strings, s)
	}
	return
}

// Ints rreturns each line in input as a int array
func Ints(url string) (ints []int) {
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

// Test prints output and compares to expected
func Test(label string, result int, expected int) {
	var extra string

	if result != expected {
		extra = fmt.Sprint("FAIL, expected: ", expected)
	} else {
		extra = "PASS"
	}

	fmt.Println(label, ":", result, extra)
}

// Run prints output
func Run(label string, result int) { fmt.Println(label, ":", result) }
