package functions

import (
	"bufio"
	"fmt"
	structs "lem-in/structs"
	"os"
	"strconv"
	"strings"
)

func ReadInput(fileName string) (structs.AntFarm, []string, error) {
	var farm structs.AntFarm; var state string; var room structs.Room; var tunnel structs.Tunnel; var exportFile []string
	file, err := os.Open(fileName)
	if err != nil {
		return structs.AntFarm{}, []string{}, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		exportFile = append(exportFile, scanner.Text())
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			switch line {
			case "##start": state = "start"
			case "##end": state = "end"
			}
			continue
		}
		parts := strings.Fields(line)
		if len(parts) == 1 && i == 0 && !strings.Contains(parts[0], "-") {
			ants, err := strconv.Atoi(parts[0])
			if err != nil || ants < 1 {
				return structs.AntFarm{}, []string{}, fmt.Errorf("invalid number of ants")
			}
			farm.Ants = ants
		} else if len(parts) == 3 {
			if strings.HasPrefix(parts[0], "L") {
				return structs.AntFarm{}, []string{}, fmt.Errorf("room name %s cannot start with 'L'", parts[0])
			}
			room, err = ParseRoom(parts)
			if err != nil {
				return structs.AntFarm{}, []string{}, fmt.Errorf("failed to parse room: %v", parts[0])
			} else if err = CheckCordonnes(farm.Rooms, room); err != nil {
				return structs.AntFarm{}, []string{}, err
			}
			if state == "start" {
				if farm.Start != (structs.Room{}) {
					return structs.AntFarm{}, []string{}, fmt.Errorf("duplicate start room defined")
				}
				farm.Start, state = room, ""
			} else if state == "end" {
				if farm.End != (structs.Room{}) {
					return structs.AntFarm{}, []string{}, fmt.Errorf("duplicate end room defined")
				}
				farm.End, state = room, ""
			}
			farm.Rooms = append(farm.Rooms, room)
		} else if len(parts) == 1 && strings.Contains(parts[0], "-") {
			if parts = strings.Split(line, "-"); len(parts) == 2 {
				if parts[0] == "" || parts[1] == "" {
					return structs.AntFarm{}, []string{}, fmt.Errorf("invalid format tunnel")
				}
				tunnel.From, tunnel.To = parts[0], parts[1]
				if tunnel.From == tunnel.To {
					return structs.AntFarm{}, []string{}, fmt.Errorf("invalid format from == to")
				} else if err = ParseTunnel(farm.Tunnels, tunnel.From, tunnel.To); err != nil {
					return structs.AntFarm{}, []string{}, err
				}
				farm.Tunnels = append(farm.Tunnels, tunnel)
			} else {
				return structs.AntFarm{}, []string{}, fmt.Errorf("invalid data format")
			}
		} else {
			return structs.AntFarm{}, []string{}, fmt.Errorf("invalid data format")
		}
	}
	if err = scanner.Err(); err != nil {
		return structs.AntFarm{}, []string{}, err
	} else if farm.Ants <= 0 || len(farm.Tunnels) <= 0 {
		return structs.AntFarm{}, []string{}, fmt.Errorf("invalid data format")
	} else if farm.Start.Name == "" || farm.End.Name == "" {
		return structs.AntFarm{}, []string{}, fmt.Errorf("invalid Path! Check Start and End Rooms")
	}
	return farm, exportFile, nil
}

func ParseTunnel(Tunnels []structs.Tunnel, From, To string) error {
	for _, tunnel := range Tunnels {
		if tunnel.From == From && tunnel.To == To || tunnel.From == To && tunnel.To == From {
			return fmt.Errorf("duplicate tunnel defined for rooms '%s' and '%s'", From, To)
		}
	}
	return nil
}

func ParseRoom(parts []string) (structs.Room, error) {
	var room structs.Room; var err, err1 error
	room.Name = parts[0]
	room.X, err = strconv.Atoi(parts[1])
	room.Y, err1 = strconv.Atoi(parts[2])
	if err != nil {
		return structs.Room{}, err
	} else if err1 != nil {
		return structs.Room{}, err1
	}
	return room, nil
}

func CheckCordonnes(rooms []structs.Room, room structs.Room) error {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].Name == room.Name {
			return fmt.Errorf("room %s is double", rooms[i].Name)
		} else if rooms[i].X == room.X && rooms[i].Y == room.Y {
			return fmt.Errorf("duplicate coordinates detected for rooms '%s' and '%s'", rooms[i].Name, room.Name)
		}
	}
	return nil
}
