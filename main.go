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

func check(mpt *[][]string, npt *[][]string, pt *[][]string, sli []string, index int) {
	count := 0
	indexs := []int{}
	for i := 0; i < len(*pt); i++ {
		for j := 1; j < len(sli)-1; j++ {
			for k := 1; k < len((*pt)[i])-1; k++ {
				if sli[j] == (*pt)[i][k] && index != i {
					indexs = append(indexs, i)
					count++
				}
			}
		}
	}
	if count == 0 {
		*npt = append(*npt, sli)
	} else if count == 1 && len(indexs) != 0 && len(sli) < len((*pt)[indexs[0]]) {
		*npt = append(*npt, sli)
	} else if count != 1 {
		*mpt = append(*mpt, sli)
	}
}

func bestpath(pt [][]string) [][]string {
	var npt [][]string; var mpt [][]string
	sort.Slice(pt, func(i, j int) bool {
		return len(pt[i]) < len(pt[j])
	})
	for i := 0; i < len(pt); i++ {
		check(&mpt, &npt, &pt, pt[i], i)
	}
	fmt.Println(mpt, "multip")
	return npt
}

func getpaths(af []Tunnel, start string, end string, pa []string) {
	pa = append(pa, start)
	for h := 0; h < len(af); h++ {
		if start == end {
			cpy := make([]string, len(pa))
			copy(cpy, pa)
			paths = append(paths, cpy)
			return
		} else if start == af[h].from && !contains(pa, af[h].to) {
			getpaths(af, af[h].to, end, pa)
		} else if start == af[h].to && !contains(pa, af[h].from) {
			getpaths(af, af[h].from, end, pa)
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
	getpaths(antFarm.tunnels, antFarm.start.name, antFarm.end.name, pt)
	fmt.Println(paths, "paths")
	best := bestpath(paths)
	fmt.Println(best, "best")
}
