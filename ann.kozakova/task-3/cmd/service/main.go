package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	ID        string  `xml:"ID,attr"`
	NumCode   int     `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   int     `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     float64 `xml:"Value"`
	VunitRate float64 `xml:"VunitRate"`
}

type ValuteInJSON struct {
	NumCode  int     `json:"num-code"`
	CharCode string  `json:"char-code"`
	Value    float64 `json:"value"`
}

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

	data = []byte(strings.ReplaceAll(string(data), ",", "."))

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

	//err = xml.Unmarshal(data, &valCurs)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}

	//fmt.Println(valCurs)

	sort.Slice(valCurs.Valute, func(i, j int) bool {
		return valCurs.Valute[i].Value < valCurs.Valute[j].Value
	})

	//dataAfterSort, err := xml.Marshal(valCurs)
	//if err != nil {
	//	panic(err)
	//}
	var vij []ValuteInJSON

	for _, valute := range valCurs.Valute {
		jsonVal := ValuteInJSON{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    valute.Value,
		}
		vij = append(vij, jsonVal)
	}

	//err = json.Unmarshal(dataAfterSort, &vij)
	//if err != nil {
	//	panic(err)
	//}
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", " ")

	//encoder := json.NewEncoder(outputFile)

	if err := encoder.Encode(vij); err != nil {
		panic(err)
	}

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
