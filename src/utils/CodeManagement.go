package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"encoding/csv"
	"os/exec"
	"strings"
	"strconv"

	"go.uber.org/zap"
	log "book-service-app/src/utils/loggerUtils"
)

func CodeManager() {
	logger := log.FileLogger()
	projectPath := "D:/harshit-golang-bookservice/src"

	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
			if err != nil {
				logger.Error("Error parsing file: %s\n", zap.Error(err))
				return err
			}
		
			csvFilePath := "D:/harshit-golang-bookservice/code-management.csv"

		
			csvFile, err := os.OpenFile(csvFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				logger.Error("Error creating CSV file:", zap.Error(err))
			}
			defer csvFile.Close()

			// Create a CSV writer
			csvWriter := csv.NewWriter(csvFile)
			defer csvWriter.Flush()

			ast.Inspect(file, func(node ast.Node) bool {
				if fd, ok := node.(*ast.FuncDecl); ok {
					funcName := fd.Name.Name
					
					filePath := path
					
					logCmd := exec.Command("git", "log", "-L", ":"+funcName+":"+filePath)

					output, err := logCmd.Output()

					if err != nil {
						logger.Error("Error executing git log command:", zap.Error(err))
					}
					lines := strings.Split(string(output), "\n")
				
					var commitCount = 0
					
					for _, line := range lines {
						line = strings.TrimSpace(line)
						if strings.HasPrefix(line, "commit") {
							commitCount ++;
						}
					}
					
					csvWriter.Write([]string{funcName, path, strconv.Itoa(commitCount)})
				}
				return true
			})
		}

		return nil
	})

	if err != nil {
		logger.Error("Error walking project directory: %s\n",zap.Error(err))
	}

}