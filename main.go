package main

import (
	"flag"
	"fmt"
	"github.com/PanchenkoYehor/GoTetrisBot/lib"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

type ToServer struct {
	Lines int
	Pitch [24]string
}

func transform(field lib.Field) [24]string {
	var ar [24]string

	for i := 0; i < 24; i++ {

		var temp = "|"
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

const help_html = `<html lang="en">
<head>
<meta charset="UTF-8" http-equiv="refresh" content="0.15">
<title>Tetris Bot by Yehor</title>
</head>
<body>
<h1>Count of destroyed lines: {{.Lines}}</h1>
<ul>
{{range $index, $element := .Pitch}}
<li>{{$element}}</li>
{{end}}
</ul>
</body>
</html>`

func modifyTemplate(w http.ResponseWriter, field ToServer) {
	/*tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}*/
	var tmpl = template.Must(template.New("t").Parse(help_html))
	if tmpl == nil {
		fmt.Printf("Wrong with tmpl\n")
		return
	}
	if err := tmpl.Execute(w, field); err != nil {
		fmt.Printf(err.Error())
	}
	//_ = tmpl.Execute(w, lib.Pane)
}

var Port int
var autoOpen bool

func processFlags() error {

	// General flags
	flag.IntVar(&Port, "Port", 8181, "Port to start the UI web server on; valid range: 0..65535")
	flag.BoolVar(&autoOpen, "autoOpen", true, "Auto-opens the UI web page in the default browser")
	flag.Parse()

	if Port < 0 || Port > 65535 {
		return fmt.Errorf("Port %d is outside of valid range", Port)
	}

	return nil
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	//mux := http.NewServeMux()
	if err := processFlags(); err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	url := fmt.Sprintf("http://localhost:%d/", Port)
	if err := open(url); err != nil {
		fmt.Println("Auto-open failed:", err)
		fmt.Printf("Open %s in your browser.\n", url)
	}

	go lib.Calculate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Calling GET")
		modifyTemplate(w, ToServer{lib.NumberOfLines, transform(lib.Pane)})
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
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
