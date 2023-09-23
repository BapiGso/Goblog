package customw

import (
	latex "github.com/aziis98/goldmark-latex"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
)

var GoldMark = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		mathjax.MathJax,
		&mermaid.Extender{},
		latex.NewLatex(),
		highlighting.NewHighlighting(
			highlighting.WithStyle("github")),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithUnsafe(),
	),
)
