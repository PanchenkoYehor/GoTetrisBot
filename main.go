package main

import (
	"TetrisBot/library"
	"fmt"
	"time"
)

type ViewData struct{

	Title string
	Users []string
}

func main() {

	library.SetFigures()
	var field = library.Field{}

	for i := 0; i < 24; i++ {
		for j := 0; j < 10; j++ {
			field[i][j] = 0
		}
	}

	for t := 0; t <= 10; t++ {
		var howLeft int
		var fig library.Block
		field, howLeft, fig = library.SetRandomFigureOnTheField(field)
		var rotate, shift = library.FindBestWay(field, howLeft, fig)
		//library.Visualize(field)
		//library.Visualize(field)
		field = library.ShiftInField(field, 3 - howLeft)
		//library.Visualize(field)
		field = library.RotateInField(field, 3, rotate)
		//library.Visualize(field)
		field = library.ShiftInField(field, shift)

		//fmt.Println("before")
		//library.Visualize(field)
		field = library.MoveDownUntilEnd(field)
		//Apply rotate and shift

		/*for {
			//MoveDown
			//Visualize field
			//Break if no changes
		}*/
		//if there is a figure in 21-th line then break
		fmt.Println("after")
		library.Visualize(field)
		time.Sleep(2 * time.Second)
	}

	/*data := ViewData{
		Title : "Users List",
		Users : []string{ "Tom", "Bob", "Sam"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)*/
}