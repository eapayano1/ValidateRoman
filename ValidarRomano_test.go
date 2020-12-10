package Kata3

import "testing"

func TestValidarArcaismoOrtografico(t *testing.T) {
	actual := ValidarNumeroRomano("iXX")
	expected := false
	if actual != expected {
		t.Errorf("FAIL: actual: %t, expected: %t", actual, expected)
	}
}
