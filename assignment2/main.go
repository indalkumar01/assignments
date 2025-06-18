package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1
	// var num int

	// fmt.Println("Enter number:=>")
	// fmt.Scan(&num)

	// isPrime := isPrimeNumber(num)
	// isPelindrom := IsPelindrome(num)

	// fmt.Println("is isPrime ", isPrime)
	// fmt.Println("is pelindrome ", isPelindrom)

	// //2
	// var num int

	// fmt.Println("Enter number for square root:=>")
	// fmt.Scan(&num)

	// fmt.Println("squre root is := ", SquareRoot(float64(num)))

	// var num1, num2 float64
	// fmt.Println("Enter 2 number for power:=>")
	// fmt.Scan(&num1, &num2)

	// fmt.Println("power is := ", Power(num1, num2))

	//3
	// matrix := [3][3]int{
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// }
	// totalSum := AdditionOfMatrix(matrix)
	// totalMultiplication := MultiplicationOfMatrix(matrix)

	// fmt.Println("total sum is := ", totalSum)
	// fmt.Println("total multiplication is:= ", totalMultiplication)

	// 4
	// fmt.Println("Enter elements of array:=>")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()

	// strArr := strings.Fields(scanner.Text())
	// intSlice := []int{}

	// for _, val := range strArr {
	// 	intData, err := strconv.Atoi(val)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		intSlice = append(intSlice, intData)
	// 	}
	// }

	// fmt.Println("after reversing the array :=> ", ReverseArray(intSlice))
	// fmt.Println("total sum of the elements in array :=> ", SumOfArray(intSlice))

	// 5
	fmt.Println("Enter elements of array:=>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	strArr := strings.Fields(scanner.Text())
	intSlice := []int{}

	for _, val := range strArr {
		intData, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
		} else {
			intSlice = append(intSlice, intData)
		}
	}

	fmt.Println("length of the array is :=>", len(intSlice))

	first, last := PrintFirstAndLastArray(intSlice)
	fmt.Println("First and Last element of Array :=>", first, last)
}

func isPrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	sqrt := math.Sqrt(float64(num))

	for i := 3; i < int(sqrt); i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func IsPelindrome(num int) bool {
	if num != ReverseNumber(num) {
		return false
	}
	return true
}

func ReverseNumber(num int) int {
	rev := 0

	for num != 0 {
		rem := num % 10
		rev = num*10 + rem
		num = num / 10
	}

	return rev
}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func SquareRoot(num float64) float64 {
	return math.Sqrt(num)
}

func ReverseArray(arr []int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--
	}

	return arr
}

func SumOfArray(arr []int) int {
	var sum int

	for _, val := range arr {
		sum += val
	}

	return sum
}

func PrintArray(arr []int) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func PrintFirstAndLastArray(arr []int) (int, int) {
	return arr[0], arr[len(arr)-1]
}

func PrintLengthOfArray(arr []int) int {
	return len(arr)
}

func AdditionOfMatrix(matrix [3][3]int) int {
	sum := 0

	for _, val := range matrix {
		for _, data := range val {
			sum += data
		}
	}

	return sum

}

func MultiplicationOfMatrix(matrix [3][3]int) int {
	total := 1

	for _, val := range matrix {
		for _, data := range val {
			total *= data
		}
	}

	return total
}
