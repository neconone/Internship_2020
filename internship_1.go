package main

import "fmt"

func findAbbre(fullname []string) {
	fmt.Println(fullname[4])
}

func main() {
	var fullname []string
	var size int
	fmt.Print("Enter the number of your input : ")
	fmt.Scanln(&size)
	fullname = make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Input number %d : ", (i + 1))
		fmt.Scanln(&fullname[i])
	}
	findAbbre(fullname)
}
