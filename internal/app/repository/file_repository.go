package repository

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/utils/common"
)

type FileRepository interface {
	Save(fileName string, file *multipart.File) (string, error)
	FindFile(c *gin.Context, filepath string, filename string)
}

type fileRepository struct {
	fileBasePath string
}

func NewFileRepository(basePath string) FileRepository {
	return &fileRepository{fileBasePath: basePath}
}

func (f *fileRepository) Save(fileName string, file *multipart.File) (string, error) {
	fileLocation := filepath.Join(f.fileBasePath, fileName)
	err := common.SaveToLocalFile(fileLocation, file)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

func (f *fileRepository) FindFile(c *gin.Context, filepath string, filename string) {
	c.FileAttachment(filepath, filename)
}
