package main

import (
	"VSCode/github_test/app"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("test")
	h := app.NewRouter()
	http.ListenAndServe(":3000", h)
}
