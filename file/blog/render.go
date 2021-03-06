package blog

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

func Render(w io.Writer, p Post) error {
	tmplt, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := tmplt.Execute(w, p); err != nil {
		return err
	}

	return nil
}
