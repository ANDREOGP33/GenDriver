package utils

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func CriarVerificarDiretorio() {
	if _, err := os.Stat(os.Getenv("uploadPath")); os.IsNotExist(err) {
		err := os.Mkdir(os.Getenv("uploadPath"), 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SalvarArquivo(w http.ResponseWriter, file multipart.File, handler *multipart.FileHeader) {

	filePath := filepath.Join(os.Getenv("uploadPath"), handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Erro ao salvar o arquivo", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Erro ao copiar o arquivo", http.StatusInternalServerError)
		return
	}

}

func ExtensoesPermitidas(handler multipart.FileHeader) bool {
	switch {
	case strings.Contains(handler.Filename, ".sql"):
		return true
	case strings.Contains(handler.Filename, ".dump"):
		return true
	case strings.Contains(handler.Filename, ".tar"):
		return true
	case strings.Contains(handler.Filename, ".gz"):
		return true
	case strings.Contains(handler.Filename, ".gzip"):
		return true
	case strings.Contains(handler.Filename, ".xtrabackup"):
		return true
	case strings.Contains(handler.Filename, ".dmp"):
		return true
	case strings.Contains(handler.Filename, ".exp"):
		return true
	case strings.Contains(handler.Filename, ".rman"):
		return true
	case strings.Contains(handler.Filename, ".bson"):
		return true
	case strings.Contains(handler.Filename, ".json"):
		return true
	case strings.Contains(handler.Filename, ".trn"):
		return true
	case strings.Contains(handler.Filename, ".csc"):
		return true
	case strings.Contains(handler.Filename, ".zip"):
		return true
	case strings.Contains(handler.Filename, ".rar"):
		return true
	case strings.Contains(handler.Filename, ".7z"):
		return true
	default:
		return false
	}
}
