package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	hm := make(map[string]bool)

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := scanner.Text()

		if len(s) == 0 {
			keys := 0

			for _, v := range hm {

				keys++
				fmt.Println(v)

			}

			total += keys

			fmt.Println("k ", keys, " total ", total)

			hm = make(map[string]bool)

		} else {
			for _, v := range s {
				hm[string(v)] = true
			}
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

	hm := make(map[string]int)

	total := 0
	n := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := scanner.Text()

		if len(s) == 0 {
			keys := 0

			for _, v := range hm {

				// keys++

				if v == n {
					keys = keys + 1
				}
				fmt.Println(v)

			}

			total += keys

			fmt.Println("k ", keys, " total ", total)

			hm = make(map[string]int)
			n = 0

		} else {
			for _, v := range s {

				if val, ok := hm[string(v)]; ok {
					hm[string(v)] = val + 1
				} else {
					hm[string(v)] = 1

				}
			}

			n++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
