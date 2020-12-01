package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	k := 2020
	hm := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if v, ok := hm[k-i]; ok {
			result := v * i
			fmt.Println(v, result)
		} else {
			hm[i] = i
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
