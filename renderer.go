package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func Renderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseFiles("./public/index.html", "public/views/event.ics")),
	}
}
