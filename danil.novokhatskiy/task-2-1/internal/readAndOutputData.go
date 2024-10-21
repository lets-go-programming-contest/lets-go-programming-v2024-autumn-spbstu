package internal

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func WriteInt(num int, out io.Writer) {
	writer := bufio.NewWriter(out)
	defer writer.Flush()
	val := strconv.Itoa(num)
	writer.WriteString(val)
	writer.WriteByte('\n')
}

func ReadData() (int, string, error) {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		return 0, "", err
	}
	var temp int
	_, err = fmt.Scan(&temp)
	if err != nil {
		return 0, "", err
	}
	return temp, str, nil
}
