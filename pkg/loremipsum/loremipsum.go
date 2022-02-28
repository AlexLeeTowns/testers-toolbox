package loremipsum

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"strings"
)

func ReadLorem(fs fs.FS, filename, method string, charcount int) (string, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result string
	switch method {
	case "word":
		words, err := scanLorem(bufio.ScanWords, charcount, file)
		if err != nil {
			return "", err
		}
		result = strings.Join(words, " ")

	case "char":
		characters, err := scanLorem(bufio.ScanBytes, charcount, file)
		if err != nil {
			return "", err
		}
		result = strings.Join(characters, "")

	case "paragraph":
		paragraphs, err := scanLorem(bufio.ScanLines, charcount, file)
		if err != nil {
			return "", err
		}
		result = strings.Join(paragraphs, "\n")

	default:
		return "", errors.New("method not allowed")
	}

	return result, nil
}

func scanLorem(f bufio.SplitFunc, count int, r io.Reader) (result []string, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(f)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
		if len(res) == count {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return make([]string, 0), err
	}

	return res, nil
}
