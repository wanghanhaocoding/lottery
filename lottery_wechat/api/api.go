package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"lottery_wechat/service"
	"net/http"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello")
}

func InitPrize(ctx *gin.Context) {
	req := service.InitPrizeReq{}
	if err := ctx.BindJSON(&req); err != nil {
		log.Errorf("InitPrize err:%v", err)
		ctx.JSON(http.StatusBadRequest, 200)
		return
	}
	if err := service.AddPrize(req.ViewPrizeList); err != nil {
		log.Errorf("api|AddPrize err:%v", err)
		ctx.JSON(http.StatusInternalServerError, 500)
		return
	}
	ctx.JSON(http.StatusOK, "success")
}

func GetPrizeInfo(ctx *gin.Context) {
	rsp := service.GetPrizeInfoRsp{}
	prizeList, err := service.GetPrizeList()
	if err != nil {
		log.Errorf("api|GetPrizeInfo err:%v", err)
		ctx.JSON(http.StatusInternalServerError, 500)
	}
	var count int = 0
	var total int64 = 0
	for _, prize := range prizeList {
		if prize.Total == 0 || (prize.Total > 0 && prize.Left > 0) {
			count++
			total += prize.Total
		}
	}
	rsp.PrizeTotal = total
	rsp.PrizeTypeNum = count
	ctx.JSON(http.StatusOK, rsp)
}

func Lottery(ctx *gin.Context) {
	res := service.GetWinner()
	ctx.JSON(http.StatusOK, res)
}
