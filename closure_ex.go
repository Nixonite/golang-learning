package main

import "fmt"

func makeOddGenerator() (func() int) {
	var i int = 1
	return func() (ret int){
		ret = i
		i += 2
		return
	}
}

func main(){
	add := func(x, y int) int{
		return x + y
	}
	fmt.Println(add(1,2))

	oddGen := makeOddGenerator()
	fmt.Println(oddGen())
	fmt.Println(oddGen())
}
