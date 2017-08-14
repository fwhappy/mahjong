package suggest

import (
	"sort"

	"github.com/fwhappy/mahjong/step"
	"github.com/fwhappy/mahjong/ting"
	"github.com/fwhappy/util"
)

// GetEffects 计算牌型所在的牌阶以及一类有效牌数量
func (ms *MSelector) GetEffects(s []int) (int, int) {
	currentStep := step.GetCardsStep(s)
	effects := calcEffects(s, currentStep)
	return currentStep, len(effects)
}

// 一类有效牌算法
// 返回所有一类有效牌
func calcEffects(tiles []int, currentStep int) []int {
	var effects = []int{}
	// 若未传入牌阶，需重新计算得到
	if currentStep == 0 {
		currentStep = step.GetCardsStep(tiles)
	}

	// 可能是一类有效的牌
	maybeFirstEffects := ting.GetMaybeTing(tiles, nil)
	sort.Ints(maybeFirstEffects)
	for _, maybeTile := range maybeFirstEffects {
		fillingStep := step.GetCardsStep(append(util.SliceCopy(tiles), maybeTile))
		if fillingStep > currentStep {
			effects = append(effects, maybeTile)
		}
	}
	return effects
}
