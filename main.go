package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy gocercies")
	todos.add("Buy Bread")
	fmt.Printf("%v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%v\n\n", todos)
}
