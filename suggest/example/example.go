package main

import (
	"fmt"

	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/mahjong/suggest"
)

func main() {
	// 初始化一个选牌器
	ms := suggest.NewMSelector()
	ms.SetTiles(card.MahjongCards32)
	tiles := ms.GetShuffleTiles()
	_ = tiles

	// 预设参数
	lack := 0
	handTiles := []int{3, 3, 5, 6, 6, 6, 17, 18, 19, 22, 22, 25, 27, 25}
	// handTiles := []int{1, 1, 5, 6, 8, 11, 12, 13, 16, 17, 21, 26, 28, 29}
	// handTiles := tiles[:14] // 5 8 11 14
	// handTiles := []int{1, 4, 7, 2, 15, 8, 3, 26, 9}
	// 明牌
	showTiles := []int{}
	// 弃牌
	discardTiles := []int{}
	// discardTiles = tiles[15:26]
	// remainTiles := []int{9, 8, 7}
	// remainTiles := tiles[5:30]
	aiLevel := suggest.AI_PLATINUM
	remainTiles := []int{}

	// 设置参数
	ms.SetAILevel(aiLevel)
	ms.SetLack(lack)
	ms.SetHandTilesSlice(handTiles)
	ms.SetShowTilesSlice(showTiles)
	ms.SetDiscardTilesSlice(discardTiles)
	if len(remainTiles) > 0 {
		ms.SetRemainTilesSlice(remainTiles)
	} else {
		ms.CalcRemaimTiles()
	}

	fmt.Println("定缺:", ms.GetLack())
	fmt.Println("手牌:", ms.ShowHandTiles())
	fmt.Println("明牌:", ms.ShowShowTiles())
	fmt.Println("弃牌:", ms.ShowDiscardTiles())
	fmt.Println("剩余的牌:", ms.ShowRemainTiles())

	// 选牌
	// tile := ms.GetSuggest()
	// fmt.Println("选牌结果:", tile)

	suggestMap := ms.GetSuggestMap(2)
	fmt.Printf("选牌结果:%#v", suggestMap)

	// weights := map[int]int{1: 10, 2: 10, 3: 5}
	// fmt.Println(suggest.GetMinWeigthTiles(weights, 3))
}
