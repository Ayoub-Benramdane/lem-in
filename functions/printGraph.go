package functions

import (
	"fmt"
	"lem-in/structs"
	"strings"
)

func PrintGraph(farm structs.AntFarm) {
	const size = 40
	var maxSize int
	rooms := []string{}
	for _, room := range farm.Rooms {
		if len(room.Name) > maxSize {
			maxSize = len(room.Name)
		}
	}
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size*4)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}
	maxX, minX, maxY, minY := farm.Rooms[0].X, farm.Rooms[0].X, farm.Rooms[0].Y, farm.Rooms[0].Y
	for _, room := range farm.Rooms {
		if room.X > maxX {
			maxX = room.X
		}
		if room.Y > maxY {
			maxY = room.Y
		}
		if room.X < minX {
			minX = room.X
		}
		if room.Y < minY {
			minY = room.Y
		}
	}
	roomPositions := make(map[string][2]int)
	for _, room := range farm.Rooms {
		rooms = append(rooms, room.Name)
		sizeRoom := len(room.Name)
		x := (room.X - minX) * (size - 2) / (maxX - minX + 1)
		y := (room.Y - minY) * (size - 2) / (maxY - minY + 1)
		roomPositions[room.Name] = [2]int{y, x * 3}
		color := ""
		if room.Name == farm.Start.Name {
			color = "\033[31m"
		} else if room.Name == farm.End.Name {
			color = "\033[32m"
		}
		grid[y][x*3] = strings.Repeat(" ", (maxSize-sizeRoom)/2) + color + room.Name + strings.Repeat(" ", (maxSize-sizeRoom-(maxSize-sizeRoom)/2)) + "\033[0m"
	}
	for _, tunnel := range farm.Tunnels {
		drawLine(grid, roomPositions[tunnel.From], roomPositions[tunnel.To], rooms)
	}
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
}

func drawLine(grid [][]string, room1, room2 [2]int, rooms []string) {
	if room1[0] == room2[0] {
		for x := min(room1[1], room2[1]); x <= max(room1[1], room2[1]); x++ {
			if grid[room1[0]][x] == " " {
				grid[room1[0]][x] = "-"
			}
		}
	} else if room1[1] == room2[1] {
		for y := min(room1[0], room2[0]); y <= max(room1[0], room2[0]); y++ {
			if grid[y][room1[1]] == " " {
				grid[y][room1[1]] = "|"
			}
		}
	} else {
		isBackslash := (room2[0]-room1[0])*(room2[1]-room1[1]) > 0
		char := "/"
		if isBackslash {
			char = "\\"
		}
		steps := max(abs(room2[1]-room1[1]), abs(room2[0]-room1[0]))
		for i := 0; i < steps; i++ {
			y := room1[0] + (room2[0]-room1[0])*i/steps
			x := room1[1] + (room2[1]-room1[1])*i/steps
			for j := 1; j < 4; j++ {
				if x+j < len(grid[y]) && grid[y][x+j] == " " {
					if Contains(rooms, grid[y][x]) {
						break
					}
					grid[y][x+j] = char
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
