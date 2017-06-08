package chiken

import (
	"github.com/fwhappy/mahjong/wall"
)

// GenChikenFB 创建前后鸡
func (mc *MChiken) GenChikenFB(tileWall *wall.Wall) {
	// 计算前后鸡索引并设置前后鸡
	if mc.fbIndex == 0 {
		// 初始化前后鸡，默认是倒数第三墩上面那个，eg：总牌数为108，则初始前后鸡的索引为102
		mc.fbIndex = tileWall.Length() - 6
	} else {
		// 成功杠牌之后，前后鸡往前走一墩
		mc.fbIndex -= 2
	}
}

// GetChikenFB 获取当前前后鸡的位置、牌面值
func (mc *MChiken) GetChikenFB(tileWall *wall.Wall) (int, int) {
	if mc.fbIndex > 0 && !tileWall.IsDrawn(mc.fbIndex) {
		return mc.fbIndex, tileWall.GetTile(mc.fbIndex)
	}
	return 0, 0
}
