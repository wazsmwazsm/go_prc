package strev 

import (
	"testing"
)

type ReverseTest struct {
	in, out string
}

var ReverseTests = []ReverseTest{
	{"ABCD", "DCBA"},
	{"CVO-AZ", "ZA-OVC"},
	{"Hello 世界", "界世 olleH"},
}


func TestReverse(t *testing.T) {
	for _, v := range ReverseTests {
		exp := Reverse(v.in)
		if v.out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", v.in, exp, v.out)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "ABCD"
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}
