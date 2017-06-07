package ting

import (
	"reflect"
	"sort"
	"testing"
)

// 测试可能听的牌
func TestMaybeTing(t *testing.T) {
	handCards := []int{1}
	showcards := []int{}
	maybeCards := []int{}

	maybeCards = GetMaybeTing(handCards, showcards)
	sort.Ints(maybeCards)
	if !reflect.DeepEqual(maybeCards, []int{1, 2}) {
		t.Error("验证可能听的牌失败1")
	}
	handCards = []int{1, 2}
	maybeCards = GetMaybeTing(handCards, showcards)
	sort.Ints(maybeCards)
	if !reflect.DeepEqual(maybeCards, []int{1, 2, 3}) {
		t.Error("验证可能听的牌失败2")
	}

	handCards = []int{6, 9}
	maybeCards = GetMaybeTing(handCards, showcards)
	sort.Ints(maybeCards)
	if !reflect.DeepEqual(maybeCards, []int{5, 6, 7, 8, 9}) {
		t.Error("验证可能听的牌失败3")
	}

	handCards = []int{6, 9}
	showcards = []int{3, 3, 3}
	maybeCards = GetMaybeTing(handCards, showcards)
	sort.Ints(maybeCards)
	if !reflect.DeepEqual(maybeCards, []int{3, 5, 6, 7, 8, 9}) {
		t.Error("验证可能听的牌失败4")
	}

	handCards = []int{11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16}
	showcards = []int{3, 3, 3}
	maybeCards = GetMaybeTing(handCards, showcards)
	sort.Ints(maybeCards)
	if !reflect.DeepEqual(maybeCards, []int{3, 11, 12, 13, 14, 15, 16, 17}) {
		t.Error("验证可能听的牌失败4")
	}

	handCards = []int{11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16}
	showcards = []int{8, 8, 8, 8}
	maybeCards = GetMaybeTing(handCards, showcards)
	sort.Ints(maybeCards)
	if !reflect.DeepEqual(maybeCards, []int{11, 12, 13, 14, 15, 16, 17}) {
		t.Error("验证可能听的牌失败5")
	}
}

// 测试能不能听牌
func TestCanTing(t *testing.T) {
	handCards := []int{1}
	showcards := []int{}
	isTing := false
	tingCards := []int{}

	isTing, tingCards = CanTing(handCards, showcards)
	if !isTing || !reflect.DeepEqual(tingCards, []int{1}) {
		t.Error("验证叫牌失败1")
	}

	handCards = []int{1, 2, 3, 4}
	showcards = []int{}
	isTing, tingCards = CanTing(handCards, showcards)
	sort.Ints(tingCards)
	if !isTing || !reflect.DeepEqual(tingCards, []int{1, 4}) {
		t.Error("验证叫牌失败2")
	}

	handCards = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	showcards = []int{8, 8, 8}
	isTing, tingCards = CanTing(handCards, showcards)
	sort.Ints(tingCards)
	if !isTing || !reflect.DeepEqual(tingCards, []int{1, 2, 4, 5, 8}) {
		t.Error("验证叫牌失败3")
	}
	handCards = []int{1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9}
	showcards = []int{}
	isTing, tingCards = CanTing(handCards, showcards)
	sort.Ints(tingCards)
	if !isTing || !reflect.DeepEqual(tingCards, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error("验证叫牌失败4")
	}

	handCards = []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 9, 11, 11, 18}
	showcards = []int{}
	isTing, tingCards = CanTing(handCards, showcards)
	sort.Ints(tingCards)
	if !isTing || !reflect.DeepEqual(tingCards, []int{18}) {
		t.Error("验证叫牌失败5")
	}
}
