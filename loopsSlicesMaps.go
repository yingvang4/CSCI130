package main

import "fmt"

func main() {
	
	fmt.Println("Even numbers from 1 to 50:")
	i := 1
	for {
		if i > 50 {
			break
		}
		if i%2 == 1 {
			i++
			continue
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	sliceOfNames := []string{"Tom", "Jerry", "Bobby", "Tony"}
	
	gradYearMap := map[string]int {
			"Tony":		2015,
			"Bobby":	2016,
			"Tom":		2012,
			"Jerry":	2009,
	}
	
	for _ , item := range sliceOfNames {
		fmt.Println(item, "is graduating in", gradYearMap[item])
	}
	
}
