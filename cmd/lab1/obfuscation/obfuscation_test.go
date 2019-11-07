package obfuscation

import (
	"testing"
)

func TestApplyMaskDighits(t *testing.T) {
	src := "0123456789"
	trg := "9876543210"
	data := "785"
	expected := "214"
	actual := applyMask(data, src, trg)
	if actual != expected {
		t.Errorf("Mask applying was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestApplyMaskAlphabeth(t *testing.T) {
	src := "abcABC"
	trg := "CBAcba"
	data := "bCaaB"
	expected := "BaCCb"
	actual := applyMask(data, src, trg)
	if actual != expected {
		t.Errorf("Mask applying was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestApplyMaskMixed(t *testing.T) {
	src := "abcABC123"
	trg := "321CBAcba"
	data := "bC.a32 1<>aB"
	expected := "2A.3ab c<>3B"
	actual := applyMask(data, src, trg)
	if actual != expected {
		t.Errorf("Mask applying was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestObfuscateDeobfuscate(t *testing.T) {
	data := "bC.a32 1<>aB"
	obf := Obfuscate(data)
	deobf := Deobfuscate(obf)
	if data != deobf {
		t.Errorf("Deobfuscation was incorrect, got: %s, want: %s.", deobf, data)
	}
}
