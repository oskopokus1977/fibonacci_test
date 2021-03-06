package main

import (
	"fmt"
	"time"
	"encoding/json"
)

func fibonacci(n int) (rez int) {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	x := 0
	y := 1
	rez = 0
	for i := 2; i < n; i++ {
		rez = x + y
		x = y
		y = rez
	}
	return rez
}

type FibonacciType struct {
	Numeral int
	Value   int
}

func main() {

	fibonacciData := FibonacciType{
		Numeral: 1,
		Value:   0,
	}

	const timeInterval = 10
	errorInputCount := 0
	correctInputCount := 0

	for correctInputCount < 10 && errorInputCount < 3 {
		time.Sleep(time.Second)
		fmt.Println("\nВведите ", fibonacciData.Numeral, " число последовательности: =>")
		//timer := time.NewTimer(time.Second * timeInterval)

		stop := make(chan bool)
		mark :=false

		go func() {
			fmt.Scan(&fibonacciData.Value)
			stop <-true
			mark=true
			if fibonacciData.Value == fibonacci(fibonacciData.Numeral) {
				fmt.Println("Верно!")
				correctInputCount++
			} else {
				fibonacciData.Value = fibonacci(fibonacciData.Numeral)
				fibEx1, _ := json.Marshal(fibonacciData)
				fmt.Println("Не верно!\n", string(fibEx1), "- верный результат")
				errorInputCount++
				correctInputCount = 0
			}
		}()
		go func() {
			time.Sleep(time.Second*5)
			//<-timer.C
			if mark==false {

			fibonacciData.Value = fibonacci(fibonacciData.Numeral)
			fibEx, _ := json.Marshal(fibonacciData)
			fmt.Println(string(fibEx))
			stop <- true
			}
		}()

		if <-stop {
			fibonacciData.Numeral++

		}
	}
}
