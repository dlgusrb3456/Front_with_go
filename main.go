package main

import (
	"VSCode/github_test/app"
	"fmt"
	"net/http"
)

func main() {
	//port := os.Getenv("PORT")
	port := "3000"
	fmt.Println("test")
	h := app.NewRouter()
	http.ListenAndServe(":"+port, h)
}
