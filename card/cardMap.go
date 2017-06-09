package card

import "sort"

// CardMap 牌面=>数量
type CardMap map[int]int

// NewCardMap 初始化一个TileMap
func NewCardMap() CardMap {
	return make(CardMap)
}

// ToSlice 将TileMap的格式转成[]int
func (cm CardMap) ToSlice() []int {
	tiles := []int{}
	for tile, cnt := range cm {
		for i := 0; i < cnt; i++ {
			tiles = append(tiles, tile)
		}
	}
	sort.Ints(tiles)
	return tiles
}
