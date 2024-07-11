package cli

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/lulzshadowwalker/art/internal/util"
)

func TestModel(t *testing.T) {
	c := NewModelCommand()

	t.Run("Name", func(t *testing.T) {
		name := c.Name()
		if name != "model" {
			t.Errorf("expected name to be model, got %s", name)
		}
	})

	t.Run("Init", func(t *testing.T) {
		err := c.Init()
		if err == nil {
			t.Error("expected error, got nil")
		}
		err = c.Init("-name", "test")
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("Run", func(t *testing.T) {
		err := c.Run()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("Template", func(t *testing.T) {
		filename := "service.tpl"
		templ, err := template.New(filename).Funcs(defaultFuncMap).ParseFiles(util.Template(filename))
		if err != nil {
			t.Errorf("failed to parse %s template because %s", filename, err)
		}

		data := []struct {
			data     map[string]interface{}
			hasError bool
		}{
			{
				map[string]interface{}{
					"Backage": "model",
				},
				true,
			},
			{
				map[string]interface{}{
					"name": "User",
				},
				true,
			},
			{
				map[string]interface{}{
					"Backage": "model",
					"name":    "User",
				},
				false,
			},
		}

		for _, d := range data {
			var buf bytes.Buffer
			err := templ.Execute(&buf, d.data)
			if (err != nil) != d.hasError {
				t.Errorf("expected error to be %t, got %v", d.hasError, err)
			}
		}
	})
}
