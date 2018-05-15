package main

import (
	"net/http"
	"github.com/rakyll/statik/fs"
	"github.com/skratchdot/open-golang/open"
	_ "github.com/fenril22/statik"
	"log"

	//files "./html"
)

func main() {
	statikFS, err := fs.New()

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	listen := make(chan bool)

	go func() {
		<-listen
		open.Run("http://localhost:8765")
		log.Println("browser start")
	}()

	listen <- true
	http.ListenAndServe(":8765", nil)
}

