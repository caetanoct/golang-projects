package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//body, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(body))

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil	
}
