package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsertSort(t *testing.T) {
	slices := []int{11, 2, 4, 5, 8, 7, 9}
	fmt.Println(slices)
	InsertSort(slices)
	fmt.Println(slices)
}

func BenchmarkInsertSort(b *testing.B) {
	var slicess []int
	for i := 0; i < N; i++ {
		slicess = append(slicess, rand.Intn(356))
	}
	for i := 0; i < b.N; i++ {
		InsertSort(slicess)
	}
}
