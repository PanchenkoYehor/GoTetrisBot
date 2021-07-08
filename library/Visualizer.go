package library

import (
	"html/template"
	"net/http"
)

type ViewData struct{

	Title string
	Users []string
}

func Visualize(field Field) {
	//mux := http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, field)
	})
	http.ListenAndServe(":8181", nil)
	
	//http.PostForm(":8181", url.Values{"key": {"val"}, "id" : {"123"}})
	//router := gin.Default()

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
