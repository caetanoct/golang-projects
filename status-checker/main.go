package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"https://google.com",
		"https://golang.org",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, link := range websites {
		go checkLink(link, c)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println("received this from channel <-c")
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	fmt.Println(link, err == nil)
	if err == nil {
		c <- link
	} else {
		c <- link
	}
}
