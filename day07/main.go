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

	//read input from file

	part1()
	part2()

}

func part1() {

	file, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// hm := make(map[string]bool)

	// total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := scanner.Text()

		if len(s) == 0 {

		} else {

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func part2() {

	file, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// hm := make(map[string]bool)

	// total := 0

	hm := make(map[string]map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := scanner.Text()

		spl := strings.Split(s, "bags contain")
		bag := make(map[string]int)

		if strings.TrimSpace(spl[1]) != "no other bags." {
			// fmt.Println(s, spl[1])

			var insideColor string
			var count int

			inBags := strings.Split(spl[1], ",")

			for _, v := range inBags {
				// count = fmt.Sprintf("%d", )
				// insideColor = strings.TrimSpace(v)

				v = strings.TrimSpace(v)

				// fmt.Println("----", v)
				splitted := strings.Split(v, " ")

				count, err = strconv.Atoi(splitted[0])

				if err != nil {
					panic(err)
				}

				insideColor = splitted[1] + " " + splitted[2]

				bag[insideColor] = count

			}

			// fmt.Println("bag ", bag)
			hm[strings.TrimSpace(spl[0])] = bag

		}

	}

	result := iter("shiny gold", hm)

	fmt.Println(len(result))

	for k, v := range hm {
		fmt.Println(k, v)
	}

	part2 := totalBags("shiny gold", hm)
	fmt.Println(part2 - 1)

	// fmt.Println(contains)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

var contains = make(map[string]int)

func iter(want string, hm map[string]map[string]int) map[string]int {

	for out, contents := range hm {
		// fmt.Println(out)
		for k := range contents {

			if k == want {
				contains[out]++
				// fmt.Println("---->", contains, out)
				iter(out, hm)
			}
		}
	}

	return contains

}

func totalBags(want string, hm map[string]map[string]int) int {

	var count int
	if val, ok := hm[want]; ok {
		for inColor, n := range val {
			// fmt.Println(hm[want])
			count += n * totalBags(inColor, hm)

		}
	}

	return count + 1
}
