package win

import (
	"fmt"
	"reflect"
	"testing"
)

// TestFindPairPos 测试找出对的位置是否正确
func TestFindPairPos(t *testing.T) {
	cards := []int{1, 1, 1, 1, 5, 8, 8, 6, 7, 11, 11}
	findPos := FindPairPos(cards)
	if !reflect.DeepEqual(findPos, []int{0, 2, 5, 9}) {
		t.Errorf("TestFindPairPos Error, findPos:%v", findPos)
	}
}

// 测试删除pair的判断
func TestRemovePair(t *testing.T) {
	cards := []int{1, 1, 1, 1, 5, 6, 7, 8, 8, 11, 11}
	// findPos := FindPairPos(cards)
	if !reflect.DeepEqual(RemovePair(cards, 9), []int{1, 1, 1, 1, 5, 6, 7, 8, 8}) {
		t.Error("removePair 验证失败")
	}
	if !reflect.DeepEqual(RemovePair(cards, 7), []int{1, 1, 1, 1, 5, 6, 7, 11, 11}) {
		t.Error("removePair 验证失败")
	}
	if !reflect.DeepEqual(RemovePair(cards, 2), []int{1, 1, 5, 6, 7, 8, 8, 11, 11}) {
		t.Error("removePair 验证失败")
	}
}

// 测试是否全部是顺或者刻
func TestIsAllShunOrKe(t *testing.T) {
	cards := make([]int, 0)

	cards = []int{1, 1, 1, 2, 3, 3, 4, 4, 5}
	if !IsAllSequenceOrTriplet(cards) {
		t.Errorf("TestIsAllShunOrKe failed1:%v", cards)
	}
	cards = []int{1, 1, 2, 2, 3, 3}
	if !IsAllSequenceOrTriplet(cards) {
		t.Errorf("TestIsAllShunOrKe failed2:%v", cards)
	}
	cards = []int{1, 1, 1, 2, 2, 3, 3, 4, 4}
	if !IsAllSequenceOrTriplet(cards) {
		t.Errorf("TestIsAllShunOrKe failed3:%v", cards)
	}
}

func TestCanWin(t *testing.T) {
	var handCards []int
	var showCards []int
	// 是否可以胡7对
	handCards = []int{1, 1, 5, 5, 9, 9, 11, 11, 12, 12, 13, 13, 19, 19}
	if !CanWin(handCards, nil) {
		t.Error("7对类型验证失败")
	}
	handCards = []int{1, 1, 5, 5, 9, 9, 11, 11, 12, 12, 13, 13, 19, 25}
	if CanWin(handCards, nil) {
		t.Error("7对类型验证失败")
	}

	// 是否可以胡地龙
	handCards = []int{1, 1, 5, 5, 9, 9, 11, 11, 12, 12, 13}
	showCards = []int{13, 13, 13}
	if !CanWin(handCards, showCards) {
		t.Error("地龙牌型判断错误")
	}
	showCards = []int{13, 14, 15}
	if CanWin(handCards, showCards) {
		t.Error("地龙牌型判断错误")
	}
	handCards = []int{1, 1, 5, 5, 9, 9, 11, 11, 12, 12, 18}
	showCards = []int{13, 13, 13}
	if CanWin(handCards, showCards) {
		t.Error("地龙牌型判断错误")
	}

	// 是否可以胡单吊
	handCards = []int{1, 1}
	if !CanWin(handCards, nil) {
		t.Error("单吊牌型验证失败")
	}
	handCards = []int{1, 2}
	if CanWin(handCards, nil) {
		t.Error("单吊牌型验证失败")
	}

	// 是否可以胡5张手牌
	handCards = []int{1, 1, 1, 2, 3}
	if !CanWin(handCards, nil) {
		t.Error("5张牌型验证失败")
	}

	// 是否可以胡5张手牌
	handCards = []int{1, 1, 1, 2, 4}
	if CanWin(handCards, nil) {
		t.Error("5张牌型验证失败")
	}

	// 是否可以胡8张手牌
	handCards = []int{1, 1, 1, 1, 2, 3, 9, 9}
	if !CanWin(handCards, nil) {
		t.Error("5张牌型验证失败")
	}
	handCards = []int{1, 1, 1, 1, 2, 3, 9, 12}
	if CanWin(handCards, nil) {
		t.Error("5张牌型验证失败")
	}
	// 是否可以胡11张手牌
	handCards = []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	if !CanWin(handCards, nil) {
		t.Error("11张牌型验证失败1")
	}
}

func TestFindFindSequenceOrTripletCnt(t *testing.T) {
	var showTiles = []int{}
	var remainTiles = []int{}
	var step int

	showTiles = []int{1, 2, 3, 4}
	step, remainTiles = FindSequenceOrTripletCnt(showTiles)
	fmt.Println(step, ",", remainTiles)
	if step != 1 {
		t.Error("验证牌阶错误1")
	}

	showTiles = []int{1, 2, 3, 4, 3}
	step, remainTiles = FindSequenceOrTripletCnt(showTiles)
	fmt.Println(step, ",", remainTiles)
	if step != 1 {
		t.Error("验证牌阶错误1")
	}

	showTiles = []int{1, 2, 3, 4, 4}
	step, remainTiles = FindSequenceOrTripletCnt(showTiles)
	fmt.Println(step, ",", remainTiles)
	if step != 1 {
		t.Error("验证牌阶错误1")
	}

	showTiles = []int{1, 2, 3, 4, 4, 4}
	step, remainTiles = FindSequenceOrTripletCnt(showTiles)
	fmt.Println(step, ",", remainTiles)
	if step != 2 {
		t.Error("验证牌阶错误1")
	}
	showTiles = []int{1, 2, 3, 4, 4, 5, 5, 6, 7}
	step, remainTiles = FindSequenceOrTripletCnt(showTiles)
	fmt.Println(step, ",", remainTiles)
	if step != 2 {
		t.Error("验证牌阶错误1")
	}
}
