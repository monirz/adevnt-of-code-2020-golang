package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {

			n, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}

			arr = append(arr, n)

		} else {
			break
		}

	}

	// fmt.Println(arr)
	badNumber := part1(arr)
	fmt.Println(badNumber)

	result := part2(arr, badNumber)

	min, max := minMax(result)

	total := min + max

	fmt.Println(total)
}

func part1(arr []int) int {
	preambleSize := 25
	badNumber := 0
	for i := range arr {
		if i+preambleSize == len(arr) {
			break
		}
		subSlice := arr[i : i+preambleSize]
		numToCheck := arr[i+preambleSize]
		if !valid(numToCheck, subSlice) {
			badNumber = numToCheck
			break
		}
	}
	return badNumber
}

func part2(nums []int, target int) []int {

	left := 0
	right := 1

	total := nums[left] + nums[right]

	for left < right && right < len(nums) {
		if total == target {
			return nums[left : right+1]
		} else if total < target {
			right++
			total += nums[right]
		} else {
			total -= nums[left]
			left++
		}

	}

	return nums
}

func valid(goal int, nums []int) bool {
	for i, num1 := range nums {
		for _, num2 := range nums[i:] {
			if num1+num2 == goal {
				return true
			}
		}
	}
	return false
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
