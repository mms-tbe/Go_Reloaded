package main

import (
	"bufio"
	"go_reloaded/internal/utils"
	"log"
	"os"
	"strings"
)

func processLine(line string) string {
	line = utils.FormatText(line)
	line = utils.FixingMods(line)
	line = utils.FixingMods(line)
	line = utils.FormatText(line)
	return line
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("<sample.file>, <result.file>")
	}

	sourceFile := os.Args[1]
	destinationFile := os.Args[2]

	if !strings.HasSuffix(destinationFile, ".txt") {
		log.Fatal("<filename.txt>")
	}

	srcFile, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal("error opening source file")
	}
	defer srcFile.Close()

	dstFile, err := os.Create(destinationFile)
	if err != nil {
		log.Fatal("error creating destination file")
	}
	defer dstFile.Close()

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		processedLine := processLine(line)
		_, err := writer.WriteString(processedLine + "\n")
		if err != nil {
			log.Fatal("error writing to file")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading from file")
	}
}