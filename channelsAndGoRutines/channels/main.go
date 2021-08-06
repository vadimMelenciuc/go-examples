// status checker

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting website status checker")
	links := []string{
		"http://google.com",
		"http://faceboook.com",
		"http://golang.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkStatus(<-c, c)
	// }

	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	go checkStatus(l, c)
	// }

	for l := range c {
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkStatus(li, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
