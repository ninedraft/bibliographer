package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	gpflag "github.com/octago/sflags/gen/gpflag"
	cobra "github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Doc         string `desc:"input document, .docx" json:"doc" yaml:"doc"`
	BiblioExcel string `desc:"bibliography table, .xlsx" json:"biblio-excel" yaml:"biblio-excel"`
	Bib         string `desc:"bibliography database, .bib" json:"bib" yaml:"bib"`
	Config      string `json:"-" yaml:"-" desc:"config file, json or yaml"`
	// flag definitions here
	// https://github.com/octago/sflags#flags-based-on-structures------
}

func (config *Config) LoadFromFile(filename string) error {
	var data, errRead = ioutil.ReadFile(filename)
	if errRead != nil {
		return errRead
	}
	switch filepath.Ext(filename) {
	case ".json":
		return json.Unmarshal(data, config)
	default:
		return yaml.Unmarshal(data, config)
	}
}

func App() *cobra.Command {
	var config = Config{}
	var cmd = &cobra.Command{
		Use: "bibliographer",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	if err := gpflag.ParseTo(&config, cmd.PersistentFlags()); err != nil {
		panic(err)
	}
	return cmd
}
