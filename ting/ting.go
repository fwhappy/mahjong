package ting

import (
	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/mahjong/win"
	"github.com/fwhappy/util"
)

// CanTing 判断牌型是否可以听牌
// 返回是否可听、听什么
func CanTing(handCards, showCards []int) (bool, []int) {
	var canTing = false
	tingCards := make([]int, 0)

	// 循环将可能听的牌，带入到手牌，再用胡牌算法检测是否可胡
	for _, t := range GetMaybeTing(handCards, showCards) {
		if win.CanWin(append([]int{t}, handCards...), showCards) {
			canTing = true
			tingCards = append(tingCards, t)
		}
	}
	return canTing, tingCards
}

// GetMaybeTing 获取哪些牌是可能听的
// 东南西北、花色等，只有自身
// 边张只有自身或者上下张的某一张
// 其他的是自身和上下张
// 如果有明牌，且明牌是3张的话，则明牌也可能是胡的
func GetMaybeTing(handCards, showCards []int) []int {
	maybeCards := card.GetSelfAndNeighborCards(handCards...)
	if len(showCards) == 3 &&
		showCards[0] == showCards[1] && showCards[1] == showCards[2] &&
		!util.IntInSlice(showCards[0], maybeCards) {
		maybeCards = append(maybeCards, showCards[0])
	}
	return maybeCards
}

// GetTingMap 获取可听的列表
// key: 打什么
// value: 听哪些
func GetTingMap(handCards, showCards []int) map[int][]int {
	tingMap := make(map[int][]int)
	for _, playCard := range util.SliceUniqueInt(handCards) {
		if ting, tingCards := CanTing(util.SliceDel(handCards, playCard), showCards); ting {
			tingMap[playCard] = tingCards
		}
	}
	return tingMap
}
