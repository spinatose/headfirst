package prose

import (
	"testing"
)

type testData struct {
	list []string
	want string 
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData {
		testData { nil, "" },
		testData { []string{"manzana"}, "manzana" },
		testData { []string{"manzanas", "bananas"}, "manzanas and bananas" },
		testData { []string{"manzanas", "bananas", "peras"}, "manzanas, bananas, and peras"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Error(errorString(test.list, got, test.want))
		}
	}
}
