package ting

import (
	"github.com/fwhappy/mahjong/tile"
	"github.com/fwhappy/mahjong/win"

	"github.com/fwhappy/util"
)

// CanTing 判断牌型是否可以听牌
// 返回是否可听、听什么
func CanTing(handTiles, showTiles []int) (bool, []int) {
	var canTing = false
	tingTiles := make([]int, 0)

	// 循环将可能听的牌，带入到手牌，再用胡牌算法检测是否可胡
	for _, t := range GetMaybeTing(handTiles, showTiles) {
		if win.CanWin(append([]int{t}, handTiles...), showTiles) {
			canTing = true
			tingTiles = append(tingTiles, t)
		}
	}
	return canTing, tingTiles
}

// GetMaybeTing 获取哪些牌是可能听的
// 东南西北、花色等，只有自身
// 边张只有自身或者上下张的某一张
// 其他的是自身和上下张
func GetMaybeTing(handTiles, showTiles []int) []int {
	return tile.GetSelfAndNeighborTiles(append([]int{showTiles[0]}, handTiles...)...)
}

// GetTingMap 获取可听的列表
// key: 打什么
// value: 听哪些
func GetTingMap(handTiles, showTiles []int) map[int][]int {
	tingMap := make(map[int][]int)
	for _, playTile := range util.SliceUniqueInt(handTiles) {
		if ting, tingTiles := CanTing(util.SliceDel(handTiles, playTile), showTiles); ting {
			tingMap[playTile] = tingTiles
		}
	}
	return tingMap
}
