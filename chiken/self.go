package chiken

// GetChikenSelf 获取本鸡
func (mc *MChiken) GetChikenSelf() int {
	return mc.GetChikenDraw()
}
