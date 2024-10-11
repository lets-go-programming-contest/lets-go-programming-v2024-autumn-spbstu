package optimaltemp

import (
	"fmt"

	myErrors "github.com/artem6554/task-2-1/myErrors"
)

func OptimalTemp() (err error) {
	var depCount int
	var emplCount int
	var op string
	var temp int
	moreThan := -1
	lessThan := 31

	if _, errDep := fmt.Scan(&depCount); errDep != nil || depCount <= 0 {
		err = new(myErrors.DepCountError)
	}
	for i := 0; i < depCount; i++ {
		if _, errEmpl := fmt.Scan(&emplCount); errEmpl != nil || emplCount <= 0 {
			err = new(myErrors.EmplCountError)
			continue
		}
	CurDep:
		for j := 0; j < emplCount; j++ {
			fmt.Scan(&op)
			fmt.Scan(&temp)
			switch op {
			case ">=":
				if temp > moreThan {
					moreThan = temp
				}
			case "<=":
				if temp < lessThan {
					lessThan = temp
				}
			default:
				fmt.Println("Incorrect input!")
				continue CurDep
			}
			switch {
			case lessThan == 31:
				fmt.Println(moreThan)
			case moreThan == -1:
				fmt.Println(lessThan)
			case moreThan <= lessThan:
				fmt.Println(moreThan)
			default:
				fmt.Println(-1)
				break CurDep
			}
		}
		moreThan = -1
		lessThan = 31
	}
	return
}
