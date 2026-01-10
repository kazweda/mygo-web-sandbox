# go-web-sandbox

A minimal sandbox repository for hackathons.

## Purpose
- Static HTML generation using Go
- Designed for GitHub Pages deployment
- Interactive features are intentionally left unimplemented

## License
MIT

## Go Web Sandbox

Go + templ を使った、ハッカソン向け Web サンドボックスです。
擬似 SSG で静的 HTML を生成し、GitHub Pages にデプロイします。

### Requirements

- Go 1.22+
- templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
