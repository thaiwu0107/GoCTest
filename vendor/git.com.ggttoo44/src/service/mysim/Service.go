package mysim

import (
	"bytes"
	"strconv"

	Init "git.com.ggttoo44/src/init"
	Context "golang.org/x/net/context"
)

// Service 單計算一次七張牌的rank
func Service(ctx Context.Context, in *IntoData) (*OutData, error) {
	var cardMap = [5][15]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var loop5 = [5]int{0, 1, 2, 3, 4}
	var loop7 = [7]int{0, 1, 2, 3, 4, 5, 6}
	var loop13 = [13]int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	var loopIndex int
	var pointNumber int
	var suitNuber int
	var suitInt int
	suitMap := make(map[int][]int, 7)
	loopIndex = 0
	for _, loopIndex = range loop7 {
		suitInt, _ = strconv.Atoi(in.Data[loopIndex])
		suitNuber = suitInt % 10
		pointNumber = (suitInt - suitNuber) / 10
		cardMap[suitNuber][pointNumber]++
		cardMap[suitNuber][0]++
		cardMap[0][pointNumber]++

		if suitMap[pointNumber] != nil {
			suitMap[pointNumber] = append(suitMap[pointNumber], suitNuber)
		} else {
			suitMap[pointNumber] = []int{suitNuber}
		}
	}
	isFlush := false
	var selectSuit int
	loopIndex = 0
	for _, loopIndex = range loop5 {
		if cardMap[loopIndex][0] >= 5 {
			isFlush = true
			selectSuit = loopIndex
			break
		}
	}
	var buffer bytes.Buffer
	loopIndex = 14
	for _, loopIndex = range loop13 {
		buffer.WriteString(strconv.Itoa(cardMap[selectSuit][loopIndex]))
	}
	keyOfRank := buffer.String()
	var rankInfo *Init.Rank7
	selectCards := make([]string, 5)
	if isFlush {
		rankInfo = Init.RankTable7CF[keyOfRank]
		loopIndex = 0
		for _, loopIndex = range loop5 {
			var point1 int
			point1 = rankInfo.CardPoint[loopIndex]
			selectCards[loopIndex] = strconv.Itoa(point1*10 + selectSuit)
		}
	} else {
		rankInfo = Init.RankTable7CNF[keyOfRank]
		loopIndex = 0
		for _, loopIndex = range loop5 {
			var point2 int
			var suit int
			point2 = rankInfo.CardPoint[loopIndex]
			suit, suitMap[point2] = suitMap[point2][len(suitMap[point2])-1], suitMap[point2][:len(suitMap[point2])-1]
			selectCards[loopIndex] = strconv.Itoa(point2*10 + suit)
		}
	}
	return &OutData{
		CardPoint: selectCards,
		Type7:     rankInfo.Type7,
		Rank:      int32(rankInfo.Rank),
		Type5Ch:   rankInfo.Type5Ch,
		Type5En:   rankInfo.Type5En,
	}, nil
}
