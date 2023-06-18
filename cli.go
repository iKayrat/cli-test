package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <string> <url>")
		os.Exit(1)
	}

	args := os.Args[1:]

	resp, err := http.Post(args[1], "text/plain", strings.NewReader(args[0]))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result:", string(body))
}
