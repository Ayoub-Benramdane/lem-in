package functions

import "sort"

func SortingPaths(path *[][]string) {
	sort.Slice(*path, func(i, j int) bool {
		return len((*path)[i]) < len((*path)[j])
	})
}

func CleanPath(bestPaths [][]string, multiplePaths *[][]string) {
	for _, path := range bestPaths {
		for j := 0; j < len((*multiplePaths)); j++ {
			for k := 1; k < len((*multiplePaths)[j])-1; k++ {
				if ContainsRoom(path, (*multiplePaths)[j][k]) {
					*multiplePaths = append((*multiplePaths)[:j], (*multiplePaths)[j+1:]...)
					j--
					break
				}
			}
		}
	}
}

func Length(shortPaths [][]string) int {
	count := 0
	for _, c := range shortPaths {
		count += len(c)
	}
	return count
}

func ContainsPath(paths, slices [][]string, room string) bool {
	for _, slice := range slices {
		for i := 1; i < len(slice)-1; i++ {
			if slice[i] != room {
				for _, path := range paths {
					if !SameSlice(slice, path) && ContainsRoom(path, slice[i]) {
						return true
					}
				}
			}
		}
	}
	return false
}

func ContainsRoom(path []string, room string) bool {
	for _, rom := range path {
		if rom == room {
			return true
		}
	}
	return false
}

func SameSlice(slice1, slice2 []string) bool {
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

func CheckSlice(paths *[][]string, slice []string) bool {
	for _, path := range *paths {
		for i := 1; i < len(path)-1; i++ {
			if ContainsRoom(slice, path[i]) {
				return true
			}
		}
	}
	return false
}
