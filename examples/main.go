package main

import "fmt"

//go:generate go run tools.go
func main() {
	fmt.Println(ItemSorting{Key: "name", Order: "asc"})
}
