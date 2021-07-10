package lib

import (
	"net/http"
	"time"
)

var Pane = Field{}

var delay = time.Millisecond * 50

func Calculate() {

	SetFigures()

	for i := 0; i < 24; i++ {
		for j := 0; j < 10; j++ {
			Pane[i][j] = 0
		}
	}

	for {
		var howLeft int
		var fig Block
		Pane, howLeft, fig = SetRandomFigureOnTheField(Pane)
		var rotate, shift = FindBestWay(Pane, howLeft, fig)

		//Pane = ShiftInField(Pane, 3 - howLeft)
		for t := 0; t < Abs(3-howLeft); t++ {
			Pane = ShiftInFieldByOne(Pane, 3-howLeft)
			callServer()
			time.Sleep(delay)
		}

		//Pane = RotateInField(Pane, 3, rotate)
		for t := 0; t < rotate; t++ {
			Pane = RotateInField(Pane, 3, 1)
			callServer()
			time.Sleep(delay)
		}

		//Pane = ShiftInField(Pane, shift)
		for t := 0; t < Abs(shift); t++ {
			Pane = ShiftInFieldByOne(Pane, shift)
			callServer()
			time.Sleep(delay)
		}

		//Pane = MoveDownUntilEnd(Pane)
		cnt := MoveDownUntilEndCount(Pane)

		for i := 0; i < 4; i++ {
			for j := 0; j < 10; j++ {
				if Pane[i][j] == 1 {
					Pane[i][j] = 2
				}
			}
		}

		for t := 0; t < cnt; t++ {
			Pane = MoveByOne(Pane)
			callServer()
			time.Sleep(delay)
		}

		for i := 0; i < 24; i++ {
			for j := 0; j < 10; j++ {
				if Pane[i][j] == 2 {
					Pane[i][j] = 1
				}
			}
		}

		Pane = DestroyLines(Pane, true)
		callServer()
		time.Sleep(delay)

		//Visualize(Pane)
		/*fmt.Println("after")
		rectImage := image.NewRGBA(image.Rect(0, 0, 200, 200))
		green := color.RGBA{0, 100, 0, 255}

		draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

		callServer()
		time.Sleep(1 * time.Second)*/
		/*callServer()
		time.Sleep(delay)*/
	}
}

func callServer() {
	http.Get("http://localhost:8181/")
}
