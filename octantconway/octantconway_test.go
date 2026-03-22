package octantconway_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/plutov/practice-go/octantconway"
)

func TestOctantConway(t *testing.T) {
	for _, test := range tests {
		conf := []byte(test.start)

		frametime := min(100*time.Millisecond, 1*time.Second/time.Duration(test.N))
		var numLines int
		if !testing.Short() {
			numLines = bytes.Count(conf, []byte("\n"))
			fmt.Println(string(conf))
		}

		for range test.N {

			if !testing.Short() {
				time.Sleep(frametime)
				fmt.Printf("\x1b[%dA\x1b[J\n", numLines+2) // clear last output
			}

			conf = octantconway.OctantConway(conf)

			if !testing.Short() {
				numLines = bytes.Count(conf, []byte("\n"))
				fmt.Println(string(conf))
			}
		}

		a := trimTrailingWhitespace([]byte(test.expected))
		b := trimTrailingWhitespace(conf)
		if !bytes.Equal(a, b) {
			t.Errorf("test failed\nexpected\n%sactual\n%s", a, b)
		}
	}
}

var regexpTrailingWhitespace = regexp.MustCompile(`[ ]+\n|\n*$`)

func trimTrailingWhitespace(s []byte) []byte {
	return regexpTrailingWhitespace.ReplaceAll(s, []byte("\n"))
}

var tests = []struct {
	start    string
	N        int
	expected string
}{
	{ // aircraft carrier
		start:    "𜴂𜴯",
		N:        1,
		expected: "𜴂𜴯",
	},
	{ // glider
		start: "" +
			"𜴩🯦",
		N: 20,
		expected: "" +
			"\n" +
			"  𜺠𜷏",
	},
	{ // anvil
		start: "" +
			"\n" +
			" 𜺠𜴧𜴧𜺣\n" +
			"  🮂▘𜴇",
		N: 43,
		expected: "" +
			"     ▀\n" +
			"\n" +
			"      🮂𜺨",
	},
	{ // diehard
		start: "" +
			"\n" +
			"   𜴣 𜴘𜴨",
		N:        130,
		expected: "",
	},
	{ // Achim's p144
		start: "" +
			"▝▘       𜵑𜶀  ▝▘\n" +
			"      𜺠𜴐𜶀𜺫𜺨\n" +
			"      𜵓 𜵅\n" +
			"    𜵑𜶀𜺫𜴁\n" +
			"🯧🯦  𜺫𜺨       🯧🯦",
		N: 144,
		expected: "" +
			"▝▘       𜵑𜶀  ▝▘\n" +
			"      𜺠𜴐𜶀𜺫𜺨    \n" +
			"      𜵓 𜵅      \n" +
			"    𜵑𜶀𜺫𜴁       \n" +
			"🯧🯦  𜺫𜺨       🯧🯦\n",
	},
}
