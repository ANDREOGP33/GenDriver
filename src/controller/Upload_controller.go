package controller

import (
	"GenDriver/src/model"
	template "GenDriver/src/templates"
	"GenDriver/src/utils"
	"GenDriver/src/view"
	"net/http"
)

func PageUploadController(w http.ResponseWriter, r *http.Request) {
	t := view.ParseFS(template.FS, "upload.html", "styles.css")
	t.Execute(w, nil)
}

func UploadController(w http.ResponseWriter, r *http.Request) {
	utils.CriarVerificarDiretorio()

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao fazer upload do arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBackup := model.FileBackup{
		File:    file,
		Handler: handler,
	}

	fileBackup.ValidarExtensao()

	utils.SalvarArquivo(w, file, handler)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
