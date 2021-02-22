package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)

	if len(sBytes) != len(tBytes) {
		return false
	}

	for _, i := range sBytes {
		alphabet[i - 'a']++
	}

	for _, i := range tBytes {
		alphabet[i - 'a']--
	}

	for _, i := range alphabet {
		if i != 0{
			return false
		}
	}

	return true
}

func main() {
	res := isAnagram("anagram", "nagaram")
	fmt.Println(res)
}
