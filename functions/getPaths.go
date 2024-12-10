package functions

import (
	structs "lem-in/structs"
	"sort"
	"strconv"
)

func ToInt(lay []string) structs.Room {
	if len(lay) != 3 {
		ErrorData()
	}
	var rtn structs.Room
	var err, err1 error
	rtn.Name = lay[0]
	rtn.X, err = strconv.Atoi(lay[1])
	rtn.Y, err1 = strconv.Atoi(lay[2])
	if err != nil {
		CheckErrors(err)
	} else if err1 != nil {
		CheckErrors(err1)
	}
	return rtn
}

func GroupPaths(groupedPaths *[][][]string, sli []string, element string) {
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

func DeleteSlice(slices *[][]string, sli []string) {
	for j := 0; j < len(*slices); j++ {
		for i := 1; i < len(sli)-1; i++ {
			if Contains((*slices)[j], sli[i]) {
				*slices = append((*slices)[:j], (*slices)[j+1:]...)
				j--
				break
			}
		}
	}
}

func BestPaths(path [][]string) ([][]string, [][]string) {
	var newPaths [][]string
	var multiplePaths [][]string
	var shortPaths [][]string
	var multiple [][]string
	var groupedPaths [][][]string
	SortingPaths(&path)
	for i := 0; i < len(path); i++ {
		Check(&multiplePaths, &newPaths, &path, path[i], i)
	}
	for i := 0; i < len(multiplePaths); i++ {
		CheckMultip(&multiplePaths, &groupedPaths, multiplePaths[i], i)
	}
	for i := 0; i < len(groupedPaths); i++ {
		i = UniqueSlices(&groupedPaths[i], &shortPaths, &multiple, i)
	}
	return append(newPaths, shortPaths...), append(newPaths, multiple...)
}

func GetPaths(af []structs.Tunnel, Start string, End string, path []string) {
	var paths = &structs.Paths
	path = append(path, Start)
	for h := 0; h < len(af); h++ {
		if Start == End {
			cpy := make([]string, len(path))
			copy(cpy, path)
			*paths = append(*paths, cpy)
			return
		} else if Start == af[h].From && !Contains(path, af[h].To) {
			GetPaths(af, af[h].To, End, path)
		} else if Start == af[h].To && !Contains(path, af[h].From) {
			GetPaths(af, af[h].From, End, path)
		}
	}
}

func SortingPaths(path *[][]string) {
	sort.Slice(*path, func(i, j int) bool {
		return len((*path)[i]) < len((*path)[j])
	})
}

func Length(shortPaths [][]string) int {
	count := 0
	for _, c := range shortPaths {
		count += len(c)
	}
	return count
}

func FinalPaths(TotalAnts *int, Paths, finalPath *[][]string, numberPaths *[]int) {
	count := 0
	for _, c := range *Paths {
		if *TotalAnts > 0 {
			if len(c)-*TotalAnts > *TotalAnts && len(c)-*TotalAnts > len((*Paths)[0]) {
				break
			}
			count++
			*finalPath = append(*finalPath, c[1:])
			*TotalAnts--
		}
	}
	*numberPaths = append(*numberPaths, count)
}

func PathAnts(TotalAnts *int, shortPaths, multiplePaths, finalPath *[][]string, numberPaths *[]int) {
	depart := *TotalAnts
	done := false
	lengthpaths := Length(*shortPaths)
	for *TotalAnts > 0 {
		if *TotalAnts == depart {
			FinalPaths(TotalAnts, shortPaths, finalPath, numberPaths)
		} else if *TotalAnts <= lengthpaths/len(*shortPaths) && !done {
			FinalPaths(TotalAnts, shortPaths, finalPath, numberPaths)
		} else {
			done = true
			FinalPaths(TotalAnts, multiplePaths, finalPath, numberPaths)
		}
	}
}
