package document

import (
	"io"
	"text/template"
)

const layout = `
<!DOCTYPE html>
<html lang="{{ .Params.Lang }}" class="has-background-light">
	<head>
		<title>{{ .Params.Title }}</title>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />

		{{ range .Params.Meta }}
			<meta {{ .Attribute }}="{{ .Value }}" />
		{{ end }}

		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css" />
		{{ range .Params.CSS }}
			<link rel="stylesheet" href="{{ . }}" />
		{{ end }}

		{{ range .Params.JS }}
			<script type="application/javascript" src="{{ . }}" defer></script>
		{{ end }}
	</head>
	<body>
		<main class="container">
			{{ .Content }}
		</main>
	</body>
</html>
`

type MetaTag struct {
	Attribute string `yaml:"attribute"`
	Value     string `yaml:"value"`
}

type PageParams struct {
	Title string    `yaml:"title"`
	Lang  string    `yaml:"lang" default:"en"`
	Meta  []MetaTag `yaml:"meta" default:"[]MetaTag{}"`
	CSS   []string  `yaml:"css" default:"[]string{}"`
	JS    []string  `yaml:"js" default:"[]string{}"`
}

type Page struct {
	Params  PageParams
	Content string
}

func (p Page) Render(w io.Writer) error {
	tmpl, err := template.New("page").Parse(layout)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, p)
}
