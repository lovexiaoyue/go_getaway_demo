package main

import "fmt"

func main() {
	num := 1
	for i:=0; i<=20; i++{
		if num >= 5{
			num = 0
		}
		fmt.Println(num%5)
		num++
	}

}
