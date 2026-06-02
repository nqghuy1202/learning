package main

import (
	"fmt"
	"math"
)

var companyName = "HL Company"
var currentYear int
var temp string

// to return 2 values use (varaiable type and error)

func bai2_sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Can't execute with number < 0")
	} else {
		return math.Sqrt(f), nil
	}
}

func checkSchedule(date string) string {
	switch date {
	case "Monday":
		return "Chest"
	case "Tuesday":
		return "Shoulders"
	case "Wednesday":
		return "Cardio"
	case "Thursday":
		return "Back"
	case "Friday":
		return "Off"
	case "Saturday":
		return ":Legs"
	case "Sunday":
		return "Cheat day"
	default:
		return "Relax"
	}
}

func findMax(a int, b int) (int, error) {
	if a > b {
		return a, nil
	} else if b > a {
		return b, nil
	} else {
		return 0, fmt.Errorf("Two numbers equals")
	}
}

func findMin(a int, b int) (int, error) {
	if a < b {
		return a, nil
	} else if b < a {
		return b, nil
	} else {
		return 0, fmt.Errorf("Two numbers equals")
	}
}

func demoArray() {
	cars := [3]string{"Toyota", "Hyundai", "Honda"}

	for index, car := range cars {
		fmt.Print(index, car)
	}
}

func demoArray2D() {
	langs := [][]string{{"C++", "C#", "C"}, {"NestJs", "NodeJs", "JavaScript"}, {"Oracle", "MySQL", "MongoDB"}}

	for _, v := range langs {
		for _, lang := range v {
			fmt.Print(lang)
		}
		fmt.Println()
	}
}

func demoSlice() {
	letters := []string{"a", "b", "c"}
	letters = append(letters, "d")
	length := len(letters)
	fmt.Println("Count: ", length)
	fmt.Println(letters[:2])
}

func main() {
	// Bai 1: Varaiable
	currentYear := 2026
	temp := "Hello World"

	s := fmt.Sprintf("%s was born in %d", companyName, currentYear)

	fmt.Println(s)

	fmt.Println(temp + "! " + s)
	// Bai 2: Codition
	var r, msg = bai2_sqrt(5)

	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Result: ", r)
	}

	var date string
	fmt.Print("Enter: ")
	_, err := fmt.Scanf("%s", &date)

	if err != nil {
		fmt.Println(err)
	} else {
		task := checkSchedule(date)
		fmt.Println(task)
	}

	var a, b int
	fmt.Print("Nhập số a: ")
	fmt.Scan(&a)
	fmt.Print("Nhập số b: ")
	fmt.Scan(&b)

	var max, err1 = findMax(a, b)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Max: ", max)
	}
	var min, err2 = findMin(a, b)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Min: ", min)
	}
	demoArray()
	demoArray2D()
	demoSlice()

}
