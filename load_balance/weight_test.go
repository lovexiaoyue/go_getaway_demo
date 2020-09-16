package load_balance

import (
	"fmt"
	"testing"
)

func TestWeightRoundBinBalance_Add(t *testing.T) {
	r := &WeightRoundBinBalance{}
	r.Add("127.0.0.1:2001","4")
	r.Add("127.0.0.1:2002","3")
	r.Add("127.0.0.1:2003","2")

	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())

}
