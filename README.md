# 麻将算法封装

* 此库主要包括了
	* 牌墙算法
	* 听牌算法
	* 胡牌算法
	* 出牌推荐算法
	* 扩展玩法设置

* 麻将牌的定义

~~~
	1 ~ 9 : 一万 ~ 9万
	11 ~ 19 : 一条 ~ 9条
	21 ~ 29 : 一筒 ~ 9筒
	31 ~ 34 : 东南西北风
	41 : 发财
	42 : 红中
	43 : 白板
	51 ~ 54 : 春夏秋冬
	61 ~ 64 : 四种花色
~~~

* 牌墙

~~~
import "github.com/fwhappy/mahjong/wall"

func main() {
	// 初始化牌墙
	w := wall.NewWall()
	w.SetTiles([]int{1,1,1,1,2,2,2,2})
	
	// 洗牌
	w.Shuffle()
	
	// 从前面抓一张
	wall.ForwardDraw()
	
	// 从前面抓多张
	wall.ForwardDrawMulti()
	
	// 从后面抓一张
	wall.BackwardDraw()
	
	// 是否已抓完
	wall.IsAllDrawn()
}
~~~

* 胡牌算法

~~~
	import "github.com/fwhappy/mahjong/win"
	
	func main() {
		// 手牌
		handTiles := []int{1,2,3,4,5,6,7,7}
		// 明牌
		showTiles := []int{}
		
		// 根据用户手牌和明牌来判断用户是否可以胡牌
		isWin := win.CanWin(handTiles, showTiles)
	}
~~~

* 听牌算法

~~~
	import "github.com/fwhappy/mahjong/ting"
	
	func main() {
		// 根据手牌、弃牌计算当前牌型，所有听牌的可能
		// 返回map[int][]int, 表示打出key，能胡value中的这些牌
		// 手牌
		handTiles := []int{1,2,3,4,5,6,7,7}
		// 明牌
		showTiles := []int{}
		ting.GetTingMap(handTiles, showTiles)
		
		// 检测当前手牌、弃牌是否已停牌
		ting.CanTing(handTiles, showTiles)
	}
~~~

* 选牌算法（AI）

~~~
	import "github.com/fwhappy/mahjong/suggest"
	
	func main() {
		ms := suggest.NewMSelector()
		
		// 根据用户手牌、明牌，弃牌（所有人的）
		// 计算出用户当前应该出什么牌
		
		// 手牌
		handTiles := []int{1,2,3,4,5,6,7,7}
		// 明牌
		showTiles := []int{}
		// 弃牌
		discardTiles := []int{}
		
		// 设置参数
		ms.SetAILevel(aiLevel) // AI等级
		ms.SetLack(lack)	// 缺的牌
		ms.SetHandTilesSlice(handTiles)
		ms.SetShowTilesSlice(showTiles)
		ms.SetDiscardTilesSlice(discardTiles)
		// 选牌
		tile := ms.GetSuggest()
	}
~~~
