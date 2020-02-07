package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	var input string
	state := true
	for state {
		fmt.Print("Enter the number 1.0 - 10.0 (less than 12 digit except point): ")
		in := bufio.NewReader(os.Stdin)
		input, _ = in.ReadString('\n')
		input = strings.TrimSpace(input)
		prime := "FALSE"
		if input == "0.0" {
			state = false
		} else if num, err := strconv.ParseFloat(input, 64); err == nil {
			if (num > 10 || num < 1 || len(input) > 13) && num != 0 {
				fmt.Println("Wrong input!!")
			} else {
				for i := 1; i <= 3 && prime == "FALSE"; i++ {
					temp := int(num * float64(math.Pow10(i)))
					if IsPrimeSqrt(temp) {
						prime = "TRUE"
					}
				}
				fmt.Println(num, " : ", prime)
			}

		} else {
			fmt.Println("Wrong input!!")
		}
	}
}
