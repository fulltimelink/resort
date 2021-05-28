package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

const N = 10

func TestBubbleSort(t *testing.T) {
	slices := []int{11, 2, 4, 5, 8, 7, 9}
	fmt.Println(slices)
	BubbleSort(slices)
	fmt.Println(slices)
}

func TestBubbleSortFor(t *testing.T) {
	slices := []int{11, 2, 4, 5, 8, 7, 9}
	fmt.Println(slices)
	BubbleSortFor(slices)
	fmt.Println(slices)
}

func BenchmarkBubbleSort(b *testing.B) {
	var slicess []int
	for i := 0; i < N; i++ {
		slicess = append(slicess, rand.Intn(356))
	}
	for i := 0; i < b.N; i++ {
		BubbleSort(slicess)
	}
}

func BenchmarkBubbleSortFor(b *testing.B) {
	var slicess []int
	for i := 0; i < N; i++ {
		slicess = append(slicess, rand.Intn(356))
	}
	for i := 0; i < b.N; i++ {
		BubbleSortFor(slicess)
	}
}
