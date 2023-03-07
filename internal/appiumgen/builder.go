package appiumgen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/w3connext/appiumgen/pkg/config"
	"github.com/w3connext/appiumgen/pkg/filex"
)

type Builder interface {
	Build(conf config.Config) error
}

type builder struct {
	FlutterGen Generator
}

// Build implements Builder
func (b *builder) Build(conf config.Config) error {
	// Get current path
	currentPath := filepath.Dir("")

	// Get files in current dir
	files, err := filex.GetFileList(currentPath, config.Extensions)
	if err != nil {
		return err
	}

	for _, file := range files {
		// Create target file
		filePath := fmt.Sprintf("%s/%s", currentPath, file.Name())
		filename := filepath.Base(filePath)
		targetFilename := strings.ReplaceAll(filename, config.SpecsSuffix[0], config.ScreenSuffix)
		targetFilename = strings.ReplaceAll(targetFilename, config.SpecsSuffix[1], config.ScreenSuffix)
		targetFilePath := filepath.Join(currentPath, targetFilename)
		if err != nil {
			fmt.Println("[ERROR] Generating file", targetFilename)
			continue
		}

		fmt.Println("[INFO] Generating file", targetFilename)

		// Read file
		fileContents, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("[ERROR] Generating file", targetFilename)
			continue
		}

		// Generate to source code
		sourceCode, err := b.FlutterGen.Generate(fileContents)

		fmt.Println(sourceCode)

		// Write to file
		targetFile, err := os.Create(targetFilePath)
		targetFile.WriteString(sourceCode)
		if err != nil {
			fmt.Println("[ERROR] Generating file", targetFilename)
			continue
		}
		defer targetFile.Close()
	}

	return nil
}

func NewBuilder(flutterGen Generator) Builder {
	return &builder{
		FlutterGen: flutterGen,
	}
}
