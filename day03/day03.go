package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	//read input from file
	file, err := os.Open("day03.in")
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

	result := part1(arr)

	fmt.Println("Part1: ", result)

	result = part2(arr)

	fmt.Println("Part2: ", result)

}

func part1(arr []string) int {

	treeCount := 0
	currentPos := 0

	row := 0
	for row < len(arr)-1 {

		currentPos += 3
		row++
		if arr[row][currentPos%len(arr[row])] == '#' {
			treeCount++
		}

	}

	return treeCount
}

func part2(arr []string) int {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1

	for _, v := range slopes {
		r := 0
		c := 0

		count := 0

		for r+1 < len(arr) {
			c += v[0]
			r += v[1]

			if arr[r][c%len(arr[r])] == '#' {
				count++
			}

		}
		result = result * count

	}

	return result
}
