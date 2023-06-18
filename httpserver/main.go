package main

import (
	"fmt"
	"io"

	"net/http"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			str, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}

			s := findSubstring(string(str))

			// response body
			fmt.Fprintln(w, s)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}

	http.HandleFunc("/api/substring", handler)

	fmt.Println("Server listening on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

}

func findSubstring(s string) string {

	if len(s) <= 1 {
		return s
	}

	// initial variables
	lastIndex := make(map[byte]int)

	start := 0
	maxLength := 0

	currentStart := 0
	currentLength := 0

	for i := 0; i < len(s); i++ {
		// current character
		ch := s[i]

		// check if ch exists and its last index is in the current substring
		if lastIdx, ok := lastIndex[ch]; ok && lastIdx >= currentStart {
			// update current start
			currentStart = lastIdx + 1
			currentLength = i - lastIdx
		} else {
			// else increment length
			currentLength++
		}

		// update last index
		lastIndex[ch] = i

		if currentLength > maxLength {
			start = currentStart
			maxLength = currentLength
		}
	}

	substr := s[start : start+maxLength]

	return substr
}
