package main

import "fmt"

func main() {
	fmt.Println("Enter day to run: ")

	// var then variable name then variable type
	var day string

	// Taking input from user
	fmt.Scanln(&day)
	switch day {
	case "1":
		fmt.Println(day1_1())
		fmt.Println(day1_2())
	case "2":
		fmt.Println(day2_1())
		fmt.Println(day2_2())
	case "3":
		fmt.Println(day3_1())
		fmt.Println(day3_2())
	case "4":
		fmt.Println(day4_1())
		fmt.Println(day4_2())
	case "5":
		fmt.Println(day5_1())
		fmt.Println(day5_2())
	case "6":
		fmt.Println(day6_1())
	case "7":
		fmt.Println(day7_1())
	case "8":
		fmt.Println(day8_1())
	case "9":
		fmt.Println(day9_1())
	case "10":
		fmt.Println(day10_1())
	case "11":
		fmt.Println(day11_1())
	case "12":
		fmt.Println(day12_1())
	case "13":
	default:
		fmt.Println(day13_1())
	}

}
