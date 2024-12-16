package functions

import (
	"strings"

	structs "lem-in/structs"
)

func GetPaths(Tunnels []structs.Tunnel, Start, End string, path []string) {
	paths := &structs.Paths
	path = append(path, Start)
	if Start == End {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*paths = append(*paths, newPath)
		return
	}
	for _, Tunnel := range Tunnels {
		if Start == Tunnel.From && !ContainsRoom(path, Tunnel.To) {
			GetPaths(Tunnels, Tunnel.To, End, path)
		} else if Start == Tunnel.To && !ContainsRoom(path, Tunnel.From) {
			GetPaths(Tunnels, Tunnel.From, End, path)
		}
	}
}

func BestPaths(paths [][]string) ([][]string, [][]string) {
	var bestPaths, multiplePaths, shortPaths, longPaths [][]string
	var groupedPaths [][][]string
	SortingPaths(&paths)
	for i := 0; i < len(paths); i++ {
		UniquePaths(&multiplePaths, &bestPaths, paths, paths[i], i) // kn9lb 3la unique path bach ykono dymn valid
	}
	CleanPath(bestPaths, &multiplePaths) // knms7 les paths li kychtarko f chi room m3a best
	for i := 0; i < len(multiplePaths); i++ {
		MultipPaths(multiplePaths, &groupedPaths, multiplePaths[i], i)
	}
	group := make(map[int]bool)
	for i := 0; i < len(groupedPaths); i++ {
		if len(groupedPaths[i]) > 0 {
			i = ShortLong(&groupedPaths[i], &shortPaths, &longPaths, i, &group)
		}
	}
	return append(bestPaths, shortPaths...), append(bestPaths, longPaths...)
}

func UniquePaths(multiplePaths, bestPaths *[][]string, paths [][]string, slice []string, index int) {
	var count int
	var room string
	slices := [][]string{slice}
	for i := 0; i < len(paths); i++ { // kn9lb 3la ch7al mn room ttchark m3a path khrin
		valid := false
		if index != i {
			for j := 1; j < len(slice)-1; j++ {
				for k := 1; k < len(paths[i])-1; k++ {
					if slice[j] == paths[i][k] {
						slices = append(slices, paths[i])
						if count == 0 {
							count++
							room = slice[j]
						} else if slice[j] != room {
							count++
						}
						valid = true
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
		*bestPaths = append(*bestPaths, slice)
	} else if count == 1 && !ContainsPath(paths, slices, room) { // kn9lb wach f dok les paths kyn chi path kydi mn chi tree9 khora f chi room khora
		same := false
		for _, path := range *bestPaths {
			for _, room := range path {
				for _, slice := range slices {
					if ContainsRoom(slice, room) {
						same = true
					}
				}
			}
		}
		if !same {
			*bestPaths = append(*bestPaths, slices[0])
		}
	} else {
		*multiplePaths = append(*multiplePaths, slice)
	}
}

func MultipPaths(paths [][]string, groupedPaths *[][][]string, slice []string, index int) {
	for i := 0; i < len(paths); i++ {
		if i != index {
			for j := 1; j < len(slice)-1; j++ {
				for k := 1; k < len(paths[i])-1; k++ {
					if slice[j] == paths[i][k] {
						GroupPaths(groupedPaths, slice, slice[j]) // kn9sm l paths lgroupat kol group kychtrko f chi room
						return
					}
				}
			}
		}
	}
}

func GroupPaths(groupedPaths *[][][]string, slice []string, room string) {
	for i, paths := range *groupedPaths {
		for _, path := range paths {
			if ContainsRoom(path, room) {
				(*groupedPaths)[i] = append((*groupedPaths)[i], slice)
				return
			}
		}
	}
	*groupedPaths = append(*groupedPaths, [][]string{slice})
}

func ShortLong(groupedPaths, shortPaths, longPaths *[][]string, index int, group *map[int]bool) int {
	result := make([][]string, 0)
	paths := make(map[string]int)
	var valid bool
	if !(*group)[index] { // knchouf f kol group les path sghar wli mkytla9aw f ta room
		(*group)[index] = true
		*shortPaths = append(*shortPaths, (*groupedPaths)[0])
		for i := 1; i < len(*groupedPaths); i++ {
			var duplicated bool
			for j := 0; j < len(*shortPaths); j++ {
				for k := 1; k < len((*shortPaths)[j])-1; k++ {
					if ContainsRoom((*groupedPaths)[i], (*shortPaths)[j][k]) {
						duplicated = true
						break
					}
				}
				if duplicated {
					break
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
				if ContainsRoom((*groupedPaths)[j], (*groupedPaths)[i][k]) { // kn9lb 3la ch7al mn path f group kytchrk m3a lpath li rani feeh f room
					paths[strings.Join((*groupedPaths)[i], ",")]++
					break
				}
			}
		}
	}
	min := paths[strings.Join((*groupedPaths)[0], ",")]
	for _, count := range paths {
		if count < min {
			min = count
		}
	}
	for path, count := range paths {
		if count == min && !CheckSlice(longPaths, strings.Split(path, ",")) {
			result = append(result, strings.Split(path, ","))
			valid = true
		}
	}
	if valid {
		SortingPaths(&result)
		if CheckLen(result, len((*shortPaths)[len(*shortPaths)-1])) && NotInShort(result, (*shortPaths)[:len(*shortPaths)-1]) { // ila kan chi 7ed mn l multiple ged short f tol kn7yd short w nde5el ga3 lmultiple
			*shortPaths = (*shortPaths)[:len(*shortPaths)-1]
			for _, path := range result {
				if !CheckSlice(shortPaths, path) {
					*shortPaths = append(*shortPaths, path)
				}
			}
		} else {
			for _, path := range result {
				if !CheckSlice(longPaths, path) {
					*longPaths = append(*longPaths, path)
				}
			}
		}
		CleanPath(result, groupedPaths) // knms7 ga3 les paths li ba9yin f had lgroup w kytla9aw f chi no9ta m3a chi w7d fles paths li l9iit
		index--
	}
	return index
}

func PathAnts(TotalAnts *int, shortPaths, longPaths, finalPath *[][]string, numberPaths *[]int) {
	depart := *TotalAnts
	done := false
	for *TotalAnts > 0 {
		if *TotalAnts == depart {
			FinalPaths(TotalAnts, shortPaths, finalPath, numberPaths)
		} else if *TotalAnts <= Length(*shortPaths)/len(*shortPaths) && !done {
			FinalPaths(TotalAnts, shortPaths, finalPath, numberPaths)
		} else {
			done = true
			FinalPaths(TotalAnts, longPaths, finalPath, numberPaths)
		}
	}
}

func FinalPaths(TotalAnts *int, Paths, finalPath *[][]string, numberPaths *[]int) {
	count := 0
	for i, path := range *Paths {
		if *TotalAnts > 0 {
			if i == 0 || check(*TotalAnts, *Paths, len(path)) {
				count++
				*finalPath = append(*finalPath, path[1:])
				*TotalAnts--
			} else {
				break
			}
		}
	}
	*numberPaths = append(*numberPaths, count)
}

func check(ants int, paths [][]string, Len int) bool {
	if Len == len(paths[0]) {
		return true
	}
	for _, path := range paths {
		if len(path)+ants <= Len {
			return false
		}
	}
	return true
}
