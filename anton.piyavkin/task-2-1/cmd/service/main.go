package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readInt(in io.Reader) (int, error) {
    reader := bufio.NewReader(in)
    input, err := reader.ReadString('\n')
    if err != nil {
        return 0, err
    }
    input = strings.TrimSpace(input)
    value, err := strconv.Atoi(input)
    if err != nil {
        return 0, err
    }
    return value, nil
}

func writeInt(value int, out io.Writer) {
    writer := bufio.NewWriter(out)
    defer writer.Flush()
    fVal := strconv.Itoa(value)
    writer.WriteString(fVal)
    writer.WriteByte('\n')
}

func getTemp(in io.Reader, out io.Writer, n int) error {
    for i := 0; i < n; i++ {
        lowerBound := 15
        upperBound := 30
        k, err := readInt(in)
        if err != nil {
            return err
        }
        for j := 0; j < k; j++ {
            var str string
            reader := bufio.NewReader(in)
            str, err = reader.ReadString('\n')
            if err != nil {
                return err
            }
            str = strings.TrimSpace(str)
            if (str[0] != '<' && str[0] != '>' && str[1] != '=') {
                return errors.New("invalid input")
            }
            num := strings.Split(str, " ")
            temp, err := strconv.Atoi(num[1])
            if err != nil || (temp > 30 || temp < 15) {
                return errors.New("invalid input2")
            }
            if str[0] == '<' {
                upperBound = min(upperBound, temp)
            } else {
                lowerBound = max(lowerBound, temp)
            }
            if (lowerBound <= upperBound) {
                writeInt(upperBound, out)
            } else {
                writeInt(-1, out)
            }
        }
    }
    return nil
}

func main() {
	n, err := readInt(os.Stdin)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    err = getTemp(os.Stdin, os.Stdout, n)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}