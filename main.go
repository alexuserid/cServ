package main

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	reN, reS *regexp.Regexp
	incorErr error
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	m := make(map[string]string)
	for k, v := range r.Form {
		str := strings.Join(v, "")
		m[k] = str
	}

	n, s, err := verify(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ans := stones(n, s)
	w.Write([]byte(strconv.FormatInt(int64(ans), 10)))
}

func verify(m map[string]string) (int, string, error) {
	var s string
	var n int

	for k, v := range m {
		switch k {
		case "n":
			if ok := reN.MatchString(v); ok {
				tmp, err := strconv.Atoi(v)
				if err != nil {
					return 0, "", incorErr
				}
				n = tmp
			}
		case "s":
			if ok := reS.MatchString(v); ok {
				s = strings.ToUpper(v)
			}
		default:
			return 0, "", incorErr
		}
	}

	l := utf8.RuneCountInString(s)
	if l == n {
		return n, s, nil
	} else {
		return 0, "", incorErr
	}
}

func stones(n int, s string) int {
	var c int
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			c++
		}
	}
	return c
}

func main() {
	incorErr = errors.New("Incorrect input data. Please, try again")

	reN = regexp.MustCompile("[0-9]{1,2}")
	reS = regexp.MustCompile("[RGBrgb]{1,50}")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
