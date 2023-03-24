package LRUCache

import "testing"

func TestLRU1(t *testing.T) {
	bank := NewLRUCache(3)
	bank.Add("A", "1")
	bank.Add("B", "2")
	bank.Add("C", "3")
	bank.Add("D", "4")

	val, _ := bank.Get("A")
	if val != "" {
		t.Fatalf("key A should be removed from cache")
	}
}

func TestLRU2(t *testing.T) {
	bank := NewLRUCache(3)
	bank.Add("A", "1")
	bank.Add("B", "2")
	bank.Add("C", "3")
	bank.Add("A", "4")

	val, _ := bank.Get("A")
	if val != "4" {
		t.Fatalf("A != 4")
	}
}

func TestLRU3(t *testing.T) {
	bank := NewLRUCache(3)
	bank.Add("A", "1")
	bank.Add("B", "2")
	bank.Add("C", "3")
	bank.Add("B", "4")

	val, _ := bank.Get("B")
	if val != "4" {
		t.Fatalf("B != 4")
	}
}

func TestLRU4(t *testing.T) {
	bank := NewLRUCache(3)
	bank.Add("A", "1")
	bank.Add("B", "2")
	bank.Add("C", "3")
	bank.Add("B", "4")

	val, _ := bank.Get("B")
	if val != "4" {
		t.Fatalf("B != 4")
	}

	bank.Remove("B")
	val, _ = bank.Get("B")
	if val != "" {
		t.Fatalf("key B should be removed from cache")
	}
	val, _ = bank.Get("C")
	if val != "3" {
		t.Fatalf("C != 3")
	}
	bank.Remove("A")
	val, _ = bank.Get("A")
	if val != "" {
		t.Fatalf("key A should be removed from cache")
	}
}

func TestLRU5(t *testing.T) {
	bank := NewLRUCache(3)
	bank.Add("A", "1")
	bank.Add("B", "2")
	bank.Add("C", "3")
	bank.Add("D", "4") // at this step A is removed

	val, _ := bank.Get("A")
	if val != "" {
		t.Fatalf("key A should be removed from cache")
	}

	bank.Remove("D") // D is removed and cache has B:2 C:3, C is head.
	val, _ = bank.Get("D")
	if val != "" {
		t.Fatalf("key D should be removed from cache")
	}

	val, _ = bank.Get("B") // cache has B:2 C:3, B is head.
	if val != "2" {
		t.Fatalf("B != 2")
	}

}
