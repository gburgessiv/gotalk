package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	fmt.Println("This is an example of how cool Go's libs are!")
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("OH NOOOOO", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("What why", err)
		return
	}
	fmt.Println("Google gave us:", string(body))
}
