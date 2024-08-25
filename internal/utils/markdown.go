package utils

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"log/slog"
)

func ConvertMarkdownToHTML(markdown string) string {
	md := goldmark.New(goldmark.WithExtensions(extension.Strikethrough, extension.Table))
	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		slog.Error("Failed convert to Markdawn", "Error", err)
	}
	return buf.String()
}
