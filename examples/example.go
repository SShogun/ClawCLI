package main

import (
	"fmt"
	"sync"
)

// Calculator performs basic mathematical operations
type Calculator struct {
	mu    sync.Mutex
	count int
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{
		count: 0,
	}
}

// Add adds two numbers and increments the operation count
func (c *Calculator) Add(a, b int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
	return a + b
}

// Subtract subtracts b from a and increments the operation count
func (c *Calculator) Subtract(a, b int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
	return a - b
}

// Multiply multiplies two numbers and increments the operation count
func (c *Calculator) Multiply(a, b int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
	return a * b
}

// GetOperationCount returns the total number of operations performed
func (c *Calculator) GetOperationCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

// DivideWithError divides two numbers and returns an error if dividing by zero
func (c *Calculator) DivideWithError(a, b int) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	c.count++
	return a / b, nil
}

// Example function demonstrating concurrent operations
func ExampleConcurrentOperations() {
	calc := NewCalculator()
	var wg sync.WaitGroup

	// Spawn multiple goroutines to perform operations
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			calc.Add(num, num*2)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Printf("Total operations: %d\n", calc.GetOperationCount())
}

func main() {
	calc := NewCalculator()

	// Perform some operations
	fmt.Println("Addition: 10 + 5 =", calc.Add(10, 5))
	fmt.Println("Subtraction: 10 - 5 =", calc.Subtract(10, 5))
	fmt.Println("Multiplication: 10 * 5 =", calc.Multiply(10, 5))

	result, err := calc.DivideWithError(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Division: 10 / 2 =", result)
	}

	fmt.Printf("Total operations: %d\n", calc.GetOperationCount())

	// Run concurrent operations example
	ExampleConcurrentOperations()
}
