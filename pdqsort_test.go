package pdqsort

import (
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkPdqsortSorted(b *testing.B) {
	var arr array
	n := 10000000
	for i := 0; i < n; i++ {
		arr.Append(i)
	}
	arr.Len()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdqsort(arr, 0, arr.Len(), 10)
	}
}

func BenchmarkSortSorted(b *testing.B) {
	var arr [10000000]int
	for i := 0; i < 10000000; i++ {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr[:])
	}
}

func BenchmarkPdqsortReversed(b *testing.B) {
	var arr array
	n := 10000000
	for i := 0; i < n; i++ {
		arr.Append(n - i)
	}
	arr.Len()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdqsort(arr, 0, arr.Len(), 10)
	}
}

func BenchmarkSortReversed(b *testing.B) {
	var arr [10000000]int
	for i := 0; i < 10000000; i++ {
		arr[i] = 10000000 - i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr[:])
	}
}

func BenchmarkPdqsortSame(b *testing.B) {
	var arr array
	n := 1000000
	for i := 0; i < n; i++ {
		arr.Append(rand.Intn(8))
	}
	arr.Len()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdqsort(arr, 0, arr.Len(), 10)
	}
}

func BenchmarkSortSame(b *testing.B) {
	var arr [1000000]int
	for i := 0; i < 1000000; i++ {
		arr[i] = rand.Intn(8)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr[:])
	}
}
func BenchmarkPdqsortShort(b *testing.B) {
	var arr array
	n := 12
	for i := 0; i < n; i++ {
		arr.Append(rand.Intn(100))
	}
	arr.Len()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdqsort(arr, 0, arr.Len(), 10)
	}
}

func BenchmarkSortShort(b *testing.B) {
	var arr [12]int
	for i := 0; i < 12; i++ {
		arr[i] = rand.Intn(100)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr[:])
	}
}


func BenchmarkPdqsortRand(b *testing.B) {
	var arr array
	n := 1000000
	for i := 0; i < n; i++ {
		arr.Append(rand.Intn(100000000))
	}
	arr.Len()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdqsort(arr, 0, arr.Len(), 10)
	}
}

func BenchmarkSortRand(b *testing.B) {
	var arr [1000000]int
	for i := 0; i < 1000000; i++ {
		arr[i] = rand.Intn(100000000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr[:])
	}
}