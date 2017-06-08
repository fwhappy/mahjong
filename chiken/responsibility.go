package chiken

// SetResponsibility 设置责任鸡
func (mc *MChiken) SetResponsibility(id int) {
	if mc.chargeBam1 == 0 && mc.responsibility == 0 {
		mc.responsibility = id
	}
}

// GetResponsibility 获取责任鸡
func (mc *MChiken) GetResponsibility() int {
	return mc.responsibility
}
