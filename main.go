package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
	"github.com/siuyin/gmap/public"
)

func main() {
	http.HandleFunc("/{$}", indexPageHandler)
	http.Handle("/", http.FileServer(http.FS(public.Content)))
	http.HandleFunc("/ger", gerHandler)

	port := dflt.EnvString("PORT", "8080")
	log.Printf("Starting server: GOOGLE_MAPS_API_KEY=**** PORT=%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(public.Content, "index.html"))
	key := dflt.EnvString("GOOGLE_MAPS_API_KEY", "your-api-key-here")
	t.ExecuteTemplate(w, "index.html", struct{ Key string }{key})
}

func gerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from /ger")
}
