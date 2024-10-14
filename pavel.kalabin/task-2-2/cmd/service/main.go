package main

import (
    "fmt"
    "log"
    "container/heap"
    
    "github.com/zafod42/task-2-2/internal/inputErrors"
    "github.com/zafod42/task-2-2/internal/dishes"
)

func main() {
    var N, k, dish int
    dishesHeap := &dishes.DishesHeap{}
    heap.Init(dishesHeap)

    _, err := fmt.Scanln(&N)
    if err != nil {
        log.Fatal(fmt.Errorf("N must be a number: %w",err))
    }
    if N < 1 || N > 10000 {
        log.Fatal(inputErrors.NRangeError{})
    }
    for range N {
        _, err = fmt.Scan(&dish)
        if err != nil {
            log.Fatal(fmt.Errorf("Dish priority must be a number: %w", err))
        }
        if dish < -10000 || dish > 10000 {
            log.Fatal(inputErrors.KRangeError{});
        }
        heap.Push(dishesHeap, dish)
    }
    _, err = fmt.Scanln(&k)
    if err != nil {
        log.Fatal(fmt.Errorf("K must be a number: %w", err))
    }
    if k > N || k < 1 {
        log.Fatal(inputErrors.KRangeError{N: N})
    }

    fmt.Println(dishes.FindKMax(dishesHeap, k))
}
