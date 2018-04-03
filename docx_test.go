package docx

import (
	"strings"
	"testing"
)

const testFile = "./TestDocument.docx"
const testFileResult = "./TestDocumentResult.docx"

func loadFile(file string) *Docx {
	r, err := ReadDocxFile(file)
	if err != nil {
		panic(err)
	}

	return r.Editable()
}

func TestReplace(t *testing.T) {
	d := loadFile(testFile)
	d.Replace("document.", "line1\r\nline2", 1)
	d.WriteToFile(testFileResult)

	d = loadFile(testFileResult)

	if strings.Contains(d.Content, "This is a word document") {
		t.Error("Missing 'This is a word document.', got ", d.Content)
	}

	if !strings.Contains(d.Content, "line1<w:br/>line2") {
		t.Error("Expected 'line1<w:br/>line2', got ", d.Content)
	}
}

func TestReplaceLink(t *testing.T) {
	d := loadFile(testFile)
	d.ReplaceLink("http://example.com/", "https://github.com/nguyenthenguyen/docx", -1)
	d.WriteToFile(testFileResult)

	d = loadFile(testFileResult)

	if strings.Contains(d.Links, "http://example.com") {
		t.Error("Missing 'http://example.com', got ", d.Links)
	}

	if !strings.Contains(d.Links, "https://github.com/nguyenthenguyen/docx") {
		t.Error("Expected 'word', got ", d.Links)
	}
}

func TestReplaceHeader(t *testing.T) {
	d := loadFile(testFile)
	d.ReplaceHeader("This is a header.", "newHeader")
	d.WriteToFile(testFileResult)

	d = loadFile(testFileResult)

	headers := d.Headers
	found := false
	for _, v := range headers {
		if strings.Contains(v, "This is a header.") {
			t.Error("Missing 'This is a header.', got ", d.Content)
		}

		if strings.Contains(v, "newHeader") {
			found = true
		}
	}
	if !found {
		t.Error("Expected 'newHeader', got ", d.Headers)
	}
}

func TestReplaceFooter(t *testing.T) {
	d := loadFile(testFile)
	d.ReplaceFooter("This is a footer.", "newFooter")
	d.WriteToFile(testFileResult)

	d = loadFile(testFileResult)

	footers := d.Footers
	found := false
	for _, v := range footers {
		if strings.Contains(v, "This is a footer.") {
			t.Error("Missing 'This is a footer.', got ", d.Content)
		}

		if strings.Contains(v, "newFooter") {
			found = true
		}
	}
	if !found {
		t.Error("Expected 'newFooter', got ", d.Headers)
	}
}
