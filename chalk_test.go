package chalk

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {

	tests := []struct {
		color  int
		result string
	}{
		{color: BLINK, result: "\033[5mWORD\033[m"},
		{color: BLACK, result: "\033[30mWORD\033[m"},
		{color: RED, result: "\033[31mWORD\033[m"},
		{color: GREEN, result: "\033[32mWORD\033[m"},
		{color: YELLOW, result: "\033[33mWORD\033[m"},
		{color: PURPLE, result: "\033[34mWORD\033[m"},
		{color: PINK, result: "\033[35mWORD\033[m"},
		{color: BLUE, result: "\033[36mWORD\033[m"},
		{color: WHITE, result: "\033[37mWORD\033[m"},
		{color: REDBG, result: "\033[41mWORD\033[m"},
		{color: GREENBG, result: "\033[42mWORD\033[m"},
		{color: YELLOWBG, result: "\033[43mWORD\033[m"},
		{color: PURPLEBG, result: "\033[44mWORD\033[m"},
		{color: PINKBG, result: "\033[45mWORD\033[m"},
		{color: BLUEBG, result: "\033[46mWORD\033[m"},
		{color: WHITEBG, result: "\033[47mWORD\033[m"},
		{color: LIGHTBLACK, result: "\033[90mWORD\033[m"},
		{color: LIGHTRED, result: "\033[91mWORD\033[m"},
		{color: LIGHTGREEN, result: "\033[92mWORD\033[m"},
		{color: LIGHTYELLOW, result: "\033[93mWORD\033[m"},
		{color: LIGHTPURPLE, result: "\033[94mWORD\033[m"},
		{color: LIGHTPINK, result: "\033[95mWORD\033[m"},
		{color: LIGHTBLUE, result: "\033[96mWORD\033[m"},
		{color: LIGHTWHITE, result: "\033[97mWORD\033[m"},
		{color: LIGHTBLACKBG, result: "\033[100mWORD\033[m"},
		{color: LIGHTREDBG, result: "\033[101mWORD\033[m"},
		{color: LIGHTGREENBG, result: "\033[102mWORD\033[m"},
		{color: LIGHTYELLOWBG, result: "\033[103mWORD\033[m"},
		{color: LIGHTPURPLEBG, result: "\033[104mWORD\033[m"},
		{color: LIGHTPINKBG, result: "\033[105mWORD\033[m"},
		{color: LIGHTBLUEBG, result: "\033[106mWORD\033[m"},
		{color: LIGHTWHITEBG, result: "\033[107mWORD\033[m"},
	}

	for _, test := range tests {

		got := Color(test.color, "WORD")
		want := test.result
		//Just to show colors for person running tests.
		fmt.Printf("%s\n", got)
		if got != want {
			t.Errorf("Got %s; Want %s", got, want)
		}

	}

}

func TestColorFunc(t *testing.T) {

	tests := []struct {
		color  int
		result string
	}{
		{color: BLINK, result: "\033[5mWORD\033[m"},
		{color: BLACK, result: "\033[30mWORD\033[m"},
		{color: RED, result: "\033[31mWORD\033[m"},
		{color: GREEN, result: "\033[32mWORD\033[m"},
		{color: YELLOW, result: "\033[33mWORD\033[m"},
		{color: PURPLE, result: "\033[34mWORD\033[m"},
		{color: PINK, result: "\033[35mWORD\033[m"},
		{color: BLUE, result: "\033[36mWORD\033[m"},
		{color: WHITE, result: "\033[37mWORD\033[m"},
		{color: REDBG, result: "\033[41mWORD\033[m"},
		{color: GREENBG, result: "\033[42mWORD\033[m"},
		{color: YELLOWBG, result: "\033[43mWORD\033[m"},
		{color: PURPLEBG, result: "\033[44mWORD\033[m"},
		{color: PINKBG, result: "\033[45mWORD\033[m"},
		{color: BLUEBG, result: "\033[46mWORD\033[m"},
		{color: WHITEBG, result: "\033[47mWORD\033[m"},
		{color: LIGHTBLACK, result: "\033[90mWORD\033[m"},
		{color: LIGHTRED, result: "\033[91mWORD\033[m"},
		{color: LIGHTGREEN, result: "\033[92mWORD\033[m"},
		{color: LIGHTYELLOW, result: "\033[93mWORD\033[m"},
		{color: LIGHTPURPLE, result: "\033[94mWORD\033[m"},
		{color: LIGHTPINK, result: "\033[95mWORD\033[m"},
		{color: LIGHTBLUE, result: "\033[96mWORD\033[m"},
		{color: LIGHTWHITE, result: "\033[97mWORD\033[m"},
		{color: LIGHTBLACKBG, result: "\033[100mWORD\033[m"},
		{color: LIGHTREDBG, result: "\033[101mWORD\033[m"},
		{color: LIGHTGREENBG, result: "\033[102mWORD\033[m"},
		{color: LIGHTYELLOWBG, result: "\033[103mWORD\033[m"},
		{color: LIGHTPURPLEBG, result: "\033[104mWORD\033[m"},
		{color: LIGHTPINKBG, result: "\033[105mWORD\033[m"},
		{color: LIGHTBLUEBG, result: "\033[106mWORD\033[m"},
		{color: LIGHTWHITEBG, result: "\033[107mWORD\033[m"},
	}

	for _, test := range tests {

		foo := ColorFunc(test.color)
		got := foo("WORD")
		want := test.result
		if got != want {
			t.Errorf("Got %s; Want %s", got, want)
		}

	}

}
