package main

import "fmt"

// 两两比较
func main() {
	//num := []int{6,3,5,1,6,78,3,43,65,2,43,65,15,47,21,321,6}
	num := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17}
	//var max int
	sum := 0
	for x:=0;x < len(num); x ++ {
		for i:=0;i < len(num)-x; i++ {
			//fmt.Println(i)
			if len(num)-1 > i {
				if num[i] > num[i+1]{
					num[i] , num[i+1]= num[i+1],num[i]
				}
			}
			fmt.Println(num)
			sum ++
		}
	}
	fmt.Println(sum)
}
