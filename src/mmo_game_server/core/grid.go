package core

import (
	"fmt"
	"sync"
)

type Point struct {
	X int
	Y int
}

type Grid struct {
	Gid       int
	Width     int
	Height    int
	UpperLeft Point
	playerIDs map[int]bool
	pIDLock   sync.RWMutex
}

func NewGrid(gid, width, height int, upperLeft Point) *Grid {
	return &Grid{
		Gid:       gid,
		Width:     width,
		Height:    height,
		UpperLeft: upperLeft,
		playerIDs: make(map[int]bool),
		pIDLock:   sync.RWMutex{},
	}
}

func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerID] = true
}

func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerID)
}

func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()

	for k := range g.playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d, UpperLeft Point: %#v, Width: %d, Height: %d, Players: %v",
		g.Gid, g.UpperLeft, g.Width, g.Height, g.playerIDs)
}
