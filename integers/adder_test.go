package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d, but got %d", expected, sum)
	}
}

func ExampleAdder() {
	sum := Add(21, 21)
	fmt.Println(sum)
	// output: 42
}
