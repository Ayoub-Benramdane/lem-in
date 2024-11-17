package functions

import (
	"fmt"
	structs "lem-in/structs"
	"os"
)

func ErrorArgs(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		os.Exit(0)
	}
}

func CheckPath(paths [][]string) {
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid path")
		os.Exit(0)
	}
}

func CheckErrors(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func CheckIsDouble(room, state string) {
	if len(room) > 0 {
		fmt.Printf("ERROR : room %s is double\n", state)
		os.Exit(0)
	}
}

func ErrorData() {
	fmt.Println("ERROR: invalid data format")
	os.Exit(0)
}

func CheckCordonnes(rooms []structs.Room, room structs.Room) {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].Name == room.Name {
			fmt.Printf("ERROR : room %s is double\n", rooms[i].Name)
			os.Exit(0)
		} else if rooms[i].X == room.X && rooms[i].Y == room.Y {
			fmt.Printf("ERROR: Duplicate coordinates detected for rooms '%s' and '%s'.\n", rooms[i].Name, room.Name)
			os.Exit(0)
		}
	}
}
