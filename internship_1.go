package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func sortName(abbreName []string) []string {
	sort.Slice(abbreName, func(i, j int) bool {
		if len(abbreName[i]) > len(abbreName[j]) {
			return true
		}
		if len(abbreName[i]) < len(abbreName[j]) {
			return false
		}
		return abbreName[i] < abbreName[j]
	})
	return abbreName
}

func findAbbre(fullName []string) []string {
	var abbreName []string
	abbreName = make([]string, len(fullName))
	isUpper := regexp.MustCompile(`[A-Z][^A-Z]*`)
	for i := 0; i < len(fullName); i++ {
		subString := isUpper.FindAllString(fullName[i], -1)
		for _, element := range subString {
			abbreName[i] = abbreName[i] + element[0:1]
		}
	}
	return abbreName
}

func main() {
	var fullName []string
	var size int
	fmt.Print("Enter the number of your input : ")
	fmt.Scanln(&size)
	fullName = make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Print("Input number ", (i + 1), " : ")
		in := bufio.NewReader(os.Stdin)
		fullName[i], _ = in.ReadString('\n')
		fullName[i] = strings.TrimSpace(fullName[i])
	}
	abbreName := findAbbre(fullName)
	sortedName := sortName(abbreName)
	fmt.Println("Size : ", size)
	fmt.Printf("Input\n\n")
	for i := 0; i < size; i++ {
		fmt.Println("	", fullName[i])
	}
	fmt.Printf("----------------------------------\n\nOutput\n\n")
	for i := 0; i < len(sortedName); i++ {
		fmt.Println("	", sortedName[i])
	}
}
