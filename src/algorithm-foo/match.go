package algorithmfoo

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Mericusta/go-stp"
	"golang.org/x/exp/slices"
)

var (
	random    = rand.New(rand.NewSource(time.Now().UnixNano()))
	rankRange = 10
	rankLevel = []int{
		0, 1, 2, 3, 4, // 青铜
		5, 6, 7, 8, 9, // 白银
		10, 11, 12, 13, 14, // 黄金
		15, 16, 17, 18, 19, // 铂金
		20, 21, 22, 23, 24, // 钻石
	}
	matchMaxDiffRank = 1
	matchLine        = make([]*group, 0, 64)
	playerCount      = 30
	playerIndex      = 0
	needSplit        = true
	matchDone        = 3
)

func calculateRank(score int) int {
	rank := rankLevel[len(rankLevel)-1]
	for _, _rank := range rankLevel {
		if score > _rank*10 {
			rank = _rank
			continue
		}
		break
	}
	return rank
}

func generatePlayer() *player {
	id := playerIndex
	score := random.Intn(rankRange * len(rankLevel))
	rank := calculateRank(score)
	split()
	fmt.Printf("new player %v score %v, rank %v\n", id, score, rank)
	split()
	playerIndex++
	return &player{id: id, score: score, rank: rank}
}

type player struct {
	id    int
	score int
	rank  int
}

type group struct {
	players []*player
}

func newGroup() *group {
	return &group{players: make([]*player, 0, matchDone)}
}

func (g *group) appendPlayer(player *player) {
	insertIndex := 0
	for index, _player := range g.players {
		if _player.score < player.score {
			continue
		}
		insertIndex = index
		break
	}
	g.players = slices.Insert(g.players, insertIndex, player)
}

func (g *group) calculateMinMaxScore() (int, int) {
	minRank, maxRank := -1, -1
	for _, player := range g.players {
		_min := player.rank - matchMaxDiffRank
		_max := player.rank + matchMaxDiffRank
		if _min <= minRank || minRank == -1 {
			minRank = _min
		}
		if _max >= maxRank || maxRank == -1 {
			maxRank = _max
		}
	}
	return minRank, maxRank
}

func split() {
	if needSplit {
		fmt.Println("--------------------------------")
		needSplit = false
	} else {
		needSplit = true
	}
}

func match(player *player) *group {
	split()
	for index, group := range matchLine {
		groupMinRank, groupMaxRank := group.calculateMinMaxScore()
		fmt.Printf("group %v min rank %v, max rank %v\n", index, groupMinRank, groupMaxRank)
		if groupMinRank <= player.rank && player.rank <= groupMaxRank {
			group.appendPlayer(player)
			fmt.Printf("append player %v to match line group %v\n", player.score, index)
			if len(group.players) == matchDone {
				matchLine = append(matchLine[:index], matchLine[index+1:]...)
				return group
			}
			return nil
		}
	}

	group := newGroup()
	group.appendPlayer(player)
	matchLine = append(matchLine, group)
	fmt.Printf("append player %v to match line new group\n", player.score)
	split()
	return nil
}

func outputMatchLine() {
	split()
	fmt.Printf("match line group count %v\n", len(matchLine))
	for index, group := range matchLine {
		players := make([]int, 0, len(group.players))
		stp.NewArray(group.players).ForEach(func(v *player, i int) {
			players = append(players, v.score)
		})
		fmt.Printf("group %v, players %v\n", index, players)
	}
	split()
}

func outputMatchDoneGroup(group *group) {
	split()
	players := make([]int, 0, len(group.players))
	stp.NewArray(group.players).ForEach(func(v *player, i int) {
		players = append(players, v.score)
	})
	fmt.Printf("match done group players %v\n", players)
	split()
}

func matchFoo() {
	for index := 0; index != playerCount; index++ {
		player := generatePlayer()
		matchDoneGroup := match(player)
		if matchDoneGroup != nil {
			outputMatchDoneGroup(matchDoneGroup)
		}
		outputMatchLine()
	}
}
