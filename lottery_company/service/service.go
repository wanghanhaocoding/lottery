package service

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var lotteryPlayers []string

var mu sync.Mutex

func SetLotteryPlayers(players []string) {
	userList := make([]string, 0, len(players))
	for _, player := range players {
		player = strings.TrimSpace(player)
		if len(player) == 0 {
			continue
		}
		userList = append(userList, player)
	}
	lotteryPlayers = userList
}
func GetWinner() string {
	mu.Lock()
	defer mu.Unlock()
	count := len(lotteryPlayers)
	if count == 0 {
		return "当前没有用户参与抽奖了"
	}
	if count == 1 {
		user := lotteryPlayers[0]
		lotteryPlayers = []string{}
		return fmt.Sprintf("当前中奖用户是%s,还剩下%d个用户", user, count-1)
	}
	seed := time.Now().UnixNano()
	index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
	user := lotteryPlayers[index]
	lotteryPlayers = append(lotteryPlayers[:index], lotteryPlayers[index+1:]...)
	return fmt.Sprintf("当前中奖用户是%s,还剩下%d个用户", user, count-1)
}
