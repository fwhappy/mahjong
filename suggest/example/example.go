package main

import (
	"fmt"

	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/mahjong/suggest"
)

var tiles36 = []int{
	// 万
	card.MAHJONG_CRAK1, card.MAHJONG_CRAK1, card.MAHJONG_CRAK1, card.MAHJONG_CRAK1,
	card.MAHJONG_CRAK2, card.MAHJONG_CRAK2, card.MAHJONG_CRAK2, card.MAHJONG_CRAK2,
	card.MAHJONG_CRAK3, card.MAHJONG_CRAK3, card.MAHJONG_CRAK3, card.MAHJONG_CRAK3,
	card.MAHJONG_CRAK4, card.MAHJONG_CRAK4, card.MAHJONG_CRAK4, card.MAHJONG_CRAK4,
	card.MAHJONG_CRAK5, card.MAHJONG_CRAK5, card.MAHJONG_CRAK5, card.MAHJONG_CRAK5,
	card.MAHJONG_CRAK6, card.MAHJONG_CRAK6, card.MAHJONG_CRAK6, card.MAHJONG_CRAK6,
	card.MAHJONG_CRAK7, card.MAHJONG_CRAK7, card.MAHJONG_CRAK7, card.MAHJONG_CRAK7,
	card.MAHJONG_CRAK8, card.MAHJONG_CRAK8, card.MAHJONG_CRAK8, card.MAHJONG_CRAK8,
}

var tiles72 = []int{
	// 筒
	card.MAHJONG_DOT1, card.MAHJONG_DOT1, card.MAHJONG_DOT1, card.MAHJONG_DOT1,
	card.MAHJONG_DOT2, card.MAHJONG_DOT2, card.MAHJONG_DOT2, card.MAHJONG_DOT2,
	card.MAHJONG_DOT3, card.MAHJONG_DOT3, card.MAHJONG_DOT3, card.MAHJONG_DOT3,
	card.MAHJONG_DOT4, card.MAHJONG_DOT4, card.MAHJONG_DOT4, card.MAHJONG_DOT4,
	card.MAHJONG_DOT5, card.MAHJONG_DOT5, card.MAHJONG_DOT5, card.MAHJONG_DOT5,
	card.MAHJONG_DOT6, card.MAHJONG_DOT6, card.MAHJONG_DOT6, card.MAHJONG_DOT6,
	card.MAHJONG_DOT7, card.MAHJONG_DOT7, card.MAHJONG_DOT7, card.MAHJONG_DOT7,
	card.MAHJONG_DOT8, card.MAHJONG_DOT8, card.MAHJONG_DOT8, card.MAHJONG_DOT8,
	card.MAHJONG_DOT9, card.MAHJONG_DOT9, card.MAHJONG_DOT9, card.MAHJONG_DOT9,

	// 条
	card.MAHJONG_BAM1, card.MAHJONG_BAM1, card.MAHJONG_BAM1, card.MAHJONG_BAM1,
	card.MAHJONG_BAM2, card.MAHJONG_BAM2, card.MAHJONG_BAM2, card.MAHJONG_BAM2,
	card.MAHJONG_BAM3, card.MAHJONG_BAM3, card.MAHJONG_BAM3, card.MAHJONG_BAM3,
	card.MAHJONG_BAM4, card.MAHJONG_BAM4, card.MAHJONG_BAM4, card.MAHJONG_BAM4,
	card.MAHJONG_BAM5, card.MAHJONG_BAM5, card.MAHJONG_BAM5, card.MAHJONG_BAM5,
	card.MAHJONG_BAM6, card.MAHJONG_BAM6, card.MAHJONG_BAM6, card.MAHJONG_BAM6,
	card.MAHJONG_BAM7, card.MAHJONG_BAM7, card.MAHJONG_BAM7, card.MAHJONG_BAM7,
	card.MAHJONG_BAM8, card.MAHJONG_BAM8, card.MAHJONG_BAM8, card.MAHJONG_BAM8,
	card.MAHJONG_BAM9, card.MAHJONG_BAM9, card.MAHJONG_BAM9, card.MAHJONG_BAM9,
}

var tiles108 = []int{
	// 万
	card.MAHJONG_CRAK1, card.MAHJONG_CRAK1, card.MAHJONG_CRAK1, card.MAHJONG_CRAK1,
	card.MAHJONG_CRAK2, card.MAHJONG_CRAK2, card.MAHJONG_CRAK2, card.MAHJONG_CRAK2,
	card.MAHJONG_CRAK3, card.MAHJONG_CRAK3, card.MAHJONG_CRAK3, card.MAHJONG_CRAK3,
	card.MAHJONG_CRAK4, card.MAHJONG_CRAK4, card.MAHJONG_CRAK4, card.MAHJONG_CRAK4,
	card.MAHJONG_CRAK5, card.MAHJONG_CRAK5, card.MAHJONG_CRAK5, card.MAHJONG_CRAK5,
	card.MAHJONG_CRAK6, card.MAHJONG_CRAK6, card.MAHJONG_CRAK6, card.MAHJONG_CRAK6,
	card.MAHJONG_CRAK7, card.MAHJONG_CRAK7, card.MAHJONG_CRAK7, card.MAHJONG_CRAK7,
	card.MAHJONG_CRAK8, card.MAHJONG_CRAK8, card.MAHJONG_CRAK8, card.MAHJONG_CRAK8,
	card.MAHJONG_CRAK9, card.MAHJONG_CRAK9, card.MAHJONG_CRAK9, card.MAHJONG_CRAK9,

	// 筒
	card.MAHJONG_DOT1, card.MAHJONG_DOT1, card.MAHJONG_DOT1, card.MAHJONG_DOT1,
	card.MAHJONG_DOT2, card.MAHJONG_DOT2, card.MAHJONG_DOT2, card.MAHJONG_DOT2,
	card.MAHJONG_DOT3, card.MAHJONG_DOT3, card.MAHJONG_DOT3, card.MAHJONG_DOT3,
	card.MAHJONG_DOT4, card.MAHJONG_DOT4, card.MAHJONG_DOT4, card.MAHJONG_DOT4,
	card.MAHJONG_DOT5, card.MAHJONG_DOT5, card.MAHJONG_DOT5, card.MAHJONG_DOT5,
	card.MAHJONG_DOT6, card.MAHJONG_DOT6, card.MAHJONG_DOT6, card.MAHJONG_DOT6,
	card.MAHJONG_DOT7, card.MAHJONG_DOT7, card.MAHJONG_DOT7, card.MAHJONG_DOT7,
	card.MAHJONG_DOT8, card.MAHJONG_DOT8, card.MAHJONG_DOT8, card.MAHJONG_DOT8,
	card.MAHJONG_DOT9, card.MAHJONG_DOT9, card.MAHJONG_DOT9, card.MAHJONG_DOT9,

	// 条
	card.MAHJONG_BAM1, card.MAHJONG_BAM1, card.MAHJONG_BAM1, card.MAHJONG_BAM1,
	card.MAHJONG_BAM2, card.MAHJONG_BAM2, card.MAHJONG_BAM2, card.MAHJONG_BAM2,
	card.MAHJONG_BAM3, card.MAHJONG_BAM3, card.MAHJONG_BAM3, card.MAHJONG_BAM3,
	card.MAHJONG_BAM4, card.MAHJONG_BAM4, card.MAHJONG_BAM4, card.MAHJONG_BAM4,
	card.MAHJONG_BAM5, card.MAHJONG_BAM5, card.MAHJONG_BAM5, card.MAHJONG_BAM5,
	card.MAHJONG_BAM6, card.MAHJONG_BAM6, card.MAHJONG_BAM6, card.MAHJONG_BAM6,
	card.MAHJONG_BAM7, card.MAHJONG_BAM7, card.MAHJONG_BAM7, card.MAHJONG_BAM7,
	card.MAHJONG_BAM8, card.MAHJONG_BAM8, card.MAHJONG_BAM8, card.MAHJONG_BAM8,
	card.MAHJONG_BAM9, card.MAHJONG_BAM9, card.MAHJONG_BAM9, card.MAHJONG_BAM9,
}

func main() {
	// 初始化一个选牌器
	ms := suggest.NewMSelector()
	ms.SetTiles(tiles72)
	tiles := ms.GetShuffleTiles()

	// 预设参数
	lack := 0
	// handTiles := []int{6, 7, 9, 9, 12, 12, 13, 14, 15, 15, 17, 26, 27, 28}
	// handTiles := []int{22, 24, 24, 25, 27, 29, 11, 13, 14, 14, 15, 16, 17, 17}
	handTiles := tiles[:14] // 5 8 11 14
	// 明牌
	showTiles := []int{}
	// 弃牌
	discardTiles := []int{}
	// discardTiles := tiles[5:30]
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
	tile := ms.GetSuggest()

	fmt.Println("选牌结果:", tile)
}
