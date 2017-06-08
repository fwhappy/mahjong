package chiken

import (
	"github.com/fwhappy/mahjong/wall"
)

// AppendChikenTumbling 创建滚筒鸡
func (mc *MChiken) AppendChikenTumbling(tileWall *wall.Wall) int {
	var appendIndex int
	if len(mc.tumblingIndexes) == 0 {
		appendIndex = tileWall.Length() - 6
	} else {
		// 最后一张前后鸡的索引
		chikenIndex := mc.tumblingIndexes[len(mc.tumblingIndexes)-1] - 2
		if chikenIndex >= 0 && !tileWall.IsDrawn(chikenIndex) {
			appendIndex = chikenIndex
		}
	}
	if appendIndex > 0 {
		mc.tumblingIndexes = append(mc.tumblingIndexes, appendIndex)
	}
	return appendIndex
}

// GetChikenTumblingMap 获取所有翻到的滚筒鸡，返回index => tile的结构
func (mc *MChiken) GetChikenTumblingMap(tileWall *wall.Wall) map[int]int {
	tumblingMap := make(map[int]int)
	for _, index := range mc.tumblingIndexes {
		if !tileWall.IsDrawn(index) {
			tumblingMap[index] = tileWall.GetTile(index)
		}
	}
	return tumblingMap
}
