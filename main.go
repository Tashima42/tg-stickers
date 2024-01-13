package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/tashima42/tg-stickers/app"
)

func main() {
	component := app.Home()
	h := Handler{}

	http.Handle("/", templ.Handler(component))
	http.Handle("/api/convert", h)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
