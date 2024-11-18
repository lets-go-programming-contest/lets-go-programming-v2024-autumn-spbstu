package solution

import (

	"fmt"
	"sync"

	"github.com/zafod42/task-4/internal/safe"
	"github.com/zafod42/task-4/internal/unsafe"
)

const GoroutinesN = 10000

func Safe() {
	var (
		wGroup sync.WaitGroup
		counter safe.Counter
	)
	fmt.Println("=== SAFE ===")
	for i := range GoroutinesN{
		wGroup.Add(1)
		go func (id int) {
			defer wGroup.Done()
			counter.Count()
		}(i)
	}
	wGroup.Wait()

	fmt.Printf("END Counter is %d\n", counter.Value())
}

func UnSafe() {
	var (
		wGroup sync.WaitGroup
		counter unsafe.UnsafeCounter
	)
	fmt.Println("=== UNSAFE ===")
	for i := range GoroutinesN{
		wGroup.Add(1)
		go func (id int) {
			defer wGroup.Done()
			counter.Count()
		}(i)
	}
	wGroup.Wait()

	fmt.Printf("END Counter is %d\n", counter.Value())

}
