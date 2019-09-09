package model

import "github.com/jinzhu/gorm"

type CreditHistory struct {
	gorm.Model
	UserId  uint   `json:"userId"`
	Change  int    `json:"change"`
	Credit  int    `json:"credit"`
	Message string `json:"message"`
}

type GameTimes struct {
	gorm.Model
	UserId  uint   `json:"userId"`
	Change  int    `json:"change"`
	Credit  int    `json:"credit"`
	Message string `json:"message"`
}

type CreditHistoriesPage struct {
	CreditHistories []*CreditHistory `json:"list"`
	TotalNum        uint             `json:"totalNum"`
	PageId          int              `json:"pageId"`
}

/*func GetCreditHistoriesByUserId(userId uint, pageId string, pageNum uint) (creditHistoriesPage CreditHistoriesPage, err error) {
	if err = db.Limit(pageNum).Where("user_id = ? AND id > ?", userId, pageId).Order("create_time desc").Find(&creditHistoriesPage.CreditHistories).Error; err == nil {
		creditHistoriesPage.PageId = creditHistoriesPage.CreditHistories
		return creditHistories, nil
	}
	return nil, err
}*/

func ModifyCreditHistory(userId uint, change int, message string) (*CreditHistory, error) {
	tx := db.Begin()
	// 根据userId（主键）查找
	user := User{}
	if err := tx.First(&user, userId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 更新User.credit
	credit := user.Credit + change
	if err := tx.Model(&user).Select("credit").Updates(map[string]interface{}{"credit": credit}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 创建积分历史记录
	creditHistory := CreditHistory{UserId: userId, Change: change, Credit: credit, Message: message}
	if err := tx.Create(creditHistory).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &creditHistory, nil
}
