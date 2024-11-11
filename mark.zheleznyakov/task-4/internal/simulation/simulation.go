package simulation

import (
	"fmt"
	"sync"
	"time"

	"github.com/mrqiz/task-4/internal/parking"
)

func Simulate(p *parking.Parking, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := p.PushCar(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("rolling in! in = %d, out = %d. going eepy sleepy for 1 sec...\n", p.Pushed(), p.Popped())
				time.Sleep(500 * time.Millisecond)
			}
		}()

	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := p.PopCar(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("rolling out! in = %d, out = %d. going eepy sleepy for .5 sec...\n", p.Pushed(), p.Popped())
			}
			time.Sleep(500 * time.Millisecond)
		}()
	}

	wg.Wait()

  fmt.Println()
}

