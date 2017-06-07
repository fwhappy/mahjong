package card

import (
	"github.com/fwhappy/util"
)

// IsSuit 是否普通牌
// 普通牌是指万、筒、条
func IsSuit(card int) bool {
	return card < MAHJONG_DOT_PLACEHOLDER
}

// GetSelfAndNeighborCards 获取自身或者相邻的一张牌, 结果需去重
// 不包括隔张
// 1条、1筒、1万只有自己和上一张
// 九条、九筒、九万只有自己和下一张
// 非万、筒、条 只有自己
func GetSelfAndNeighborCards(cards ...int) []int {
	result := []int{}
	for _, card := range cards {
		result = append(result, card)
		// 非普通牌、只有自身
		if !IsSuit(card) {
			continue
		}
		switch card {
		case MAHJONG_CRAK1, MAHJONG_BAM1, MAHJONG_DOT1:
			result = append(result, card+1)
		case MAHJONG_CRAK9, MAHJONG_BAM9, MAHJONG_DOT9:
			result = append(result, card-1)
		default:
			result = append(result, card-1, card+1)
		}
	}
	return util.SliceUniqueInt(result)
}

// GetRelationTiles 获取有关联的牌
// 包括自己、相邻的、跳张
func GetRelationTiles(cards ...int) []int {
	result := []int{}
	for _, card := range cards {
		result = append(result, card)
		// 非普通牌、只有自身
		if !IsSuit(card) {
			continue
		}
		switch card {
		case MAHJONG_CRAK1, MAHJONG_BAM1, MAHJONG_DOT1:
			result = append(result, card+1, card+2)
		case MAHJONG_CRAK2, MAHJONG_BAM2, MAHJONG_DOT2:
			result = append(result, card+1, card+2, card-1)
		case MAHJONG_CRAK8, MAHJONG_BAM8, MAHJONG_DOT8:
			result = append(result, card-1, card-2, card+1)
		case MAHJONG_CRAK9, MAHJONG_BAM9, MAHJONG_DOT9:
			result = append(result, card-1, card-2)
		default:
			result = append(result, card-1, card-2, card+1, card+2)
		}
	}
	return util.SliceUniqueInt(result)
}

// IsSameType 检查两张牌是否同类
func IsSameType(checkCard, lackCard int) bool {
	return checkCard/10 == lackCard/10
}
