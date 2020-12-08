package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}

	}

	_, result := part1(arr)

	fmt.Println("result ", result)

	fmt.Println(part2(arr))
}

func part1(arr []string) (bool, int) {

	l := len(arr)

	// pos := 0
	acc := 0
	// jmp := 0

	hm := make(map[int]bool)

	for i := 0; i < l; i++ {

		if _, ok := hm[i]; ok {
			return true, acc
		}
		hm[i] = true

		splitted := strings.Split(arr[i], " ")

		if strings.TrimSpace(splitted[0]) == "nop" {
			continue
		}

		n, err := strconv.Atoi(splitted[1])

		if err != nil {
			panic(err)
		}

		if strings.TrimSpace(splitted[0]) == "acc" {
			acc += n
			// fmt.Println("---> ", n, acc)

		}

		if strings.TrimSpace(splitted[0]) == "jmp" {

			i += n - 1
		}

		// pos +=
	}

	return false, acc

}

func part2(arr []string) int {

	ok, acc := part1(arr)

	for i := 0; i < len(arr); i++ {
		tmp := make([]string, len(arr))
		copy(tmp, arr)

		if arr[i][:3] == "jmp" {
			tmp[i] = "nop" + arr[i][3:]
			// fmt.Println(tmp)

		} else if arr[i][:3] == "nop" {
			tmp[i] = "jmp" + arr[i][3:]
		} else {
			continue
		}

		ok, acc = part1(tmp)

		// fmt.Println(ok, acc)
		if !ok {
			break
		}
	}

	return acc
}
