package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/KRYST4L614/task-4/internal/duplicateremover"
	"github.com/KRYST4L614/task-4/internal/user"
)

func main() {
	safe := flag.Bool("safe", true, "use safe method")
	flag.Parse()

	channel := make(chan user.User)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		user.GenerateUsers(channel)
	}()
	dr := duplicateremover.NewDuplicateRemover()
	var result []user.User
	go func() {
		defer wg.Done()
		if *safe {
			result = dr.GetUniqueSafe(channel)
		} else {
			result = dr.GetUnique(channel)
		}
	}()
	wg.Wait()
	if len(result) != user.UniqueUsers {
		fmt.Printf("Incorrect result. Expected: %v, actual: %v", user.UniqueUsers, len(result))
	} else {
		fmt.Printf("Count of unique users: %v", user.UniqueUsers)
	}
}
