package main

import (
	"fmt"
	"net/http"
)

func checkStatus(link string, c chan string) {
	res, _ := http.Get(link)
	if res.Status == "200 OK"{
		c <- link+ " is ok."
		return
	}
	c <-link+ " is down."
}

func main()  {
	urls := []string{
		"http://bigzoo.me",
		"http://google.com",
		"http://vof.andela.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)

	for _, link := range urls{
		go checkStatus(link, c)
	}
	for i := 0; i <len(urls); i++{
		fmt.Println(<-c)
	}
}