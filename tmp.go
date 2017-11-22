package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total/float64(len(xs))
}

func f() (int, int) {
	return 5,6
}

func main(){
	var total float64 = 0
	x := [5]float64{98, 93, 77, 82, 83}
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
