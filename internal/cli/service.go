package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/lulzshadowwalker/art/internal/util"
)

type ServiceCommand struct {
	fs   *flag.FlagSet
	name string
}

func NewServiceCommand() *ServiceCommand {
	sc := &ServiceCommand{
		fs: flag.NewFlagSet("service", flag.ContinueOnError),
	}

	sc.fs.StringVar(&sc.name, "name", "", "Name of the service to create")
	return sc
}

func (c *ServiceCommand) Name() string {
	return c.fs.Name()
}

func (c *ServiceCommand) Init(args ...string) error {
	c.fs.Parse(args)

	if c.name == "" {
		return fmt.Errorf("service name is required")
	}

	return nil
}

func (c *ServiceCommand) Run() error {
	filename := "service.tpl"
	templ, err := template.
		New(filename).
		Funcs(defaultFuncMap).
		ParseFiles(util.Template(filename))
	if err != nil {
		return fmt.Errorf("failed to parse %s template because %w", filename, err)
	}

	data := map[string]interface{}{
		"Backage": "service",
		"name":    c.name,
	}
	if err := templ.Execute(os.Stdout, data); err != nil {
		return fmt.Errorf("failed to execute %s template because %w", filename, err)
	}

	return nil
}
