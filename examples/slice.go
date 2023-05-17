package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("s:%#v\n", s)
	fmt.Printf("s1:%#v\n", s[:])
	fmt.Printf("s2:%#v\n", s[1:])
	fmt.Printf("s3:%#v\n", s[:1])
}
