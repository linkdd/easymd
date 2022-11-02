# easymd

This program is a very simple HTTP server used to render Markdown documents to
HTML, to share easily some informations with others.

## Features

 - parse frontmatter
 - render Markdown documents to HTML using [Goldmark](https://github.com/yuin/goldmark)
 - tables and syntax highlighting support

Frontmatter:

```yaml
---
title: "Hello world"
lang: en
meta:
 - attribute: keywords
   value: hello,world
css:
 - https://url/to/style.css
js:
 - https://url/to/script.js
---
```

## Usage

```
easymd - a simple server rendering markdown documents to HTML

Usage:
  easymd [flags]

Flags:
  -b, --bind ip       IP address to listen on (default 0.0.0.0)
  -h, --help          help for easymd
  -p, --port int      Port to listen on (default 8000)
  -r, --root string   Root directory to scan for markdown documents (default ".")
```

If you have the following structure:

```
|-+ docs/
  |-- _index.md
  |-- hello.md
  |-+ foo/
    |-- bar.md
```

And run the command:

```
$ easymd -h 127.0.0.1 -p 8000 -r ./docs/
```

The following URLs will be available:

 - http://127.0.0.1:8000/
 - http://127.0.0.1:8000/hello/
 - http://127.0.0.1:8000/hello
 - http://127.0.0.1:8000/foo/bar/
 - http://127.0.0.1:8000/foo/bar

## License

This project is released under the terms of the [MIT License](./LICENSE.txt).
