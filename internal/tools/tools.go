package tools

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"
)

func FileToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func NormalizeFilename(filename string) string {
	// TODO: Check if there is an extension. If not, what should we do?
	return fmt.Sprintf("%d%s", time.Now().UTC().UnixMicro(), filepath.Ext(filename))
}

type FolderType string

const (
	ClientFolderType      FolderType = "client"
	AsesorFolderType      FolderType = "asesor"
	InstitutionFolderType FolderType = "institution"
)

func FolderTypePath(docType FolderType, id int32) string {
	return fmt.Sprintf("%s:%d", docType, id)
}
