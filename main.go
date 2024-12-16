package main

import (
	"fmt"
	"os"
	"strings"

	functions "lem-in/functions"
	structs "lem-in/structs"
)

func main() {
	path, finalPath, numberPaths := []string{}, [][]string{}, []int{}
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	antFarm, file, err := functions.ReadInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	functions.GetPaths(antFarm.Tunnels, antFarm.Start.Name, antFarm.End.Name, path)
	if len(structs.Paths) == 0 {
		fmt.Println("ERROR: invalid path")
		return
	}
	fmt.Println(strings.Join(file, "\n") + "\n")
	// functions.PrintGraph(antFarm)
	shortPaths, longPaths := functions.BestPaths(structs.Paths)
	if len(shortPaths) == 0 && len(longPaths) == 0 {
		fmt.Println("ERROR: invalid path")
		return
	} else if len(shortPaths) >= len(longPaths) {
		longPaths = shortPaths
	}
	functions.SortingPaths(&shortPaths)
	functions.SortingPaths(&longPaths)
	functions.PathAnts(&antFarm.Ants, &shortPaths, &longPaths, &finalPath, &numberPaths)
	functions.PrintAnt(finalPath, numberPaths)
}
