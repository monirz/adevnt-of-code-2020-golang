package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//read input from file
	file, err := os.Open("input.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passport := make(map[string]string)

	scanner := bufio.NewScanner(file)

	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	eyeColor := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	count := 0

	for scanner.Scan() {

		s := scanner.Text()
		s = strings.TrimSpace(s)

		if len(s) == 0 {

			valid := true

			for _, v := range required {
				if _, ok := passport[v]; !ok {
					valid = false
				}
			}

			for k, v := range passport {

				if k == "byr" {

					i, err := strconv.Atoi(v)
					if err != nil {
						panic(err)
					}

					if i < 1920 || i > 2002 {
						valid = false
					}

				}

				if k == "iyr" {
					i, err := strconv.Atoi(v)
					if err != nil {
						panic(err)
					}

					if i < 2010 || i > 2020 {
						valid = false
					}
				}

				if k == "eyr" {
					i, err := strconv.Atoi(v)
					if err != nil {
						panic(err)
					}

					if i < 2020 || i > 2030 {
						valid = false
					}
				}

				if k == "hgt" {

					if strings.HasSuffix(v, "cm") {
						v = strings.TrimSuffix(v, "cm")

						i, err := strconv.Atoi(v)
						if err != nil {
							panic(err)
						}

						if i < 150 || i > 193 {
							valid = false
						}

					} else if strings.HasSuffix(v, "in") {
						v = strings.TrimSuffix(v, "in")

						i, err := strconv.Atoi(v)
						if err != nil {
							panic(err)
						}

						if i < 59 || i > 76 {
							valid = false
						}
					} else {
						valid = false
					}

				}

				if k == "hcl" {

					var isStringAlphabetic = regexp.MustCompile(`^#[a-f0-9]{6}$`).MatchString

					if !isStringAlphabetic(v) {
						valid = false
					}

				}

				if k == "ecl" {

					for x, y := range eyeColor {

						if v == y {
							break
						}

						if x == len(eyeColor)-1 && v != y {
							valid = false
						}

					}
				}

				if k == "pid" {

					var isStringAlphabetic = regexp.MustCompile(`^\d{9}$`).MatchString

					if !isStringAlphabetic(v) {
						valid = false
					}

				}

			}

			if valid {
				count++
			}

			passport = make(map[string]string)

		} else {

			splitted := strings.Split(s, " ")

			for _, v := range splitted {

				spl := strings.Split(v, ":")
				// fmt.Println(spl)
				passport[spl[0]] = spl[1]

			}

		}

	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
