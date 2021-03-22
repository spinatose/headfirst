package prose

import (
	"fmt"
	"testing"
)

const (
	twoFruitResult   string = "manzanas and bananas"
	threeFruitResult string = "manzanas, bananas, and peras"
)


func TestJoinWithCommasTwoElements(t *testing.T) {
	fruits := []string{"manzanas", "bananas"}
	want := twoFruitResult
	got := JoinWithCommas(fruits)
	if got != want {
		//t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", fruits, got, want)
		t.Error(errorString(fruits, got, want))
	}
}


func TestJoinWithCommasThreeElements(t *testing.T) {
	fruits := []string{"manzanas", "bananas", "peras"}
	want := threeFruitResult
	got := JoinWithCommas(fruits)
	if got != want {
		t.Error(errorString(fruits, got, want))
	}
}

func errorString (list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}