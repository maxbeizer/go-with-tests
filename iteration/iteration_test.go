package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeat 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat 2", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 3)
	fmt.Println(result)
	// Output: bbb
}
