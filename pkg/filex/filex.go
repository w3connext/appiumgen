package filex

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/w3connext/appiumgen/pkg/common"
)

func GetFileList(path string, exts []string) ([]fs.FileInfo, error) {
	files := []fs.FileInfo{}

	// Get the current directory
	dir, err := filepath.Abs(path)
	if err != nil {
		return files, err
	}

	// Open the directory
	dirHandle, err := os.Open(dir)
	if err != nil {
		return files, err
	}
	defer dirHandle.Close()

	// Read the directory entries
	fileInfos, err := dirHandle.Readdir(0)
	if err != nil {
		return files, err
	}

	// Get file by extensions
	for _, fileInfo := range fileInfos {
		if common.Contains(exts, filepath.Ext(fileInfo.Name())) {
			files = append(files, fileInfo)
		}
	}

	return files, nil
}
