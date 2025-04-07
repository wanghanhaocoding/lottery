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
