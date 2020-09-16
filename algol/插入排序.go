package main

import "fmt"

func main() {
	num := []int{6,3,5,1,6,78,3,43,65,2,43,65,15,47,21,321,6}
	sum := 0
	for j:=1; j < len(num); j ++ {
		current := num[j]
		preIndex := j - 1
		for i:= 1; i <= j ; i++ {
			if num[preIndex] > current && preIndex >= 0{
				num[preIndex + 1] = num[preIndex]
				preIndex --
			}
			sum++
		}
		num[preIndex+1] = current
		fmt.Println(num)
	}
	fmt.Println(sum)
}
