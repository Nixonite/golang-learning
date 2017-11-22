package main

import "fmt"

func main(){
	for i:=1; i<=50; i++{
		if i%5==0 && i%3==0 {
			fmt.Println("fizzbuzz: ", i)
		} else if i%3==0 {
			fmt.Println("fizz: ", i)
		} else if i%5==0 {
			fmt.Println("buzz: ", i)
		} else {
			fmt.Println(i)
		}
	}
}
