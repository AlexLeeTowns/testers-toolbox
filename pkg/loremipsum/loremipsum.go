package loremipsum

import (
	"bufio"
	"io/fs"
	"strings"
)

func ReadLoremByCharacterCount(fs fs.FS, filename string, charcount int) (result string, err error) {
	file, err := fs.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	b := make([]byte, charcount)
	fml, err := file.Read(b)

	if err != nil {
		return "", err
	}

	return string(b[:fml]), nil

}

func ReadLoremByWordCount(fs fs.FS, filename string, wordcount int) (result string, err error) {
	file, err := fs.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
		if len(res) == wordcount {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(res, " "), nil
}

func ReadLoremByParagraph(fs fs.FS, filename string, paragraphCount int) (string, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
		if len(res) == paragraphCount {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(res, "\n"), nil
}
