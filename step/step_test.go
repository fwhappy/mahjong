package step

import "testing"

func TestGetTilesStep(t *testing.T) {
	var handTiles = []int{}

	handTiles = []int{1, 2, 3, 4, 4}
	if step := GetTilesStep(handTiles); step != 2 {
		t.Errorf("计算牌阶错误,step:%v,%v", step, handTiles)
	}

	handTiles = []int{1, 2, 3, 4, 5}
	if step := GetTilesStep(handTiles); step != 1 {
		t.Errorf("计算牌阶错误,step:%v,%v", step, handTiles)
	}

	handTiles = []int{1, 1, 1, 2, 3, 4, 5, 6}
	if step := GetTilesStep(handTiles); step != 3 {
		t.Errorf("计算牌阶错误,step:%v,%v", step, handTiles)
	}

	handTiles = []int{1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9}
	if step := GetTilesStep(handTiles); step != 4 {
		t.Errorf("计算牌阶错误,step:%v,%v", step, handTiles)
	}

}
