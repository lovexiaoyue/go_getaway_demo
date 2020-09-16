package main

import (
	"fmt"
	"strconv"
)

type HandlerFunc func(data, res string)

func (f HandlerFunc)ServerHttp(data,res int)  {
	//fmt.Println(string(data),res)
	f(strconv.Itoa(data),strconv.Itoa(res))
}

func main() {
	hf := HandlerFunc(Hello)
	hf.ServerHttp(1,2)
}

func Hello(data,res string)  {
	fmt.Println("Hello",data)
	fmt.Println("Hello",res)
}