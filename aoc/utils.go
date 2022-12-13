package aoc

import (
	"io"
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

	body, err = io.ReadAll(res.Body)

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
