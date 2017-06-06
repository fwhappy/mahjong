package wall

import (
	"github.com/fwhappy/util"
)

// Wall 牌墙
type Wall struct {
	tiles    []int // 所有牌
	forward  int   // 前游标
	backward int   // 后游标
}

// NewWall 新建一个牌墙
func NewWall() *Wall {
	wall := &Wall{}
	return wall
}

// GetForward 获取前游标
func (wall *Wall) GetForward() int {
	return wall.forward
}

// GetBackwoad 获取后游标
func (wall *Wall) GetBackwoad() int {
	return wall.backward
}

// SetTiles 设置牌墙的牌
func (wall *Wall) SetTiles(tiles []int) {
	wall.tiles = tiles
}

// GetTile 取某个索引的牌面值
func (wall *Wall) GetTile(index int) int {
	return wall.tiles[index]
}

// GetTiles 获取牌墙所有的牌
func (wall *Wall) GetTiles() []int {
	return wall.tiles
}

// Shuffle 洗牌
func (wall *Wall) Shuffle() {
	wall.tiles = util.ShuffleSliceInt(wall.tiles)
}

// Length 获取牌墙长度
func (wall *Wall) Length() int {
	return len(wall.tiles)
}

// RemainLength 牌墙剩余张数
func (wall *Wall) RemainLength() int {
	return wall.Length() - wall.forward - wall.backward
}
