package load_balance

import (
	"fmt"
	"testing"
)

func TestPolling(t *testing.T)  {
	res := &Polling{}
	res.Add("127.0.0.1:2001")
	res.Add("127.0.0.1:2002")
	res.Add("127.0.0.1:2003")
	res.Add("127.0.0.1:2004")
	res.Add("127.0.0.1:2005")
	res.Add("127.0.0.1:2006")

	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
	fmt.Println(res.Next())
}
