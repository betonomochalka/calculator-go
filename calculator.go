package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clearCMD() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func writeNumbers() [2]int {
	var answerNumbers [2]int

	fmt.Println("Write two numbers: ")
	_, err := fmt.Scan(&answerNumbers[0])
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = fmt.Scan(&answerNumbers[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	rightAnswerNumbers(answerNumbers[0], answerNumbers[1])

	writeOperator(answerNumbers[0], answerNumbers[1])

	return answerNumbers
}

func writeOperator(number1 int, number2 int) {
	fmt.Println("What math operation do you want me to do?", "\n",
		"1. +", "\n",
		"2. -", "\n",
		"3. *", "\n",
		"4. /")
	var operator string
	_, err := fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Error:", err)
	}

	rightAnswerOperator(operator, number1, number2)

	var result int

	switch operator {
	case "+":
		result = number1 + number2
		fmt.Println("Answer:", result)
	case "-":
		result = number1 - number2
		fmt.Println("Answer:", result)
	case "*":
		result = number1 * number2
		fmt.Println("Answer:", result)
	case "/":
		result = number1 / number2
		fmt.Println("Answer:", result)
	}
}

func rightAnswerOperator(operator string, number1 int, number2 int) {
	clearCMD()
	var answerRight string

	fmt.Println(operator, "Right? (y/n)")

	_, err := fmt.Scan(&answerRight)
	if err != nil {
		fmt.Println("Error:", err)
	}

	answerRight = strings.ToLower(answerRight)

	if answerRight == "yes" {

	} else if answerRight == "y" {

	} else if answerRight == "no" {
		clearCMD()
		writeOperator(number1, number2)
	} else if answerRight == "n" {
		clearCMD()
		writeOperator(number1, number2)
	} else {
		fmt.Println("No option", answerRight)
		rightAnswerOperator(operator, number1, number2)
	}

	clearCMD()
}

func rightAnswerNumbers(number1 int, number2 int) {
	clearCMD()
	var answerRight string

	fmt.Println(number1, "and", number2, "Right? (y/n)")

	_, err := fmt.Scan(&answerRight)
	if err != nil {
		fmt.Println("Error:", err)
	}

	answerRight = strings.ToLower(answerRight)

	if answerRight == "yes" {

	} else if answerRight == "y" {

	} else if answerRight == "no" {
		clearCMD()
		writeNumbers()
	} else if answerRight == "n" {
		clearCMD()
		writeNumbers()
	} else {
		fmt.Println("No option", answerRight)
		rightAnswerNumbers(number1, number2)
	}

	clearCMD()
}

func exit() {
	fmt.Println("Do you want to exit? (y/n)")
	var answerExit string
	_, err := fmt.Scan(&answerExit)
	if err != nil {
		fmt.Println("Error:", err)
	}

	answerExit = strings.ToLower(answerExit)

	if answerExit == "yes" {
		clearCMD()
		fmt.Println("Bye")
	} else if answerExit == "y" {
		clearCMD()
		fmt.Println("Bye")
	} else if answerExit == "no" {
		clearCMD()
		main()
	} else if answerExit == "n" {
		clearCMD()
		main()
	} else {
		fmt.Println("No option", answerExit)
		exit()
	}
}

func main() {
	clearCMD()
	writeNumbers()
	exit()
}
