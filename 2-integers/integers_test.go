package integers

import "testing"
import "fmt"

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("\nExpected: %d\nGot: %d", expected, sum)
	}
}
