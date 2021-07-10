package lib

import (
	"math/rand"
	"time"
)

type Field [25][10]int
type Block [4][4]int

var Figures = [7]Block{}

var NumberOfLines = 0

func SetFigures() {
	Figures[0] = Block{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}} // I
	Figures[1] = Block{{1, 0, 0, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}} // J
	Figures[2] = Block{{0, 0, 1, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}} // L
	Figures[3] = Block{{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}} // O
	Figures[4] = Block{{0, 1, 1, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}} // S
	Figures[5] = Block{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}} // T
	Figures[6] = Block{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}} // Z
}

func ApplyRotate(figure Block, rotate int) Block {
	for t := 0; t < rotate; t++ {
		var temp = Block{}
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				temp[3-j][i] = figure[i][j]
			}
		}
		figure = temp
	}
	return figure
}

func Diameter(figure Block) int {
	var mx = 0
	for i := 0; i < 4; i++ {
		var cnt = 0
		for j := 0; j < 4; j++ {
			cnt += figure[i][j]
		}
		if cnt > mx {
			mx = cnt
		}
	}
	return mx
}

func MoveLeft(figure Block) Block {
	var mnj = 4

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if figure[i][j] == 1 && j < mnj {
				mnj = j
			}
		}
	}

	var temp = Block{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if figure[i][j] == 1 {
				temp[i][j-mnj] = 1
			}
		}
	}

	return temp
}

func SetRandomFigureOnTheField(field Field) (Field, int, Block) {
	rand.Seed(time.Now().UTC().UnixNano())
	var randIndex = int(rand.Int31n(7))
	var randRotate = int(rand.Int31n(4))
	var figure = ApplyRotate(Figures[randIndex], randRotate)
	figure = MoveLeft(figure)
	var diameter = Diameter(figure)
	var left = int(rand.Int31n(int32(11 - diameter)))
	/*fmt.Println(figure)
	fmt.Println(diameter)
	fmt.Println(left)*/

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if left+j < 10 {
				field[i][left+j] = figure[i][j]
			}
		}
	}

	return field, left, figure
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
