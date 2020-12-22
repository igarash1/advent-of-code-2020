package day22

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			"Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10",
			306},
		{
			"Player 1:\n1\n10\n28\n29\n13\n11\n35\n7\n43\n8\n30\n25\n4\n5\n17\n32\n22\n39\n50\n46\n16\n26\n45\n38\n21\n\nPlayer 2:\n19\n40\n2\n12\n49\n23\n34\n47\n9\n14\n20\n24\n42\n37\n48\n44\n27\n6\n33\n18\n15\n3\n36\n41\n31\n",
			31314},
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
			"Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10",
			291},
		{
			"Player 1:\n1\n10\n28\n29\n13\n11\n35\n7\n43\n8\n30\n25\n4\n5\n17\n32\n22\n39\n50\n46\n16\n26\n45\n38\n21\n\nPlayer 2:\n19\n40\n2\n12\n49\n23\n34\n47\n9\n14\n20\n24\n42\n37\n48\n44\n27\n6\n33\n18\n15\n3\n36\n41\n31\n",
			32760},
	}

	for i, test := range tests {
		actual := part2(test.in)
		if actual != test.want {
			t.Errorf("#%d (part2): actual = %d, but want %d", i, actual, test.want)
		}
	}
}
