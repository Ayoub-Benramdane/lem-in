package main

import (
	"fmt"
	functions "lem-in/functions"
	structs "lem-in/structs"
	"os"
	"strings"
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
	fmt.Println("File:\n" + strings.Join(file, "\n") + "\n\nLemin:")
	// functions.PrintGraph(antFarm)
	shortPaths, multiplePaths := functions.BestPaths(structs.Paths)
	if len(shortPaths) == 0 && len(multiplePaths) == 0 {
		fmt.Println("ERROR: invalid path")
		return
	} else if len(shortPaths) >= len(multiplePaths) {
		multiplePaths = shortPaths
	}
	functions.PathAnts(&antFarm.Ants, &shortPaths, &multiplePaths, &finalPath, &numberPaths)
	functions.PrintAnt(finalPath, numberPaths)
}
