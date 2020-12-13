package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// arr := make([]int, 0)

	// scanner := bufio.NewScanner(os.Stdin)

	// timeStamp := 0
	// countLine := 0
	// var err error
	// busIDs := []int{}

	// for {
	// 	scanner.Scan()
	// 	text := scanner.Text()
	// 	if len(text) != 0 {

	// 		countLine++

	// 		if countLine == 1 {
	// 			timeStamp, err = strconv.Atoi(text)
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 		}

	// 		if countLine == 2 {

	// 			splitted := strings.Split(text, ",")
	// 			for _, v := range splitted {
	// 				if v == "x" {
	// 					continue
	// 				}

	// 				n, err := strconv.Atoi(v)
	// 				if err != nil {
	// 					panic(err)
	// 				}

	// 				busIDs = append(busIDs, n)
	// 			}
	// 		}

	// 	} else {
	// 		break
	// 	}

	// }

	// fmt.Println(busIDs, timeStamp)

	// part1(timeStamp, busIDs)
	part2()
}

func part1(timestamp int, arr []int) int {

	bid, maxToWait := 0, 999999999

	for _, b := range arr {
		toWait := b - (timestamp % b)

		if toWait < maxToWait {
			bid = b
			maxToWait = toWait
		}
	}

	fmt.Println(bid, maxToWait, bid*maxToWait)
	return 0
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	timeStamp := 0
	countLine := 0
	var err error
	busIDs := make(map[int]int)
	bid := 0

	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {

			countLine++

			if countLine == 1 {
				timeStamp, err = strconv.Atoi(text)
				if err != nil {
					panic(err)
				}

			}

			if countLine == 2 {

				splitted := strings.Split(text, ",")
				for _, v := range splitted {
					if v != "x" {
						n, err := strconv.Atoi(v)
						if err != nil {
							panic(err)
						}
						busIDs[bid] = n
					}

					bid++

				}

			}

		} else {
			break
		}

	}

	ts, ntot := 0, 1

	fmt.Println("busIDs", busIDs)
	for bid, bmod := range busIDs {
		inv := modInverse(ntot, bmod)

		res := (bmod - (ts+bid)%bmod) % bmod
		n := (inv * res) % bmod
		ts = ts + ntot*n

		ntot = ntot * bmod

	}

	fmt.Println(timeStamp, busIDs, "ts ", ts)
}

func modInverse(b, m int) int {

	g, x, _ := gcd(b, m)

	if g != 1 {
		return 0
	}

	return x % m

}

func gcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	g, x1, y1 := gcd(b%a, a)

	return g, y1 - (b/a)*x1, x1

}
