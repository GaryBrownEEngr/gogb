package gogb

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewWorkerPool(t *testing.T) {
	jobCount := 10

	results := make([]int, jobCount)

	wp := NewWorkerPool(3)

	for jobIndex := 0; jobIndex < jobCount; jobIndex++ {
		localIndex := jobIndex
		wp.AddTask(func() {
			sum := 0
			for i := 0; i <= localIndex; i++ {
				sum += i
			}

			results[localIndex] = sum
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
		})
	}

	wp.WaitForCompletion()
	//                0  1  2  3  4   5   6   7   8   9
	expected := []int{0, 1, 3, 6, 10, 15, 21, 28, 36, 45}
	require.Equal(t, expected, results)
}
