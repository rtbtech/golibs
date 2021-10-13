package load

import (
	"bufio"
	"os"
)

func ByLine(path string, parse func(line string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parse(scanner.Text())
	}

	return scanner.Err()
}

func ByLineMax(path string, maxSize int, parse func(line string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    buf := make([]byte, maxSize)
    scanner.Buffer(buf, cap(buf))

	for scanner.Scan() {
		parse(scanner.Text())
	}

	return scanner.Err()
}
