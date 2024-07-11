package cli

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/lulzshadowwalker/art/internal/util"
)

func TestServiceCommand(t *testing.T) {
	s := NewServiceCommand()

	t.Run("Name", func(t *testing.T) {
		name := s.Name()
		if name != "service" {
			t.Errorf("expected name to be service, got %s", name)
		}
	})

	t.Run("Init", func(t *testing.T) {
		err := s.Init()
		if err == nil {
			t.Error("expected error, got nil")
		}
		err = s.Init("-name", "test")
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("Run", func(t *testing.T) {
		err := s.Run()
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
					"Backage": "service",
				},
				true,
			},
			{
				map[string]interface{}{
					"name": "Authentication",
				},
				true,
			},
			{
				map[string]interface{}{
					"Backage": "service",
					"name":    "Authentication",
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
