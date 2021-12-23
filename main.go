package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {

		screen.MoveTopLeft()
		currentTime := time.Now()

		clock := [...]placeholder{
			digits[currentTime.Hour()/10], digits[currentTime.Hour()%10],
			colon,
			digits[currentTime.Minute()/10], digits[currentTime.Minute()%10],
			colon,
			digits[currentTime.Second()/10], digits[currentTime.Second()%10],
		}

		for i := range zero {
			for _, digit := range clock {
				next := digit[i]
				if digit == colon && (currentTime.Second()%2 == 0) {
					next = "   "
				}
				fmt.Print(next + "  ")
			}
			fmt.Println()
		}

		fmt.Println()

		time.Sleep(time.Second)

	}

}
