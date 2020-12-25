package day25

import "testing"

func TestDetransform(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{5764801, 8},
		{17807724, 11},
	}

	for i, test := range tests {
		actual := detransform(7, test.in)
		if actual != test.want {
			t.Errorf("#%d (detransform): actual = %d, but want %d", i, actual, test.want)
		}
	}
}

func TestHandshake(t *testing.T) {
	actual, _ := handshake(8, 11, 5764801, 17807724)
	want := 14897079
	if actual != want {
		t.Errorf("#%d (handshake): actual = %d, but want %d", 0, actual, want)
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"5764801\n17807724", 14897079},
		{"12578151\n5051300\n", 296776},
	}

	for i, test := range tests {
		actual := part1(test.in)
		if actual != test.want {
			t.Errorf("#%d (part1): actual = %d, but want %d", i, actual, test.want)
		}
	}
}
