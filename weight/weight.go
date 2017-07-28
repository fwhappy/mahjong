package weight

import (
	"sort"

	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/util"
)

// GetCardsWeight 获取牌的权重列表
// 如果指定了specified，表示只计算specified列出的这些，其他的不计算
func GetCardsWeight(tiles []int, specified []int) map[int]int {
	return GetCardsMapWeight(util.SliceToMap(tiles), specified)
}

// GetCardsMapWeight 获取牌的权重列表
// 如果指定了specified，表示只计算specified列出的这些，其他的不计算
func GetCardsMapWeight(m map[int]int, specified []int) map[int]int {
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

// GetMinWeigthTiles 获取权重最低的length张牌
func GetMinWeigthTiles(originTiles []int, specified []int, length int) []int {
	weights := GetCardsWeight(originTiles, specified)
	tiles := []int{}
	// 若总张数少于要求的张数，直接返回所有的
	if length >= len(weights) {
		for tile := range weights {
			tiles = append(tiles, tile)
		}
		return tiles
	}

	weightValues := []int{}
	weightTiles := make(map[int][]int)
	for k, v := range weights {
		weightValues = append(weightValues, v)
		vs, exists := weightTiles[v]
		if exists {
			vs = append(vs, k)
		} else {
			vs = []int{k}
		}
		weightTiles[v] = vs
	}

	weightValues = util.SliceUniqueInt(weightValues)
	sort.Ints(weightValues)
	for _, v := range weightValues {
		tiles = append(tiles, weightTiles[v]...)
		// 取出了足够的牌
		if len(tiles) >= length {
			return tiles[:length]
		}
	}
	return tiles
}
