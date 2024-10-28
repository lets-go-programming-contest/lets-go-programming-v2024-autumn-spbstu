package temp

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/Piyavva/task-2-1/internal/input"
	"github.com/Piyavva/task-2-1/internal/output"
)

func GetTemp(in io.Reader, out io.Writer, n int) error {
    for i := 0; i < n; i++ {
        lowerBound := 15
        upperBound := 30
        k, err := input.ReadInt(in)
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
                output.WriteInt(upperBound, out)
            } else {
                output.WriteInt(-1, out)
            }
        }
    }
    return nil
}