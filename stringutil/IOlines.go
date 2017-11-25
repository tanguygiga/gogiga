package stringutil

import (
	"bufio"
	"io"
	"os"
	"sort"
)

func readLines(file string) (lines []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		const delim = '\n'
		line, err := r.ReadString(delim)
		if err == nil || len(line) > 0 {
			if err != nil {
				line += string(delim)
			}
			lines = append(lines, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return lines, nil
}

func writeLines(file string, lines []string) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		if "\n" == line {
			continue
		}
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

// Sort sort the line of the file
func Sort(file string) (err error) {
	// tri des lignes
	lines, err := readLines(file)
	if err != nil {
		return err
	}
	sort.Strings(lines)
	err = writeLines(file, lines)
	if err != nil {
		return err
	}
	return nil
}
