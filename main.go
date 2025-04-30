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

	fmt.Println(style.Render("  ____ _          _       ____ ___    ____ _           _ _ \n" +
		" / ___(_)_ __ ___| | ___ / ___|_ _|  / ___| |__   __ _| | | ___ _ __   __ _  ___ \n" +
		"| |   | | '__/ __| |/ _ \\ |    | |  | |   | '_ \\ / _` | | |/ _ \\ '_ \\ / _` |/ _ \\ \n" +
		"| |___| | | | (__| |  __/ |___ | |  | |___| | | | (_| | | |  __/ | | | (_| |  __/\n" +
		" \\____|_|_|  \\___|_|\\___|\\____|___|  \\____|_| |_|\\__,_|_|_|\\___|_| |_|\\__, |\\___|\n" +
		"                                                                      |___/      "))

	fmt.Println("Starting http server...")

	http.HandleFunc("/challenge", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Request received\n")

		w.Header().Add("Content-Type", "text/json")
		w.Write([]byte("{\"response\": \"Challenge accepted!\"}"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
