package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lottery_company/service"
	"net/http"
)

type UploadLotteryPlayerReq struct {
	LotteryPlayers []string `json:"lottery_players"`
}

// 仅有一个抽奖动作，导入 人名单
func UploadLotteryPlayers(ctx *gin.Context) {
	req := UploadLotteryPlayerReq{}
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	service.SetLotteryPlayers(req.LotteryPlayers)
	ctx.JSON(http.StatusOK, "success")
}

// 抽取动作，返回 单个人名称
func GetWinner(ctx *gin.Context) {
	msg := service.GetWinner()
	ctx.JSON(http.StatusOK, msg)
}
