package lib

import (
	"fmt"
	"net/http"
	"time"
)

type ViewData struct {
	Title string
	Users []string
}

var Pane = Field{}

func Calculate() {

	SetFigures()

	for i := 0; i < 24; i++ {
		for j := 0; j < 10; j++ {
			Pane[i][j] = 0
		}
	}

	for t := 0; t <= 10; t++ {
		var howLeft int
		var fig Block
		Pane, howLeft, fig = SetRandomFigureOnTheField(Pane)
		var rotate, shift = FindBestWay(Pane, howLeft, fig)
		Pane = ShiftInField(Pane, 3-howLeft)
		Pane = RotateInField(Pane, 3, rotate)
		Pane = ShiftInField(Pane, shift)
		Pane = MoveDownUntilEnd(Pane)

		/*for {
			//MoveDown
			//Visualize Pane
			//Break if no changes
		}*/
		//if there is a figure in 21-th line then break
		fmt.Println("after")
		callServer()
		//Visualize(Pane)
		time.Sleep(1 * time.Second)
	}
}

func callServer() {
	http.Get("http://localhost:8181/")
}
