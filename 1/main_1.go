package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var (
	dedicate              = 255
	octal                 = 0377
	hexadecimal           = 0xFF
	pi          float32   = 3.1415926
	hi                    = "Hello World!"
	isActive              = true
	z           complex64 = 3 + 4i
)

func main() {

	fmt.Printf("%T\n", dedicate)
	fmt.Printf("%T\n", octal)
	fmt.Printf("%T\n", hexadecimal)
	fmt.Printf("%T\n", pi)
	fmt.Printf("%T\n", hi)
	fmt.Printf("%T\n", isActive)
	fmt.Printf("%T\n", z)

	const salt = "go-2024"
	runeSlice := []rune(AnyToString(dedicate, octal, hexadecimal, pi, hi, isActive, z))

	salted := InsertSaltMiddle(runeSlice, salt)
	hash := HashSHA256WithSalt(runeSlice, salt)

	fmt.Printf("original:  %s\n", string(runeSlice))
	fmt.Printf("salted:    %s\n", string(salted))
	fmt.Printf("sha256:    %x\n", hash)
}

func AnyToString(v ...any) string {
	var b strings.Builder

	for _, i := range v {
		fmt.Fprintf(&b, "%v", i)
	}
	return b.String()
}

func InsertSaltMiddle(runes []rune, salt string) []rune {
	mid := len(runes) / 2
	saltRunes := []rune(salt)

	salted := make([]rune, 0, len(runes)+len(salt))
	salted = append(salted, runes[:mid]...)
	salted = append(salted, saltRunes...)
	salted = append(salted, runes[mid:]...)
	return salted
}

func HashSHA256WithSalt(runes []rune, salt string) string {
	salted := InsertSaltMiddle(runes, salt)
	data := []byte(string(salted))
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
