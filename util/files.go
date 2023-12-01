package util

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) []string {
	inputFile, err := os.Getwd()
	inputFile = inputFile + "/data/" + fileName
	CheckError(err)

	readFile, err := os.Open(inputFile)
	CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	return fileLines
}
