package main

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main() {

	fileName := os.Args[1]

	filePath := fmt.Sprintf("tests/step2/%s",fileName)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	data, _ := io.ReadAll(file)
	// stats, _ := file.Stat()

	tokens := tokenize(&data, newJson())

	fmt.Printf("%d %d %d %d comma: %d newline: %d\n", tokens.openingBracket, tokens.closingBracket, tokens.quote, tokens.colon, tokens.comma, tokens.newLine)
	
	fmt.Printf("result: %d\n", parse(tokens, &data))
}

func parse(t *tokens, data *[]byte) int {
	if t.count() < 2 {
		return 1
	}

	if !strings.Contains(string((*data)[0]), "{") || !strings.Contains(string((*data)[len(*data)-1]), "}") {
		return 1
	}

	if t.comma > t.newLine {
		return 1
	}

	if t.comma > 0 && t.comma != t.newLine - 2 {
		return 1
	}

	return 0 // valid
}

func tokenize(data *[]byte, tokens *tokens) *tokens {
	for _, char := range *data {
		fmt.Printf("%s\n", string(char))
		switch char {
			case '{': // Use single quotes for rune literals
				tokens.incrementOpenB()
			case '}': // Use single quotes for rune literals
				tokens.incrementCloseB()
			case '"': // Use single quotes for rune literals
				tokens.incrementQ()
			case ':':
				tokens.incrementColon()
			case ',':
				tokens.incrementComma()
			case '\n':
				tokens.incrementNewline()
			}
	}
	return tokens
}

type tokens struct {
	openingBracket int
	closingBracket int
	quote int
	colon int
	comma int
	newLine int
}

func newJson() *tokens {
	return &tokens{0,0,0,0,0, 0}
}

func (t *tokens) count() int {
	return t.openingBracket + t.closingBracket
}

func (t *tokens) incrementOpenB() {
	t.openingBracket++
}
func (t *tokens) incrementCloseB() {
	t.closingBracket++
}
func (t *tokens) incrementQ() {
	t.quote++
}
func (t *tokens) incrementColon() {
	t.colon++
}
func (t *tokens) incrementComma() {
	t.comma++
}
func (t *tokens) incrementNewline() {
	t.newLine++
}

