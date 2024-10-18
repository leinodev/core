package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"strings"
)

type CoreConfig struct {
	DB DBConfig `json:"db"`
}

func NewConfig(path string) (*CoreConfig, error) {
	config := new(CoreConfig)
	if err := config.readConfig(path); err != nil {
		return nil, err
	}
	return config, nil
}

func (config *CoreConfig) readConfig(path string) error {
	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return err
	}
	return nil
}

func (config *CoreConfig) String() string {
	var builder strings.Builder
	builder.WriteString("DB: {\n")
	builder.WriteString(fmt.Sprintf("\tconnection: %s,\n", config.DB.Connection.Dialect))
	builder.WriteString("\tschemas: [\n")
	for i, s := range config.DB.Schemas {
		builder.WriteString("\t\t{\n")
		builder.WriteString(fmt.Sprintf("\t\t\tname: %s,\n", s.Name))
		builder.WriteString("\t\t\ttables: [\n")
		for j, t := range s.Tables {
			builder.WriteString("\t\t\t\t\t{\n")
			builder.WriteString(fmt.Sprintf("\t\t\t\t\t\t---- %s ----,\n", t.Name))
			for _, f := range t.Fields {
				builder.WriteString(fmt.Sprintf("\t\t\t\t\t\t%s\t\t%s\t\t%s,\n", f.Name, f.Type.Type, strings.Join(f.Tags, " ")))
			}
			builder.WriteString("\t\t\t\t\t}")
			if j == len(s.Tables)-1 {
				builder.WriteString(",\n")
			} else {
				builder.WriteString("\n")
			}
		}
		builder.WriteString("\t\t\t]\n")
		// TODO print indexes
		builder.WriteString("\t\t}")
		if i == len(config.DB.Schemas)-1 {
			builder.WriteString(",\n")
		} else {
			builder.WriteString("\n")
		}
	}
	builder.WriteString("\t]\n")
	builder.WriteString("}\n")

	return builder.String()
}
