package chiken

import (
	"github.com/fwhappy/util"
)

// GetWeekChikens 获取星期鸡
func (mc *MChiken) GetWeekChikens() []int {
	week := util.GetChinaWeekDay()
	return []int{week, week + 10, week + 20}
}
