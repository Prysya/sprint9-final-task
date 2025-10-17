package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size == 0 {
		return []int{}
	}

	arr := make([]int, size)
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for idx := range size {
		arr[idx] = random.Int()
	}

	return arr
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := math.MinInt

	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup

	maxInChunkList := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i += 1 {
		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize

		if i == CHUNKS-1 {
			endIdx = len(data)
		}

		go func(idx, start, end int) {
			defer wg.Done()

			chunkMax := data[start]

			for j := start + 1; j < end; j++ {
				if data[j] > chunkMax {
					chunkMax = data[j]
				}
			}

			maxInChunkList[idx] = chunkMax

		}(i, startIdx, endIdx)
	}

	wg.Wait()

	return maximum(maxInChunkList)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(elements)
	duration := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, duration.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(elements)
	duration = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, duration.Microseconds())
}
