package caesar

import (
	"testing"
)

func TestEncode(t *testing.T) {
	msg := "Attack at dawn"
	t.Logf("testing message %q", msg)
	if Encode(msg) != "Dwwdfn dw gdzq" {
		t.Fail()
	}
}

func TestRepeatedEncode(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	orig := "Attack at dawn"
	msg := orig
	for i := 0; i < 1e6; i++ {
		msg = Encode(msg)
	}
	for i := 0; i < 1e6; i++ {
		msg = Decode(msg)
	}
	if msg != orig {
		t.Errorf("expected %q but got %q", orig, msg)
	}
}

func TestEncodeTable(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"abcxyz", "defabc"},
		{"ABCXYZ", "DEFABC"},
		{"1234567890", "4567890123"},
		{"!@#$%^&*()", "!@#$%^&*()"},
	}
	for _, test := range tests {
		result := Encode(test.in)
		if result != test.out {
			t.Errorf("encode %q: expected %q, got %q", test.in, test.out, result)
		}
	}
}

func TestCoder(t *testing.T) {
	coder := Coder{Key: 1, Ranges: []RuneRange{{'a', 'z'}}}
	t.Run("Encode", func(t *testing.T) {
		if coder.Encode("abc") != "bcd" {
			t.Fail()
		}
	})
	t.Run("Decode", func(t *testing.T) {
		if coder.Decode("bcd") != "abc" {
			t.Fail()
		}
	})
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode("attack at dawn")
	}
}
