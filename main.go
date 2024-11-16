package main

import (
	functions "lem-in/functions"
	structs "lem-in/structs"
	"os"
)

func main() {
	var path []string
	finalPath := [][]string{}
	numberPaths := []int{}
	functions.ErrorArgs(os.Args)
	antFarm := functions.ReadInput(os.Args[1])
	functions.GetPaths(antFarm.Tunnels, antFarm.Start.Name, antFarm.End.Name, path)
	functions.CheckPath(structs.Paths)
	shortPaths, multiplePaths := functions.BestPaths(structs.Paths)
	functions.PathAnts(&antFarm.Ants, &shortPaths, &multiplePaths, &finalPath, &numberPaths)
	functions.PrintAnt(finalPath, numberPaths)
}
