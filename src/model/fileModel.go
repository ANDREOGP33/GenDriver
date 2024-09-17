package model

import (
	"mime/multipart"
	"strings"
)

type FileBackup struct {
	File            multipart.File
	Handler         *multipart.FileHeader
}

func (fb *FileBackup) ValidarExtensao() bool {
	switch {
	case strings.Contains(fb.Handler.Filename, ".sql"):
		return true
	case strings.Contains(fb.Handler.Filename, ".dump"):
		return true
	case strings.Contains(fb.Handler.Filename, ".tar"):
		return true
	case strings.Contains(fb.Handler.Filename, ".gz"):
		return true
	case strings.Contains(fb.Handler.Filename, ".gzip"):
		return true
	case strings.Contains(fb.Handler.Filename, ".xtrabackup"):
		return true
	case strings.Contains(fb.Handler.Filename, ".dmp"):
		return true
	case strings.Contains(fb.Handler.Filename, ".exp"):
		return true
	case strings.Contains(fb.Handler.Filename, ".rman"):
		return true
	case strings.Contains(fb.Handler.Filename, ".bson"):
		return true
	case strings.Contains(fb.Handler.Filename, ".json"):
		return true
	case strings.Contains(fb.Handler.Filename, ".trn"):
		return true
	case strings.Contains(fb.Handler.Filename, ".csc"):
		return true
	case strings.Contains(fb.Handler.Filename, ".zip"):
		return true
	case strings.Contains(fb.Handler.Filename, ".rar"):
		return true
	case strings.Contains(fb.Handler.Filename, ".7z"):
		return true
	default:
		return false
	}
}
