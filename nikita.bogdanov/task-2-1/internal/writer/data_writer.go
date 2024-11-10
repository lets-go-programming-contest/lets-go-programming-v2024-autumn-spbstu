package writer

import (
	"bufio"
	"io"
	"strconv"
)

func Write_optional_temperature(out io.Writer, value int) {
	output := bufio.NewWriter(out)
	defer output.Flush()
	str_value := strconv.Itoa(value)
	output.WriteString(str_value)
	output.WriteByte('\n')
}
