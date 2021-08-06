package main

import "fmt"

func main() {
	fmt.Println("--------------------------------")

	// key is a string and value is a string
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	// }

	// starts a map with zero values
	// var colors map[string]string

	// starts a map with zero values
	colors := make(map[string]string)

	// colors := map[string]string

	colors["red"] = "#ff0000"

	fmt.Printf("%+v\n", colors)
	fmt.Println(colors)

	delete(colors, "red")

	fmt.Println(colors)

	fmt.Println("--------------------------------")
	fmt.Println("Map iteration")

	colors["white"] = "#FFF"
	colors["black"] = "#000"
	colors["grey"] = "#777"

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value)
	}
}
