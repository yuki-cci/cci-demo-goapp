package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color("#049949")).
		PaddingTop(2).
		PaddingBottom(2).
		PaddingLeft(4).
		PaddingRight(4)

	fmt.Println(style.Render("              _ \n" +
		"__      _____| | ___ ___  _ __ ___   ___ \n" +
		"\\ \\ /\\ / / _ \\ |/ __/ _ \\| '_ ` _ \\ / _ \\ \n" +
		" \\ V  V /  __/ | (_| (_) | | | | | |  __/_\n" +
		"  \\_/\\_/ \\___|_|\\___\\___/|_| |_| |_|\\___(_)\n" +
		" "))

	fmt.Println("Starting http server...")

	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Request received\n")

		w.Header().Add("Content-Type", "text/json")
		w.Write([]byte("{\"response\": \"Request received!\"}"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
