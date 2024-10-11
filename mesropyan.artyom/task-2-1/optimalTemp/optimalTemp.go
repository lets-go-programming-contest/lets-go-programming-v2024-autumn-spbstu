package optimaltemp

import "fmt"

func OptimalTemp() {
	var depCount int
	var emplCount int
	var op string
	var temp int
	moreThan := -10
	lessThan := 40

	fmt.Scan(&depCount)
	for i := 0; i < depCount; i++ {
		fmt.Scan(&emplCount)
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
			}
		}
		moreThan = -1
		lessThan = 31
	}
}
