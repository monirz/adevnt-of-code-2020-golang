package main

import "fmt"

func part2(arr []int, k int) int {

	// fmt.Println(arr, k)
	var result int
	hm := make(map[int]int)

	for i := 0; i < len(arr); i++ {

		for j := 1; j < len(arr); j++ {

			target := k - (arr[i] + arr[j])

			if _, ok := hm[target]; ok {
				fmt.Println("ok")
				return arr[i] * arr[j] * target
			}
			hm[arr[i]] = arr[i]

		}
	}

	return result
}
