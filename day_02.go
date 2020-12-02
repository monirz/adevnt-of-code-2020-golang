package main

import (
	"log"
	"strconv"
	"strings"
)

func validPassword(arr [][]string) int {

	count := 0
	//

	for _, v := range arr {

		hm := make(map[string]int)

		// fmt.Println(k, v)
		for _, ch := range v[2] {

			if _, ok := hm[string(ch)]; ok {
				hm[string(ch)]++
			} else {
				hm[string(ch)] = 1
			}
		}

		// fmt.Println(hm)

		char := strings.TrimRight(v[1], ":")
		splittedNum := strings.Split(v[0], "-")

		firstNum, err := strconv.Atoi(splittedNum[0])

		//todo return the error
		if err != nil {
			log.Fatal(err)
		}
		secondNum, err := strconv.Atoi(splittedNum[1])

		if err != nil {
			log.Fatal(err)
		}

		if val, ok := hm[char]; ok {
			if val <= secondNum && val >= firstNum {
				count++
			}

		}
	}

	return count

}
