package main

import (
	art "ASCIIArt/asciiArt"
	_const "ASCIIArt/asciiArt/constants"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error loading template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.PostFormValue("text")
	style := r.PostFormValue("style")

	bannerFilename := _const.DefaultBannerName + ".txt"
	if style == "shadow" || style == "thinkertoy" {
		bannerFilename = style + ".txt"
	}

	asciiArt := art.СonvertToAsciiArt(text, _const.ResourcesPath+bannerFilename)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(asciiArt))
}
