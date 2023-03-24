package main

import LRUCache "LRUCache/internal/cache"

func main() {
	c := LRUCache.NewLRUCache(3)

	c.Add("A", "1")
	c.Add("B", "2")
	c.Add("C", "3")
	c.Add("D", "4")

	c.Get("A")

	c.Remove("D")

	c.Get("B")
}
