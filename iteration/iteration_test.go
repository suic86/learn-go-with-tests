package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat \"a\" five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectResult(t, expected, repeated)
	})
	t.Run("count = 0 returns an empty string", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectResult(t, expected, repeated)
	})
	t.Run("count < 0 returns an empty string", func(t *testing.T) {
		repeated := Repeat("a", -2)
		expected := ""
		assertCorrectResult(t, expected, repeated)
	})
}

func assertCorrectResult(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("b", 6))
	//Output: bbbbbb
}
