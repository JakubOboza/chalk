package main

import (
	"fmt"
	"github.com/JakubOboza/chalk"
)

func main() {

	samples := []int{
		chalk.BLINK,
		chalk.BLACK,
		chalk.RED,
		chalk.GREEN,
		chalk.YELLOW,
		chalk.PURPLE,
		chalk.PINK,
		chalk.BLUE,
		chalk.WHITE,
		chalk.REDBG,
		chalk.GREENBG,
		chalk.YELLOWBG,
		chalk.PURPLEBG,
		chalk.PINKBG,
		chalk.BLUEBG,
		chalk.WHITEBG,
		chalk.LIGHTBLACK,
		chalk.LIGHTRED,
		chalk.LIGHTGREEN,
		chalk.LIGHTYELLOW,
		chalk.LIGHTPURPLE,
		chalk.LIGHTPINK,
		chalk.LIGHTBLUE,
		chalk.LIGHTWHITE,
		chalk.LIGHTBLACKBG,
		chalk.LIGHTREDBG,
		chalk.LIGHTGREENBG,
		chalk.LIGHTYELLOWBG,
		chalk.LIGHTPURPLEBG,
		chalk.LIGHTPINKBG,
		chalk.LIGHTBLUEBG,
		chalk.LIGHTWHITEBG,
	}

	for _, sample := range samples {
		fmt.Println(chalk.Color(sample, "I'm so fab now!!!"))
	}

	fmt.Println(chalk.Color(chalk.LIGHTBLUE, "I'm blue nananana..."))
	fmt.Printf("Normal %s Normal\n", chalk.Color(chalk.LIGHTRED, "RED"))

	yellow := chalk.ColorFunc(chalk.YELLOWBG)
	grey := chalk.ColorFunc(chalk.LIGHTBLACKBG)

	badImage := [][]int{
		[]int{0, 1, 1, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 0, 1, 0, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 0, 1, 0},
	}

	n := len(badImage)
	m := len(badImage[0])

	for i := 0; i < n; i++ {

		for j := 0; j < m; j++ {
			if badImage[i][j] == 1 {
				fmt.Print(yellow("  "))
			} else {
				fmt.Print(grey("  "))
			}
		}
		fmt.Println()
	}

	badfaceImage := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 1, 1, 0, 1, 1, 0, 0},
		[]int{0, 0, 1, 1, 0, 1, 1, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Println(chalk.Color(chalk.LIGHTBLUEBG, "                  "))

	n = len(badfaceImage)
	m = len(badfaceImage[0])

	for i := 0; i < n; i++ {

		for j := 0; j < m; j++ {
			if badfaceImage[i][j] == 1 {
				fmt.Print(yellow("  "))
			} else {
				fmt.Print(grey("  "))
			}
		}
		fmt.Println()
	}

}
