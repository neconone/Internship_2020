package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkInArray(str string, list []string, ans []string, score int, wrong []string, wrongNum int) ([]string, int, []string, int, bool) {
	state := false
	for i := 0; i < len(list); i++ {
		if str == list[i] {
			if ans[i] == "_" {
				ans[i] = list[i]
				score++
				state = true
			} else {
				return ans, score, wrong, wrongNum, state
			}
		}
	}
	if state == false {
		state = true
		for i := 0; i < wrongNum; i++ {
			if str == wrong[i] {
				state = false
			}
		}
		if state == true {
			wrong[wrongNum] = str
			wrongNum++
		}
	}
	return ans, score, wrong, wrongNum, state
}

func main() {

	// Take the input
	in := bufio.NewReader(os.Stdin)
	state := false
	var num []int
	num = make([]int, 12)
	var numString []string
	for state == false {
		state = true
		fmt.Print("Enter the number 0 - 9 : ")
		input, _ := in.ReadString('\n')
		input = strings.TrimSpace(input)
		isNumber := regexp.MustCompile(` `)
		temp := isNumber.Split(input, -1)
		var err error
		for i := 0; i < len(temp); i++ {
			num[i], err = strconv.Atoi(temp[i])
			if err != nil || num[i] > 9 || num[i] < 0 || len(temp) != 12 {
				i = len(temp) + 1
				fmt.Println("Wrong input!!")
				state = false
			}
		}
		numString = temp
	}

	// Game zone
	if state == true {
		fmt.Println("_ _ _ _ _ _ _ _ _ _ _ _ ||\n")
		score := 0
		ans := []string{"_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_"}
		var wrong []string
		wrong = make([]string, 10)
		wrongNum := 0
		for i := 0; i < 5; {
			fmt.Print("Enter your answer : ")
			input, _ := in.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			if err != nil || num > 9 || num < 0 {
				fmt.Println("Wrong input!!")
			} else {
				i++
				var state bool
				ans, score, wrong, wrongNum, state = checkInArray(input, numString, ans, score, wrong, wrongNum)
				if state == true {
					for k := 0; k < 12; k++ {
						fmt.Print(ans[k], " ")
					}
					fmt.Print("|| ")
					for k := 0; k < wrongNum; k++ {
						fmt.Print(wrong[k], " ")
					}
				} else {
					fmt.Println("Please, don't use the same answer.")
				}
			}
			fmt.Println("\nscore : ", score, "\n")
		}
	}
}
