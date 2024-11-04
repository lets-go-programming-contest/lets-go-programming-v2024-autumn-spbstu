package ioconfig

import (
    "errors"
    "gopkg.in/yaml.v3"
)

type Config struct {
    InputFile string `yaml:"input-file"`
    OutputFile string `yaml:"output-file"`
}

func (c *Config) Parse(configuration []byte) error {
    err := yaml.Unmarshal(configuration, c)
    if err != nil {
        return err
    }
    if c.InputFile == "" {
        return errors.New("Unable to parse input file")
    }
    if c.OutputFile == "" {
        return errors.New("Unable to parse output file")
    }
    return err
}
