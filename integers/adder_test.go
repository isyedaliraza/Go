package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("Wanted %d but got %d", want, got)
	}
}

func ExampleAdder() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
