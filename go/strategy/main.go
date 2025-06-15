package main

import (
	"fmt"

	"github.com/asunaro276/design-patterns/go/strategy/internal"
)


func main() {
	lfu := &cache.Lfu{}
	cache := cache.InitCache(lfu)
	cache.Add("a", "1")
	cache.Add("b", "2")
	fmt.Println(cache.Get("a"))
	fmt.Println(cache.Get("b"))
}
