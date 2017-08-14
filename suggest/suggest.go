package suggest

import (
	"github.com/fwhappy/mahjong/card"
	"github.com/fwhappy/mahjong/step"
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
		relationCnt := ms.getRemainTilesCnt(card.GetRelationTiles(tile))
		if relationCnt < minRelationCnt {
			minRelationCnt = relationCnt
			minRelationTile = tile
		}
	}
	return minRelationTile
}

// GetSuggestMap 获取推荐的列表
// 这里默认外面已经考虑过叫牌了，不再考虑叫牌的情况
// 有缺的时候，不给任何推荐，用户必须先打缺
// 如果有孤张，优先返回孤张
// 孤一对，不拆做提示
// 最多返回maxLength个结果，如果maxLength=0，表示返回所有
// maxLength: 最多返回多少个孤张
func (ms *MSelector) GetSuggestMap(maxLength int) map[int][]int {
	suggestMap := make(map[int][]int)
	// 优先判断是否有缺
	if ms.hasLack() {
		return suggestMap
	}
	// 用户手牌
	handTiles := util.MapToSlice(ms.handTiles)
	// 如果有孤张，则优先提示孤张
	// 孤张按权重来选，优先边张
	if gTiles := ms.getGuTiles(); len(gTiles) > 0 {
		for _, v := range weight.GetMinWeigthTiles(handTiles, gTiles, maxLength) {
			suggestMap[v] = []int{}
		}
		return suggestMap
	}
	// 按一类有效牌剩余数推荐
	// 获取所有的孤一对
	gpTiles := ms.getGuPairTiles()
	// 计算手牌当前牌阶
	currentStep := step.GetCardsStep(handTiles)

	// 循环删除一张手牌后，计算一类有效牌的数量
	for playTile := range ms.handTiles {
		// 孤对不推荐
		if util.IntInSlice(playTile, gpTiles) {
			continue
		}
		tiles := util.SliceDel(handTiles, playTile)
		// 如果打出后，牌阶比之前的还要低，肯定不能这么打
		playedStep := step.GetCardsStep(tiles)
		if playedStep < currentStep {
			continue
		}
		// 计算一类有效牌数量
		effects := calcEffects(tiles, currentStep)
		if len(effects) > 0 {
			suggestMap[playTile] = effects
		}
	}

	// 如果未找到一类有效牌，则继续查找二类有效牌
	if len(suggestMap) == 0 {
		// 循环删除一张手牌后，计算一类有效牌的数量
		for playTile := range ms.handTiles {
			sEffects := []int{}
			// 孤对不推荐
			if util.IntInSlice(playTile, gpTiles) {
				continue
			}
			tiles := util.SliceDel(handTiles, playTile)
			// 如果打出后，牌阶比之前的还要低，肯定不能这么打
			playedStep := step.GetCardsStep(tiles)
			if playedStep < currentStep {
				continue
			}

			// 获取跟playTile有关系的牌
			for _, rTile := range card.GetRelationTiles(playTile) {
				effects := calcEffects(append([]int{rTile}, tiles...), currentStep)
				if len(effects) > 0 {
					sEffects = append(sEffects, rTile)
				}
			}
			if len(sEffects) > 0 {
				suggestMap[playTile] = sEffects
			}
		}
	}

	// 如果找到了推荐的牌，则找出有效进张数最多的那些
	remainTiles := []int{}
	maxCnt := 0
	for tile, effects := range suggestMap {
		remainCnt := ms.getRemainTilesCnt(effects)
		if remainCnt > maxCnt {
			remainTiles = []int{tile}
			maxCnt = remainCnt
		} else if remainCnt == maxCnt {
			remainTiles = append(remainTiles, tile)
		}
	}
	if len(remainTiles) > maxLength {
		remainTiles = weight.GetMinWeigthTiles(handTiles, remainTiles, maxLength)
	}
	for k := range suggestMap {
		if !util.IntInSlice(k, remainTiles) {
			delete(suggestMap, k)
		}
	}

	return suggestMap
}
