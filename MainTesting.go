package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var run bool
	run = true
	for run {
		fmt.Println("Hello, enter a string")
		scanthing := bufio.NewScanner(os.Stdin)
		scanthing.Scan()
		var textUser string
		textUser = scanthing.Text()
		for true {
			fmt.Println("You typed '", textUser, "'\nIf you would like to enter another string, enter (0). If not, enter (1) to exit.")

			scanthing.Scan()
			userChoice := scanthing.Text()

			if userChoice == "0" {
				break
			} else if userChoice == "1" {
				run = false
				break
			} else {
				fmt.Println("Invalid input! Please try again!")
			}
		}
	}
	fmt.Println("EXIT!")
}
