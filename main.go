package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
	"github.com/siuyin/gmap/public"
)

var t *template.Template

func main() {
	t = template.Must(template.ParseFS(public.Content, "tmpl/*"))
	http.HandleFunc("/myloc", myLocationPageHandler)
	http.HandleFunc("/ger", gerHandler)
	http.HandleFunc("/{$}", indexHandler)
	http.HandleFunc("/placepicker",placePickerHandler)
	http.HandleFunc("/placepickermap",placePickerMapHandler)
	http.HandleFunc("/datalayers",dataLayersHandler)
	http.HandleFunc("/index.html", indexHandler)

	http.Handle("/", http.FileServer(http.FS(public.Content)))

	port := dflt.EnvString("PORT", "8080")
	log.Printf("Starting server: GOOGLE_MAPS_API_KEY=**** PORT=%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index.html", nil)
}
func myLocationPageHandler(w http.ResponseWriter, r *http.Request) {
	key := dflt.EnvString("GOOGLE_MAPS_API_KEY", "your-api-key-here")
	t.ExecuteTemplate(w, "myloc.html", struct{ Key string }{key})
}

func gerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from /ger")
}

func placePickerHandler(w http.ResponseWriter, r *http.Request) {
	key := dflt.EnvString("GOOGLE_MAPS_API_KEY", "your-api-key-here")
	t.ExecuteTemplate(w, "placepicker.html", struct{ Key string }{key})
}

func placePickerMapHandler(w http.ResponseWriter, r *http.Request) {
	key := dflt.EnvString("GOOGLE_MAPS_API_KEY", "your-api-key-here")
	t.ExecuteTemplate(w, "placepickermap.html", struct{ Key string }{key})
}

func dataLayersHandler(w http.ResponseWriter, r *http.Request) {
	key := dflt.EnvString("GOOGLE_MAPS_API_KEY", "your-api-key-here")
	t.ExecuteTemplate(w, "datalayers.html", struct{ Key string }{key})
	}