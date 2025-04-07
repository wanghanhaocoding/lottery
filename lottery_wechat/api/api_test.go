package api

import (
	"encoding/json"
	"lottery_wechat/internal/pkg/constant"
	"lottery_wechat/service"
	"testing"
)

func TestAddPrize(t *testing.T) {
	prizeList := make([]*service.ViewPrize, 4)
	prizeCoin := service.ViewPrize{
		ID:             1,
		Name:           "Q币",
		Pic:            "http://q.qlogo.cn/g?b=qq&nk=1&s=100&nk2=1&s2=100",
		Link:           "http://q.qq.com",
		Type:           constant.PrizeTypeCoin,
		Data:           "100Q币",
		Total:          20000,
		Left:           20000,
		IsUse:          1,
		Probability:    5000, // 百分之百中奖
		ProbabilityMax: 0,
		ProbabilityMin: 0,
	}
	prizeList[0] = &prizeCoin

	prizeSmallEntity := service.ViewPrize{
		ID:             2,
		Name:           "充电宝",
		Pic:            "",
		Link:           "",
		Type:           constant.PrizeTypeSmallEntity,
		Data:           "",
		Total:          100,
		Left:           100,
		IsUse:          1,
		Probability:    100, // 百分之1中奖
		ProbabilityMax: 0,
		ProbabilityMin: 0,
	}
	prizeList[1] = &prizeSmallEntity

	prizeTypeLargeEntity := service.ViewPrize{
		ID:             3,
		Name:           "iphone14",
		Pic:            "",
		Link:           "",
		Type:           constant.PrizeTypeLargeEntity,
		Data:           "",
		Total:          10,
		Left:           10,
		IsUse:          1,
		Probability:    10, // 千分之1中奖
		ProbabilityMax: 0,
		ProbabilityMin: 0,
	}
	prizeList[2] = &prizeTypeLargeEntity

	prizeTypeCoupon := service.ViewPrize{
		ID:             4,
		Name:           "优惠券满100减10元",
		Pic:            "",
		Link:           "",
		Type:           constant.PrizeTypeCoupon,
		Data:           "黄焖鸡外卖优惠券",
		Total:          5000,
		Left:           5000,
		IsUse:          1,
		Probability:    3000,
		ProbabilityMax: 0,
		ProbabilityMin: 0,
	}
	prizeList[3] = &prizeTypeCoupon

	var start int64 = 0
	for _, prize := range prizeList {
		if prize.IsUse != constant.PrizeInUse {
			continue
		}
		prize.ProbabilityMin = start
		prize.ProbabilityMax = start + prize.Probability
		if prize.ProbabilityMax >= constant.ProbabilityLimit {
			prize.ProbabilityMax = constant.ProbabilityLimit
			start = 0
		} else {
			start += prize.Probability
		}
	}

	viewPrizeList := []*service.ViewPrize{
		&prizeCoin, &prizeSmallEntity, &prizeTypeLargeEntity, &prizeTypeCoupon,
	}
	addPrizeReq := service.InitPrizeReq{
		ViewPrizeList: viewPrizeList,
	}
	bytesData, err := json.Marshal(&addPrizeReq)
	if err != nil {
		t.Errorf("Marshal req err:%v", err)
	}
	t.Log(string(bytesData))
}
