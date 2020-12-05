package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(arr)
}

func part1(arr []string) {

	var n int

	count := 0

	l := make(map[int]bool)

	for _, v := range arr {

		left := 0
		right := 127
		for i, j := range v {
			mid := (left + right) / 2

			if j == 'F' {
				right = mid
			} else if j == 'B' {
				left = mid + 1
			}

			if j == 'L' || j == 'R' {
				str := v[i:]

				n = search(str)

				break

			}
		}

		result := left*8 + n

		if result < 1 {
			fmt.Println(result)
		}

		l[result] = true

		// fmt.Println(result)
		if result > count {
			count = result
		}

	}

	total := 127*8 + 7

	for i := 0; i < total; i++ {
		if l[i-1] == true && l[i] == false && l[i+1] {
			fmt.Println("seat >>>>>>>>>>>", i)
		}
	}

	fmt.Println(count)
	// fmt.Println(left, right, result)

}

func search(str string) int {

	left := 0
	right := 7

	for _, j := range str {
		mid := (left + right) / 2

		// fmt.Println(mid)

		if j == 'L' {
			right = mid
		} else if j == 'R' {
			left = mid + 1
		}

	}

	return left
}
