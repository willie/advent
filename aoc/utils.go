package aoc

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func bodyFromURL(url string, sessionCookie string) (body []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})
	req.Header.Add("User-Agent", "willie@pobox.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return
}

// AtoI converts string to int
func AtoI(a string) (i int) {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// BtoI converts bool to int
func BtoI(b bool) int {
	if b {
		return 1
	}
	return 0
}

// byte to int
func byteToInt(c byte) (val uint8) {
	switch c {
	case '0':
		val = 0
	case '1':
		val = 1
	case '2':
		val = 2
	case '3':
		val = 3
	case '4':
		val = 4
	case '5':
		val = 5
	case '6':
		val = 6
	case '7':
		val = 7
	case '8':
		val = 8
	case '9':
		val = 9

	default:
		log.Fatal("what?")
	}
	return
}
