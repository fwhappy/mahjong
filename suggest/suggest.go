package suggest

import (
	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/mahjong/ting"
	"github.com/fwhappy/mahjong/weight"
	"github.com/fwhappy/util"
)

// GetSuggest 根据当前牌型，推荐一张牌
func (ms *MSelector) GetSuggest() int {
	handTiles := ms.ShowHandTiles()
	// 明牌
	showTiles := ms.ShowShowTiles()

	// 如果有定缺的牌，则推荐缺的牌
	// 因为大的牌在右手边，所以优先推荐，便于用户操作
	if ms.lack > 0 {
		for i := len(handTiles) - 1; i >= 0; i-- {
			if card.IsSameType(handTiles[i], ms.lack) {
				return handTiles[i]
			}
		}
	}

	// 如果当前牌型已经叫牌，推荐剩余张数最多的牌
	if tingMap := ting.GetTingMap(handTiles, showTiles); len(tingMap) > 0 {
		selectTile := 0
		selectTileTingCnt := -1
		for tile, tingTiles := range tingMap {
			tingCnt := ms.getRemainTilesCnt(tingTiles)
			if tingCnt > selectTileTingCnt {
				selectTileTingCnt = tingCnt
				selectTile = tile
			}
		}
		return selectTile
	}

	// 如果当前牌未叫牌，则根据选牌算法计算
	switch ms.aiLevel {
	case AI_BRASS: // 英勇黄铜
		fallthrough
		// return ms.suggestByAIBrass(s)
	case AI_SLIVER: // 不屈白银
		fallthrough
		// return ms.suggestByAISliver(s)
	case AI_GOLD: // 荣耀黄金
		fallthrough
		// return ms.suggestByAIGold(s)
	case AI_PLATINUM: // 华贵铂金
		fallthrough
		// return ms.suggestByAIPlatinum(s)
	case AI_DIAMOND: // 璀璨钻石
		fallthrough
		// return ms.suggestByAIDiamond(s)
	case AI_MASTER: // 非凡大师
		fallthrough
		// return ms.suggestByAIMaster(s)
	case AI_KING: // 最强王者
		fallthrough
		// return ms.suggestByAIKing(s)
	default:
		return ms.suggestByAIPlatinum(handTiles)
	}
}

// 华贵铂金
func (ms *MSelector) suggestByAIPlatinum(s []int) int {
	// fixme 暂时没空实现，还是根据权重选择
	tilesWeight := weight.GetCardsWeight(s, nil)
	return ms.suggestByWeightAndRemain(tilesWeight)

	/*
		// 计算当前手牌所处的阶段
		tilesStep := step.GetCardsStep(s)

		if tilesStep < 3 {
			tilesWeight := weight.GetCardsWeight(s, nil)
			return ms.suggestByWeightAndRemain(tilesWeight)
		}


			// 最多一类有效牌张数
			maxEffectTileCnt := 0
			// 最多一类有效牌列表
			maxEffectTiles := []int{}
			maxEffectList := map[int][]int{}
			maxEffectTotalWeights := map[int]int{}

			// 循环删除某一张手牌，计算一类有效牌的数量
			for _, playTile := range util.SliceUniqueInt(s) {
				tiles := util.SliceDel(s, playTile)
				// 计算删除后的牌阶，如果小于当前牌阶，跳过计算
				currentStep := step.GetCardsStep(tiles)
				if currentStep < tilesStep {
					continue
				}
				effects, weight := calcEffectsAndRemainWeight(tiles, currentStep)
				effectsLen := len(effects)
				if effectsLen > maxEffectTileCnt {
					maxEffectTileCnt = effectsLen
					maxEffectTiles = []int{playTile}
					maxEffectList = map[int][]int{}
					maxEffectList[playTile] = effects
					maxEffectTotalWeights = map[int]int{}
					maxEffectTotalWeights[playTile] = GetCardsWeightSum(effects)
				}

				step, effects, totalWeight := calcEffectsAndRemainWeight(tiles)
				if step >= unPlayStep {
					sort.Ints(effects)
					showDebug("打出:%v,手牌:%v, 一类有效牌:%v(%v)(remain:%v)------------", playTile, tiles, effects, len(effects), ms.getRemainTilesCnt(effects))

					effectsLen := len(effects)
					if effectsLen > maxEffectTileCnt {
						maxEffectTileCnt = effectsLen
						maxEffectTiles = []int{playTile}
						maxEffectList = map[int][]int{}
						maxEffectList[playTile] = effects
						maxEffectTotalWeights = map[int]int{}
						maxEffectTotalWeights[playTile] = totalWeight
					} else if len(effects) == maxEffectTileCnt {
						maxEffectTiles = append(maxEffectTiles, playTile)
						maxEffectList[playTile] = effects
						maxEffectTotalWeights[playTile] = totalWeight
					}
				}
			}

			// 如果存在相同的一类有效牌，则根据权重再取一次
			showDebug("maxEffectTiles:%v", maxEffectTiles)
			if len(maxEffectTiles) > 1 {
				showDebug("存在多张有效牌相同的打法，根据剩余牌权重重新筛选一次:%v", maxEffectTotalWeights)

				// 读取权重最大的牌
				maxWeightTiles, _ := getMaxValueSlice(maxEffectTotalWeights)
				showDebug("权重最大的牌:%v", maxWeightTiles)

				// 找出权重最小的牌中，关联牌最少的一张
				maxRemainCnt := -1
				var maxRemainTile int
				for _, tile := range maxWeightTiles {
					remainCnt := ms.getRemainTilesCnt([]int{tile})
					if maxRemainCnt < remainCnt {
						maxRemainCnt = remainCnt
						maxRemainTile = tile
					}
				}
				return maxRemainTile
			}

			return maxEffectTiles[0]
	*/
}

// 根据权重筛选
func (ms *MSelector) suggestByWeightAndRemain(tilesWeight map[int]int) int {
	// 读取权重最小的牌
	_, minWeightTiles := util.GetMapMinValue(tilesWeight)

	// 找出权重最小的牌中，关联牌最少的一张
	minRelationCnt := 1000000
	minRelationTile := 0
	for _, tile := range minWeightTiles {
		relationCnt := ms.getRemainTilesCnt(ms.getRelationTiles(tile))
		if relationCnt < minRelationCnt {
			minRelationCnt = relationCnt
			minRelationTile = tile
		}
	}
	return minRelationTile
}
