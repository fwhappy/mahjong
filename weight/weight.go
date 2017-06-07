package weight

import (
	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/util"
)

// GetCardsWeight 获取牌的权重列表
// 如果指定了specified，表示只计算specified列出的这些，其他的不计算
func GetCardsWeight(tiles []int, specified []int) map[int]int {
	m := util.SliceToMap(tiles)
	tilesWeight := map[int]int{}

	// 统计出出牌的优先级
	for tile, cnt := range m {
		// 指定只统计某些牌的权重
		if specified != nil && !util.IntInSlice(tile, specified) {
			continue
		}

		score := 0
		// 计算牌自身的分
		if !card.IsSuit(tile) {
			// 非普通牌不加分
		} else if util.IntInSlice(tile, card.SideCards) {
			// 边张的分
			score += 10
		} else if util.IntInSlice(tile, card.SideNeighborCards) {
			// 邻边张的分
			score += 20
		} else {
			// 中间张的分
			score += 30
		}
		// 邻张的分
		if m[tile-1] > 0 {
			score += 50
		}
		if m[tile+1] > 0 {
			score += 50
		}
		// 隔张的分
		if !util.IntInSlice(tile, card.LeftSideCards) && m[tile+2] > 0 {
			score += 40
		}
		if !util.IntInSlice(tile, card.RightSideCards) && m[tile-2] > 0 {
			score += 40
		}
		// 同张的分
		if cnt >= 2 {
			score += 80 * (cnt - 1)
		}
		tilesWeight[tile] = score
	}
	return tilesWeight
}

// GetCardsWeightSum 获取牌的权重和
func GetCardsWeightSum(tiles []int, specified []int) int {
	tilesWeight := GetCardsWeight(tiles, specified)
	weight := 0
	for _, score := range tilesWeight {
		weight += score
	}
	return weight
}
