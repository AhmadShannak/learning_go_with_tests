package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array of 5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers[:])
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
	t.Run("slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum All", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllNew(t *testing.T) {
	t.Run("Sum All New", func(t *testing.T) {
		got := SumAllNew([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAppend(t *testing.T) {
	got := SumAppend([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAppend(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	slice3 := []int{11, 12, 13, 14, 15}
	slice4 := []int{16, 17, 18, 19, 20}

	for i := 0; i < b.N; i++ {
		SumAppend(slice1, slice2, slice3, slice4)
	}
}

func BenchmarkSumNew(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	slice3 := []int{11, 12, 13, 14, 15}
	slice4 := []int{16, 17, 18, 19, 20}

	for i := 0; i < b.N; i++ {
		SumAllNew(slice1, slice2, slice3, slice4)
	}
}

func BenchmarkSum(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	slice3 := []int{11, 12, 13, 14, 15}
	slice4 := []int{16, 17, 18, 19, 20}

	for i := 0; i < b.N; i++ {
		SumAll(slice1, slice2, slice3, slice4)
	}
}

func TestSumTail(t *testing.T) {
	t.Run("Test Sum Tail", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}

		got := SumTail(slice1, slice2)
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Empty slice", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{9}

		got := SumTail(slice1, slice2)
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
