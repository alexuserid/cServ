package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("s")

	s, err := verify(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%d\n", stones(s))
}

func verify(param string) (string, error) {
	var s string

	if n := len(param); n == 0 || n > 50 {
		return "", errors.New("Incorrect parameter. Wrong size.")
	}

	s = strings.ToLower(param)
	for i := range s {
		if s[i] != 'b' && s[i] != 'g' && s[i] != 'r' {
			return "", errors.New("Incorrect parameter. Use only 'r', 'g', 'b'.")
		}
	}
	return s, nil
}

func stones(s string) int {
	var c int
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {c++}
	}
	return c
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
