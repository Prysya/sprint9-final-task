package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	result := generateRandomElements(0)
	assert.Equal(t, 0, len(result), "generateRandomElements(0) должен вернуть пустой слайс")

	size := 100
	result = generateRandomElements(size)
	require.Equal(t, size, len(result), "generateRandomElements(%d) должен вернуть слайс длины %d", size, size)
}

func TestMaximum(t *testing.T) {
	result := maximum([]int{})
	assert.Equal(t, 0, result, "maximum([]int{}) должен вернуть 0")

	result = maximum([]int{5})
	assert.Equal(t, 5, result, "maximum([5]) должен вернуть 5")

	result = maximum([]int{1, 5, 3, 9, 2})
	assert.Equal(t, 9, result, "maximum([1,5,3,9,2]) должен вернуть 9")

	result = maximum([]int{-1, -5, -3, -9})
	assert.Equal(t, -1, result, "maximum([-1,-5,-3,-9]) должен вернуть -1")

	result = maximum([]int{-5, 10, -3, 25, 0})
	assert.Equal(t, 25, result, "maximum([-5,10,-3,25,0]) должен вернуть 25")
}

func TestMaxChunks(t *testing.T) {
	result := maxChunks([]int{})
	assert.Equal(t, 0, result, "maxChunks([]int{}) должен вернуть 0")

	result = maxChunks([]int{42})
	assert.Equal(t, 42, result, "maxChunks([42]) должен вернуть 42")

	result = maxChunks([]int{1, 5, 3})
	assert.Equal(t, 5, result, "maxChunks([1,5,3]) должен вернуть 5")

	arrCount := CHUNKS * 3
	data := make([]int, arrCount)
	for i := range arrCount {
		data[i] = i + 1
	}
	result = maxChunks(data)
	assert.Equal(t, arrCount, result, "maxChunks(%v) должен вернуть %d", data, arrCount)

	arrCount = CHUNKS * 10
	data = make([]int, arrCount)
	for i := range arrCount {
		data[i] = i + 1
	}
	result = maxChunks(data)
	assert.Equal(t, arrCount, result, "maxChunks(%v) должен вернуть %d", data, arrCount)
}

func TestMaxChunksEdgeCases(t *testing.T) {
	data := []int{math.MaxInt, 1, 2, 3, 4, 5, 6, 7}
	result := maxChunks(data)
	assert.Equal(t, math.MaxInt, result, "maxChunks([MaxInt,1,2,3,4,5,6,7]) должен вернуть MaxInt")

	data = []int{1, 2, 3, 4, 5, 6, 7, math.MaxInt}
	result = maxChunks(data)
	assert.Equal(t, math.MaxInt, result, "maxChunks([1,2,3,4,5,6,7,MaxInt]) должен вернуть MaxInt")

	data = []int{math.MinInt, -1, -2, -3}
	result = maxChunks(data)
	assert.Equal(t, -1, result, "maxChunks([MinInt,-1,-2,-3]) должен вернуть -1")
}
