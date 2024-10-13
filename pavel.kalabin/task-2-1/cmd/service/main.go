package main

import (
    "fmt"
    "log"

    "github.com/zafod42/task-2-1/util/condition"
)


func main() {
    var (
        N, K, T int
        sign string
        cond condition.Condition
    )
	_, err := fmt.Scanln(&N)
	if err != nil {
        log.Fatalf("Number required: %v", err)
	}
    for i := 0; i < N; i++ {
        _, err = fmt.Scanln(&K)
		if err != nil {
            log.Fatalf("Number required: %v", err)
		}
        cond.Init() 
        for j := 0; j < K; j++ {
            _, err = fmt.Scan(&sign)
            if err != nil {
                log.Fatalf("String required: %v", err)
            }
            _, err = fmt.Scanln(&T)
            if err != nil {
                log.Fatalf("Number required: %v", err)
            }
            err = cond.Set(sign, T)
            if err != nil {
                log.Fatal(err)
            }
            optimal := cond.GetOptimal()
            fmt.Println(optimal)
        }
    }
}
