package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type MultilineType int

const (
	None MultilineType = iota
	Paragraph
	Quote
	Code
	UnorderedList
	OrderedList
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readMarkdown(path string) ([]string, error) {
	mdData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	mdText := string(mdData)
	return strings.Split(mdText, "\n"), nil
}

// Appends the appropriate closing tag to the passed html string based on the multiline type,
// and returns true if a closing tag was appended - false otherwise.
func appendClosingTagIfNeeded(html *string, mType *MultilineType) bool {

	switch *mType {
	case Paragraph:
		*html += "</p>"
	case Quote:
		*html += "</blockquote>"
	default:
		return false
	}

	*mType = None
	return true

}

// checks to see if the passed markdown line is a proper header (i.e. 1-6 # symbols followed by a space)
// returns the correct html conversion + true if it is, otherwise returns the original string + false
func checkAndConvertHeader(mdLine string) (string, bool) {
	trimmed := strings.TrimSpace(mdLine)

	if strings.HasPrefix(trimmed, "#") {
		headerLevel := 0
		for headerLevel < len(trimmed) && trimmed[headerLevel] == '#' {
			headerLevel++
		}

		if headerLevel >= 1 && headerLevel <= 6 && headerLevel < len(trimmed) && trimmed[headerLevel] == ' ' {
			headerContent := strings.TrimSpace(trimmed[headerLevel+1:])
			return fmt.Sprintf("<h%d>%v</h%d>", headerLevel, headerContent, headerLevel), true
		}
	}
	return mdLine, false

}

func convertMarkdownToHTML(lines []string) string {

	var html string
	var mType MultilineType = None

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			appendClosingTagIfNeeded(&html, &mType)
			continue
		}

		if trimmedLine == "---" {
			appendClosingTagIfNeeded(&html, &mType)
			html += "<hr>"
			continue
		}

		headerText, isHeader := checkAndConvertHeader(trimmedLine)
		if isHeader {
			appendClosingTagIfNeeded(&html, &mType)
			html += headerText
			continue
		}

		if strings.HasPrefix(trimmedLine, "> ") {

			if mType != Quote {
				appendClosingTagIfNeeded(&html, &mType)
				html += "<blockquote>"
				mType = Quote
			} else {
				html += " "
			}
			html += strings.TrimSpace(strings.TrimPrefix(trimmedLine, "> "))
			continue
		}

		if mType == None {
			html += "<p>"
			mType = Paragraph
		} else {
			html += " "
		}
		html += trimmedLine

	}
	appendClosingTagIfNeeded(&html, &mType)

	return html
}

func wrapWithHTMLBoilerplate(body string, title string) string {
	finalTitle := title
	if title == "" {
		finalTitle = "md2html"
	}

	return fmt.Sprintf("<!DOCTYPE html><html><head><title>%v</title></head><body>%v</body></html>",
		finalTitle, body)
}

func writeHTMLToFile(fullHtml string, path string) error {
	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	err := os.WriteFile(path, []byte(fullHtml), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("successfully wrote to: %v\n", path)
	return nil
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("no input markdown file specified")
	}

	inputPath := args[0]
	if !strings.HasSuffix(inputPath, ".md") {
		inputPath += ".md"
	}

	mdLines, err := readMarkdown(inputPath)
	check(err)
	htmlText := convertMarkdownToHTML(mdLines)

	fullHtml := wrapWithHTMLBoilerplate(htmlText, "")

	outputFile := "output/output.html"
	if len(args) < 2 {
		fmt.Println("no output file specified, defaulting to output/output.html")
	} else {
		outputFile = args[1]
		if !strings.HasSuffix(outputFile, ".html") {
			outputFile += ".html"
		}
	}
	if filepath.Dir(outputFile) == "." {
		outputFile = "output/" + outputFile
	}

	err = writeHTMLToFile(fullHtml, outputFile)
	check(err)
}
