package main

import (
	"bufio"
	"fmt"
	"gogiga/controller"
	"gogiga/stringutil"
	"os"
	"os/user"
	"strings"
)

func main() {

	tc := controller.NewTodoController("txt")
	tc.GetAll()
	tc.Get(5)

	u, err := user.Current()
	if err != nil {
		fmt.Print(err)
	}
	home := u.HomeDir

	var (
		// implying the todo file is in current home directory
		todo = home + "/todo.txt"
		//Port = ":8088"
	)

	write, err := os.OpenFile(todo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print(err)
	}
	defer write.Close()

	stdin := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(": ")
		input, _ := stdin.ReadString('\n')
		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)

		if "q" == input {
			os.Exit(0)
		} else {
			_, err = write.WriteString("\n" + input)
			if err != nil {
				fmt.Print(err)
			}
			err = stringutil.Sort(todo)
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
