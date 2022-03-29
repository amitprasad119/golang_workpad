package add

import "testing"

func testAdd(t *testing.T) {
	result := addTwo(5,6)

	expected := 11
   
    if result != expected {
		t.Errorf("got %q, wanted %q ", result,expected)
	}
}


func addTwo(x, y int) int {
	return x + y
}