package controller

import (
	"GenDriver/src/model"
	template "GenDriver/src/templates"
	"GenDriver/src/view"
	"net/http"
)

func PageLoginController(w http.ResponseWriter, r *http.Request) {
	t := view.ParseFS(template.FS, "login.html", "styles.css")
	t.Execute(w, nil)
}

func getUser(r *http.Request) model.User {
	clinica := r.FormValue("clinica")
	cnpj := r.FormValue("cnpj")
	senha := r.FormValue("senha")
	return model.User{
		Clinica: clinica,
		Cnpj:    cnpj,
		Senha:   senha,
	}
}
