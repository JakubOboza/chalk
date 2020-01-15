//Package chalk lets you easily color strings for output in console
package chalk

import "fmt"

//Consts to make use of lib much easier and intuitive esepcially with IDEs
const (
	BLINK         = 5
	BLACK         = 30
	RED           = 31
	GREEN         = 32
	YELLOW        = 33
	PURPLE        = 34
	PINK          = 35
	BLUE          = 36
	WHITE         = 37
	REDBG         = 41
	GREENBG       = 42
	YELLOWBG      = 43
	PURPLEBG      = 44
	PINKBG        = 45
	BLUEBG        = 46
	WHITEBG       = 47
	LIGHTBLACK    = 90
	LIGHTRED      = 91
	LIGHTGREEN    = 92
	LIGHTYELLOW   = 93
	LIGHTPURPLE   = 94
	LIGHTPINK     = 95
	LIGHTBLUE     = 96
	LIGHTWHITE    = 97
	LIGHTBLACKBG  = 100
	LIGHTREDBG    = 101
	LIGHTGREENBG  = 102
	LIGHTYELLOWBG = 103
	LIGHTPURPLEBG = 104
	LIGHTPINKBG   = 105
	LIGHTBLUEBG   = 106
	LIGHTWHITEBG  = 107
)

//Color is a basic package function that colorizes for terminal a string
func Color(color int, str string) string {
	return fmt.Sprintf("\033[%dm%s\033[m", color, str)
}

//ColorFunc produces a clousure with already bound color
func ColorFunc(color int) func(string) string {
	return func(str string) string {
		return Color(color, str)
	}
}
