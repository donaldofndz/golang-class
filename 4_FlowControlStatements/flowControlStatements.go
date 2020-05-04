package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	forExplanation()
	ifExplanation()
	switchExplanation()
	deferExplanetion()
}

func forExplanation() {
	//classicFor()
	//forAsWhile()
	//inifinitLoop()
	//breakAndContinue()
}

func classicFor() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func forAsWhile() {
	var i int
	for i < 5 {
		fmt.Print(i)
		i = i + 1
	}
}

func inifinitLoop() {
	for {
		fmt.Println("loop")
		time.Sleep(2 * time.Second)
	}
}

func breakAndContinue() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 5 {
			break
		}
		println(i)

	}
}

func ifExplanation() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//The small statement con only be call inside the condition
	if num := 9; num < 0 {
		fmt.Println("The number is lower than 0")
	} else if num < 10 {
		fmt.Println("The number has one digit")
	} else {
		fmt.Println("The number has multiple digits")
	}

	//fmt.Println(num)
}

// The if function can have a small statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func switchExplanation() {
	//normalSwitch()
	//switchEvaluationOrder()
	switchWithoutCondition()
}

func normalSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%v.\n", os)
	}
}

func switchEvaluationOrder() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithoutCondition() {
	// this will stop in the first true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferExplanetion() {
	//simpleDefer()
	multiDefer()
}

func simpleDefer() {
	defer fmt.Println("World!")

	fmt.Println("Hello ")
}

func multiDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
