package main

import "fmt"

func main() {
	fmt.Println(sqrtDir(float64(2), 1.0, &sqrtCalc{}))
	fmt.Println(sqrtRes(float64(2), 1.0, &sqrtCalc{}))
}

func sqrtDir(num float64, guess float64, calc SqrtCalcInterface) float64 {
	for {
		if calc.IsGoodEnough(num, guess) {
			break
		}
		guess = calc.Improve(num, guess)
	}

	return guess
}

func sqrtRes(num float64, guess float64, calc SqrtCalcInterface) float64 {
	if calc.IsGoodEnough(num, guess) {
		return guess
	}
	return sqrtRes(num, calc.Improve(num, guess), calc)
}

type SqrtCalcInterface interface {
	IsGoodEnough(float64, float64) bool
	Improve(float64, float64) float64
}

type sqrtCalc struct{}

func (s *sqrtCalc) IsGoodEnough(num float64, res float64) bool {
	diff := res*res - num
	if diff < 0 {
		diff = diff * -1
	}

	if diff < 0.1 {
		return true
	}
	return false
}

func (s *sqrtCalc) Improve(num float64, res float64) float64 {
	return (res + num/res) / 2
}
