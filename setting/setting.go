package setting

// MSetting 游戏设置
type MSetting struct {
	lack     bool // 是否支持定缺
	pinghu   bool // 是否支持平胡
	exchange bool // 是否支持换牌

	// 按顺序，0：满堂鸡；1：连庄；2：上下鸡；3：乌骨鸡；4：前后鸡；5：星期鸡；6：意外鸡；7：吹风鸡；8：滚筒鸡；9：麻将人数；10：麻将张数；11：本鸡
	// 满堂鸡: 打出去的牌也算鸡
	// 上下鸡: 翻出来的牌，上一张也算鸡，eg: 翻出来5条，如果开启了上下鸡，则4条也算鸡
	// 乌骨鸡: 8筒
	// 前后鸡: 开局时倒数第三墩上面翻开的鸡，每杠一次，在前面一墩翻开一张并关闭当前墩
	// 星期鸡: 今天星期几，则几条、几筒、几万都算鸡，周日算7
	// 吹风鸡: 5筒，翻到5筒，所以的鸡、杠、胡都不算，直接开始下一局
	// 滚筒鸡: 开局时倒数第三墩上面翻开的鸡，每杠一次，在前面一墩翻开一张，当前墩的牌不关闭
	setting []int
}

// NewMSetting 新建一个MSetting的引用
func NewMSetting() *MSetting {
	mSetting := &MSetting{lack: false, pinghu: false, exchange: false, setting: []int{}}
	return mSetting
}

// SetSetting 设置需要定缺
func (ms *MSetting) SetSetting(setting []int) {
	ms.setting = setting
}

// GetSetting 获取设置内容
func (ms *MSetting) GetSetting() []int {
	return ms.setting
}

// SetEnableLack 设置需要定缺
func (ms *MSetting) SetEnableLack() {
	ms.lack = true
}

// IsEnableLack 是否支持定缺
func (ms *MSetting) IsEnableLack() bool {
	return ms.lack
}

// SetEnableExchange 设置支持换牌
func (ms *MSetting) SetEnableExchange() {
	ms.exchange = true
}

// IsEnableExchange 是否支持定缺
func (ms *MSetting) IsEnableExchange() bool {
	return ms.exchange
}

// SetEnablePinghu 设置需要定缺
func (ms *MSetting) SetEnablePinghu() {
	ms.pinghu = true
}

// IsEnablePinghu 是否支持定缺
func (ms *MSetting) IsEnablePinghu() bool {
	return ms.pinghu
}

// GetSettingPlayerCnt 获取设置的玩家人数
func (ms *MSetting) GetSettingPlayerCnt() int {
	return ms.setting[9]
}

// GetSettingTileCnt 获取设置的麻将张数
func (ms *MSetting) GetSettingTileCnt() int {
	return ms.setting[10]
}

// IsSettingAllChikenDraw 是否设置了满堂鸡
func (ms *MSetting) IsSettingAllChikenDraw() bool {
	return ms.setting[0] == 1
}

// IsSettingRemainDealer 是否设置了连庄
func (ms *MSetting) IsSettingRemainDealer() bool {
	return ms.setting[1] == 1
}

// IsSettingChikenUD 是否设置了上下鸡
func (ms *MSetting) IsSettingChikenUD() bool {
	return ms.setting[2] == 1
}

// IsSettingChikenDot8 是否设置了乌骨鸡
func (ms *MSetting) IsSettingChikenDot8() bool {
	return ms.setting[3] == 1
}

// IsSettingChikenFB 是否设置了前后鸡
func (ms *MSetting) IsSettingChikenFB() bool {
	return ms.setting[4] == 1
}

// IsSettingChikenWeekday 是否设置了星期鸡
func (ms *MSetting) IsSettingChikenWeekday() bool {
	return ms.setting[5] == 1
}

// IsSettingChikenUnexpect 是否设置了意外鸡
func (ms *MSetting) IsSettingChikenUnexpect() bool {
	return ms.setting[6] == 1
}

// IsSettingChikenWind 是否支持吹风鸡
func (ms *MSetting) IsSettingChikenWind() bool {
	return ms.setting[7] == 1
}

// IsSettingChikenTumbling 是否支持滚筒鸡
func (ms *MSetting) IsSettingChikenTumbling() bool {
	return ms.setting[8] == 1
}

// IsSettingChikenSelf 是否支持本鸡
func (ms *MSetting) IsSettingChikenSelf() bool {
	return ms.setting[11] == 1
}

// IsSettingChikenRock 是否设置了滚鸡
func (ms *MSetting) IsSettingChikenRock() bool {
	return ms.IsSettingChikenFB() || ms.IsSettingChikenTumbling()
}
