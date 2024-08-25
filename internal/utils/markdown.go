package utils

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func ConvertMarkdownToHTML(markdown string) (string, error) {
	md := goldmark.New(goldmark.WithExtensions(extension.Strikethrough, extension.Table))
	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
