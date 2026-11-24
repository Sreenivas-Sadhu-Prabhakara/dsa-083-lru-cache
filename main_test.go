package main

import "testing"

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if v, ok := c.Get(1); !ok || v != 1 {
		t.Errorf("Get(1) = (%d,%v), want 1", v, ok)
	}
	c.Put(3, 3) // evicts 2 (1 was just used)
	if _, ok := c.Get(2); ok {
		t.Error("2 should have been evicted")
	}
	c.Put(4, 4) // evicts 1
	if _, ok := c.Get(1); ok {
		t.Error("1 should have been evicted")
	}
	if v, _ := c.Get(3); v != 3 {
		t.Errorf("Get(3) = %d, want 3", v)
	}
	if v, _ := c.Get(4); v != 4 {
		t.Errorf("Get(4) = %d, want 4", v)
	}
}

func TestUpdateRefreshes(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(1, 10) // update must refresh recency and value
	c.Put(3, 3)  // evicts 2, not 1
	if _, ok := c.Get(2); ok {
		t.Error("2 should have been evicted")
	}
	if v, _ := c.Get(1); v != 10 {
		t.Errorf("Get(1) = %d, want 10", v)
	}
}

func TestCapacityOne(t *testing.T) {
	c := NewLRUCache(1)
	c.Put(1, 1)
	c.Put(2, 2)
	if _, ok := c.Get(1); ok {
		t.Error("1 should have been evicted at capacity 1")
	}
	if v, _ := c.Get(2); v != 2 {
		t.Errorf("Get(2) = %d, want 2", v)
	}
}
