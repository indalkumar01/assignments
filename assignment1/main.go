package main

import (
	"fmt"
	"math"
)

func main() {
	// var year int
	// fmt.Println("Enter year to know leap year:-")
	// _, err := fmt.Scan(&year)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(IsLeapYear(uint32(year)))

	// // 2
	// var num int
	// fmt.Println("Enter number:-")
	// _, err := fmt.Scan(&num)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// EvenOdd(num)

	// if IsNegativeNumber(num) {
	// 	fmt.Println("this is negative number")
	// } else {
	// 	fmt.Println("this is not a negative number")
	// }

	// fmt.Println("this number digit count is :- ", CountDigit(num))

	// // 3
	// var firstNum, secondNum int
	// fmt.Println("Enter number:-")
	// _, err := fmt.Scan(&firstNum, &secondNum)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("sum is :==>", Sum(firstNum, secondNum))
	// fmt.Println("minus is :==>", Minus(firstNum, secondNum))
	// fmt.Println("division is :==>", Division(firstNum, secondNum))
	// fmt.Println("multiplication is :==>", Multiplication(firstNum, secondNum))

	// 4
	// var str string
	// fmt.Println("Enter string:-")
	// _, err := fmt.Scan(&str)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("lenght of the string is:=>", PintStrLength(str))
	// fmt.Println("last charecter in string is:=>", PrintLastCharOfString(str))
	// fmt.Println("Ascii value of each string charecter is:=>", PrintAsciiChar(str))
	// fmt.Println("reverse string is:=>", ReverseString(str))

	var str1, str2 string
	fmt.Println("Enter two string:-")
	_, err := fmt.Scan(&str1, &str2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("concatinated string is :=>", ConcatinateString(str1, str2))
}

func IsLeapYear(year uint32) string {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return "This is a leap year"
	}

	return "This is not a leap year"
}

func EvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("this is even number")
	} else {
		fmt.Println("this is odd number")
	}
}

func IsNegativeNumber(num int) bool {
	if num < 0 {
		return true
	}

	return false
}

func CountDigit(num int) int {
	if num == 0 {
		return 1
	}

	absNum := int(math.Abs(float64(num)))

	count := 0

	for absNum > 0 {
		absNum /= 10

		count++
	}

	return count
}

func Sum(num ...int) int {
	sum := 0
	for n := range num {
		sum += n
	}
	return sum
}

func Minus(firstNum, secondNum int) int {
	return firstNum - secondNum
}

func Division(firstNum, secondNum int) int {
	return firstNum / secondNum
}

func Multiplication(firstNum, secondNum int) int {
	return firstNum * secondNum
}

func ReverseString(str string) string {
	runes := []rune(str)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]

		left++
		right--
	}

	return string(runes)
}

func PrintLastCharOfString(str string) string {
	runes := []rune(str)
	return string(runes[len(runes)-1])
}

func PrintAsciiChar(str string) map[string]rune {
	runes := []rune(str)
	mapData := make(map[string]rune)

	for _, val := range runes {
		mapData[string(val)] = val
	}

	return mapData
}

func PintStrLength(str string) int {
	return len(str)
}

func ConcatinateString(str ...string) string {
	var conStr string

	for _, val := range str {
		conStr += val
	}

	return conStr
}
