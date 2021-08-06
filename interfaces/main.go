package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreetings() string
	// getBotVersion() int
}

func main() {
	fmt.Println("----------------------------")
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)

}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string {
	return "hello"
}

func (spanishBot) getGreetings() string {
	return "hola"
}
