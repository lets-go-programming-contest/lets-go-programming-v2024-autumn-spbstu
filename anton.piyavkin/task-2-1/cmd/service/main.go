package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    for i := 0; i < n; i++ {
        lowerBound := 15
        upperBound := 30
        var k int
        _, err = fmt.Scan(&k)
        if err != nil {
            fmt.Fprint(os.Stderr, err)
            os.Exit(1)
        }
        for j := 0; j < k; j++ {
            var str string
            _, err = fmt.Scan(&str)
            if err != nil {
                fmt.Fprint(os.Stderr, err)
                os.Exit(1)
            }
            if (str[0] != '<' && str[0] != '>' && str[1] != '=') || len(str) != 2 {
                fmt.Fprint(os.Stderr, "invalid input\n")
                os.Exit(2)
            }
            var temp int
            _, err = fmt.Scan(&temp)
            if err != nil || (temp > 30 || temp < 15) {
                fmt.Fprint(os.Stderr, "invalid input\n")
                os.Exit(3)
            }
            if str[0] == '<' {
                upperBound = min(upperBound, temp)
            } else {
                lowerBound = max(lowerBound, temp)
            }
            if (lowerBound < upperBound) {
                fmt.Println(upperBound)
            } else {
                fmt.Println(-1)
            }
        }
    }
}