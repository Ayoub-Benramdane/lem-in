package functions

import (
	"bufio"
	"fmt"
	structs "lem-in/structs"
	"os"
	"strings"
)

func ReadInput(fileName string) structs.AntFarm {
	var af structs.AntFarm
	var state string
	var rm structs.Room
	var tn structs.Tunnel
	file, err := os.Open(fileName)
	CheckErrors(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
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
		parts := strings.Fields(line)
		if len(parts) == 1 && i == 0 && !strings.Contains(parts[0], "-") {
			_, err := fmt.Sscanf(line, "%d", &af.Ants)
			CheckErrors(err)
		} else if len(parts) == 3 {
			if len(parts[0]) > 0 && parts[0][0] == 'L' {
				ErrorRoom(parts[0])
			}
			rm = ToInt(parts)
			CheckErrors(err)
			CheckCordonnes(af.Rooms, rm)
			if state == "start" {
				CheckIsDouble(af.Start.Name, state)
				af.Start = rm
				state = ""
			} else if state == "end" {
				CheckIsDouble(af.End.Name, state)
				af.End = rm
				state = ""
			}
			af.Rooms = append(af.Rooms, rm)
		} else if len(parts) == 1 && strings.Contains(parts[0], "-") {
			if parts = strings.Split(line, "-"); len(parts) == 2 {
				tn.From = parts[0]
				tn.To = parts[1]
				af.Tunnels = append(af.Tunnels, tn)
				if tn.From == tn.To {
					ErrorData()
				}
			} else {
				ErrorData()
			}
		} else {
			ErrorData()
		}
	}
	CheckErrors(scanner.Err())
	if af.Ants <= 0 || len(af.Tunnels) <= 0 {
		ErrorData()
	} else if af.Start.Name == "" || af.End.Name == "" {
		ErrorStartEndRoom()
	}
	return af
}
