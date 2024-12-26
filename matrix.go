// Package main provides matrix multiplication implementation
package main

import (
	"log"
	"math/rand"
	"time"
)

// Matrix generates two random matrices of size 1000x1000 for testing purposes.
// Returns two matrices (MatrixA, MatrixB) filled with random integers between 0 and 99.
// The matrices are of equal size to ensure they can be multiplied together.
func Matrix() ([][]int, [][]int) {
	// Initialize random seed with current time
	rand.Seed(time.Now().UnixNano())
	
	// Define matrix size
	const size = 1000
	
	// Initialize matrices with the given size
	MatrixA := make([][]int, size)
	MatrixB := make([][]int, size)
	
	// Fill matrices with random values
	for i := 0; i < size; i++ {
		MatrixA[i] = make([]int, size)
		MatrixB[i] = make([]int, size)
		for j := 0; j < size; j++ {
			MatrixA[i][j] = rand.Intn(100) // Random integers from 0 to 99
			MatrixB[i][j] = rand.Intn(100)
		}
	}
	log.Println("matrices generation completed")
	
	return MatrixA, MatrixB
}