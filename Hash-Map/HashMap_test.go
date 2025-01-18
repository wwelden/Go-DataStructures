package hashmap

import (
    "testing"
)

func TestHashMapCreation(t *testing.T) {
    h := New[string, int](10)
    if !h.IsEmpty() {
        t.Error("New HashMap should be empty")
    }
    if h.Size() != 0 {
        t.Error("New HashMap should have size 0")
    }
}

func TestPutAndGet(t *testing.T) {
    h := New[string, int](10)

    // Test putting and getting single value
    h.Put("one", 1)
    if value, exists := h.Get("one"); !exists || value != 1 {
        t.Error("Failed to get value after Put")
    }

    // Test updating existing value
    h.Put("one", 2)
    if value, exists := h.Get("one"); !exists || value != 2 {
        t.Error("Failed to update existing value")
    }

    // Test getting non-existent value
    if _, exists := h.Get("two"); exists {
        t.Error("Get should return false for non-existent key")
    }

    // Test multiple values
    testData := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    for k, v := range testData {
        h.Put(k, v)
    }

    for k, v := range testData {
        if value, exists := h.Get(k); !exists || value != v {
            t.Errorf("Expected value %v for key %v, got %v", v, k, value)
        }
    }
}

func TestRemove(t *testing.T) {
    h := New[string, int](10)

    // Test remove on empty map
    if h.Remove("one") {
        t.Error("Remove should return false on empty map")
    }

    // Test remove existing value
    h.Put("one", 1)
    if !h.Remove("one") {
        t.Error("Remove should return true for existing key")
    }
    if _, exists := h.Get("one"); exists {
        t.Error("Value should not exist after removal")
    }

    // Test remove with multiple values
    h.Put("one", 1)
    h.Put("two", 2)
    h.Put("three", 3)

    h.Remove("two")
    if _, exists := h.Get("two"); exists {
        t.Error("Value should not exist after removal")
    }
    if value, exists := h.Get("one"); !exists || value != 1 {
        t.Error("Remove should not affect other values")
    }
}

func TestClear(t *testing.T) {
    h := New[string, int](10)
    h.Put("one", 1)
    h.Put("two", 2)

    h.Clear()
    if !h.IsEmpty() {
        t.Error("HashMap should be empty after Clear")
    }
    if h.Size() != 0 {
        t.Error("HashMap size should be 0 after Clear")
    }
}

func TestWithDifferentTypes(t *testing.T) {
    // Test with int keys
    h1 := New[int, string](10)
    h1.Put(1, "one")
    if value, exists := h1.Get(1); !exists || value != "one" {
        t.Error("Failed with int keys")
    }

    // Test with float values
    h2 := New[string, float64](10)
    h2.Put("pi", 3.14)
    if value, exists := h2.Get("pi"); !exists || value != 3.14 {
        t.Error("Failed with float64 values")
    }
}

func TestCollisions(t *testing.T) {
    // Create a small HashMap to force collisions
    h := New[string, int](1)

    // These will definitely collide as there's only one bucket
    h.Put("one", 1)
    h.Put("two", 2)
    h.Put("three", 3)

    // Verify all values are still accessible
    testData := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    for k, v := range testData {
        if value, exists := h.Get(k); !exists || value != v {
            t.Errorf("Expected value %v for key %v, got %v", v, k, value)
        }
    }
}