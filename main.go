package main

import (
	"VSCode/github_test/app"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("test")
	h := app.NewRouter()
	http.ListenAndServe(":"+port, h)
}
