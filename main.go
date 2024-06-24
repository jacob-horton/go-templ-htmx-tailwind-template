package main

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"go-templ-htmx-tailwind/components"
)

func ServeFavicon(w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := context.Background()
	components.Home().Render(ctx, w)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /favicon.ico", ServeFavicon)
	mux.HandleFunc("GET /static/", ServeStaticFiles)
	mux.HandleFunc("GET /", Home)

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
