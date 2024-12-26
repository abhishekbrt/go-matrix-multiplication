

# Parallel Matrix Multiplication in Go

This project implements both sequential and parallel matrix multiplication algorithms in Go, demonstrating the performance benefits of concurrent processing.

## Features

- Sequential matrix multiplication implementation
- Parallel matrix multiplication using worker pools
- Configurable number of Workers
- Performance comparison between sequential and parallel approaches
- Built-in timing measurements and speedup calculations


## Installation

```bash
git clone https://github.com/abhishekbrt/go-matrix-multiplication.git
cd go-matrix-multiplication
```

## Usage

Run the program with default settings (8 workers):
```bash
go run .
```

Specify custom number of workers:
```bash
go run . -workers 10
```

## Implementation Details

### Sequential Implementation
- Traditional three-loop matrix multiplication algorithm
- Time complexity: O(n³)

### Parallel Implementation
- Uses Go's concurrency features (goroutines and channels)
- Worker pool pattern for parallel processing
- Job distribution through channels

## Performance

The program automatically calculates and displays:
- Sequential execution time
- Parallel execution time
- Speedup ratio (Sequential time / Parallel time)


     		 		 	
	
     				   
	
     		 	   
	
