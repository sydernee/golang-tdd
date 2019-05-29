package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("Expected '%s' but got '%s'", expected, repeated)
		}
	}

	t.Run("Repeat 5 times the test", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
