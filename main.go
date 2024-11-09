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
						return AntFarm{}, fmt.Errorf("error: invalid tunnel")
					}
				} else {
					return AntFarm{}, fmt.Errorf("ERROR: bad Tunnel Format")
				}
			} else {
				return AntFarm{}, fmt.Errorf("ERROR: bad data Format")
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

func check(mpt *[][]string, npt *[][]string, pt *[][]string, sli []string, index int) int {
	count := 0
	valid := false
	slices := [][]string{sli}
	for i := 0; i < len(*pt); i++ {
		for j := 1; j < len(sli)-1; j++ {
			for k := 1; k < len((*pt)[i])-1; k++ {
				if sli[j] == (*pt)[i][k] && index != i {
					slices = append(slices, (*pt)[i])
					count++
					valid = true
				}
			}
			if valid {
				break
			}
		}
	}
	if count == 0 {
		*npt = append(*npt, sli)
	} else if count == 1 && len(slices) != 0 && len(sli) < len(slices[1]) {
		*npt = append(*npt, sli)
	} else if count >= 2 {
		*mpt = append(*mpt, sli)
	}
	return index
}

func checkSlice(gpt *[][][]string, sli []string) bool {
	for _, group := range *gpt {
		for _, slices := range group {
			if slicesEqual(slices, sli) {
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

func groupPaths(gpt *[][][]string, sli []string, element string) {
	for i, group := range *gpt {
		for _, slices := range group {
			for _, elem := range slices {
				if elem == element {
					(*gpt)[i] = append((*gpt)[i], sli)
					return
				}
			}
		}
	}
	*gpt = append(*gpt, [][]string{sli})
}

func checkMultip(pt *[][]string, gpt *[][][]string, sli []string, index, ants int) {
	if checkSlice(gpt, sli) {
		return
	}
	for i := 0; i < len(*pt); i++ {
		for j := 1; j < len(sli)-1; j++ {
			for k := 1; k < len((*pt)[i])-1; k++ {
				if sli[j] == (*pt)[i][k] && index != i {
					groupPaths(gpt, sli, sli[j])
					return
				}
			}
		}
	}
}

func uniqueSlices(gpt [][]string) [][]string {
	result := make([][]string, 0)
	pt := []string{}
	pt1 := []string{}
	paths := make(map[string]int)
	for i := 0; i < len(gpt); i++ {
		for j := 0; j < len(gpt); j++ {
			if i != j {
				if len(gpt[i]) < len(gpt[i]) {
					pt = gpt[i]
					pt1 = gpt[j]
				} else {
					pt = gpt[j]
					pt1 = gpt[i]
				}
				for k := 1; k < len(pt)-1; k++ {
					if contains(pt1, pt[k]) {
						key := strings.Join(gpt[i], ",")
						paths[key]++
						break
					}
				}
			}
		}
	}
	for path, count := range paths {
		if count < len(gpt)-1 {
			result = append(result, strings.Split(path, ","))
		}
	}
	return result
}

func bestPaths(pt [][]string, ants int) [][]string {
	var npt [][]string
	var mpt [][]string
	var res [][]string
	var gpt [][][]string
	sort.Slice(pt, func(i, j int) bool {
		return len(pt[i]) < len(pt[j])
	})
	for i := 0; i < len(pt); i++ {
		i = check(&mpt, &npt, &pt, pt[i], i)
	}
	for i := 0; i < len(mpt); i++ {
		checkMultip(&mpt, &gpt, mpt[i], i, ants)
	}
	for _, gr := range gpt {
		res = append(res, uniqueSlices(gr)...)
	}
	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})
	return append(npt, res...)
}

func getPaths(af []Tunnel, start string, end string, pa []string) {
	pa = append(pa, start)
	for h := 0; h < len(af); h++ {
		if start == end {
			cpy := make([]string, len(pa))
			copy(cpy, pa)
			paths = append(paths, cpy)
			return
		} else if start == af[h].from && !contains(pa, af[h].to) {
			getPaths(af, af[h].to, end, pa)
		} else if start == af[h].to && !contains(pa, af[h].from) {
			getPaths(af, af[h].from, end, pa)
		}
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
	}
	var pt []string
	getPaths(antFarm.tunnels, antFarm.start.name, antFarm.end.name, pt)
	best := bestPaths(paths, antFarm.ants)
	fmt.Println(best, "best")
}
