package main

import (
	functions "lem-in/functions"
	structs "lem-in/structs"
	"os"
)

func main() {
	path, finalPath, numberPaths := []string{}, [][]string{}, []int{}
	functions.ErrorArgs(os.Args)
	antFarm := functions.ReadInput(os.Args[1])
	functions.GetPaths(antFarm.Tunnels, antFarm.Start.Name, antFarm.End.Name, path)
	functions.CheckPath(structs.Paths)
	functions.PrintGraph(antFarm)
	shortPaths, multiplePaths := functions.BestPaths(structs.Paths)
	functions.PathAnts(&antFarm.Ants, &shortPaths, &multiplePaths, &finalPath, &numberPaths)
	functions.PrintAnt(finalPath, numberPaths)
}
