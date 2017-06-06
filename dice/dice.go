package dice

import (
	"github.com/fwhappy/util"
)

// DiceRange 骰子的范围
var diceRange = [6]int{1, 2, 3, 4, 5, 6}

// GetDiceRange 获取骰子范围
func GetDiceRange() [6]int {
	return diceRange
}

// GenerateDiceList 生成筛子
func GenerateDiceList() [2]int {
	var dices [2]int
	for i := 0; i < 2; i++ {
		randIndex := util.RandIntn(len(diceRange))
		dices[i] = diceRange[randIndex]
	}
	return dices
}
