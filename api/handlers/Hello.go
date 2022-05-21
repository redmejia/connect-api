package handlers

import (
	"fmt"
	"net/http"
)

func (a *App) SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s ", "World")
	a.Info.Println("testing info")
}
