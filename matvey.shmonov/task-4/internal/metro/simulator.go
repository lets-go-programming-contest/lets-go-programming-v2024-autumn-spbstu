package metro

import (
	"fmt"
	"sync"

	"golang.org/x/exp/rand"
)

func SimulateWithoutSync(turnstileCollection *TurnstileCollection, numPeople int) {
	for i := 0; i < numPeople; i++ {
		index := rand.Intn(len(turnstileCollection.turnstiles))
		go turnstileCollection.turnstiles[index].RawEnter()
	}
	fmt.Printf("SimulateWithoutSync: %d/%d\n", turnstileCollection.GetCount(), numPeople)
}

func SimulateWithSync(turnstileCollection *TurnstileCollection, numPeople int) {
	var wg sync.WaitGroup
	for i := 0; i < numPeople; i++ {
		wg.Add(1)
		index := rand.Intn(len(turnstileCollection.turnstiles))
		go turnstileCollection.turnstiles[index].Enter(&wg)
	}
	wg.Wait()
	fmt.Printf("SimulateWithSync: %d/%d\n", turnstileCollection.GetCount(), numPeople)
}
