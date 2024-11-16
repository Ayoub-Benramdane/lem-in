package functions

import "strings"

func Contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
func Check(multiplePaths *[][]string, newPaths *[][]string, path *[][]string, sli []string, index int) {
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

func CheckMultip(path *[][]string, groupedPaths *[][][]string, sli []string, index int) {
	if CheckSlice(groupedPaths, sli) {
		return
	}
	for i := 0; i < len(*path); i++ {
		for j := 1; j < len(sli)-1; j++ {
			for k := 1; k < len((*path)[i])-1; k++ {
				if sli[j] == (*path)[i][k] && index != i {
					GroupPaths(groupedPaths, sli, sli[j])
					return
				}
			}
		}
	}
}

func CheckSlice(groupedPaths *[][][]string, sli []string) bool {
	for _, group := range *groupedPaths {
		for _, slices := range group {
			if SlicesEqual(slices, sli) {
				return true
			}
		}
	}
	return false
}

func CheckSlice1(groupedPaths *[][]string, sli []string) bool {
	for _, slice := range *groupedPaths {
		for i := 1; i < len(slice)-1; i++ {
			if Contains(sli, slice[i]) {
				return true
			}
		}
	}
	return false
}

func SlicesEqual(slice1, slice2 []string) bool {
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

func UniqueSlices(groupedPaths *[][]string, shortPaths, multiplePaths *[][]string, index int) int {
	result := make([][]string, 0)
	paths := make(map[string]int)
	var valid bool
	if len(*shortPaths) == index {
		*shortPaths = append(*shortPaths, (*groupedPaths)[0])
		for i := 1; i < len(*groupedPaths); i++ {
			var duplicated bool
			for j := 0; j < len(*shortPaths); j++ {
				for k := 1; k < len((*shortPaths)[j])-1; k++ {
					if Contains((*groupedPaths)[i], (*shortPaths)[j][k]) {
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
				if Contains((*groupedPaths)[j], (*groupedPaths)[i][k]) {
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
			if count == min && !CheckSlice1(multiplePaths, strings.Split(path, ",")) {
				result = append(result, strings.Split(path, ","))
				valid = true
			}
		}
	}
	SortingPaths(&result)
	if valid {
		index--
		*multiplePaths = append(*multiplePaths, result[0])
		DeleteSlice(groupedPaths, result[0])
	}
	return index
}
