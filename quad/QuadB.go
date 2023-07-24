package quad

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				z01.PrintRune('/')
			} else if i == 1 && j == x {
				z01.PrintRune('\\')
			} else if i == y && j == 1 {
				z01.PrintRune('\\')
			} else if i == y && j == x {
				z01.PrintRune('/')
			} else if i == 1 || i == y || j == x || j == 1 {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
