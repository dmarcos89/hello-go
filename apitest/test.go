package main

import (
	"net/http"
	// "net/url"
	"io/ioutil"
	"fmt"
	"strings"
)

// keep first n lines
func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func main() {
	// We can use GET form to get result.
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("get:\n", string(body))
}
