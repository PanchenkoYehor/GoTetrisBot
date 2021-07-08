package library

import (
	"math"
)

var costPit = 26.9
var costWell = 15.8
var costPitInCol = 27.6
var costPitInRow = 30.1
var costTop = 25.6

func evaluate(field Field) float64 { //Returns how bad this field is

	var cost = 0.0
	var cntPit = 0

	/*for i := 0; i < 21; i++ {
		for j := 0; j < 10; j++ {
			if field[i][j] == 1 {
				continue
			}
			for ii := i + 1; ii < 21; ii++ {
				if field[ii][j] == 1 {
					cntPit++
					break
				}
			}
		}
	}*/

	for i := 23; i >= 4; i-- {
		for j := 0; j < 10; j++ {
			if field[i][j] == 1 {
				continue
			}
			for ii := i; ii >= 4; ii-- {
				if field[ii][j] == 1 {
					cntPit++
					break
				}
			}
		}
	}

	cost += costPit * float64(cntPit)

	var cntWell = 0
	for i := 23; i >= 4; i-- {
		for j := 0; j < 10; j++ {
			if field[i][j] == 1 {
				continue
			}
			if (j == 0 || field[i][j - 1] == 1) && (j + 1 == 10 || field[i][j + 1] == 1) {
				cntWell++
			}
		}
	}

	cost += costWell * float64(cntWell)

	var cntPitInCol = 0
	for i := 23; i >= 4; i-- {
		for j := 0; j < 10; j++ {
			if field[i][j] == 1 {
				continue
			}
			if i == 4 || field[i - 1][j] == 1 {
				cntPitInCol++
			}
			if i == 23 || field[i + 1][j] == 1 {
				cntPitInCol++
			}
		}
	}

	cost += costPitInCol * float64(cntPitInCol)

	var cntPitInRow = 0
	for i := 23; i >= 4; i-- {
		for j := 0; j < 10; j++ {

			if field[i][j] == 1 {
				continue
			}
			if j == 0 || field[i][j - 1] == 1 {
				cntPitInRow++
			}
			if j + 1 == 10 || field[i][j + 1] == 1 {
				cntPitInRow++
			}
		}
	}

	cost += costPitInRow * float64(cntPitInRow)

	var top = 0
	for i := 4; i < 24; i++ {
		var is1 = false
		for j := 0; j < 10; j++ {
			if field[i][j] == 1 {
				is1 = true
			}
		}
		if is1 {
			top = 24 - i
			break
		}
	}

	cost += costTop * float64(top)

	return cost
}

func FindBestWay(field Field, howLeft int, fig Block) (int, int) {

	var bestRotate, bestShift, bestScore = 0, 0, float64(math.MaxInt32) //Shift: -1 - to left, +1 - to right

	//Move figure in center
	var startShift = 3 - howLeft
	field = ShiftInField(field, startShift)

	for rotate := 0; rotate < 4; rotate++ {
		var newField = field
		newField = RotateInField(newField, 3, rotate)
		var leftBound, rightBound = Bounds(newField)

		for shift := -leftBound; shift < 10 - rightBound; shift++ {
			newField = ShiftInField(newField, shift)
			//Move down
			var fieldAfter = MoveDownUntilEnd(newField)
			//Calculate score of field
			var score = evaluate(fieldAfter)
			//Update bestScore
			if score < bestScore {
				bestScore = score
				bestRotate = rotate
				bestShift = shift
			}
			newField = ShiftInField(newField, -shift)
		}

	}

	return bestRotate, bestShift

}