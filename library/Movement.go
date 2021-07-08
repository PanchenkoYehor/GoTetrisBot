package library

func RotateInField(field Field, left int, rotates int) Field {

	var temp = Block{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			temp[i][j] = field[i][j + left]
		}
	}

	temp = ApplyRotate(temp, rotates)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			field[i][j + left] = temp[i][j]
		}
	}

	return field
}

func ShiftInField(field Field, shift int) Field {
	var temp = [4][10]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j + shift >= 0 && j + shift < 10 {
				temp[i][j + shift] = field[i][j]
			}
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			field[i][j] = temp[i][j]
		}
	}

	return field

}

func Bounds(field Field) (int, int) {
	var right, left = 0, 9

	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {

			if field[i][j] == 1 {
				if j < left {
					left = j
				}
				if j > right {
					right = j
				}
			}

		}
	}

	return left, right
}

func DestroyLines(field Field) Field {
	for i := 23; i >= 4; i-- {
		var sum = 0
		for j := 0; j < 10; j++ {
			sum += field[i][j]
		}

		if sum == 10 {

			for ii := i; ii >= 4; ii-- {
				for j := 0; j < 10; j++ {
					field[ii][j] = field[ii - 1][j]
				}
			}

			for j := 0; j < 10; j++ {
				field[3][j] = 0
			}
			i++
		}
	}

	return field
}

func MoveDownUntilEnd(field Field) Field {

	for j := 0; j < 10; j++ {
		field[24][j] = 1
	}
	var mn = 200

	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {

			if field[i][j] == 1 {
				for ii := 4; ii < 25; ii++ {

					if field[ii][j] == 1 {
						if ii - i < mn {
							mn = ii - i
						}

					}
				}
			}

		}
	}

	for i := 3; i >= 0; i-- {
		for j := 0; j < 10; j++ {
			if field[i][j] == 1 {
				field[i + mn - 1][j] = 1
				field[i][j] = 0
			}
		}
	}

	field = DestroyLines(field)

	return field

}

