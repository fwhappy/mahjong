package card

import (
	"sort"
	"sync"
)

// CMap 牌面=>数量
type CMap struct {
	Mux   *sync.RWMutex
	tiles map[int]int
}

// NewCMap 初始化一个TileMap
func NewCMap() *CMap {
	return &CMap{
		Mux:   &sync.RWMutex{},
		tiles: make(map[int]int),
	}
}

// SetTiles 初始化手牌
func (cm *CMap) SetTiles(tiles []int) {
	cm.Mux.Lock()
	defer cm.Mux.Unlock()
	for _, tile := range tiles {
		cm.tiles[tile]++
	}
}

// GetTileMap 读取所有牌的列表
// 这里不会主动加锁，在外面用的话，如果用于range，需要手动加锁
func (cm *CMap) GetTileMap() map[int]int {
	return cm.tiles
}

// AddTile 添加手牌
func (cm *CMap) AddTile(tile, cnt int) {
	cm.Mux.Lock()
	defer cm.Mux.Unlock()
	cm.tiles[tile] += cnt
}

// DelTile 删除手牌
func (cm *CMap) DelTile(tile, cnt int) bool {
	cm.Mux.Lock()
	defer cm.Mux.Unlock()
	if cm.tiles[tile] > cnt {
		cm.tiles[tile] -= cnt
	} else if cm.tiles[tile] == cnt {
		delete(cm.tiles, tile)
	} else {
		return false
	}
	return true
}

// ToSlice 转成slice
func (cm *CMap) ToSlice() []int {
	cm.Mux.RLock()
	defer cm.Mux.RUnlock()
	tiles := []int{}
	for tile, cnt := range cm.tiles {
		for i := 0; i < cnt; i++ {
			tiles = append(tiles, tile)
		}
	}
	sort.Ints(tiles)
	return tiles
}

// ToSortedSlice 转成slice并排序
func (cm *CMap) ToSortedSlice() []int {
	tiles := cm.ToSlice()
	sort.Ints(tiles)
	return tiles
}

// GetUnique 获取独立的牌
func (cm *CMap) GetUnique() []int {
	cm.Mux.RLock()
	defer cm.Mux.RUnlock()
	tiles := []int{}
	for tile := range cm.tiles {
		tiles = append(tiles, tile)
	}
	return tiles
}

// GetTileCnt 获取某张牌的数量
func (cm *CMap) GetTileCnt(tile int) int {
	cm.Mux.RLock()
	defer cm.Mux.RUnlock()
	return cm.tiles[tile]
}
