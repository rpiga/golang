package main

import (
	"fmt"
	"strings"
)

func main() {
	var s_in string

	fmt.Printf("Enter string: ")
	fmt.Scan(&s_in)
	s_in = strings.ToLower(s_in)
	fmt.Printf("string is %s.\n", s_in)

	if strings.HasPrefix(s_in, "i") && strings.HasSuffix(s_in, "n") && strings.Contains(s_in, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
