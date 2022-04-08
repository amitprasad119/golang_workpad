package main

import (
	"testing"
)

type inputData struct {
	num1   int
	num2   int
	result int
}

func TestSubtract(t *testing.T) {
	inputs := []inputData{inputData{20, 10, 10},
		inputData{50, 60, -10}, {70, 40, 30}}
	for _, v := range inputs {
		finalResult := subTwo(v.num1, v.num2)
		if finalResult != v.result {
			t.Error("expected ", v.result, "got ", finalResult)
		}
	}

}
