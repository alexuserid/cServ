package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var (
	incorErr error
)

func handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("s")

	s, err := verify(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(strconv.FormatInt(int64(stones(s)), 10)))
}

func verify(param string) (string, error) {
	var s string

	if n := len(param); 1 > n || n > 50 {
		return "", incorErr
	} else {
		s = strings.ToLower(param)
		for i := range s {
			if s[i] != 98 && s[i] != 103 && s[i] != 114 {
				return "", incorErr
			}
		}
		return s, nil
	}
}

func stones(s string) int {
	var c int
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			c++
		}
	}
	return c
}

func main() {
	incorErr = errors.New("Incorrect input data. Please, try again")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
