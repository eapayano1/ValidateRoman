package KataValidarRomano

import "testing"

func TestArcaismoOrtografico(t *testing.T) {
	actual := ValidarNumeroRomano("iX")
	expected := false
	if actual != expected {
		t.Errorf("Fail: actual %t | expected: %t", actual, expected)
	}
}
func TestCaracterMenorIgualtres(t *testing.T) {
	actual := ValidarNumeroRomano("IIII")
	expected := false
	if actual != expected {
		t.Errorf("Fail: actual %t | expected: %t", actual, expected)
	}
}
func TestCaracteresParticularesMenor2(t *testing.T) {
	actual := ValidarNumeroRomano("VVX")
	expected := false
	if actual != expected {
		t.Errorf("Fail: actual %t | expected: %t", actual, expected)
	}
}
