package routes

import (
	"GenDriver/src/controller"
	template "GenDriver/src/templates"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/upload", controller.PageLoginController)
	r.Get("/", controller.PageUploadController)
	r.Post("/upload", controller.UploadController)
	r.Handle("/templates/*", http.StripPrefix("/templates/", http.FileServer(http.FS(template.FS))))
	http.ListenAndServe(os.Getenv("port"), r)
}
