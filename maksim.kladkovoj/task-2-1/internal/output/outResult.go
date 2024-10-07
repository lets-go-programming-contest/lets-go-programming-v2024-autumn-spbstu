package output

import "github.com/Mmmakskl/task-2-1/internal/calculate"

func OutResult(n int) {
	for i := 0; i < n; i++ {
		calculate.OptimalTemp(n)
	}
}
