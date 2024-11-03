package exhibit

import (
	"fmt"

	visitorcounter "anastasiya.soboleva/task-4/internal/visitorCounter"
)

type Exhibit struct {
	name           string
	visitorCounter *visitorcounter.VisitorCounter
}

func NewExhibit(name string) *Exhibit {
	return &Exhibit{name: name, visitorCounter: visitorcounter.NewVisitorCounter()}
}

func (e *Exhibit) SimulateVisitorSafe(visitorCount int) {
	for i := 0; i < visitorCount; i++ {
		e.visitorCounter.Increment()
	}
}

func (e *Exhibit) SimulateVisitorUnsafe(visitorCount int) {
	for i := 0; i < visitorCount; i++ {
		e.visitorCounter.UnsafeIncrement()
	}
}

func (e *Exhibit) ShowInfo() {
	count := e.visitorCounter.GetCount()
	fmt.Printf("%s: %d visitors\n", e.name, count)
}
