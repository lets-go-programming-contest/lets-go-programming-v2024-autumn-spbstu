package app

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/EmptyInsid/task-4/internal/interfaces/postSocialNet"
	"github.com/EmptyInsid/task-4/internal/structs/postSocialNetSync"
	"github.com/EmptyInsid/task-4/internal/structs/postSocialNetUnsync"
)

func Run(flag string) {

	post := chooseVersion(flag)
	printDocumentContent(&post)

	logChannel := make(chan string, 10)
	var wg sync.WaitGroup
	var logCounter int
	var counterMutex sync.Mutex

	getNextLogID := func() int {
		counterMutex.Lock()
		defer counterMutex.Unlock()
		logCounter++
		return logCounter
	}

	performOperation := func(operationName string, i int, action func()) {
		defer wg.Done()
		action()

		logID := getNextLogID()
		logChannel <- fmt.Sprintf("Log %d: %s by copywriter %d\nCurrent Content:\n%s", logID, operationName, i, post.GetTextContent())
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go performOperation("Added", i, func() {
			post.AddTextContent("\nContent from copywriter " + strconv.Itoa(i))
		})
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go performOperation("Removed 'from'", i, func() {
			post.RemoveTextContent("from")
		})
	}

	wg.Wait()
	close(logChannel)

	writeLogsToFile(logChannel)

	fmt.Println("\nFinal Document Content:")
	printDocumentContent(&post)
}

func printDocumentContent(post *postSocialNet.PostSocialNet) {
	content := (*post).GetTextContent()
	fmt.Println("\nDocument Content:")
	fmt.Println(content)
}

func chooseVersion(flag string) postSocialNet.PostSocialNet {
	var post postSocialNet.PostSocialNet

	switch flag {
	case "sync":
		post = &postSocialNetSync.PostSocialNetSync{ID: 1, TextContent: "Initial content"}
		fmt.Println("Synchronized version start")
	case "unsync":
		post = &postSocialNetUnsync.PostSocialNetUnsync{ID: 1, TextContent: "Initial content"}
		fmt.Println("Unsynchronized version start")
	default:
		panic(fmt.Errorf("unknown flag"))
	}
	return post
}

func writeLogsToFile(logChannel chan string) {
	logFile, err := os.OpenFile("../../logs/operation_logs.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Errorf("error opening log file: %w", err))
	}
	defer logFile.Close()

	for log := range logChannel {
		if _, err := logFile.WriteString("==========" + log + "\n"); err != nil {
			panic(fmt.Errorf("error writing to log file: %w", err))
		}
	}
}
