package win

import (
	"sort"

	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/util"
)

// CanWin 判断当前牌型是否是胡牌牌型(7对或4A+2B)
// 需要根据手牌和明牌去判断是否胡牌
func CanWin(handCards, showCards []int) bool {
	var sortedCards = util.SliceCopy(handCards)
	// 升序排列
	sort.Ints(sortedCards)
	// 找到所有的对
	var pos = FindPairPos(sortedCards)
	// 找不到对，无法胡牌
	if len(pos) == 0 {
		return false
	}

	// 7对(目前版本只有手中为7个对才可以胡)
	if len(pos) == 7 {
		return true
	}

	// 地龙
	// 手牌有5对;明牌3张;明牌三张相同;且手牌的孤张与明牌相同
	if len(pos) == 5 &&
		len(showCards) == 3 &&
		showCards[0] == showCards[1] && showCards[0] == showCards[2] &&
		util.IntInSlice(showCards[0], handCards) {
		return true
	}

	// 遍历所有对，因为胡必须有对
	var lastPairTile int // 上次做为对的牌
	for _, v := range pos {
		// 避免有4张同样手牌时，多判断一次
		if sortedCards[v] == lastPairTile {
			continue
		} else {
			lastPairTile = sortedCards[v]
		}
		cards := removePair(sortedCards, v)
		if isAllSequenceOrTriplet(cards) {
			return true
		}
	}
	return false
}

// FindPairPos 找出所有对牌的位置
// 传入的牌需要是已排序的
func FindPairPos(sortedCards []int) []int {
	var pos = []int{}
	length := len(sortedCards) - 1
	for i := 0; i < length; i++ {
		if sortedCards[i] == sortedCards[i+1] {
			pos = append(pos, i)
			i++
		}
	}
	return pos
}

// removePair 从已排序的牌中，移除一对
func removePair(sortedCards []int, pos int) []int {
	remainCards := make([]int, 0, len(sortedCards)-2)
	remainCards = append(remainCards, sortedCards[:pos]...)
	remainCards = append(remainCards, sortedCards[pos+2:]...)
	return remainCards
}

// isAllSequenceOrTriplet 是否全部顺或者刻
// 传入的牌需要是已排序的
func isAllSequenceOrTriplet(sortedCards []int) bool {
	cardsLen := len(sortedCards)
	for i := 0; i < cardsLen/3; i++ {
		find := findAndRemoveTriplet(&sortedCards)
		if !find {
			find = findAndRemoveSequence(&sortedCards)
		}
		if !find {
			return false
		}
	}
	return len(sortedCards) == 0
}

// findAndRemoveTriplet 从已排序的[]int中移除排头的刻子
func findAndRemoveTriplet(sortedCards *[]int) bool {
	var v = *sortedCards
	if isTriplet(v[0], v[1], v[2]) {
		*sortedCards = append([]int{}, v[3:]...)
		return true
	}
	return false
}

// findAndRemoveSequence 从已排序的[]int中移除排头的顺子
func findAndRemoveSequence(sortedCards *[]int) bool {
	var v = *sortedCards
	var tmp = make([]int, 0)
	for i := 1; i < len(v); i++ {
		switch {
		case v[i] == v[i-1]:
			tmp = append(tmp, v[i])
		case v[i] == v[i-1]+1:
			if v[i]-v[0] == 2 {
				tmp = append(tmp, v[i+1:]...)
				*sortedCards = tmp
				return true
			}
		default:
			return false
		}
	}
	return false
}

// isSequence 是否顺子
// 传入的牌必须是已排序的
// 非万、筒、条肯定不是顺
func isSequence(tileA, tileB, tileC int) bool {
	if !card.IsSuit(tileA) || !card.IsSuit(tileB) || !card.IsSuit(tileC) {
		return false
	}
	if tileB == tileA+1 && tileC == tileB+1 {
		return true
	}
	return false
}

// isTriplet 是否刻子
func isTriplet(tileA, tileB, tileC int) bool {
	if tileB == tileA && tileC == tileB {
		return true
	}
	return false
}
