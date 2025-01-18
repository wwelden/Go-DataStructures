package main

import (
	"fmt"
	"sort"
	"time"

	BST "github.com/wwelden/Go-DataStructures/BST"
	HashMap "github.com/wwelden/Go-DataStructures/Hash-Map"
	LinkedList "github.com/wwelden/Go-DataStructures/Linked-List"
	QueueArr "github.com/wwelden/Go-DataStructures/Queue-Arr"
	QueueLL "github.com/wwelden/Go-DataStructures/Queue-LL"
	StackArr "github.com/wwelden/Go-DataStructures/Stack-Arr"
	StackLL "github.com/wwelden/Go-DataStructures/Stack-LL"
)

func main() {
	fmt.Println("\n=== Linked List ===")
	ll := LinkedList.LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Print()

	fmt.Println("\n=== Stack (Array-based) ===")
	stackArr := StackArr.Stack[int]{}
	stackArr.Push(1)
	stackArr.Push(2)
	stackArr.Push(3)
	stackArr.Print()

	fmt.Println("\n=== Stack (Linked List-based) ===")
	stackLL := StackLL.Stack[int]{}
	stackLL.Push(1)
	stackLL.Push(2)
	stackLL.Push(3)
	stackLL.Print()

	fmt.Println("\n=== Queue (Array-based) ===")
	queueArr := QueueArr.Queue[int]{}
	queueArr.Enqueue(1)
	queueArr.Enqueue(2)
	queueArr.Enqueue(3)
	queueArr.Print()

	fmt.Println("\n=== Queue (Linked List-based) ===")
	queueLL := QueueLL.Queue[int]{}
	queueLL.Enqueue(1)
	queueLL.Enqueue(2)
	queueLL.Enqueue(3)
	queueLL.Print()

	fmt.Println("\n=== Binary Search Tree ===")
	bst := BST.BST{}
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)
	bst.Print()

	fmt.Println("\n=== Hash Map ===")
	hm := HashMap.New[string, int](10)
	hm.Put("one", 1)
	hm.Put("two", 2)
	hm.Put("three", 3)
	hm.Print()

	runBenchmarks()
}

func runBenchmarks() {
	// fmt.Println("\n=== Performance Benchmarks ===")

	// Test parameters
	n := 100000 // Number of operations

	// Create a slice to store benchmark results
	type BenchmarkResult struct {
		operation string
		duration  time.Duration
	}
	var results []BenchmarkResult

	// Linked List Benchmark
	start := time.Now()
	ll := LinkedList.LinkedList[int]{}
	for i := 0; i < n; i++ {
		ll.Add(i)
	}
	results = append(results, BenchmarkResult{"Linked List Add", time.Since(start)})

	// Stack (Array-based) Benchmark
	start = time.Now()
	stackArr := StackArr.Stack[int]{}
	for i := 0; i < n; i++ {
		stackArr.Push(i)
	}
	results = append(results, BenchmarkResult{"Stack (Array) Push", time.Since(start)})

	// Stack (Linked List-based) Benchmark
	start = time.Now()
	stackLL := StackLL.Stack[int]{}
	for i := 0; i < n; i++ {
		stackLL.Push(i)
	}
	results = append(results, BenchmarkResult{"Stack (LL) Push", time.Since(start)})

	// Queue (Array-based) Benchmark
	start = time.Now()
	queueArr := QueueArr.Queue[int]{}
	for i := 0; i < n; i++ {
		queueArr.Enqueue(i)
	}
	results = append(results, BenchmarkResult{"Queue (Array) Enqueue", time.Since(start)})

	// Queue (Linked List-based) Benchmark
	start = time.Now()
	queueLL := QueueLL.Queue[int]{}
	for i := 0; i < n; i++ {
		queueLL.Enqueue(i)
	}
	results = append(results, BenchmarkResult{"Queue (LL) Enqueue", time.Since(start)})

	// BST Benchmark
	start = time.Now()
	bst := BST.BST{}
	for i := 0; i < n; i++ {
		bst.Insert(i)
	}
	results = append(results, BenchmarkResult{"BST Insert", time.Since(start)})

	// HashMap Benchmark
	start = time.Now()
	hm := HashMap.New[string, int](n)
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key%d", i)
		hm.Put(key, i)
	}
	results = append(results, BenchmarkResult{"HashMap Put", time.Since(start)})

	// Search Operations
	// fmt.Println("\n=== Search Operations ===")
	var searchResults []BenchmarkResult

	// BST Search
	start = time.Now()
	for i := 0; i < n; i++ {
		bst.Search(i)
	}
	searchResults = append(searchResults, BenchmarkResult{"BST Search", time.Since(start)})

	// HashMap Get
	start = time.Now()
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key%d", i)
		hm.Get(key)
	}
	searchResults = append(searchResults, BenchmarkResult{"HashMap Get", time.Since(start)})

	// Sort and display insertion results
	fmt.Println("\nInsertion Operations (Ranked):")
	fmt.Printf("Testing with %d operations\n", n)
	fmt.Println("-----------------------------------")
	sort.Slice(results, func(i, j int) bool {
		return results[i].duration < results[j].duration
	})
	for i, result := range results {
		fmt.Printf("#%d: %-20s %v\n", i+1, result.operation, result.duration)
	}

	// Sort and display search results
	fmt.Println("\nSearch Operations (Ranked):")
	fmt.Printf("Testing with %d operations\n", n)
	fmt.Println("-----------------------------------")
	sort.Slice(searchResults, func(i, j int) bool {
		return searchResults[i].duration < searchResults[j].duration
	})
	for i, result := range searchResults {
		fmt.Printf("#%d: %-20s %v\n", i+1, result.operation, result.duration)
	}
}
