package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Solution struct {
}

func (s *Solution) Encode(strs []string) string {
	strBuilder := strings.Builder{}

	for _, str := range strs {
		strLen := len(str)
		strBuilder.WriteString(fmt.Sprintf("%d-%s", strLen, str))
	}

	return strBuilder.String()
}

func (s *Solution) Decode(encoded string) []string {
	strs := make([]string, 0)
	strReader := strings.NewReader(encoded)
	nextStrBuilder := strings.Builder{}

	for {
		r, _, err := strReader.ReadRune()

		if err != nil {
			break // EOF
		}

		if r == '-' {
			nextStrLen, _ := strconv.Atoi(nextStrBuilder.String())
			buf := make([]byte, nextStrLen)
			_, err = io.ReadAtLeast(strReader, buf, nextStrLen)
			nextStrBuilder.Reset()
			nextStrBuilder.Write(buf)
			strs = append(strs, nextStrBuilder.String())
			nextStrBuilder.Reset()
		} else {
			nextStrBuilder.WriteRune(r)
		}
	}

	return strs
}

func main() {
	input := []string{"hello", "world"}
	solution := Solution{}
	encoded := solution.Encode(input)
	decoded := solution.Decode(encoded)
	fmt.Printf("encode(%v) => %s\n", input, encoded)
	fmt.Printf("decode(%s) => %v\n", encoded, decoded)

	input = []string{""}
	encoded = solution.Encode(input)
	decoded = solution.Decode(encoded)
	fmt.Printf("encode(%v) => %s\n", input, encoded)
	fmt.Printf("decode(%s) => %v\n", encoded, decoded)

	input = []string{" "}
	encoded = solution.Encode(input)
	decoded = solution.Decode(encoded)
	fmt.Printf("encode(%v) => %s\n", input, encoded)
	fmt.Printf("decode(%s) => %v\n", encoded, decoded)

	input = []string{"  "}
	encoded = solution.Encode(input)
	decoded = solution.Decode(encoded)
	fmt.Printf("encode(%v) => %s\n", input, encoded)
	fmt.Printf("decode(%s) => %v\n", encoded, decoded)

	input = []string{"bela", "undok", "yes"}
	encoded = solution.Encode(input)
	decoded = solution.Decode(encoded)
	fmt.Printf("encode(%v) => %s\n", input, encoded)
	fmt.Printf("decode(%s) => %v\n", encoded, decoded)
}
