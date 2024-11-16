package structs

var Paths [][]string

type Room struct {
	Name string
	X    int
	Y    int
}

type Tunnel struct {
	From string
	To   string
}

type AntFarm struct {
	Rooms   []Room
	Tunnels []Tunnel
	Start   Room
	End     Room
	Ants    int
}
