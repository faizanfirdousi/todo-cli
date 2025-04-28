package main

import "fmt"

func main() {
	mytodos := Todos{}

	mytodos.add("buy groceries")
	mytodos.add("do homework")
	fmt.Println(mytodos)
	mytodos.delete(1)
	fmt.Println()
	mytodos.toggle(0)
	fmt.Println(mytodos)
}
