package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// create a new thread and run checkLink in it
		go checkLink(link, c)
	}

	// fmt.Println(<-c) // this is a blocking call, Main routine is put to sleep until message is received

	// for i := 0; i < len(links); i++ {

	// new coolness: using range on a channel, loop executes for every time channel emits a message
	for l := range c {
		// fmt.Println(<-c)
		// go checkLink(l, c)
		// this is a function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yup it's up"
	c <- link
}
