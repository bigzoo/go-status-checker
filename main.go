package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkStatus(link string, c chan string) {
	res, err := http.Get(link)
	if err != nil || res.Status != "200 OK"{
		fmt.Println(link, "is down.")
	}else {
		fmt.Println(link, "is ok.")
	}
	c <- link
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
	for l := range c {
		go func(link string){
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}