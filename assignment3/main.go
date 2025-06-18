package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//1
	fmt.Println("Enter elements for array:=>")
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

	PrintEvenNum(intSlice)
	PrintOddNum(intSlice)

	//2
	// fmt.Println("Enter elements for array:=>")
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

	// var searchElement int
	// fmt.Println("Enter search element:=>")
	// fmt.Scan(&searchElement)

	// index := LenearSearch(intSlice, searchElement)
	// fmt.Println("index is :=>", index)

	//3
	// fmt.Println("Enter elements for array:=>")
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

	// totalSum := SumOfArray(intSlice)
	// totalMultiplication := MultiplicationOfArray(intSlice)

	// fmt.Println("total sum is :=", totalSum)
	// fmt.Println("total multiplication is :=", totalMultiplication)

	//4
	// fmt.Println("Enter elements of slices:=")
	// reader := bufio.NewReader(os.Stdin)
	// strData, _ := reader.ReadString('\n')

	// sliStr := strings.Split(strData, " ")

	// intSlice := []int{}

	// for _, val := range sliStr {
	// 	intData, err := strconv.Atoi(strings.TrimSpace(val))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		intSlice = append(intSlice, intData)
	// 	}
	// }

	// fmt.Println("Enter element to search :=")
	// var searchElem int
	// fmt.Scan(&searchElem)

	// index := BinarySearch(intSlice, searchElem)

	// fmt.Println("sorted slice:=>", intSlice)
	// fmt.Println("element index is :=>", index)

	// 5
	// emp := map[string]interface{}{}

	// var employeeName string
	// var employeeID int

	// fmt.Println("Enter Employee ID:=>")
	// fmt.Scan(&employeeID)

	// if employeeID == -999 {
	// 	os.Exit(1)
	// }

	// if employeeID != 0 {
	// 	emp["employee_id"] = employeeID
	// }

	// fmt.Println("Enter Employee Name:=>")
	// fmt.Scan(&employeeName)

	// if employeeName != "" {
	// 	emp["employee_name"] = employeeName
	// }

	// fmt.Println("Employee Data:- ", emp)

}

func PrintOddNum(num []int) {
	for _, val := range num {
		if val%2 != 0 {
			fmt.Printf("Odd number %d \n", val)
		}
	}
}

func PrintEvenNum(num []int) {
	for _, val := range num {
		if val%2 == 0 {
			fmt.Printf("Event number %d \n", val)
		}
	}
}

func SearchElement(arr []int, element int) int {
	for i, val := range arr {
		if val == element {
			return i
		}
	}

	return -1
}

func SumOfArray(arr []int) int {
	var sum int

	for _, val := range arr {
		sum += val
	}

	return sum
}

func MultiplicationOfArray(arr []int) int {
	var sum int = 1

	for _, val := range arr {
		sum *= val
	}

	return sum
}

func SortSlice(sli []int) []int {
	sort.Ints(sli)
	return sli
}

func SortSliceStr(sli []string) []string {
	sort.Strings(sli)
	return sli
}

func LenearSearch(sli []int, element int) int {
	for i, val := range sli {
		if val == element {
			return i
		}
	}

	return -1
}

func BinarySearch(sli []int, element int) int {
	sort.Ints(sli)

	left, right := 0, len(sli)-1

	for left <= right {
		mid := left + (right-left)/2

		if element == sli[mid] {
			return mid
		} else if sli[mid] > element {
			left = mid - 1
		} else if sli[mid] < element {
			right = mid + 1
		}
	}

	return -1
}
