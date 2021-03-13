package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller
type Controller interface {
	GetShapeByName(w http.ResponseWriter, r *http.Request)
	//SetShapeName(w http.ResponseWriter, r *http.Request)
}

// New returns router instance which is used in main package to register handlers.
func New(controller Controller) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/get/shapes/{shape_name}", controller.GetShapeByName).Methods(http.MethodGet).Name("get_shape")
	//r.HandleFunc("/set/shapes/{shape_name}", controller.SetShapeName).Methods(http.MethodGet).Name("put_shape")

	return r
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
