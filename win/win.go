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
		cards := RemovePair(sortedCards, v)
		if IsAllSequenceOrTriplet(cards) {
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

// RemovePair 从已排序的牌中，移除一对
func RemovePair(sortedCards []int, pos int) []int {
	remainCards := make([]int, 0, len(sortedCards)-2)
	remainCards = append(remainCards, sortedCards[:pos]...)
	remainCards = append(remainCards, sortedCards[pos+2:]...)
	return remainCards
}

// IsAllSequenceOrTriplet 是否全部顺或者刻
// 传入的牌需要是已排序的
func IsAllSequenceOrTriplet(sortedCards []int) bool {
	cardsLen := len(sortedCards)
	for i := 0; i < cardsLen/3; i++ {
		find := FindAndRemoveTriplet(&sortedCards)
		if !find {
			find = FindAndRemoveSequence(&sortedCards)
		}
		if !find {
			return false
		}
	}
	return len(sortedCards) == 0
}

// FindAndRemoveTriplet 从已排序的[]int中移除排头的刻子
func FindAndRemoveTriplet(sortedCards *[]int) bool {
	var v = *sortedCards
	if IsTriplet(v[0], v[1], v[2]) {
		*sortedCards = append([]int{}, v[3:]...)
		return true
	}
	return false
}

// FindAndRemoveSequence 从已排序的[]int中移除排头的顺子
func FindAndRemoveSequence(sortedCards *[]int) bool {
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

// IsSequence 是否顺子
// 传入的牌必须是已排序的
// 非万、筒、条肯定不是顺
func IsSequence(tileA, tileB, tileC int) bool {
	if !card.IsSuit(tileA) || !card.IsSuit(tileB) || !card.IsSuit(tileC) {
		return false
	}
	if tileB == tileA+1 && tileC == tileB+1 {
		return true
	}
	return false
}

// IsTriplet 是否刻子
func IsTriplet(tileA, tileB, tileC int) bool {
	if tileB == tileA && tileC == tileB {
		return true
	}
	return false
}

// FindSequenceOrTripletCnt 找出当前牌中所有刻和顺的数量
// 返回数量和抽完剩余的牌
func FindSequenceOrTripletCnt(sortedCards []int) (int, []int) {
	var cnt = 0
	var remain = []int{}
	for {
		if len(sortedCards) <= 2 {
			remain = append(remain, sortedCards...)
			break
		}
		find := FindAndRemoveTriplet(&sortedCards)
		if !find {
			find = FindAndRemoveSequence(&sortedCards)
		}
		if find {
			cnt++
		} else {
			remain = append(remain, sortedCards[0])
			sortedCards = sortedCards[1:]
		}
	}
	return cnt, remain
}
