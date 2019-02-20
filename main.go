package main

import (
	"fmt"
	"net/http"
)

func checkStatus(link string) bool {
	res, err := http.Get(link)
	if err != nil {
		return false
	}
	if res.Status == "200 OK"{
		return true
	}
	return false
}

func main()  {
	urls := []string{
		"http://vof-sandbox.andela.com",
		"http://mail.google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range urls{
		if checkStatus(link){
			fmt.Println(link, "is ok.")
		}else {
			fmt.Println(link, "is down.")
		}
	}
}