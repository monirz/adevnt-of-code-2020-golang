package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// part1()
	part2()
}

var mem = make(map[int]int)

func part1() {

	maskSet := 0
	maskClr := 0

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {

			// n, err := strconv.Atoi(text)
			// if err != nil {
			// 	panic(err)
			// }

			if strings.Contains(text, "mask =") {
				splitted := strings.Split(text, "mask =")

				bitStr := strings.TrimSpace(splitted[1])
				maskSet = 0
				maskClr = 0

				for _, v := range bitStr {
					maskSet = maskSet * 2
					maskClr = maskClr * 2
					if v == '0' {
						maskClr++
					} else if v == '1' {
						maskSet++
					}
				}

			} else {
				spl := strings.Split(text, "mem[")
				spl2 := strings.Split(spl[1], "]")

				adr, err := strconv.Atoi(spl2[0])

				if err != nil {
					panic(err)
				}

				daStr := strings.Split(spl2[1], "=")

				da, err := strconv.Atoi(strings.TrimSpace(daStr[1]))

				fmt.Println(adr, da)
				mem[adr] = (da | maskSet) & ^maskClr

			}

		} else {
			break
		}

	}

	total := 0

	for _, da := range mem {
		total = total + da
	}

	fmt.Println("total ", total)

}

var maskz = make(map[int]bool)

func part2() {

	maskSet := 0
	maskClr := 0
	// bitn := 0

	scanner := bufio.NewScanner(os.Stdin)
	for {

		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {

			if strings.Contains(text, "mask =") {
				splitted := strings.Split(text, "mask =")

				bitStr := strings.TrimSpace(splitted[1])
				maskSet = 0
				maskClr = 0
				maskz = make(map[int]bool)
				bitn := 35

				for _, v := range bitStr {
					maskSet = maskSet * 2
					maskClr = maskClr * 2
					if v == '0' {
						maskClr++
					} else if v == '1' {
						maskSet++
					} else if v == 'X' {
						maskz[bitn] = true
					}

					bitn--
				}

			} else {
				spl := strings.Split(text, "mem[")
				spl2 := strings.Split(spl[1], "]")

				adr, err := strconv.Atoi(spl2[0])

				if err != nil {
					panic(err)
				}

				daStr := strings.Split(spl2[1], "=")

				da, err := strconv.Atoi(strings.TrimSpace(daStr[1]))

				fmt.Println(adr, da)
				adr = adr | maskSet
				setAll(0, adr, da)

			}

		} else {
			break
		}

	}

	total := 0
	for _, da := range mem {
		total = total + da
	}
	// fmt.Println("mem ", mem)
	fmt.Println("total ", total)

}

func setAll(bitn int, adr, da int) {
	if bitn == 36 {
		mem[adr] = da
		return
	}

	if maskz[bitn] {
		bitU := bitn
		adr = adr & ^(1 << uint64(bitU))
		setAll(bitn+1, adr|(1<<uint64(bitU)), da)
		setAll(bitn+1, adr, da)
	} else {
		setAll(bitn+1, adr, da)
	}

}
