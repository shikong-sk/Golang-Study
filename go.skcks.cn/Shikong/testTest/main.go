package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1,2,3,4,5,6"
	fmt.Println(s)
	fmt.Println(Split(s, ","))
}

func Split(s, sep string) (res []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		res = append(res, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	res = append(res, s)
	return
}
