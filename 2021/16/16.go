package main

import (
	"github.com/willie/advent/aoc"
)

const (
	sumType = iota
	productType
	minimumType
	maximumType
	literalType
	greaterThanType
	lessThanType
	equalToType
)

type binarystring struct {
	bin string
	pc  int
}

func BinaryString(in string) *binarystring            { return &binarystring{bin: in} }
func BinaryStringFromHex(in string) *binarystring     { return BinaryString(aoc.HexToBin(in)) }
func (b *binarystring) Header() (version, typeID int) { return b.Int(3), b.Int(3) }
func (b *binarystring) Int(bits int) (out int)        { return int(b.Int64(bits)) }
func (b *binarystring) Int64(bits int) (out int64)    { return aoc.BinToDec(b.Bits(bits)) }
func (b *binarystring) Literal() (out int)            { return int(b.Literal64()) }

func (b *binarystring) Bits(bits int) (out string) {
	out = b.bin[b.pc : b.pc+bits]
	b.pc += bits
	return
}

func (b *binarystring) Literal64() (out int64) {
	v := ""

	for {
		s := b.Bits(5)
		v += s[1:]

		if s[:1] == "0" {
			break
		}
	}

	return aoc.BinToDec(v)
}

func (b *binarystring) Value() (out int) {
	_, typeID := b.Header()

	switch typeID {
	case literalType:
		literal := b.Literal()
		out = literal

	default:
		lengthTypeID := b.Int(1)
		switch lengthTypeID {
		case 0: // total length in bits of the sub-packets
			length := b.Int(15)
			start := b.pc

			for b.pc < start+length {
				out += b.Value()
			}

		case 1: // number of sub-packets immediately contained
			subpackets := b.Int(11)
			for i := 0; i < subpackets; i++ {
				out += b.Value()
			}
		}
	}

	return
}

func (b *binarystring) Parse() (p *packet) {
	p = &packet{version: b.Int(3), typeID: b.Int(3)}

	switch p.typeID {
	case literalType:
		p.literal = b.Literal()

	default:
		lengthTypeID := b.Int(1)
		switch lengthTypeID {
		case 0: // total length in bits of the sub-packets
			length := b.Int(15)
			start := b.pc

			for b.pc < start+length {
				p.subpackets = append(p.subpackets, b.Parse())
			}

		case 1: // number of sub-packets immediately contained
			subpackets := b.Int(11)
			for i := 0; i < subpackets; i++ {
				p.subpackets = append(p.subpackets, b.Parse())
			}
		}
	}

	return
}

type packet struct {
	version    int
	typeID     int
	literal    int
	subpackets []*packet
}

func (p *packet) Values() (values []int) {
	for _, s := range p.subpackets {
		values = append(values, s.Value())
	}
	return
}

func (p *packet) Value() (out int) {
	switch p.typeID {
	case literalType:
		return p.literal

	case sumType:
		out = aoc.Sum(p.Values()...)

	case productType:
		out = aoc.Product(p.Values()...)

	case minimumType:
		out = aoc.Min(p.Values()...)

	case maximumType:
		out = aoc.Max(p.Values()...)

	case lessThanType:
		if p.subpackets[0].Value() < p.subpackets[1].Value() {
			out = 1
		}

	case greaterThanType:
		if p.subpackets[0].Value() > p.subpackets[1].Value() {
			out = 1
		}

	case equalToType:
		if p.subpackets[0].Value() == p.subpackets[1].Value() {
			out = 1
		}
	}

	return
}

func (p *packet) VersionTotal() (out int) {
	out = p.version
	for _, sub := range p.subpackets {
		out += sub.VersionTotal()
	}
	return
}

func part1(in string) (result int) {
	return BinaryStringFromHex(in).Parse().VersionTotal()
}

func part2(in string) (result int) {
	return BinaryStringFromHex(in).Parse().Value()
}

const day = "https://adventofcode.com/2021/day/16"

func main() {
	println(day)

	aoc.Test("test1", part1("D2FE28"), 6)
	aoc.Test("test1", part1("38006F45291200"), 9)
	aoc.Test("test1", part1("EE00D40C823060"), 14)
	aoc.Test("test1", part1("8A004A801A8002F478"), 16)
	aoc.Test("test1", part1("620080001611562C8802118E34"), 12)
	aoc.Test("test1", part1("C0015000016115A2E0802F182340"), 23)
	aoc.Test("test1", part1("A0016C880162017C3686B18A3D4780"), 31)

	aoc.Test("test2", part2("C200B40A82"), 3)
	aoc.Test("test2", part2("04005AC33890"), 54)
	aoc.Test("test2", part2("880086C3E88112"), 7)
	aoc.Test("test2", part2("CE00C43D881120"), 9)
	aoc.Test("test2", part2("D8005AC2A8F0"), 1)
	aoc.Test("test2", part2("F600BC2D8F"), 0)
	aoc.Test("test2", part2("9C005AC2F8F0"), 0)
	aoc.Test("test2", part2("9C0141080250320F1802104A08"), 1)

	println("-------")

	aoc.Run("part1", part1(aoc.String(day)))
	aoc.Run("part2", part2(aoc.String(day)))
}
