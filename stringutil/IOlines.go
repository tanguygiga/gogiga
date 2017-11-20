package stringutil

import (
    "bufio"
    "os"
    "io"
)

func ReadLines(file string) (lines []string, err error) {
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

func WriteLines(file string, lines []string) (err error) {
    f, err := os.Create(file)
    if err != nil {
        return err
    }
    defer f.Close()
    w := bufio.NewWriter(f)
    defer w.Flush()
    for _, line := range lines {
        _, err := w.WriteString(line)
        if err != nil {
            return err
        }
    }
    return nil
}
