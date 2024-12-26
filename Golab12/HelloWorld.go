package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Print(ReadTxtFromTerminal())
	PrintEven(2, 7)
	PrintEven(1, 0)
	PrintEven(5, 101)
	fmt.Println(Apply(9, 10, "*"))
	fmt.Println(Apply(100, 0, "/"))
	fmt.Println(Apply(10, 22, "0"))
}

func ReadTxtFromTerminal() string {
	fmt.Print("hello user, write one word here:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return line
}

func PrintEven(a int, b int) {
	var name string
	var trigger bool
	var low int
	if a > b {
		trigger = true
		fmt.Print("A > B, range set wrong")
	}
	low = a / 2
	low = low * 2
	if low < a {
		low = low + 2
	}
	for !trigger {
		name = name + strconv.Itoa(low) + " "
		low = low + 2
		if low > b {
			trigger = true
		}
	}
	fmt.Println(name)
}

func Apply(a int, b int, symbol string) (int, error) {
	var output int
	var symbols [5]string = [5]string{"*", "+", "-", "/", ":"}
	if symbol == symbols[0] {
		output = a * b
	} else if symbol == symbols[1] {
		output = a + b
	} else if symbol == symbols[2] {
		output = a - b
	} else if symbol == symbols[3] {
		if b == 0 {
			return 0, errors.New("Do NOT divide by zero")
		}
		output = a / b
	} else if symbol == symbols[4] {
		output = a / b
	} else {
		return 0, errors.New("No correct symbol been given")
	}
	return output, nil
}
