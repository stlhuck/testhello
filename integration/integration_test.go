package integration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectResult(t, got, want)
	})
	t.Run("repeat 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		assertCorrectResult(t, got, want)
	})
}

func assertCorrectResult(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
func ExampleRepeat() {
	got := Repeat("x", 5)
	fmt.Println(got)
	// Output: xxxxx
}
