package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
)

var cache *lru.Cache[int, any]

func init() {
	cache, _ = lru.NewWithEvict(2,
		func(key int, value any) {
			fmt.Printf("Evicted: key=%v value=%v\n", key, value)
		},
	)
}

func main() {
	cache.Add(1, "a")
	cache.Add(2, "b")

	fmt.Println(cache.Get(1))

	cache.Add(3, "c")

	fmt.Println(cache.Get(2))
}
