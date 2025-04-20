package repo

import (
	log "github.com/sirupsen/logrus"
	"lottery_wechat/internal/gormcli"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/pkg/constant"
)

func AddPrize(prizeList []*model.Prize) error {
	db := gormcli.GetDb()
	if err := db.Model(&model.Prize{}).Create(prizeList).Error; err != nil {
		log.Errorf("repo|add prize err:%v", err)
		return err
	}
	log.Infof("repo|add prize success")
	return nil
}

func GetPrizeList() ([]*model.Prize, error) {
	db := gormcli.GetDb()
	var prizeList []*model.Prize
	err := db.Model(&model.Prize{}).Where("is_use = ?", constant.PrizeInUse).Find(&prizeList).Error
	if err != nil {
		log.Errorf("repo|GetPrizeList err:%v", err)
		return nil, err
	}
	return prizeList, nil
}

func SavePrize(prize *model.Prize) error {
	db := gormcli.GetDb()
	if err := db.Model(&model.Prize{}).Where("id=?", prize.ID).Save(prize).Error; err != nil {
		log.Errorf("repo|SavePrize err:%v", err)
		return err
	}
	return nil
}
