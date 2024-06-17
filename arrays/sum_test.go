package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum 10 numbers", func(t *testing.T) {
		numbers := []int{5, 5, 4, 4, 3, 3, 2, 2, 1, 1}
		got := Sum(numbers)
		want := 30

		assertCorrectResult(t, got, want, numbers)
	})
	t.Run("sum 4 numbers", func(t *testing.T) {
		numbers := []int{2, 2, 2, 1}
		got := Sum(numbers)
		want := 7

		assertCorrectResult(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {
	numbers := []int{2, 3, 4}
	numbers2 := []int{5, 6, 7, 8}

	got := SumAll(numbers, numbers2)
	want := []int{9, 26}

	match := reflect.DeepEqual(got, want)
	if !match {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestSumAlTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		match := reflect.DeepEqual(got, want)
		if !match {
			t.Errorf("got %v but want %v", got, want)
		}
	}

	t.Run("2 slice test", func(t *testing.T) {
		numbers := []int{2, 3, 4}
		numbers2 := []int{5, 6, 7, 8}

		got := SumAllTails(numbers, numbers2)
		want := []int{7, 21}

		checkSums(t, got, want)
		t.Run("empty slice", func(t *testing.T) {
			numbers := []int{2, 3, 4}
			numbers2 := []int{}

			got := SumAllTails(numbers, numbers2)
			want := []int{7, 0}

			checkSums(t, got, want)
		})

	})
}

func assertCorrectResult(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but want %v given %v", got, want, numbers)
	}
}
