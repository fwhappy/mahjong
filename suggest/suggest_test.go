package suggest

import (
	"testing"

	"github.com/fwhappy/mahjong/card"
)

func BenchmarkSuggestMap(b *testing.B) {
	b.StopTimer()
	ms := NewMSelector()
	ms.SetTiles(card.MahjongCards108)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		handTiles := []int{6, 7, 9, 9, 12, 12, 13, 14, 15, 15, 17, 26, 27, 28}
		ms.SetHandTilesSlice(handTiles)
		ms.CalcRemaimTiles()
		ms.GetSuggestMap(2)
	}
}
