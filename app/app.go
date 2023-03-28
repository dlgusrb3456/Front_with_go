package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render = render.New()

func getMainPage(w http.ResponseWriter, r *http.Request) {
	rd.JSON(w, http.StatusOK, nil)
}

func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", getMainPage)

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}
