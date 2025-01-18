package hashmap

import "fmt"

type Node[K comparable, V any] struct {
	key   K
	value V
	next  *Node[K, V]
}

type HashMap[K comparable, V any] struct {
	buckets []*Node[K, V]
	size    int
}

func New[K comparable, V any](capacity int) *HashMap[K, V] {
	return &HashMap[K, V]{
		buckets: make([]*Node[K, V], capacity),
		size:    0,
	}
}

func (h *HashMap[K, V]) hash(key K) int {
	hash := fmt.Sprintf("%v", key)
	sum := 0
	for _, ch := range hash {
		sum += int(ch)
	}
	return sum % len(h.buckets)
}

func (h *HashMap[K, V]) Put(key K, value V) {
	index := h.hash(key)

	if h.buckets[index] == nil {
		h.buckets[index] = &Node[K, V]{key: key, value: value}
		h.size++
		return
	}

	current := h.buckets[index]
	for current != nil {
		if current.key == key {
			current.value = value
			return
		}
		if current.next == nil {
			break
		}
		current = current.next
	}

	current.next = &Node[K, V]{key: key, value: value}
	h.size++
}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	index := h.hash(key)
	current := h.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}

	var zero V
	return zero, false
}

func (h *HashMap[K, V]) Remove(key K) bool {
	index := h.hash(key)

	if h.buckets[index] == nil {
		return false
	}

	if h.buckets[index].key == key {
		h.buckets[index] = h.buckets[index].next
		h.size--
		return true
	}

	current := h.buckets[index]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			h.size--
			return true
		}
		current = current.next
	}

	return false
}

func (h *HashMap[K, V]) Size() int {
	return h.size
}

func (h *HashMap[K, V]) IsEmpty() bool {
	return h.size == 0
}

func (h *HashMap[K, V]) Clear() {
	h.buckets = make([]*Node[K, V], len(h.buckets))
	h.size = 0
}

func (h *HashMap[K, V]) Print() {
	for i, bucket := range h.buckets {
		if bucket != nil {
			fmt.Printf("Bucket %d: ", i)
			current := bucket
			for current != nil {
				fmt.Printf("[%v: %v] ", current.key, current.value)
				current = current.next
			}
			fmt.Println()
		}
	}
}
