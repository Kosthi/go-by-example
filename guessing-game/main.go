package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxNum := 100
	// 现在Go可以自动播种了
	secretNum := rand.Intn(maxNum)
	// fmt.Println("The secret num is", secretNum)

	fmt.Println("Please input your guess")

	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("A error occur when reading.")
			return
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			return
		}

		if guess > secretNum {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNum {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
