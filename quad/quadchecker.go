package quad

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/01-edu/z01"
)

func PrintError() {
	s := "Not a quad function"

	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}

func PrintValue(value string) {
	z01.PrintRune('[')
	for i := 0; i < len(value); i++ {
		z01.PrintRune(rune(value[i]))
	}
	z01.PrintRune(']')
}

func PrintInt(n int) {
	res := []int{}

	for i := n; i > 0; i /= 10 {
		res = append(res, i%10)
	}

	z01.PrintRune('[')
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(rune(res[i]) + '0')
	}
	z01.PrintRune(']')
}

func QuadChecker() {
	input, _ := ioutil.ReadAll(os.Stdin)
	x, y := 0, 0

	for _, value := range input {
		if value == byte('\n') {
			y++
		} else {
			x++
		}
	}

	if x == 0 || y == 0 {
		PrintError()
		return
	}

	x = x / y
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	ans := false

	for _, value := range quads {
		cmd := exec.Command("./"+value, strconv.Itoa(x), strconv.Itoa(y))
		output, _ := cmd.Output()

		if string(input) == string(output) && ans {
			z01.PrintRune(' ')
			z01.PrintRune('|')
			z01.PrintRune('|')
			z01.PrintRune(' ')
			PrintValue(value)
			z01.PrintRune(' ')
			PrintInt(x)
			z01.PrintRune(' ')
			PrintInt(y)
		}

		if string(input) == string(output) && !ans {
			PrintValue(value)
			z01.PrintRune(' ')
			PrintInt(x)
			z01.PrintRune(' ')
			PrintInt(y)
			ans = true
		}
	}

	if !ans {
		PrintError()
		return
	}

	z01.PrintRune('\n')
}
