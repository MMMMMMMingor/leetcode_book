package main

import (
	"fmt"
)

func isHappy(n int) bool {
	res, num, record := 0, n, map[int]int{}

	for {
		for num > 0 {
			res += (num % 10) * (num % 10)
			num = num / 10
		}

		if _, ok := record[res]; !ok{
			if res == 1{
				return true
			}
			record[res] = res
			num = res
			res = 0
		}else{
			return false
		}
	}
    
}

func main() {
	res := isHappy(0)
	fmt.Println(res)
}
