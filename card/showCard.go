package card

import (
	"fmt"

	"github.com/fwhappy/mahjong/fbs/Common"
)

// ShowCard 明牌
type ShowCard struct {
	opCode int   // 操作类型，对应吃、碰、杠对应的操作类型id
	target int   // 明牌对象，吃、碰、杠的牌是谁打出来的
	tiles  []int // 关联的牌
	free   bool  // 是否付费，用于转弯杠，暂时用不上了
}

// NewShowCard 生成一个明牌
func NewShowCard(opCode, target int, tiles []int, free bool) *ShowCard {
	showCard := &ShowCard{opCode: opCode, target: target, tiles: tiles, free: free}
	return showCard
}

func (s *ShowCard) string() string {
	return fmt.Sprintf("[明牌]code:%v,target:%v,tiles:%v,free:%v", s.opCode, s.target, s.tiles, s.free)
}

// GetOpCode 获取明牌类型
func (s *ShowCard) GetOpCode() int {
	return s.opCode
}

// GetTiles 获取明牌类型
func (s *ShowCard) GetTiles() []int {
	return s.tiles
}

// GetTile 返回明牌中的牌是什么
// 至于这个showCard是不是吃，需要外面的逻辑判断
func (s *ShowCard) GetTile() int {
	return s.tiles[0]
}

// GetTarget 获取明牌对象
func (s *ShowCard) GetTarget() int {
	return s.target
}

// IsFree 是否免费
func (s *ShowCard) IsFree() bool {
	return s.free
}

// GetTilesLen 牌的数量
func (s *ShowCard) GetTilesLen() int {
	return len(s.tiles)
}

// ModifyPongToKong 将碰修改成杠
func (s *ShowCard) ModifyPongToKong(kongCode int, free bool) {
	s.opCode = kongCode
	s.free = free
	s.tiles = append(s.tiles, s.tiles[0])
}

// ModifyQiangKong 将kong设置为被抢的状态
func (s *ShowCard) ModifyQiangKong() {
	s.tiles = append([]int{}, s.tiles[0:s.GetTilesLen()-1]...)
}

// IsPong 明牌是否是pong
func (s *ShowCard) IsPong() bool {
	return s.opCode == Common.OperationCodePONG
}

// IsPongTile 明牌是否是pong了这个牌
func (s *ShowCard) IsPongTile(tile int) bool {
	return s.opCode == Common.OperationCodePONG && s.tiles[0] == tile
}
