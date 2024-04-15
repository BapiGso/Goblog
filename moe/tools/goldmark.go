package tools

import (
	"bytes"
	latex "github.com/aziis98/goldmark-latex"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/mermaid"
)

func init() {
	a := goldmark.DefaultParser()
	a.AddOptions()
	b := goldmark.DefaultRenderer()
	b.AddOptions()
	goldmark.WithParserOptions()
	//c.Parser().AddOptions()
}

var GoldMark = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Linkify,
		mathjax.MathJax,
		&mermaid.Extender{
			MermaidURL: "/assets/blog/js/mermaid.js", //自定义js静态资源路径
		},
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

var (
	tocResultKey = parser.NewContextKey()
	tocEnableKey = parser.NewContextKey()
)

type tocTransformer struct {
	r renderer.Renderer
}

func (t *tocTransformer) Transform(n *ast.Document, reader text.Reader, pc parser.Context) {
	if b, ok := pc.Get(tocEnableKey).(bool); !ok || !b {
		return
	}

	var (
		level       int
		row         = -1
		inHeading   bool
		headingText bytes.Buffer
	)

	ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		if n.Kind() == ast.KindHeading {
			if inHeading && !entering {
				return s, nil
			}

			inHeading = true
		}

		if !(inHeading && entering) {
			return s, nil
		}

		switch n.Kind() {
		case ast.KindHeading:
			heading := n.(*ast.Heading)
			level = heading.Level

			if level == 1 || row == -1 {
				row++
			}

		case
			ast.KindCodeSpan,
			ast.KindLink,
			ast.KindImage,
			ast.KindEmphasis:
			err := t.r.Render(&headingText, reader.Source(), n)
			if err != nil {
				return s, err
			}

			return ast.WalkSkipChildren, nil
		case
			ast.KindAutoLink,
			ast.KindRawHTML,
			ast.KindText,
			ast.KindString:
			err := t.r.Render(&headingText, reader.Source(), n)
			if err != nil {
				return s, err
			}
		}

		return s, nil
	})
}

type tocExtension struct {
	options []renderer.Option
}

func (e *tocExtension) Extend(m goldmark.Markdown) {
	r := goldmark.DefaultRenderer()
	r.AddOptions(e.options...)
	m.Parser().AddOptions(parser.WithASTTransformers(util.Prioritized(&tocTransformer{
		r: r,
	}, 10)))
}
