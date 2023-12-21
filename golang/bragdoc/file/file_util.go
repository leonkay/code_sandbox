package file

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/leonkay/code_sandbox/golang/bragdoc/config"
)

func BragDir(config *config.Config) (string, error) {
	bragPath := filepath.Join(config.Brag.Home, config.Brag.Dir)
	if _, err := os.Stat(bragPath); os.IsNotExist(err) {
		err := os.MkdirAll(bragPath, 0755)
		if err != nil {
			log.Fatal("Error creating Directory:", bragPath, err)
			return "", errors.New("Brag Directory Not Found or Creatable")
		}
		log.Println("Brag Dir created successfully", bragPath)
	}
	return bragPath, nil
}

func CheckFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}
