package card

import (
	"testing"
)

func TestCardMap(t *testing.T) {
	initTiles := []int{3, 2, 1, 7, 9, 8, 7, 3, 1}
	cm := NewCMap()
	cm.SetTiles(initTiles)

	t.Logf("cm.tiles:%#v", cm.GetTileMap())
	t.Logf("cm.GetUnique:%v", cm.GetUnique())
	t.Logf("cm.ToSlice:%v", cm.ToSlice())
	t.Logf("cm.ToSortedSlice:%v", cm.ToSortedSlice())
	cm.DelTile(1, 1)
	t.Logf("cm.ToSortedSlice:%v", cm.ToSortedSlice())
	cm.DelTile(1, 1)
	t.Logf("cm.ToSortedSlice:%v", cm.ToSortedSlice())
	cm.DelTile(3, 2)
	t.Logf("cm.ToSortedSlice:%v", cm.ToSortedSlice())
	flag := cm.DelTile(3, 1)
	t.Logf("cm.DelTile:%v", flag)
	cm.AddTile(3, 1)
	t.Logf("cm.ToSortedSlice:%v", cm.ToSortedSlice())
	cm.AddTile(3, 1)
	t.Logf("cm.ToSortedSlice:%v", cm.ToSortedSlice())
	cm.AddTile(1, 2)
	t.Logf("cm.ToSortedSlice:%v", cm.ToSortedSlice())
	t.Logf("cm.GetTileCnt:%v", cm.GetTileCnt(1))
	t.Logf("cm.GetTileCnt:%v", cm.GetTileCnt(11))
}
