package core

import (
	"fmt"
)

type AOIManger struct {
	Region
	CntX  int
	CntY  int
	Grids []*Grid
}

func NewAOIManger(width, height, upperLeftX, upperLeftY, cntX, cntY int) *AOIManger {
	aoiMgr := &AOIManger{
		Region: Region{
			UpperLeftX: upperLeftX,
			UpperLeftY: upperLeftY,
			Width:      width,
			Height:     height,
		},
		CntX:  cntX,
		CntY:  cntY,
		Grids: make([]*Grid, cntX*cntY),
	}

	gridWidth, gridHeight := width/cntX, height/cntY
	for y := 0; y < cntY; y++ {
		for x := 0; x < cntX; x++ {
			gID := y*cntX + x
			aoiMgr.Grids[gID] = NewGrid(gID, gridWidth, gridHeight, x*gridWidth, y*gridHeight)
		}
	}
	return aoiMgr
}

func (m *AOIManger) String() string {
	s := fmt.Sprintf("AOIManger: \n Region : %#v, CntX: %d, CntY: %d\n",
		m.Region, m.CntX, m.CntY)

	for _, grid := range m.Grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

func (m *AOIManger) GetSurroundGridByGid(gID, rangeX, rangeY int) (grids []*Grid) {
	idX, idY := gID%m.CntX, gID/m.CntX
	leftId, upperId := idX-(rangeX-1)/2, idY-(rangeY-1)/2
	for y := upperId; y < upperId+rangeY; y++ {
		for x := leftId; x < leftId+rangeX; x++ {
			if x < 0 || x >= m.CntX || y < 0 || y >= m.CntY {
				continue
			}
			surroundGid := y*m.CntX + x
			grids = append(grids, m.Grids[surroundGid])
		}
	}
	return
}

func (m *AOIManger) GetGidByPos(x, y float32) (gID int) {
	gridWidth, gridHeight := m.Grids[0].Width, m.Grids[0].Height
	idX, idY := x/float32(gridWidth), y/float32(gridHeight)
	gID = int(idY)*m.CntX + int(idX)
	return
}

func (m *AOIManger) GetPidsByPos(x, y float32, rangeX, rangeY int) (playersIDs []int) {
	gID := m.GetGidByPos(x, y)
	grids := m.GetSurroundGridByGid(gID, rangeX, rangeY)

	for _, grid := range grids {
		playersIDs = append(playersIDs, grid.GetPlayerIDs()...)
	}
	return
}

func (m *AOIManger) AddPidToGridByGid(pID, gID int) {
	m.Grids[gID].Add(pID)
}

func (m *AOIManger) RemovePidFromGridByGid(pID, gID int) {
	m.Grids[gID].Remove(pID)
}

func (m *AOIManger) GetPidsByGid(gID int) []int {
	return m.Grids[gID].GetPlayerIDs()
}

func (m *AOIManger) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	m.AddPidToGridByGid(pID, gID)
}

func (m *AOIManger) RemovePidFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	m.RemovePidFromGridByGid(pID, gID)
}
