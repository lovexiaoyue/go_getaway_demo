package main

import "fmt"

func main() {
	num := []int{6,3,5,1,6,78,3,43,65,2,43,65,15,47,21,321,6}
	//var max int
	sum := 0

	var cur int
	for x:=0; x < len(num); x ++{
		min := num[x]
		for i:=0;i < len(num)-x ; i ++ {
			if min > num[i+x]{
				min = num[i+x]
				cur = i+x
			}
			sum ++
		}
		num[x],num[cur] = num[cur], num[x]
	}
	fmt.Println(sum)
	fmt.Println(num)
}
