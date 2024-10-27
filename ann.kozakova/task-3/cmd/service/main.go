package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

var (
	configPath string
)

func initFlag() {
	flag.StringVar(&configPath, "config", "config.yml", "The path to the configuration file")
	flag.Parse()
}

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []Valute `xml:"Valute"`
}

type Valute struct {
	ID        string `xml:"ID,attr"`
	NumCode   string `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

//float?????

func main() {
	initFlag()
	configFile, err := os.OpenFile(configPath, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	buffer := make([]byte, 512)
	n, err := configFile.Read(buffer)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(buffer[:n], &config)
	if err != nil {
		panic(err)
	}

	//fmt.Println("input: ", config.Input)
	//fmt.Println("output: ", config.Output)

	inputFile, err := os.OpenFile(config.Input, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	dir := filepath.Dir(config.Output)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	outputFile, err := os.OpenFile(config.Output, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	defer outputFile.Close()

	buf := make([]byte, 512)
	var data []byte
	for {
		n, err := inputFile.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		data = append(data, buf[:n]...)
	}

	//fmt.Println(string(data))

	valCurs := new(ValCurs)
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&valCurs)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(valCurs)
	//decoder := xml.NewDecoder(inputFile)
	//decoder.CharsetReader = charset.NewReaderLabel
	//
	//err = decoder.Decode(&data)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}

	//err = xml.Unmarshal(data, &valCurs)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//fmt.Println(data)

	//decoder := xml.NewDecoder(charmap.Windows1251.NewDecoder().Reader(bytes.NewReader(data)))
	//
	//err = decoder.Decode(&valCurs)
	//if err != nil {
	//	panic(err)
	//}
}

// go run main.go -config D:\lets-go-programming-v2024-autumn-spbstu\ann.kozakova\task-3\config\config.yml
