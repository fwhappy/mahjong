package chiken

import (
	"testing"

	"reflect"

	"github.com/fwhappy/mahjong/setting"
	"github.com/fwhappy/mahjong/wall"
)

// 翻牌鸡测试用例
func TestChiken(t *testing.T) {
	tiles := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	tileWall := wall.NewWall()
	tileWall.SetTiles(tiles)
	ms := setting.NewMSetting()
	ms.SetSetting([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 20, 0})
	mc := NewMChiken()

	// 冲锋鸡
	mc.SetChargeBam1(100)
	mc.SetChargeBam1(200)
	mc.SetResponsibility(300)
	if mc.GetChargeBam1() != 100 {
		t.Error("冲锋鸡验证失败1")
	}
	if mc.GetResponsibility() != 0 {
		t.Error("责任鸡验证失败1")
	}
	mc = NewMChiken()
	mc.SetResponsibility(300)
	mc.SetChargeBam1(100)
	if mc.GetResponsibility() != 300 || mc.GetChargeBam1() != 0 {
		t.Error("责任鸡验证失败2")
	}
	mc.SetChargeDot8(500)
	if mc.GetChargeDot8() != 500 {
		t.Error("冲锋乌骨验证失败1")
	}

	// 翻牌鸡
	mc.GenChikenDraw(tileWall)
	if mc.GetChikenDraw() != 1 {
		t.Error("翻牌鸡验证失败1")
	}
	tileWall.ForwardDrawMulti(10)
	mc = NewMChiken()
	mc.GenChikenDraw(tileWall)
	if mc.GetChikenDraw() != 11 {
		t.Error("翻牌鸡验证失败2")
	}

	tileWall.ForwardDrawMulti(10)
	mc = NewMChiken()
	mc.GenChikenDraw(tileWall)
	if mc.GetChikenDraw() != 0 {
		t.Error("翻牌鸡验证失败2")
	}
	if mc.GetChikenSelf() != mc.GetChikenDraw() {
		t.Error("本鸡验证失败")
	}

	// 星期鸡
	// if !reflect.DeepEqual(mc.GetWeekChikens(), []int{4, 14, 24}) {
	// 	t.Error("翻牌鸡验证失败2")
	// }

	// 前后鸡
	var fbIndex, fbTile int
	tileWall = wall.NewWall()
	tileWall.SetTiles(tiles)
	mc = NewMChiken()
	mc.GenChikenFB(tileWall)
	fbIndex, fbTile = mc.GetChikenFB(tileWall)
	if fbIndex != 14 || fbTile != 15 {
		t.Error("前后鸡验证失败1")
	}
	tileWall.ForwardDrawMulti(1)
	fbIndex, fbTile = mc.GetChikenFB(tileWall)
	if fbIndex != 14 || fbTile != 15 {
		t.Error("前后鸡验证失败2")
	}
	tileWall.ForwardDrawMulti(13)
	fbIndex, fbTile = mc.GetChikenFB(tileWall)
	if fbIndex != 14 || fbTile != 15 {
		t.Error("前后鸡验证失败3")
	}
	tileWall.ForwardDraw()
	fbIndex, fbTile = mc.GetChikenFB(tileWall)
	if fbIndex != 0 || fbTile != 0 {
		t.Error("前后鸡验证失败3")
	}
	tileWall = wall.NewWall()
	tileWall.SetTiles(tiles)
	mc = NewMChiken()
	mc.GenChikenFB(tileWall)
	// 后面连抓5次
	for i := 0; i < 5; i++ {
		tileWall.BackwardDraw()
	}
	fbIndex, fbTile = mc.GetChikenFB(tileWall)
	if fbIndex != 0 || fbTile != 0 {
		t.Error("前后鸡验证失败3")
	}
	// 滚筒鸡
	tileWall = wall.NewWall()
	tileWall.SetTiles(tiles)
	mc = NewMChiken()

	// 连滚3次
	for i := 0; i < 3; i++ {
		mc.AppendChikenTumbling(tileWall)
	}
	if !reflect.DeepEqual(mc.GetChikenTumblingMap(tileWall), map[int]int{14: 15, 12: 13, 10: 11}) {
		t.Error("滚筒鸡验证失败1")
	}
	// 后面连抓5次
	for i := 0; i < 5; i++ {
		tileWall.BackwardDraw()
	}
	if !reflect.DeepEqual(mc.GetChikenTumblingMap(tileWall), map[int]int{12: 13, 10: 11}) {
		t.Error("滚筒鸡验证失败2")
	}
	// 前面抓11张
	tileWall.ForwardDrawMulti(11)
	if !reflect.DeepEqual(mc.GetChikenTumblingMap(tileWall), map[int]int{12: 13}) {
		t.Error("滚筒鸡验证失败3")
	}

}
