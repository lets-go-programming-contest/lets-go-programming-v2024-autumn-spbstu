package output

import (
	"bufio"
	"io"
	"strconv"
)

func writeInt(value int, out io.Writer) {
    writer := bufio.NewWriter(out)
    defer writer.Flush()
    fVal := strconv.Itoa(value)
    writer.WriteString(fVal)
    writer.WriteByte('\n')
}