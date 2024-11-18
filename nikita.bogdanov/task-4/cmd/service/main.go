package main

import (
	"fmt"

	"github.com/solomonalfred/task-4/internal/mulithread_programms"
)

func main() {
	fmt.Println("Sync multithreading metro turnstile")
	mulithread_programms.SyncMT()
	fmt.Println("\n\nUnsync multithreading metro turnstile")
	mulithread_programms.UnsyncMT()
}
