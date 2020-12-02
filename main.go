package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	k := 2020
	var arr []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		arr = append(arr, i)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//part one
	result := part1(arr, k)
	fmt.Println("Part 1 result:", result)

	//part two input
	result = part2(arr, k)
	fmt.Println("Part two: ", result)

	//day 2 part 01

	//TODO refactor this code

	file, err = os.Open("day_02.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	var strArr [][]string
	for scanner.Scan() {

		s := scanner.Text()
		splittedStr := strings.Split(s, " ")
		strArr = append(strArr, splittedStr)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result = validPassword(strArr)
	fmt.Println("total valid password: ", result)
}

func part1(arr []int, k int) int {
	hm := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if v, ok := hm[k-arr[i]]; ok {
			return v * arr[i]
		}
		hm[arr[i]] = arr[i]

	}

	return 0

}
