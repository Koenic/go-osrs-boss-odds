package monsters

import (
	"fmt"
	"math"
	"testing"
)

func TestBarrowsUniqueItemOdds(t *testing.T) {
	odds := BarrowsUniqueItemOdds(0, 1)

	expected := 7 * (1.0 / 102) * math.Pow((101.0/102), 6)

	fmt.Println(odds, expected)
	if odds != expected {
		t.Fatalf(`BarrowsUniqueItemOdds(0,1) = %f`, odds)
	}

	odds = BarrowsUniqueItemOdds(0, 2)

	expected = BinCoeff(7, 2) * math.Pow(1.0/102, 2) * math.Pow((101.0/102), 5)

	fmt.Println(odds, expected)
	if odds != expected {
		t.Fatalf(`BarrowsUniqueItemOdds(0,2) = %f`, odds)
	}
}
