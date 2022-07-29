package unittest

import (
	"Zserver/src/mmo_game_server/core"
	"fmt"
	"log"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoiMgr := core.NewAOIManger(250, 250, 0, 0, 5, 5)
	fmt.Println(aoiMgr)
}

func TestAOIManagerSurroundGridsByGid(t *testing.T) {
	aoiMgr := core.NewAOIManger(250, 250, 0, 0, 5, 5)
	for gid, _ := range aoiMgr.Grids {
		grids := aoiMgr.GetSurroundGridByGid(gid, 3, 3)
		log.Println("gid = ", gid, "grids len: ", len(grids))
		for _, grid := range grids {
			log.Printf("%#v\n", grid)
		}
	}
}
