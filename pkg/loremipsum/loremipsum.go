package loremipsum

import (
	_ "embed"
	"strings"
)

//go:embed lorem.txt
var f string

func GetLorem(method string, charcount int) string {
	var result string
	switch method {
	case "word":
		words := strings.Split(f, " ")[:charcount]
		result = strings.Join(words, " ")
	case "char":
		characters := strings.Split(f, "")[:charcount]
		result = strings.Join(characters, "")
	case "paragraph":
		paragraphs := strings.Split(f, "\r\n\r")[:charcount]
		result = strings.Join(paragraphs, "\r\n\r")
	}

	return result
}
