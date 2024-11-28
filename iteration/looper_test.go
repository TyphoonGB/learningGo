package iteration

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	looped := Loop("x", 3)
	expected := "xxx"

	if looped != expected {
		t.Errorf("expected '%q' but got '%q'", expected, looped)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop("x", 3)
	}
}

func ExampleLoop() {
	looped := Loop("x", 6)
	fmt.Printf("%q", looped)
	// Output: "xxxxxx"
}
