package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) //go => creates a new thread go routine
		// fmt.Println(<-c)
	}

	// for {  //infinite loop
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }

	for l := range c {  //infinite loop
		// fmt.Println(<-c)
		go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	} else {
		fmt.Println(link, "is up!")
		c <- link
	}
}
