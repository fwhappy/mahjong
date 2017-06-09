package chiken

import (
	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/mahjong/setting"
	"github.com/fwhappy/util"
)

// MChiken 鸡
type MChiken struct {
	draw            int   // 翻牌鸡
	fbIndex         int   // 前后鸡的位置
	tumblingIndexes []int // 滚筒鸡的位置
	chargeBam1      int   // 冲锋幺鸡
	chargeDot8      int   // 冲锋乌骨
	responsibility  int   // 责任鸡
	weekDay         int   // 周几
}

// NewMChiken 生成一个新的鸡的容器
func NewMChiken() *MChiken {
	mc := &MChiken{tumblingIndexes: make([]int, 0), weekDay: util.GetChinaWeekDay()}
	return mc
}

// GetChikensByGivenTiles 获取牌对应的鸡
func (mc *MChiken) GetChikensByGivenTiles(s *setting.MSetting, tiles ...int) []int {
	chikens := []int{}
	for _, originTile := range tiles {
		if originTile == 0 {
			continue
		}
		// 下一张牌肯定是鸡
		if tile := card.GetBehindCardCycle(originTile); tile > 0 {
			chikens = append(chikens, tile)
		}
		// 如果支持上下鸡，需要加入上一张牌
		if s.IsSettingChikenUD() {
			if tile := card.GetFrontCardCycle(originTile); tile > 0 {
				chikens = append(chikens, tile)
			}
		}
	}
	return util.SliceUniqueInt(chikens)
}
