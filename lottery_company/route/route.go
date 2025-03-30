package route

import (
	"github.com/gin-gonic/gin"
	"lottery_company/api"
)

func SetRoute() *gin.Engine {
	r := gin.Default()
	group := r.Group("/lottery_system")
	group.POST("/upload_lottery_players", api.UploadLotteryPlayers)
	group.GET("/lottery", api.GetWinner)
	return r
}
