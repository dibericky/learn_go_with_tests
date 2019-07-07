package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {

	assertRepeatCharacter := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	}

	t.Run("repeat character a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertRepeatCharacter(t, repeated, expected)
	})

	t.Run("repeat character a 7 times", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"

		assertRepeatCharacter(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 4)
	fmt.Println(repeated)
	// Output: bbbb
}
