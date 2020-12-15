package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	part2()
}

func part1() {

	nums := make(map[int]int)
	t := 1

	turns := make(map[int]int)
	// nextn := 0
	prevTurn := 0

	lastnum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)

			splitted := strings.Split(text, ",")

			for _, v := range splitted {
				n, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				nums[n] = t
				turns[t] = n
				lastnum = n
				t = t + 1

			}

		} else {
			break
		}

	}

	fmt.Println("nums ----------", nums)
	for i := t; i <= 2020; i++ {
		thisn := 0

		if prevTurn > 0 {
			thisn = i - 1 - prevTurn
			// fmt.Println(thisn)

		} else {
			thisn = 0
		}

		lastnum = thisn
		prevTurn = nums[thisn]
		nums[thisn] = i
		turns[i] = thisn
	}

	fmt.Println(lastnum)

}

//
func part2() {

	nums := make(map[int]int)
	t := 1

	turns := make(map[int]int)
	// nextn := 0
	prevTurn := 0

	lastnum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)

			splitted := strings.Split(text, ",")

			for _, v := range splitted {
				n, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				nums[n] = t
				turns[t] = n
				lastnum = n
				t = t + 1

			}

		} else {
			break
		}

	}

	fmt.Println("nums ----------", nums)
	for i := t; i <= 30000000; i++ {
		thisn := 0

		if prevTurn > 0 {
			thisn = i - 1 - prevTurn
			// fmt.Println(thisn)

		} else {
			thisn = 0
		}

		lastnum = thisn
		prevTurn = nums[thisn]
		nums[thisn] = i
		turns[i] = thisn
	}

	fmt.Println(lastnum)

}
