package main

import (
	art "ASCIIArt/asciiArt"
	_const "ASCIIArt/asciiArt/constants"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ASCIIRequest struct {
	Text  string `json:"text"`
	Style string `json:"style"`
}

type ASCIIResponse struct {
	Art string `json:"art"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("Running server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Render the template, don't worry about the data to fill it yet
	tmpl.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON from the request body
	var requestData ASCIIRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Get ASCII art
	bannerFilename := _const.DefaultBannerName + ".txt"
	if requestData.Style == "shadow" || requestData.Style == "thinkertoy" {
		bannerFilename = requestData.Style + ".txt"
	}
	asciiArt := art.СonvertToAsciiArt(requestData.Text, _const.ResourcesPath+bannerFilename)
	// Respond with ASCII art in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ASCIIResponse{Art: asciiArt})
}
