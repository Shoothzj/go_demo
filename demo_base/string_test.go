package demo_base

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestStringSub(t *testing.T) {
	str := "testing"
	u := str[0]
	fmt.Println(u)
}

func TestPrintSample(t *testing.T) {
	sample := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Println()
	fmt.Println("==================")
	fmt.Printf("%x\n", sample)
	fmt.Printf("% x\n", sample)
	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)
}

func TestPrintStr(t *testing.T) {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func TestAsRange(t *testing.T) {
	const nihongo = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func TestRangeLibrary(t *testing.T) {
	const nihongo = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98日本語"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
