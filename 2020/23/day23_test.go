package day23

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			"389125467",
			67384529},
		{
			"364289715",
			98645732},
	}

	for i, test := range tests {
		actual := part1(test.in)
		if actual != test.want {
			t.Errorf("#%d (part1): actual = %d, but want %d", i, actual, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			"389125467",
			149245887792},
		{
			"364289715",
			689500518476},
	}

	for i, test := range tests {
		actual := part2(test.in)
		if actual != test.want {
			t.Errorf("#%d (part2): actual = %d, but want %d", i, actual, test.want)
		}
	}
}
