package functions

import "fmt"

func PrintAnt(finalPath [][]string, path []int) {
	count := 0
	for i := 0; i < len(path); i++ {
		count += path[i]
		for j := 0; j < count; j++ {
			if 0 < len(finalPath[j]) {
				fmt.Print("L", j+1, "-", finalPath[j][0], " ")
				finalPath[j] = finalPath[j][1:]
			}
		}
		fmt.Println()
	}
	max := len(finalPath[0])
	for j := 0; j < len(finalPath); j++ {
		if max < len(finalPath[j]) {
			max = len(finalPath[j])
		}
	}
	for j := 0; j < max; j++ {
		for j := 0; j < count; j++ {
			if 0 < len(finalPath[j]) {
				fmt.Print("L", j+1, "-", finalPath[j][0], " ")
				finalPath[j] = finalPath[j][1:]
			}
		}
		fmt.Println()
	}
}
