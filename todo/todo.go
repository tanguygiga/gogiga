package main

import (
	"fmt"
	"gogiga/controller"
)

func main() {

	tc := controller.NewTodoController("txt")
	fmt.Println("tc.GetAll()")
	tc.GetAll()
	fmt.Println("tc.Get(5)")
	tc.Get(5)
	fmt.Println("tc.Delete(5)")
	tc.Delete(5)
	fmt.Println(`tc.Set(5, "tc.Set(5)")`)
	tc.Set(5, "tc.Set(5)")
	fmt.Println(`tc.Set(0, "tc.Set(0)")`)
	tc.Set(0, "tc.Set(0)")
	fmt.Println(`tc.Set(22, "tc.Set(22)")`)
	tc.Set(22, "tc.Set(22)")
}
