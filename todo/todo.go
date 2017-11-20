package main

import (
	"bufio"
	"fmt"
	"gogiga/json"
	"gogiga/stringutil"
	"os"
	"sort"
	"strings"
)

type config struct {
	Path string `json:"path"`
	Port int    `json:"port"`
}

func main() {

	// parsing du fichier de config json dans la struct
	// pour avoir le chemin du fichier
	wd, _ := os.Getwd()
	path := wd + "/config.json"
	var c config
	json.Parser(path, &c)
	path = c.Path

	write, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print(err)
	}
	defer write.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		// lecture du fichier dans la sortie standard
		read, err := os.Open(path)
		if err != nil {
			fmt.Print(err)
		}
		scanner := bufio.NewScanner(read)
		i := 0
		for scanner.Scan() {
			i++
			fmt.Println(i, scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			fmt.Print(err)
		}

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if "q" == text {
			os.Exit(0)
		} else {
			_, err = write.WriteString("\n" + text)
			if err != nil {
				fmt.Print(err)
			}

			// tentative de tri des lignes
			lines, err := stringutil.ReadLines(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			sort.Strings(lines)
			err = stringutil.WriteLines(path, lines)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

	}

}
