package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Data struct {
	Elment, Value, Color, Downlaod string
}

var output Data

func Print(w http.ResponseWriter, r *http.Request) { // w for send data from server to user and r for take data from user
	tmp, err := template.ParseFiles("./templet/index.html") // pionter in file html
	if r.URL.Path != "/" {                                  // handel if url was not valide
		http.NotFound(w, r) // enter to func for print not found
		return
	}
	if err != nil {
		http.Error(w, "server down", http.StatusInternalServerError) // hundul if was file html not
		return
	}
	tmp = template.Must(tmp, err)
	fmt.Println(output.Color)
	tmp.Execute(w, output)
	if r.Method == "GET" {
		output.Downlaod = ""
	} // send data to showing in page in form respons and creat ascii art
}

func Download(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(output.Elment)))

	// ")
	// //		"Content-Type",
	// //		"Content-Type",
	// //		"Content-Length",

	// w.Header().Set("")
	w.Write([]byte(output.Elment))
}

func Handel_input(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	font := r.FormValue("select")
	user_input := r.FormValue("user_input")
	output.Color = r.FormValue("color")

	if len(font) == 0 || len(user_input) == 0 {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	mapDraw := Font(font)
	if mapDraw == nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	Stock(user_input, mapDraw, &output)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Stock(s string, mapDraw map[int][]string, r *Data) {
	r.Value = s
	r.Elment = SplitAndPrint(s, mapDraw)
	r.Downlaod = "download"
}
