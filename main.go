package main

import (
	"bufio"
	"fmt"
	"os"
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

// Read the input file and initialize the ant farm structure
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

// Display the ant farm information
// func (af AntFarm) display() {
// 	fmt.Println(af.ants)
// 	for _, room := range af.rooms {
// 		fmt.Printf("%s %d %d\n", room.name, room.x, room.y)
// 	}
// 	for _, tunnel := range af.tunnels {
// 		fmt.Printf("%s-%s\n", tunnel.from, tunnel.to)
// 	}
// }

// func howpath(rrm []Tunnel,star string) int {
// 	rtn := 0
// 	// var pa []string
// 	for i:= 0 ;i<len(rrm);i++{
// 		if star == rrm[i].from || star == rrm[i].to {
// 			rtn++
// 		}
// 	}
// 	return rtn
// }

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

var paths [][]string

func smalest(pt *[][]string) []string {
	s := 0
	ix := -1
	for i := 0; i < len(*pt); i++ {
		fmt.Println(pt, (*pt)[i], i, "looool")
		if s == 0 && len((*pt)[i]) > 0 && check(pt, (*pt)[i]) == 1 {
			s = len((*pt)[i])
			ix = i
		}
		if len((*pt)[i]) < s && len((*pt)[i]) != 0 {
			s = len((*pt)[i])
			ix = i
		}
	}
	if ix == -1 {
		return nil
	}
	sm := make([]string, len((*pt)[ix]))
	copy(sm, (*pt)[ix])
	(*pt)[ix] = []string{}
	return sm
}

func check(npt *[][]string, sli []string) int {
	count := 0
	for i := 0; i < len(*npt); i++ {
		for j := 1; j < len(sli)-1; j++ {
			for k := 1; k < len((*npt)[i])-1; k++ {
				if sli[j] == (*npt)[i][k] {
					count++
				}
			}
		}
	}
	fmt.Println(count, sli)
	return count
}

func bestpath(pt [][]string) [][]string {
	var npt [][]string
	for i := 0; i < len(pt); i++ {
		sli := smalest(&pt)
		fmt.Println(sli, "lool")
		if sli == nil {
			return npt
		}
		// if check(&npt, sli) == 1 {
			cp := make([]string, len(sli))
			copy(cp, sli)
			npt = append(npt, cp)
		// }
	}
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

func total(nbr int, best [][]string) []int {
	var ants []int
	for o := 0; o < len(best); o++ {
		ants = append(ants, 0)
	}

	for i := 0; i < nbr; i++ {
		c := 0
		for l := 0; l < len(best); l++ {
			if ants[l]+len(best[l]) < ants[c]+len(best[c]) {
				c = l
			}
		}
		ants[c]++
	}
	return ants
}

func antrun(antf AntFarm, best [][]string) {
	ants := total(antf.ants, best)
	fmt.Println(ants, best)

	// first := 1

	// for i := antf.ants;i>0;{
	// 	a := (antf.ants - i) + 1
	// 	for j:= 0;j < len(best) ; j++{
	// 		lol := len(best)
	// 		for an:= 0;an <  {

	// 		}
	// 		// fmt.Println(len(ants))
	// 		// os.Exit(0)
	// 		// fmt.Println(ants)
	// 		for u:= ants[j][0];u> ants[j][1] && u>= 1 ;u-- {
	// 			fmt.Printf("L%d-%s ",a,best[j][u])
	// 		}
	// 		ants[j][0]++
	// 		ants[j][1]++	// first := 1

	// for i := antf.ants;i>0;{
	// 	a := (antf.ants - i) + 1
	// 	for j:= 0;j < len(best) ; j++{
	// 		lol := len(best)
	// 		for an:= 0;an <  {

	// 		}
	// 		// fmt.Println(len(ants))
	// 		// os.Exit(0)
	// 		// fmt.Println(ants)
	// 		for u:= ants[j][0];u> ants[j][1] && u>= 1 ;u-- {
	// 			fmt.Printf("L%d-%s ",a,best[j][u])
	// 		}
	// 		ants[j][0]++
	// 		ants[j][1]++
	// 		if ants[j][0] == len(best[j])  {
	// 			ants[j][0]--
	// 			i--
	// 			// a := (antf.ants - i) + 1
	// 		}
	// 		// os.Exit(0)
	// 	}
	// // 	fmt.Println()
	// }
	// 		if ants[j][0] == len(best[j]name)  {
	// 			ants[j][0]--
	// 			i--
	// 			// a := (antf.ants - i) + 1
	// 		}
	// 		// os.Exit(0)
	// 	}
	// // 	fmt.Println()
	// }
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
	fmt.Println(paths)
	best := bestpath(paths)
	fmt.Println(best)
	antrun(antFarm, best)
}
