package chiken

// GetWeekChikens 获取星期鸡
func (mc *MChiken) GetWeekChikens() []int {
	return []int{mc.weekDay, mc.weekDay + 10, mc.weekDay + 20}
}
