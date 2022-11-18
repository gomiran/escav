// Package cat extracts text from multiple document formats
package escav

import (
	"io/ioutil"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gomiran/escav/docxtxt"
	"github.com/gomiran/escav/odtxt"
	"github.com/gomiran/escav/plaintxt"
	"github.com/gomiran/escav/rtftxt"
)

// File reads a .odt, .docx, .rtf or plaintext file and returns the content as a string
func File(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return FromBytes(content)
}

// FromBytes converts a []bytes representation of a document to text
func FromBytes(data []byte) (string, error) {
	mime := mimetype.Detect(data)
	switch mime.String() {
	case "application/vnd.oasis.opendocument.text":
		return odtxt.BytesToStr(data)
	case "text/rtf":
		return rtftxt.BytesToStr(data)
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return docxtxt.BytesToStr(data)
	default:
		return plaintxt.BytesToStr(data)
	}
}
