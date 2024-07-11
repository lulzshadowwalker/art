package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/lulzshadowwalker/art/internal/util"
)

type ModelCommand struct {
	fs   *flag.FlagSet
	name string
}

func NewModelCommand() *ModelCommand {
	mc := &ModelCommand{
		fs: flag.NewFlagSet("model", flag.ContinueOnError),
	}

	mc.fs.StringVar(&mc.name, "name", "", "Name of the model to create")
	return mc
}

func (c *ModelCommand) Name() string {
	return c.fs.Name()
}

func (c *ModelCommand) Init(args ...string) error {
	c.fs.Parse(args)

	if c.name == "" {
		return fmt.Errorf("model name is required")
	}

	return nil
}

func (c *ModelCommand) Run() error {
	filename := "model.tpl"
	data := map[string]interface{}{
		"Backage": "model",
		"Model":   c.name,
	}

	templ, err := template.
		New(filename).
		Funcs(defaultFuncMap).
		ParseFiles(util.Template(filename))
	if err != nil {
		return err
	}

	return templ.Execute(os.Stdout, data)
}
