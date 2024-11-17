package ioconfig

import (
    "os"
    "errors"

    "gopkg.in/yaml.v3"
)

type Config struct {
    InputFile string `yaml:"input-file"`
    OutputFile string `yaml:"output-file"`
}

func (c* Config) ReadFromFile(FilePath string) error {
    contents, err := os.ReadFile(FilePath)
    if err != nil {
        return err
    }
    err = c.Parse(contents)
    return err;
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
