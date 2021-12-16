package aoc

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func HexToBin(in string) (out string) {
	b, err := hex.DecodeString(in)
	if err != nil {
		log.Fatalln(err)
	}

	for _, bb := range b {
		out += fmt.Sprintf("%08b", bb)
	}
	return
}

func BinToDec(in string) (out int64) {
	out, _ = strconv.ParseInt(in, 2, 64)
	return
}
