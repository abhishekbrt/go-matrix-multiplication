package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

// order struct to hold row and column indices
type order struct {
	row int
	col int
}

func main() {
	// Parse command-line flags
	worker := flag.Int("worker", 8, "Number of default worker")
	flag.Parse()
	log.Printf("number of worker used: %d\n", *worker)

	// Initialize matrices
	matrixA, matrixB := Matrix()

	// Sequential multiplication with timing
	start := time.Now()
	res := matrixMultiplication(matrixA, matrixB)
	sequential := time.Since(start)
	fmt.Printf("Sequential result: %v\n", res)

	// Parallel multiplication with timing
	var wg sync.WaitGroup
	numWorkers := *worker

	// Channel for sending data to workers
	jobs := make(chan order, 15)

	// Initialize result matrix
	resultRow := len(matrixA)
	resultCol := len(matrixB[0])
	result := make([][]int, resultRow)
	for i := range result {
		result[i] = make([]int, resultCol)
	}

	// Start worker pool
	start = time.Now()
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go multiplyWorkers(&wg, jobs, &result, matrixA, matrixB)
	}

	// Send data to the channel
	go func() {
		for i := 0; i < resultRow; i++ {
			for j := 0; j < resultCol; j++ {
				orders := order{
					row: i,
					col: j,
				}
				jobs <- orders
			}
		}
		close(jobs)
	}()

	// Wait for all workers to finish
	wg.Wait()
	parallel := time.Since(start)

	// Print results and speedup
	fmt.Printf("Parallel result: %v\n", result)
	fmt.Printf("Sequential time: %v\n", sequential)
	fmt.Printf("Parallel time: %v\n", parallel)
	fmt.Printf("Speedup: %.2fx\n", float64(sequential)/float64(parallel))
}

// Worker function to perform matrix multiplication
func multiplyWorkers(wg *sync.WaitGroup, jobs <-chan order, result *[][]int, matA, matB [][]int) {
	defer wg.Done()

	for job := range jobs {
		sum := 0
		for k := 0; k < len(matA[0]); k++ {
			sum += matA[job.row][k] * matB[k][job.col]
		}
		(*result)[job.row][job.col] = sum
	}
}

// Function to perform sequential matrix multiplication
func matrixMultiplication(matA [][]int, matB [][]int) [][]int {
	if len(matA[0]) != len(matB) {
		fmt.Println("matrix multiplication cannot be performed")
		return nil
	}

	result := make([][]int, len(matA))
	for i := range result {
		result[i] = make([]int, len(matB[0]))
	}

	for i := 0; i < len(matA); i++ {
		for j := 0; j < len(matB[0]); j++ {
			for k := 0; k < len(matB); k++ {
				result[i][j] += matA[i][k] * matB[k][j]
			}
		}
	}
	return result
}
