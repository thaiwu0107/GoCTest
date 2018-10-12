package flopplayer2

import (
	"bytes"
	"sort"
	"strconv"

	"git.com.ggttoo44/src/config"
	Init "git.com.ggttoo44/src/init"
	Context "golang.org/x/net/context"
)

// Service 計算保險五張牌
func Service(ctx Context.Context, in *IntoData) (*OutData, error) {
	var lessAllPokers = []int{
		21, 22, 23, 24, 31, 32, 33, 34, 41, 42, 43, 44, 51, 52, 53, 54, 61, 62, 63, 64, 71, 72,
		73, 74, 81, 82, 83, 84, 91, 92, 93, 94, 101, 102, 103, 104, 111, 112, 113, 114, 121, 122,
		123, 124, 131, 132, 133, 134, 141, 142, 143, 144,
	}
	var removePoker = []int{}
	for _, publicPoker := range in.PublicPoker {
		removePoker = append(removePoker, int(publicPoker))
	}
	for _, playerPoker := range in.Data {
		removePoker = append(removePoker, int(playerPoker.Pokers[0]))
		removePoker = append(removePoker, int(playerPoker.Pokers[1]))
	}

	// 移除掉已經使用掉的牌
	for _, n := range removePoker {
		j := 0
		for _, nn := range lessAllPokers {
			if n != nn {
				lessAllPokers[j] = nn
				j++
			}
		}
		lessAllPokers = lessAllPokers[:j]
	}
	playerMap := make(map[string][5]int, 2)
	playerMap[in.Data[0].Id] = [5]int{
		int(in.PublicPoker[0]),
		int(in.PublicPoker[1]),
		int(in.PublicPoker[2]),
		int(in.Data[0].Pokers[0]),
		int(in.Data[0].Pokers[1]),
	}
	playerMap[in.Data[1].Id] = [5]int{
		int(in.PublicPoker[0]),
		int(in.PublicPoker[1]),
		int(in.PublicPoker[2]),
		int(in.Data[1].Pokers[0]),
		int(in.Data[1].Pokers[1]),
	}
	// 整理出誰目前贏面大
	var leaderID string
	ranks := make(map[string]int, 2)
	for index := 0; index < 2; index++ {
		ranks[in.Data[index].Id], _ = calculator5(playerMap[in.Data[index].Id])
	}
	sortRanks := sortMapByValue(ranks)
	leaderID = sortRanks[1].Key

	// 區分目前贏面最大的使用者跟牌面輸掉的使用者
	var loserID = ""
	for index4 := 0; index4 < 2; index4++ {
		if in.Data[index4].Id != leaderID {
			loserID = in.Data[index4].Id
		}
	}

	// 算出會輸的牌型
	overCardNumber := 0
	var leaderPoker [6]int
	for index1, poker1 := range playerMap[leaderID] {
		leaderPoker[index1] = poker1
	}
	var loserPoker [6]int
	for index2, poker2 := range playerMap[loserID] {
		loserPoker[index2] = poker2
	}
	for _, n := range config.ArrayC45_1 {
		pickPoker := lessAllPokers[n[0]]
		leaderPoker[5] = pickPoker
		loserPoker[5] = pickPoker
		var leader = calculator6(leaderPoker)
		var loser = calculator6(loserPoker)
		if leader < loser {
			overCardNumber++
		}
	}
	if overCardNumber == 0 {
		overCardNumber = -1
	}
	return &OutData{
		OverCardNumber: int32(overCardNumber),
		LeaderId:       leaderID,
	}, nil
}

// calculator5 單計算一次五張牌的rank
func calculator5(pokers [5]int) (int, error) {
	var cardMap = [5][15]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var loop5 = [5]int{0, 1, 2, 3, 4}
	var loop13 = [13]int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	var loopIndex int
	var pointNumber int
	var suitNuber int
	var suitInt int
	loopIndex = 0
	for _, loopIndex = range loop5 {
		suitInt = pokers[loopIndex]
		suitNuber = suitInt % 10
		pointNumber = (suitInt - suitNuber) / 10
		cardMap[suitNuber][pointNumber]++
		cardMap[suitNuber][0]++
		cardMap[0][pointNumber]++
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
	var rankInfo *Init.Rank5
	if isFlush {
		rankInfo = Init.RankTable5CF[keyOfRank]
	} else {
		rankInfo = Init.RankTable5CNF[keyOfRank]
	}
	return rankInfo.Rank, nil
}

// calculator6 單計算一次六張牌的rank
func calculator6(pokers [6]int) int {
	var cardMap = [5][15]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var loop5 = [5]int{0, 1, 2, 3, 4}
	var loop6 = [6]int{0, 1, 2, 3, 4, 5}
	var loop13 = [13]int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	var loopIndex int
	var pointNumber int
	var suitNuber int
	var suitInt int
	loopIndex = 0
	for _, loopIndex = range loop6 {
		suitInt = pokers[loopIndex]
		suitNuber = suitInt % 10
		pointNumber = (suitInt - suitNuber) / 10
		cardMap[suitNuber][pointNumber]++
		cardMap[suitNuber][0]++
		cardMap[0][pointNumber]++
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
	var rankInfo *Init.Rank6
	if isFlush {
		rankInfo = Init.RankTable6CF[keyOfRank]
	} else {
		rankInfo = Init.RankTable6CNF[keyOfRank]
	}
	return rankInfo.Rank
}

func omitArray(sArray []int, removeArray []int) (outArray []int) {
	for _, n := range removeArray {
		j := 0
		for _, nn := range sArray {
			if n != nn {
				sArray[j] = nn
				j++
			}
		}
		sArray = sArray[:j]
	}
	return sArray
}

// Pair A slice of Pairs that implements sort.Interface to sort by Value.
type Pair struct {
	Key   string
	Value int
}

// PairList A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
