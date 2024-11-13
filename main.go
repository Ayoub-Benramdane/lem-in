package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	name string
	x    int
	y    int
}

type Tunnel struct {
	from string
	to   string
}

type AntFarm struct {
	rooms   []Room
	tunnels []Tunnel
	start   Room
	end     Room
	ants    int
}

func tointeg(lay []string) (Room, error) {
	var rtn Room
	var err, err1 error
	rtn.name = lay[0]
	rtn.x, err = strconv.Atoi(lay[1])
	rtn.y, err1 = strconv.Atoi(lay[2])
	if err != nil {
		return Room{}, err
	} else if err1 != nil {
		return Room{}, err
	}
	return rtn, nil
}

func readInput(filename string) (AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return AntFarm{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var af AntFarm
	var state string
	var rm Room
	var tn Tunnel
	for i := 0; scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "##") {
			if line == "##start" {
				state = "start"
			} else if line == "##end" {
				state = "end"
			}
			continue
		} else if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		switch state {
		case "start":
			if len(af.start.name) > 0 {
				return AntFarm{}, fmt.Errorf("error : double start room")
			}
			af.start, err = tointeg(strings.Split(line, " "))
			state = ""
		case "end":
			if len(af.end.name) > 0 {
				return AntFarm{}, fmt.Errorf("error : double end room")
			}
			af.end, err = tointeg(strings.Split(line, " "))
			state = ""
		default:
			parts := strings.Fields(line)
			if len(parts) == 1 && i == 0 && !strings.Contains(parts[0], "-") {
				_, err := fmt.Sscanf(line, "%d", &af.ants)
				if err != nil {
					return AntFarm{}, err
				}
			} else if len(parts) == 3 {
				rm, err = tointeg(parts)
				af.rooms = append(af.rooms, rm)
			} else if len(parts) == 1 && strings.Contains(parts[0], "-") {
				if parts = strings.Split(line, "-"); len(parts) == 2 {
					tn.from = parts[0]
					tn.to = parts[1]
					af.tunnels = append(af.tunnels, tn)
					if tn.from == tn.to {
						return AntFarm{}, fmt.Errorf("ERROR: invalid data format")
					}
				} else {
					return AntFarm{}, fmt.Errorf("ERROR: invalid data format")
				}
			} else {
				return AntFarm{}, fmt.Errorf("ERROR: invalid data format")
			}
		}
		if err != nil {
			return AntFarm{}, err
		}
	}
	if scanner.Err() != nil {
		return AntFarm{}, scanner.Err()
	}
	return af, nil
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

var paths [][]string

func check(multiplePaths *[][]string, newPaths *[][]string, path *[][]string, sli []string, index int) {
	count := 0
	slices := [][]string{sli}
	for i := 0; i < len(*path); i++ {
		valid := false
		if index != i {
			for j := 1; j < len(sli)-1; j++ {
				for k := 1; k < len((*path)[i])-1; k++ {
					if sli[j] == (*path)[i][k] {
						slices = append(slices, (*path)[i])
						count++
						valid = true
					}
					if valid {
						break
					}
				}
				if valid {
					break
				}
			}
		}
	}
	if count == 0 {
		*newPaths = append(*newPaths, sli)
	} else if count == 1 && len(slices) != 0 && len(sli) <= len(slices[1]) {
		*newPaths = append(*newPaths, sli)
	} else if count >= 2 {
		*multiplePaths = append(*multiplePaths, sli)
	}
}

func checkSlice(groupedPaths *[][][]string, sli []string) bool {
	for _, group := range *groupedPaths {
		for _, slices := range group {
			if slicesEqual(slices, sli) {
				return true
			}
		}
	}
	return false
}

func checkSlice1(groupedPaths *[][]string, sli []string) bool {
	for _, slice := range *groupedPaths {
		for i := 1; i < len(slice)-1; i++ {
			if contains(sli, slice[i]) {
				return true
			}
		}
	}
	return false
}

func slicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func groupPaths(groupedPaths *[][][]string, sli []string, element string) {
	for i, group := range *groupedPaths {
		for _, slices := range group {
			for _, elem := range slices {
				if elem == element {
					(*groupedPaths)[i] = append((*groupedPaths)[i], sli)
					return
				}
			}
		}
	}
	*groupedPaths = append(*groupedPaths, [][]string{sli})
}

func checkMultip(path *[][]string, groupedPaths *[][][]string, sli []string, index int) {
	if checkSlice(groupedPaths, sli) {
		return
	}
	for i := 0; i < len(*path); i++ {
		for j := 1; j < len(sli)-1; j++ {
			for k := 1; k < len((*path)[i])-1; k++ {
				if sli[j] == (*path)[i][k] && index != i {
					groupPaths(groupedPaths, sli, sli[j])
					return
				}
			}
		}
	}
}

func uniqueSlices(groupedPaths *[][]string, shortPaths, multiplePaths *[][]string, index int) int {
	result := make([][]string, 0)
	paths := make(map[string]int)
	var valid bool
	if len(*shortPaths) == index {
		*shortPaths = append(*shortPaths, (*groupedPaths)[0])
		for i := 1; i < len(*groupedPaths); i++ {
			var duplicated bool
			for j := 0; j < len(*shortPaths); j++ {
				for k := 1; k < len((*shortPaths)[j])-1; k++ {
					if contains((*groupedPaths)[i], (*shortPaths)[j][k]) {
						duplicated = true
						break
					}
				}
			}
			if !duplicated {
				*shortPaths = append(*shortPaths, (*groupedPaths)[i])
			}
		}
	}
	for i := 0; i < len(*groupedPaths); i++ {
		for j := 0; j < len(*groupedPaths); j++ {
			for k := 1; k < len((*groupedPaths)[i])-1; k++ {
				if contains((*groupedPaths)[j], (*groupedPaths)[i][k]) {
					key := strings.Join((*groupedPaths)[i], ",")
					paths[key]++
					break
				}
			}
		}
	}
	if len(*groupedPaths) != 0 {
		min := paths[strings.Join((*groupedPaths)[0], ",")]
		for _, count := range paths {
			if count < min {
				min = count
			}
		}
		for path, count := range paths {
			if count == min && !checkSlice1(multiplePaths, strings.Split(path, ",")) {
				result = append(result, strings.Split(path, ","))
				valid = true
			}
		}
	}
	sortingPaths(&result)
	if valid {
		index--
		*multiplePaths = append(*multiplePaths, result[0])
		deleteSlice(groupedPaths, result[0])
	}
	return index
}

func deleteSlice(slices *[][]string, sli []string) {
	for j := 0; j < len(*slices); j++ {
		for i := 1; i < len(sli)-1; i++ {
			if contains((*slices)[j], sli[i]) {
				*slices = append((*slices)[:j], (*slices)[j+1:]...)
				j--
				break
			}
		}
	}
}

func bestPaths(path [][]string) ([][]string, [][]string) {
	var newPaths [][]string
	var multiplePaths [][]string
	var shortPaths [][]string
	var multiple [][]string
	var groupedPaths [][][]string
	sortingPaths(&path)
	for i := 0; i < len(path); i++ {
		check(&multiplePaths, &newPaths, &path, path[i], i)
	}
	for i := 0; i < len(multiplePaths); i++ {
		checkMultip(&multiplePaths, &groupedPaths, multiplePaths[i], i)
	}
	for i := 0; i < len(groupedPaths); i++ {
		i = uniqueSlices(&groupedPaths[i], &shortPaths, &multiple, i)
	}
	return append(newPaths, shortPaths...), append(newPaths, multiple...)
}

func getPaths(af []Tunnel, start string, end string, path []string) {
	path = append(path, start)
	for h := 0; h < len(af); h++ {
		if start == end {
			cpy := make([]string, len(path))
			copy(cpy, path)
			paths = append(paths, cpy)
			return
		} else if start == af[h].from && !contains(path, af[h].to) {
			getPaths(af, af[h].to, end, path)
		} else if start == af[h].to && !contains(path, af[h].from) {
			getPaths(af, af[h].from, end, path)
		}
	}
}

func sortingPaths(path *[][]string) {
	sort.Slice(*path, func(i, j int) bool {
		return len((*path)[i]) < len((*path)[j])
	})
}

func length(shortPaths [][]string) int {
	count := 0
	for _, c := range shortPaths {
		count += len(c)
	}
	return count
}

func finalPaths(totalAnts *int, Paths, finalPath *[][]string, numberPaths *[]int) {
	count := 0
	for _, c := range *Paths {
		if *totalAnts > 0 {
			if len(c)-*totalAnts > *totalAnts && len(c)-*totalAnts > len((*Paths)[0]) {
				break
			}
			count++
			*finalPath = append(*finalPath, c[1:])
			*totalAnts--
		}
	}
	*numberPaths = append(*numberPaths, count)
}

func printAnt(finalPath [][]string, path []int) {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	filename := os.Args[1]
	antFarm, err := readInput(filename)
	if err != nil {
		fmt.Println(err)
		return
	} else if antFarm.ants <= 0 {
		fmt.Println("invalid number of Ants")
		return
	} else if len(antFarm.tunnels) <= 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	var path []string
	getPaths(antFarm.tunnels, antFarm.start.name, antFarm.end.name, path)
	if len(paths) <= 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	shortPaths, multiplePaths := bestPaths(paths)
	depart := antFarm.ants
	finalPath := [][]string{}
	numberPaths := []int{}
	done := false
	lengthpaths := length(shortPaths)
	for antFarm.ants > 0 {
		if antFarm.ants == depart {
			finalPaths(&antFarm.ants, &shortPaths, &finalPath, &numberPaths)
		} else if antFarm.ants <= lengthpaths/len(shortPaths) && !done {
			finalPaths(&antFarm.ants, &shortPaths, &finalPath, &numberPaths)
		} else {
			done = true
			finalPaths(&antFarm.ants, &multiplePaths, &finalPath, &numberPaths)
		}
	}
	printAnt(finalPath, numberPaths)
}
