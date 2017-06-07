package step

import (
	"sort"

	"github.com/fwhappy/mahjong/win"
	"github.com/fwhappy/util"
)

// GetCardsStep 计算当前手牌所处牌阶
// fixme 暂未考虑7对情形
func GetCardsStep(showTiles []int) int {
	var cardsStep int
	var sortedCards = util.SliceCopy(showTiles)
	// 升序排列
	sort.Ints(sortedCards)

	// 计算不抽对的时候的牌阶
	cardsStep, _ = win.FindSequenceOrTripletCnt(showTiles)

	// 找到所有的对
	var pos = win.FindPairPos(sortedCards)

	// 遍历所有的对牌，找出离胡最近的抽法，算出牌处于第几阶
	var lastPairTile int // 上次做为对的牌
	for _, v := range pos {
		// 避免有4张同样手牌时，多判断一次
		if sortedCards[v] == lastPairTile {
			continue
		} else {
			lastPairTile = sortedCards[v]
		}
		cards := win.RemovePair(sortedCards, v)
		cnt, _ := win.FindSequenceOrTripletCnt(cards)
		if cnt >= cardsStep {
			cardsStep = cnt + 1
		}
	}
	return cardsStep
}
