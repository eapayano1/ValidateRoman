package KataValidarRomano

import "testing"

func TestArcaismoOrtografico(t *testing.T) {
	actual := ValidarNumeroRomano("iX")
	expected := false
	if actual != expected {
		t.Errorf("Fail: actual %t | expected: %t", actual, expected)
	}
}
