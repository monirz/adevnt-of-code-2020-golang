package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	fmt.Println(arr)

	result := part1(arr)

	fmt.Println(result)

	part2(arr)

}

func part1(arr []int) int {

	sort.Ints(arr)

	jdiff := [4]int{0, 0, 1, 1}
	jLast := 0

	for _, v := range arr {
		jdiff[v-jLast] = jdiff[v-jLast] + 1
		jLast = v
	}

	fmt.Println(jdiff[1], jdiff[3], jdiff[1]*jdiff[3], jdiff)
	return jdiff[1] * jdiff[3]
}

func part2(adapters []int) {

	sort.Ints(adapters)

	for _, v := range adapters {
		hasAdapter[v] = true
		if v > max {
			max = v
		}

	}

	result := nways(0)

	fmt.Println("part2: ", result)
}

var nwaysMem = make(map[int]int)
var hasAdapter = make(map[int]bool)
var max = 0

func nways(jolts int) int {

	if _, ok := nwaysMem[jolts]; ok {
		return nwaysMem[jolts]
	}

	ways := 0

	for i := jolts + 1; i <= jolts+3; i++ {
		if i == max+3 {
			ways++
		} else if _, ok := hasAdapter[i]; ok {
			fmt.Println("okkkk")
			ways = ways + nways(i)
		}
	}

	nwaysMem[jolts] = ways

	return ways

}
