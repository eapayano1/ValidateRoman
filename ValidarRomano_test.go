package Kata3

import "testing"

func TestDetectarArcaismoOrtografico(t *testing.T) {
	actual := ValidarNumeroRomano("Xi")
	expected := false
	if actual != expected {
		t.Errorf("FAIL: actual: %t, expected: %t", actual, expected)
	}
}
