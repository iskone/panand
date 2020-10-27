package cache

import (
	"fmt"
	"testing"
)

func TestNode_Add(t *testing.T) {
	n := &Node{
		Parent:    nil,
		Id:        "xsxsxs",
		Ttl:       0,
		Name:      "fffffff",
		ChildNode: nil,
	}
	n.Add(&Node{
		Parent:    nil,
		Id:        "xsxsxs",
		Ttl:       0,
		Name:      "abc",
		ChildNode: nil,
	})
	cache.Add(n)
	fmt.Println(cache)
	fmt.Println(n)
	fmt.Println(cache.Get())
	fmt.Println(cache.Get("fffffff", "abc"))
	fmt.Println(cache.Get("fffffff"))
	fmt.Println(cache.Get("sss"))
}
