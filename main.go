package main

import (
	"html/template"
	"os"

	"github.com/goccy/go-yaml"
)

func main() {
	buf, err := os.ReadFile("test.yml")
	if err != nil {
		panic(err)
	}

	var cfg config
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		panic(err)
	}

	t, err := template.New("test").Parse(cfg.Template)
	if err != nil {
		panic(err)
	}

	m := model{Name: "success"}
	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}

type config struct {
	Template string `yaml:"template"`
}

type model struct {
	Name string `yaml:"name"`
}
