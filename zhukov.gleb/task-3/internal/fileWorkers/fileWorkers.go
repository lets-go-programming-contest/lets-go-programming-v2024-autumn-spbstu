package fileWorkers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"task-3/internal/structs"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "", "Path to cfg file")
	flag.Parse()
}

type FilesWorker struct{}

func NewFileWorker() *FilesWorker {
	return &FilesWorker{}
}

func (p *FilesWorker) CfgParse() (inOutFilePaths structs.Cfg, err error) {
	cfgFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return structs.Cfg{}, err
	}

	cfg := structs.Cfg{}
	if err = yaml.Unmarshal(cfgFile, &cfg); err != nil {
		return structs.Cfg{}, err
	}

	return cfg, nil
}

func (p *FilesWorker) InputFileParse(inFilePath string) ([]structs.Currency, error) {
	file, err := os.ReadFile(inFilePath)
	if err != nil {
		return nil, err
	}

	data := []byte(strings.ReplaceAll(string(file), ", ", "."))
	data = []byte(strings.ReplaceAll(string(data), ",", "."))

	fileReader := bytes.NewReader(data)
	decoder := xml.NewDecoder(fileReader)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs structs.ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, err
	}

	sort.Slice(valCurs.Currencies, func(i, j int) bool {
		return valCurs.Currencies[i].Value > valCurs.Currencies[j].Value
	})

	return valCurs.Currencies, nil
}

func (p *FilesWorker) OutputFileParse(outFilePath string, data []structs.Currency) error {
	dir := filepath.Dir(outFilePath)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
