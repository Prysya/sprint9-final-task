package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		size    int
		message string
	}{
		{0, "пустой слайс"},
		{1, "один элемент"},
		{50, "несколько элементов"},
	}

	for _, tc := range testCases {
		result := generateRandomElements(tc.size)
		assert.Equal(t, tc.size, len(result), "generateRandomElements для %s", tc.message)
	}
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
		message  string
	}{
		{[]int{}, 0, "пустой слайс"},
		{[]int{5}, 5, "один элемент"},
		{[]int{1, 5, 3, 9, 2}, 9, "несколько элементов"},
		{[]int{1, math.MaxInt, 3, 5}, math.MaxInt, "с максимальным int"},
		{[]int{7, 7, 7, 7}, 7, "все элементы одинаковые"},
		{[]int{10, 5, 3, 8}, 10, "максимум в начале"},
		{[]int{1, 5, 3, 15}, 15, "максимум в конце"},
	}

	for _, tc := range testCases {
		result := maximum(tc.input)
		assert.Equal(t, tc.expected, result, "maximum для %s", tc.message)
	}
}

func TestMaxChunks(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
		message  string
	}{
		{[]int{}, 0, "пустой слайс"},
		{[]int{5}, 5, "один элемент"},
		{[]int{1, 5, 3, 9, 2}, 9, "несколько элементов"},
		{[]int{1, math.MaxInt, 3, 5}, math.MaxInt, "с максимальным int в центре"},
		{[]int{math.MaxInt, 1, 3, 5}, math.MaxInt, "с максимальным int в начале"},
		{[]int{1, 3, 5, math.MaxInt}, math.MaxInt, "с максимальным int в конце"},
		{[]int{7, 7, 7, 7}, 7, "все элементы одинаковые"},
		{[]int{10, 5, 3, 8}, 10, "максимум в начале"},
		{[]int{1, 5, 3, 15}, 15, "максимум в конце"},
	}

	for _, tc := range testCases {
		result := maximum(tc.input)
		assert.Equal(t, tc.expected, result, "maxChunks для %s", tc.message)
	}
}

func TestMaxChunksGeneratedList(t *testing.T) {
	testCases := []struct {
		count int
	}{
		{count: CHUNKS * 3},
		{count: CHUNKS * 10},
	}

	for _, tc := range testCases {
		data := make([]int, tc.count)
		for i := range tc.count {
			data[i] = i + 1
		}
		result := maxChunks(data)
		assert.Equal(t, tc.count, result, "maxChunks для %d элементов должен вернуть %d", tc.count, tc.count)
	}
}
