package wall

import "testing"
import "fmt"
import "reflect"

func TestWall(t *testing.T) {
	tiles := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	tileWall := NewWall()
	tileWall.SetTiles(tiles)

	if tileWall.Length() != len(tiles) {
		t.Error("牌墙长度计算错误")
	}

	if tileWall.GetTile(10) != tiles[10] {
		t.Error("牌墙索引错误")
	}

	tileWall.Shuffle()
	fmt.Println("洗牌结果:", tileWall.GetTiles())
}

func TestDraw(t *testing.T) {
	tiles := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	tileWall := NewWall()
	tileWall.SetTiles(tiles)

	// 测试从前面抓牌
	if !reflect.DeepEqual([]int{1, 2, 3, 4, 5}, tileWall.ForwardDrawMulti(5)) {
		t.Error("批量抓牌测试失败")
	}
	// 测试读取下一张牌的值
	if tileWall.GetFrowrdNextTile() != 6 {
		t.Error("读取下一个牌面值失败")
	}
	// 测试从后面抓牌
	if tileWall.BackwardDraw() != 19 {
		t.Error("从后面抓第一张牌失败")
	}
	if tileWall.BackwardDraw() != 20 {
		t.Error("从后面抓第二张牌失败")
	}
	if tileWall.BackwardDraw() != 17 {
		t.Error("从后面抓第三张牌失败")
	}

	// 重置
	tileWall.forward = 0
	tileWall.backward = 0

	// 从前面抓19张
	tileWall.ForwardDrawMulti(19)

	// 测试从后面抓牌
	if tileWall.BackwardDraw() != 20 {
		t.Error("从后面抓最后一张牌失败")
	}
}

func TestDrawn(t *testing.T) {
	tiles := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	tileWall := NewWall()
	tileWall.SetTiles(tiles)

	// 初始状态
	if tileWall.IsDrawn(0) || tileWall.IsDrawn(1) {
		t.Error("判断牌有没有被抓过失败")
	}
	if tileWall.IsDrawn(19) || tileWall.IsDrawn(18) {
		t.Error("判断牌有没有被抓过失败")
	}

	// 前面抓过了
	tileWall.forward = 5
	if !tileWall.IsDrawn(4) || !tileWall.IsDrawn(0) {
		t.Error("判断牌有没有被抓过失败")
	}
	if tileWall.IsDrawn(5) {
		t.Error("判断牌有没有被抓过失败")
	}

	// 后面抓过奇数张
	tileWall.backward = 1
	if tileWall.IsDrawn(19) || tileWall.IsDrawn(17) || tileWall.IsDrawn(16) {
		t.Error("判断牌有没有被抓过失败")
	}
	if !tileWall.IsDrawn(18) {
		t.Error("判断牌有没有被抓过失败")
	}

	// 后面抓过偶数张
	tileWall.backward = 2
	if tileWall.IsDrawn(17) || tileWall.IsDrawn(16) || tileWall.IsDrawn(15) {
		t.Error("判断牌有没有被抓过失败")
	}
	if !tileWall.IsDrawn(18) || !tileWall.IsDrawn(19) {
		t.Error("判断牌有没有被抓过失败")
	}
}
