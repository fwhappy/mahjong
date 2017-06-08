package chiken

// SetChargeBam1 设置冲锋幺鸡
func (mc *MChiken) SetChargeBam1(id int) {
	if mc.chargeBam1 == 0 && mc.responsibility == 0 {
		mc.chargeBam1 = id
	}
}

// GetChargeBam1 获取冲锋幺鸡
func (mc *MChiken) GetChargeBam1() int {
	return mc.chargeBam1
}

// SetChargeDot8 设置冲锋乌骨
func (mc *MChiken) SetChargeDot8(id int) {
	if mc.chargeDot8 == 0 {
		mc.chargeDot8 = id
	}
}

// GetChargeDot8 获取冲锋乌骨
func (mc *MChiken) GetChargeDot8() int {
	return mc.chargeDot8
}
