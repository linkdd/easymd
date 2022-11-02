package document

import (
	"bytes"
	"fmt"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
)

func RenderPath(path string) (Page, bool) {
	content, err := os.ReadFile(path)
	if err != nil {
		return getErrorPage(err), false
	}

	var params PageParams
	content, err = frontmatter.Parse(bytes.NewReader(content), &params)
	if err != nil {
		return getErrorPage(err), false
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					html.WithLineNumbers(true),
				),
			),
		),
	)
	var buf bytes.Buffer
	if err := markdown.Convert(content, &buf); err != nil {
		return getErrorPage(err), false
	}
	content = buf.Bytes()

	page := Page{
		Params: params,
		Content: fmt.Sprintf(
			`
				<div class="panel is-radiusless m-5">
					<div class="panel-block has-background-white">
						<div class="container">
							<article class="content">%s</article>
						</div>
					</div>
				</div>
			`,
			content,
		),
	}
	return page, true
}

func getErrorPage(err error) Page {
	return Page{
		Params: PageParams{
			Title: "Internal Server Error",
			Lang:  "en",
			Meta:  []MetaTag{},
			CSS:   []string{},
			JS:    []string{},
		},
		Content: fmt.Sprintf(
			`
				<div class="message is-danger m-5">
					<div class="message-header"><p>Internal Server Error</p></div>
					<div class="message-body">ould not render document: %s</div>
				</div>
			`,
			err,
		),
	}
}
