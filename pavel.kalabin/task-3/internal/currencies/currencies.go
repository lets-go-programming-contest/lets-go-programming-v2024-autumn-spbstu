package currencies

import (
    "encoding/xml"
    "strings"
    "golang.org/x/net/html/charset"
    "sort"
)

type Entrie struct {
    NumCode   int     `xml:"NumCode"`     
    CharCode  string  `xml:"CharCode"`
    Nominal   int     `xml:"Nominal"`
    Name      string  `xml:"Name"`
    Value     float32 `xml:"Value"`
    VunitRate float32 `xml:"VunitRate"`
}

type Currencies struct {
    XMLName xml.Name `xml:"ValCurs"`
    Date    string   `xml:"Date,attr"`
    Name    string   `xml:"name,attr"`
    Entries []Entrie `xml:"Valute"`
}

func (c *Currencies) Parse(contents []byte) error {
    decoder := xml.NewDecoder(strings.NewReader(string(contents)))
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(c)
    return err
}

func (c *Currencies) Sort() {
    sort.Slice(c.Entries, func(i, j int) bool {
        return c.Entries[i].Value > c.Entries[j].Value
    })
}
