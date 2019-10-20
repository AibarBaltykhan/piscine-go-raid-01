package student

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	x--
	y--
	leftup := 'A'
	leftbot := 'C'
	rightup := 'C'
	rightbot := 'A'
	vertical := 'B'
	middlsec := ' '
	midllfir := 'B'

	if (x < 0) || (y < 0) {

		return
	}
	for j := 0; j <= y; j++ {
		for i := 0; i <= x; i++ {

			if i > 0 && j > 0 && i < x && j < y {

				z01.PrintRune(middlsec)
			}

			if j > 0 && j < y && (i == 0 || i == x) {

				z01.PrintRune(vertical)
			}
			if i > 0 && i < x && (j == 0 || j == y) {

				z01.PrintRune(midllfir)
			}

			if i == 0 {
				if j == 0 {

					z01.PrintRune(leftup)
				} else if j == y {

					z01.PrintRune(leftbot)
				}

			} else if i == x {

				if j == 0 {
					z01.PrintRune(rightup)
				} else if j == y {
					z01.PrintRune(rightbot)
				}

			}
		}
		z01.PrintRune(10)

	}

}
