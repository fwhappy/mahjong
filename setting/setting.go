package setting

// MSetting 游戏设置
type MSetting struct {
	lack          bool // 是否支持定缺
	pinghu        bool // 是否支持平胡
	exchange      bool // 是否支持换牌
	EnableKongTXZ bool // 明杠是否算通行证

	// 积分倍数
	Multiple int // 积分倍数

	// 牌型支持
	EnableShuangLongQiDui bool // 是否支持双龙七对
	EnableBianKaDiao      bool // 是否支持边卡吊
	EnableDaKuanZhang     bool // 是否支持大宽张
	EnableDi7Dui          bool // 是否支持地七对
	EnablePinghuZimo      bool // 是否支持平胡自摸
	EnableDanDiao         bool // 是否支持平胡自摸
	EnableDoubleDealer    bool // 是否支持庄家胡牌翻倍
	EnableKongAfterDraw   bool // 是否支持杠上开花，这里是指把杠上开花独立作为一个牌型，杠上开花时，不计其他牌型

	// 按顺序，0：满堂鸡；1：连庄；2：上下鸡；3：乌骨鸡；4：前后鸡；5：星期鸡；6：意外鸡；7：吹风鸡；8：滚筒鸡；9：麻将人数；10：麻将张数；11：本鸡
	// 12：站鸡；13：翻倍鸡；14：首圈鸡；15：清一色奖励三分；16：自摸翻倍；17：一扣二；18：通三；19：大牌翻倍
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
	mSetting := &MSetting{
		lack:                  false,
		pinghu:                false,
		exchange:              false,
		EnableKongTXZ:         true,
		Multiple:              1,
		setting:               []int{},
		EnableShuangLongQiDui: true,
		EnableBianKaDiao:      false,
		EnableDaKuanZhang:     false,
		EnableDi7Dui:          false,
		EnablePinghuZimo:      false,
		EnableDanDiao:         true,
		EnableDoubleDealer:    false,
		EnableKongAfterDraw:   false,
	}
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
	return len(ms.setting) >= 12 && ms.setting[11] == 1
}

// IsSettingChikenRock 是否设置了滚鸡
func (ms *MSetting) IsSettingChikenRock() bool {
	return ms.IsSettingChikenFB() || ms.IsSettingChikenTumbling()
}

// IsSettingStandChiken 是否支持站鸡
func (ms *MSetting) IsSettingStandChiken() bool {
	return len(ms.setting) >= 13 && ms.setting[12] == 1
}

// IsSettingDoubleChiken 是否支持翻倍鸡
func (ms *MSetting) IsSettingDoubleChiken() bool {
	return len(ms.setting) >= 14 && ms.setting[13] == 1
}

// IsSettingFirstCycleChiken 是否支持首圈鸡
func (ms *MSetting) IsSettingFirstCycleChiken() bool {
	return len(ms.setting) >= 15 && ms.setting[14] == 1
}

// IsSettingQE 是否支持清一色加3分
func (ms *MSetting) IsSettingQE() bool {
	return len(ms.setting) >= 16 && ms.setting[15] == 1
}

// IsSettingDoubleZM 是否支持自摸翻倍
func (ms *MSetting) IsSettingDoubleZM() bool {
	return len(ms.setting) >= 17 && ms.setting[16] == 1
}

// IsSettingZME 是否支持自摸加一分
func (ms *MSetting) IsSettingZME() bool {
	return len(ms.setting) >= 18 && ms.setting[17] == 1
}

// IsSettingTS 是否支持通三
func (ms *MSetting) IsSettingTS() bool {
	return len(ms.setting) >= 19 && ms.setting[18] == 1
}

// IsSettingDoubleDP 是否支持大牌翻倍
func (ms *MSetting) IsSettingDoubleDP() bool {
	return len(ms.setting) >= 20 && ms.setting[19] == 1
}
