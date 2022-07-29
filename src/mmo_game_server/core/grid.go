package core

import (
	"fmt"
	"sync"
)

type Region struct {
	UpperLeftX int
	UpperLeftY int
	Width      int
	Height     int
}

type Grid struct {
	Gid int
	Region
	playerIDs map[int]bool
	pIDLock   sync.RWMutex
}

func NewGrid(gid, width, height, upperLeftX, upperLeftY int) *Grid {
	return &Grid{
		Gid: gid,
		Region: Region{
			UpperLeftX: upperLeftX,
			UpperLeftY: upperLeftY,
			Width:      width,
			Height:     height,
		},
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
	return fmt.Sprintf("Grid id: %d, Region: %#v, Players: %v",
		g.Gid, g.Region, g.playerIDs)
}
