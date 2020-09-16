package load_balance

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T)  {
	res := NewConsistentHashBalance(10,nil)
	res.Add("127.0.0.1:2001")
	res.Add("127.0.0.1:2002")
	res.Add("127.0.0.1:2003")
	res.Add("127.0.0.1:2004")
	res.Add("127.0.0.1:2005")
	res.Add("127.0.0.1:2006")

	fmt.Println(res.Get("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(res.Get("http://127.0.0.1/f/getinfo1"))
	fmt.Println(res.Get("http://127.0.0.1:2004/base/getinfo"))
	fmt.Println(res.Get("http://127.0.0.1/bafse/getinfo1"))
	fmt.Println(res.Get("http://127.0.0.1:2003/base/getinfo1"))
	fmt.Println(res.Get("http://127.0.0.1:2004/base/getinfo"))
}
