package main

import (
	"TetrisBotKost/lib"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
)

type ViewData struct {
	Title string
	Users []string
}

type ToServer struct {
	Lines int
	Pitch [24]string
}

func transform(field lib.Field) [24]string {
	var ar [24]string

	for i := 0; i < 24; i++ {

		var temp string = "|"
		for j := 0; j < 10; j++ {
			if field[i][j] == 0 {
				//temp += "----"
				temp += "⬜"
			} else {
				temp += "⬛"
				//temp += "0"
			}
		}
		ar[i] = temp + "|"

	}

	return ar
}

func modifyTemplate(w http.ResponseWriter, field ToServer) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	_ = tmpl.Execute(w, field)
	//_ = tmpl.Execute(w, lib.Pane)
}

var port int
var autoOpen bool

func processFlags() {

	// General flags
	flag.IntVar(&port, "port", 8181, "Port to start the UI web server on; valid range: 0..65535")
	flag.BoolVar(&autoOpen, "autoOpen", true, "Auto-opens the UI web page in the default browser")
	flag.Parse()
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	//mux := http.NewServeMux()
	processFlags()
	url := fmt.Sprintf("http://localhost:%d/", port)
	open(url)

	go lib.Calculate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Calling GET")
		modifyTemplate(w, ToServer{lib.NumberOfLines, transform(lib.Pane)})
	})
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)

	return exec.Command(cmd, args...).Start()
}
